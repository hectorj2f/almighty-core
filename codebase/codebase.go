package codebase

import (
	"time"

	"golang.org/x/net/context"

	"github.com/almighty/almighty-core/errors"
	"github.com/almighty/almighty-core/gormsupport"

	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	errs "github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

// CodebaseContent defines all parameters those are useful to associate Che Editor's window to a WI
type CodebaseContent struct {
	Repository string `json:"repository"`
	Branch     string `json:"branch"`
	FileName   string `json:"filename"`
	LineNumber int    `json:"linenumber"`
}

// Following keys define attribute names in the map of Codebase
const (
	RepositoryKey = "repository"
	BranchKey     = "branch"
	FileNameKey   = "filename"
	LineNumberKey = "linenumber"
)

// ToMap converts CodebaseContent to a map of string->Interface{}
func (c *CodebaseContent) ToMap() map[string]interface{} {
	res := make(map[string]interface{})
	res[RepositoryKey] = c.Repository
	res[BranchKey] = c.Branch
	res[FileNameKey] = c.FileName
	res[LineNumberKey] = c.LineNumber
	return res
}

// IsValid perform following checks
// Repository value is mandatory
func (c *CodebaseContent) IsValid() error {
	if c.Repository == "" {
		return errors.NewBadParameterError("system.codebase", RepositoryKey+" is mandatory")
	}
	return nil
}

// NewCodebaseContent builds CodebaseContent instance from input Map.
func NewCodebaseContent(value map[string]interface{}) (CodebaseContent, error) {
	cb := CodebaseContent{}
	validKeys := []string{RepositoryKey, BranchKey, FileNameKey, LineNumberKey}
	for _, key := range validKeys {
		if v, ok := value[key]; ok {
			switch key {
			case RepositoryKey:
				cb.Repository = v.(string)
			case BranchKey:
				cb.Branch = v.(string)
			case FileNameKey:
				cb.FileName = v.(string)
			case LineNumberKey:
				switch v.(type) {
				case int:
					cb.LineNumber = v.(int)
				case float64:
					y := v.(float64)
					cb.LineNumber = int(y)
				}
			}
		}
	}
	err := cb.IsValid()
	if err != nil {
		return cb, err
	}
	return cb, nil
}

// NewCodebaseContentFromValue builds CodebaseContent from interface{}
func NewCodebaseContentFromValue(value interface{}) (*CodebaseContent, error) {
	if value == nil {
		return nil, nil
	}
	switch value.(type) {
	case CodebaseContent:
		result := value.(CodebaseContent)
		return &result, nil
	case map[string]interface{}:
		result, err := NewCodebaseContent(value.(map[string]interface{}))
		if err != nil {
			return nil, err
		}
		return &result, nil
	default:
		return nil, nil
	}
}

// Codebase describes a single codebase
type Codebase struct {
	gormsupport.Lifecycle
	ID                uuid.UUID `sql:"type:uuid default uuid_generate_v4()" gorm:"primary_key"` // This is the ID PK field
	SpaceID           uuid.UUID `sql:"type:uuid"`
	Type              string
	URL               string
	StackID           string
	LastUsedWorkspace string
}

// Repository describes interactions with codebases
type Repository interface {
	Create(ctx context.Context, u *Codebase) error
	Save(ctx context.Context, codebase *Codebase) (*Codebase, error)
	List(ctx context.Context, spaceID uuid.UUID, start *int, limit *int) ([]*Codebase, uint64, error)
	Load(ctx context.Context, id uuid.UUID) (*Codebase, error)
}

// NewCodebaseRepository creates a new storage type.
func NewCodebaseRepository(db *gorm.DB) Repository {
	return &GormCodebaseRepository{db: db}
}

// GormCodebaseRepository is the implementation of the storage interface for Codebases.
type GormCodebaseRepository struct {
	db *gorm.DB
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m *GormCodebaseRepository) TableName() string {
	return "codebases"
}

// Create creates a new record.
func (m *GormCodebaseRepository) Create(ctx context.Context, codebase *Codebase) error {
	defer goa.MeasureSince([]string{"goa", "db", "codebase", "create"}, time.Now())
	if codebase.ID == uuid.Nil {
		codebase.ID = uuid.NewV4()
	}

	if err := m.db.Create(codebase).Error; err != nil {
		goa.LogError(ctx, "error adding Codebase", "error", err.Error())
		return errs.WithStack(err)
	}

	return nil
}

// Save a single codebase
func (m *GormCodebaseRepository) Save(ctx context.Context, codebase *Codebase) (*Codebase, error) {
	c := Codebase{}
	tx := m.db.Where("id=?", codebase.ID).First(&c)
	if tx.RecordNotFound() {
		// treating this as a not found error: the fact that we're using number internal is implementation detail
		return nil, errors.NewNotFoundError("codebase", codebase.ID.String())
	}
	if err := tx.Error; err != nil {
		return nil, errors.NewInternalError(err.Error())
	}

	tx = tx.Save(codebase)
	if err := tx.Error; err != nil {
		return nil, errors.NewInternalError(err.Error())
	}

	return codebase, nil
}

// List all codebases related to a single item
func (m *GormCodebaseRepository) List(ctx context.Context, spaceID uuid.UUID, start *int, limit *int) ([]*Codebase, uint64, error) {
	defer goa.MeasureSince([]string{"goa", "db", "codebase", "query"}, time.Now())

	db := m.db.Model(&Codebase{}).Where("space_id = ?", spaceID)
	orgDB := db
	if start != nil {
		if *start < 0 {
			return nil, 0, errors.NewBadParameterError("start", *start)
		}
		db = db.Offset(*start)
	}
	if limit != nil {
		if *limit <= 0 {
			return nil, 0, errors.NewBadParameterError("limit", *limit)
		}
		db = db.Limit(*limit)
	}
	db = db.Select("count(*) over () as cnt2 , *")

	rows, err := db.Rows()
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	result := []*Codebase{}
	columns, err := rows.Columns()
	if err != nil {
		return nil, 0, errors.NewInternalError(err.Error())
	}

	// need to set up a result for Scan() in order to extract total count.
	var count uint64
	var ignore interface{}
	columnValues := make([]interface{}, len(columns))

	for index := range columnValues {
		columnValues[index] = &ignore
	}
	columnValues[0] = &count
	first := true

	for rows.Next() {
		value := &Codebase{}
		db.ScanRows(rows, value)
		if first {
			first = false
			if err = rows.Scan(columnValues...); err != nil {
				return nil, 0, errors.NewInternalError(err.Error())
			}
		}
		result = append(result, value)

	}
	if first {
		// means 0 rows were returned from the first query (maybe becaus of offset outside of total count),
		// need to do a count(*) to find out total
		orgDB := orgDB.Select("count(*)")
		rows2, err := orgDB.Rows()
		defer rows2.Close()
		if err != nil {
			return nil, 0, err
		}
		rows2.Next() // count(*) will always return a row
		rows2.Scan(&count)
	}
	return result, count, nil
}

// Load a single codebase regardless of parent
func (m *GormCodebaseRepository) Load(ctx context.Context, id uuid.UUID) (*Codebase, error) {
	defer goa.MeasureSince([]string{"goa", "db", "codebase", "get"}, time.Now())
	var obj Codebase

	tx := m.db.Where("id=?", id).First(&obj)
	if tx.RecordNotFound() {
		return nil, errors.NewNotFoundError("codebase", id.String())
	}
	if tx.Error != nil {
		return nil, errors.NewInternalError(tx.Error.Error())
	}
	return &obj, nil
}
