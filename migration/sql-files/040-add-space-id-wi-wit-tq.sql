-- Alter the table work_items
ALTER TABLE work_items ADD space_id uuid DEFAULT '{{index . 0}}' NOT NULL;
-- Once we set the values to the default. We drop this default constraint
ALTER TABLE work_items ALTER space_id DROP DEFAULT;
ALTER TABLE work_items ADD FOREIGN KEY (space_id) REFERENCES spaces(id) ON DELETE CASCADE;

-- Create indexes

-- Alter the table work_item_types
ALTER TABLE work_item_types ADD space_id uuid DEFAULT '{{index . 0}}' NOT NULL;
-- Once we set the values to the default. We drop this default constraint
ALTER TABLE work_item_types ALTER space_id DROP DEFAULT;
ALTER TABLE work_item_types ADD FOREIGN KEY (space_id) REFERENCES spaces(id) ON DELETE CASCADE;

-- Create indexes


-- Alter the table tracker_queries
ALTER TABLE tracker_queries ADD space_id uuid DEFAULT '{{index . 0}}' NOT NULL;
-- Once we set the values to the default. We drop this default constraint
ALTER TABLE tracker_queries ALTER space_id DROP DEFAULT;
ALTER TABLE tracker_queries ADD FOREIGN KEY (space_id) REFERENCES spaces(id) ON DELETE CASCADE;

-- Create indexes
