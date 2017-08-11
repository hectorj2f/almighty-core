package migration_test

import (
	"fmt"
	"strings"
)

var __044_insert_test_data_sql = []byte(`-- users
INSERT INTO
   users(created_at, updated_at, id, email, full_name, image_url, bio, url, context_information)
VALUES
   (
      now(), now(), '01b291cd-9399-4f1a-8bbc-d1de66d76192', 'testone@example.com', 'test one', 'https://www.gravatar.com/avatar/testone', 'my test bio one', 'http://example.com', '{"key": "value"}'
   ),
   (
      now(), now(), '0d19928e-ef61-46fd-9bdc-71d1ecbce2c7', 'testtwo@example.com', 'test two', 'http://https://www.gravatar.com/avatar/testtwo', 'my test bio two', 'http://example.com', '{"key": "value"}'
   )
;
-- identities
INSERT INTO
   identities(created_at, updated_at, id, username, provider_type, user_id, profile_url)
VALUES
   (
      now(), now(), '01b291cd-9399-4f1a-8bbc-d1de66d76192', 'testone', 'github', '01b291cd-9399-4f1a-8bbc-d1de66d76192', 'http://example-github.com'
   ),
   (
      now(), now(), '5f946975-ff47-4c4a-b5dc-778f0b7e476c', 'testwo', 'rhhd', '0d19928e-ef61-46fd-9bdc-71d1ecbce2c7', 'http://example-rhd.com'
   )
;
-- spaces
INSERT INTO
   spaces (created_at, updated_at, id, version, name, description, owner_id)
VALUES
   (
      now(), now(), '86af5178-9b41-469b-9096-57e5155c3f31', 0, 'test.space.one', 'space desc one', '01b291cd-9399-4f1a-8bbc-d1de66d76192'
   )
;
-- work_item_types
INSERT INTO
   work_item_types(created_at, updated_at, id, name, version, fields, space_id)
VALUES
   (
      now(), now(), 'bbf35418-04b6-426c-a60b-7f80beb0b624', 'Test item type 1', 1.0, '{}', '2e0698d8-753e-4cef-bb7c-f027634824a2'
   )
;
INSERT INTO
   work_item_types(created_at, updated_at, id, name, version, path, fields, space_id)
VALUES
   (
      now(), now(), '86af5178-9b41-469b-9096-57e5155c3f31', 'Test item type 2', 1.0, 'bbf35418_04b6_426c_a60b_7f80beb0b624.86af5178_9b41_469b_9096_57e5155c3f31', '{}', '86af5178-9b41-469b-9096-57e5155c3f31'
   )
;
-- trackers
INSERT INTO
   trackers(created_at, updated_at, id, url, type)
VALUES
   (
      now(), now(), 1, 'http://example.com', 'github'
   ),
   (
      now(), now(), 2, 'http://example-jira.com', 'jira'
   )
;
-- tracker_queries id | query | schedule | tracker_id | space_id
INSERT INTO
   tracker_queries(created_at, updated_at, id, query, schedule, tracker_id, space_id)
VALUES
   (
      now(), now(), 1, 'SELECT * FROM', 'schedule', 1, '86af5178-9b41-469b-9096-57e5155c3f31'
   ),
   (
      now(), now(), 2, 'SELECT * FROM', 'schedule', 2, '86af5178-9b41-469b-9096-57e5155c3f31'
   )
;

-- space_resources
INSERT INTO
   space_resources(created_at, updated_at, id, space_id, resource_id, policy_id, permission_id)
VALUES
   (
      now(), now(), '2e0698d8-753e-4cef-bb7c-f027634824a2', '86af5178-9b41-469b-9096-57e5155c3f31', 'resource_id', 'policy_id', 'permission_id'
   ),
   (
      now(), now(), '71171e90-6d35-498f-a6a7-2083b5267c18', '86af5178-9b41-469b-9096-57e5155c3f31', 'resource_id', 'policy_id', 'permission_id'
   )
;
-- areas created_at | updated_at | deleted_at | id | space_id | version | path | name
INSERT INTO
   areas(created_at, updated_at, id, space_id, version, path, name)
VALUES
   (
      now(), now(), '2e0698d8-753e-4cef-bb7c-f027634824a2', '86af5178-9b41-469b-9096-57e5155c3f31', 0, 'path', 'area test one'
   ),
   (
      now(), now(), '71171e90-6d35-498f-a6a7-2083b5267c18', '86af5178-9b41-469b-9096-57e5155c3f31', 0, '', 'area test two'
   )
;
-- iterations
INSERT INTO
   iterations(created_at, updated_at, id, space_id, start_at, end_at, name, description, state)
VALUES
   (
      now(), now(), '71171e90-6d35-498f-a6a7-2083b5267c18', '86af5178-9b41-469b-9096-57e5155c3f31', now(), now(), 'iteration test one', 'description', 'new'
   ),
   (
      now(), now(), '2e0698d8-753e-4cef-bb7c-f027634824a2', '86af5178-9b41-469b-9096-57e5155c3f31', now(), now(), 'iteration test two', 'description', 'start'
   )
;
-- comments
INSERT INTO
   comments(created_at, updated_at, id, parent_id, body, created_by, markup)
VALUES
   (
      now(), now(), '71171e90-6d35-498f-a6a7-2083b5267c18', '2e0698d8-753e-4cef-bb7c-f027634824a2', 'body test one', '01b291cd-9399-4f1a-8bbc-d1de66d76192', 'PlainText'
   ),
   (
      now(), now(), '2e0698d8-753e-4cef-bb7c-f027634824a2', '2e0698d8-753e-4cef-bb7c-f027634824a2', 'body test two', '01b291cd-9399-4f1a-8bbc-d1de66d76192', 'PlainText'
   )
;
-- comment_revisions
INSERT INTO
   comment_revisions(id, revision_time, revision_type, modifier_id, comment_id, comment_body, comment_markup, comment_parent_id)
VALUES
   (
      '71171e90-6d35-498f-a6a7-2083b5267c18', now(), 1, '5f946975-ff47-4c4a-b5dc-778f0b7e476c', '71171e90-6d35-498f-a6a7-2083b5267c18', 'comment body test one', 'comment markup test one', '71171e90-6d35-498f-a6a7-2083b5267c18'
   ),
   (
      '2e0698d8-753e-4cef-bb7c-f027634824a2', now(), 1, '5f946975-ff47-4c4a-b5dc-778f0b7e476c', '71171e90-6d35-498f-a6a7-2083b5267c18', 'comment body test two', 'comment markup test two', '71171e90-6d35-498f-a6a7-2083b5267c18'
   )
;
-- work_item_link_categories
INSERT INTO
   work_item_link_categories(created_at, updated_at, id, version, name, description)
VALUES
   (
      now(), now(), '71171e90-6d35-498f-a6a7-2083b5267c18', 1, 'name test one', 'description one'
   ),
   (
      now(), now(), '2e0698d8-753e-4cef-bb7c-f027634824a2', 1, 'name test two', 'description two'
   )
;
-- work_item_link_types
INSERT INTO
   work_item_link_types(created_at, updated_at, id, version, name, description, forward_name, reverse_name, topology, link_category_id, space_id, source_type_id, target_type_id)
VALUES
   (
      now(), now(), '2e0698d8-753e-4cef-bb7c-f027634824a2', 1, 'test one', 'desc', 'forward test one', 'reverser test one', 'dependency', '71171e90-6d35-498f-a6a7-2083b5267c18', '2e0698d8-753e-4cef-bb7c-f027634824a2', '86af5178-9b41-469b-9096-57e5155c3f31', '86af5178-9b41-469b-9096-57e5155c3f31'
   ),
   (
      now(), now(), '71171e90-6d35-498f-a6a7-2083b5267c18', 1, 'test two', 'desc', 'forward test two', 'reverser test two', 'network', '2e0698d8-753e-4cef-bb7c-f027634824a2', '2e0698d8-753e-4cef-bb7c-f027634824a2', '86af5178-9b41-469b-9096-57e5155c3f31', '86af5178-9b41-469b-9096-57e5155c3f31'
   )
;
-- work_items
INSERT INTO
   work_items(created_at, updated_at, type, version, space_id, fields)
VALUES
   (
      now(), now(), 'bbf35418-04b6-426c-a60b-7f80beb0b624', 1.0, '86af5178-9b41-469b-9096-57e5155c3f31', '{}'
   ),
   (
      now(), now(), 'bbf35418-04b6-426c-a60b-7f80beb0b624', 2.0, '86af5178-9b41-469b-9096-57e5155c3f31', '{}'
   )
;
-- work_item_revisions
INSERT INTO
   work_item_revisions(id, revision_time, revision_type, modifier_id, work_item_id, work_item_type_id, work_item_version, work_item_fields)
VALUES
   (
      '2e0698d8-753e-4cef-bb7c-f027634824a2', now(), 1, '01b291cd-9399-4f1a-8bbc-d1de66d76192', 1, '2e0698d8-753e-4cef-bb7c-f027634824a2', 1, '{}'
   ),
   (
      '71171e90-6d35-498f-a6a7-2083b5267c18', now(), 1, '01b291cd-9399-4f1a-8bbc-d1de66d76192', 1, '2e0698d8-753e-4cef-bb7c-f027634824a2', 1, '{}'
   )
;
-- work_item_links
INSERT INTO
   work_item_links(created_at, updated_at, id, version, link_type_id)
VALUES
   (
      now(), now(), '2e0698d8-753e-4cef-bb7c-f027634824a2', 1, '2e0698d8-753e-4cef-bb7c-f027634824a2'
   ),
   (
      now(), now(), '71171e90-6d35-498f-a6a7-2083b5267c18', 1, '71171e90-6d35-498f-a6a7-2083b5267c18'
   )
;
-- work_item_link_revisions
INSERT INTO
   work_item_link_revisions(id, revision_time, revision_type, modifier_id, work_item_link_id, work_item_link_version, work_item_link_source_id, work_item_link_target_id, work_item_link_type_id)
VALUES
   (
      '71171e90-6d35-498f-a6a7-2083b5267c18', now(), 1, '01b291cd-9399-4f1a-8bbc-d1de66d76192', '71171e90-6d35-498f-a6a7-2083b5267c18', 1, 1, 2, '2e0698d8-753e-4cef-bb7c-f027634824a2'
   ),
   (
      '2e0698d8-753e-4cef-bb7c-f027634824a2', now(), 2, '01b291cd-9399-4f1a-8bbc-d1de66d76192', '71171e90-6d35-498f-a6a7-2083b5267c18', 1, 2, 1, '2e0698d8-753e-4cef-bb7c-f027634824a2'
   )
;
-- tracker_items
INSERT INTO
   tracker_items(created_at, updated_at, id, remote_item_id, item, batch_id, tracker_id)
VALUES
   (
      now(), now(), 1, 'remote_id', 'test one', 'batch_id', 1
   ),
   (
      now(), now(), 2, 'remote_id', 'test two', 'batch_id', 2
   )
;
`)

func _044_insert_test_data_sql() ([]byte, error) {
	return __044_insert_test_data_sql, nil
}

var __045_update_work_items_sql = []byte(`-- work_items
UPDATE work_items SET execution_order=100000 WHERE type='bbf35418-04b6-426c-a60b-7f80beb0b624';
UPDATE work_items SET execution_order=200000 WHERE type='bbf35418-04b6-426c-a60b-7f80beb0b624';
`)

func _045_update_work_items_sql() ([]byte, error) {
	return __045_update_work_items_sql, nil
}

var __046_insert_oauth_states_sql = []byte(`-- oauth_state_references
INSERT INTO
   oauth_state_references(created_at, updated_at, id, referrer)
VALUES
   (
      now(), now(), '2e0698d8-753e-4cef-bb7c-f027634824a2', 'test referrer one text'
   )
;
INSERT INTO
   oauth_state_references(created_at, updated_at, id, referrer)
VALUES
   (
      now(), now(), '71171e90-6d35-498f-a6a7-2083b5267c18', 'test referrer two text'
   )
;
`)

func _046_insert_oauth_states_sql() ([]byte, error) {
	return __046_insert_oauth_states_sql, nil
}

var __047_insert_codebases_sql = []byte(`-- codebases
INSERT INTO
   codebases(created_at, updated_at, id, space_id, type, url)
VALUES
   (
      now(), now(), '2e0698d8-753e-4cef-bb7c-f027634824a2', '86af5178-9b41-469b-9096-57e5155c3f31', 'type test one', 'http://example-jira.com'
   )
;
INSERT INTO
   codebases(created_at, updated_at, id, space_id, type, url)
VALUES
   (
      now(), now(), '71171e90-6d35-498f-a6a7-2083b5267c18', '86af5178-9b41-469b-9096-57e5155c3f31', 'type test two', 'http://example-jira.com'
   )
;
`)

func _047_insert_codebases_sql() ([]byte, error) {
	return __047_insert_codebases_sql, nil
}

var __048_unique_idx_failed_insert_sql = []byte(`-- insert two iterations one will fail due to invalid iterations_name_space_id_path_unique
INSERT INTO
   iterations(created_at, updated_at, id, space_id, start_at, end_at, name, description, state, path)
VALUES
   (
      now(), now(), '86af5178-9b41-469b-9096-57e5155c3f31', '86af5178-9b41-469b-9096-57e5155c3f31', now(), now(), 'iteration test one', 'description', 'new', '/'
   )
;

INSERT INTO
   iterations(created_at, updated_at, id, space_id, start_at, end_at, name, description, state, path)
VALUES
   (
      now(), now(), '0a24d3c2-e0a6-4686-8051-ec0ea1915a28', '86af5178-9b41-469b-9096-57e5155c3f31', now(), now(), 'iteration test one', 'description', 'new', '/'
   )
;
`)

func _048_unique_idx_failed_insert_sql() ([]byte, error) {
	return __048_unique_idx_failed_insert_sql, nil
}

var __050_users_add_column_company_sql = []byte(`-- Set company value to the existing users
UPDATE users SET company='RedHat Inc.' WHERE full_name='test one' OR full_name='test two';
`)

func _050_users_add_column_company_sql() ([]byte, error) {
	return __050_users_add_column_company_sql, nil
}

var __053_edit_username_sql = []byte(`-- users
INSERT INTO
   users(created_at, updated_at, id, email, full_name, image_url, bio, url, context_information)
VALUES
   (
      now(), now(), 'f03f023b-0427-4cdb-924b-fb2369018ab7', 'test2@example.com', 'test1', 'https://www.gravatar.com/avatar/testtwo2', 'my test bio one', 'http://example.com/001', '{"key": "value"}'
   ),
   (
      now(), now(), 'f03f023b-0427-4cdb-924b-fb2369018ab6', 'test3@example.com', 'test2', 'http://https://www.gravatar.com/avatar/testtwo3', 'my test bio two', 'http://example.com/002', '{"key": "value"}'
   )
;
-- identities
INSERT INTO
   identities(created_at, updated_at, id, username, provider_type, user_id, profile_url)
VALUES
   (
      now(), now(), '2a808366-9525-4646-9c80-ed704b2eebbe', 'test1', 'github', 'f03f023b-0427-4cdb-924b-fb2369018ab7', 'http://example-github.com/001'
   ),
   (
      now(), now(), '2a808366-9525-4646-9c80-ed704b2eebbb', 'test2', 'rhhd', 'f03f023b-0427-4cdb-924b-fb2369018ab6', 'http://example-rhd.com/002'
   )
;
`)

func _053_edit_username_sql() ([]byte, error) {
	return __053_edit_username_sql, nil
}

var __054_add_stackid_to_codebase_sql = []byte(`UPDATE codebases set stack_id ='java-centos';
`)

func _054_add_stackid_to_codebase_sql() ([]byte, error) {
	return __054_add_stackid_to_codebase_sql, nil
}

var __055_assign_root_area_if_missing_sql = []byte(`insert into spaces (id, name) values ('11111111-2222-0000-0000-000000000000', 'test');
insert into areas (id, name, path, space_id) values ('11111111-3333-0000-0000-000000000000', 'test area', '', '11111111-2222-0000-0000-000000000000');
insert into work_item_types (id, name, space_id) values ('11111111-4444-0000-0000-000000000000', 'Test WIT','11111111-2222-0000-0000-000000000000');
insert into work_items (id, space_id, type, fields) values (12345, '11111111-2222-0000-0000-000000000000', '11111111-4444-0000-0000-000000000000', '{"system.title":"Title"}'::json);`)

func _055_assign_root_area_if_missing_sql() ([]byte, error) {
	return __055_assign_root_area_if_missing_sql, nil
}

var __056_assign_root_iteration_if_missing_sql = []byte(`insert into spaces (id, name) values ('11111111-2222-bbbb-0000-000000000000', 'test');
insert into iterations (id, name, path, space_id) values ('11111111-3333-bbbb-0000-000000000000', 'test area', '', '11111111-2222-bbbb-0000-000000000000');
insert into work_item_types (id, name, space_id) values ('11111111-4444-bbbb-0000-000000000000', 'Test WIT','11111111-2222-bbbb-0000-000000000000');
insert into work_items (id, space_id, type, fields) values (12346, '11111111-2222-bbbb-0000-000000000000', '11111111-4444-bbbb-0000-000000000000', '{"system.title":"Title"}'::json);`)

func _056_assign_root_iteration_if_missing_sql() ([]byte, error) {
	return __056_assign_root_iteration_if_missing_sql, nil
}

var __057_add_last_used_workspace_to_codebase_sql = []byte(`UPDATE codebases set last_used_workspace ='java-centos-last-workspace';
`)

func _057_add_last_used_workspace_to_codebase_sql() ([]byte, error) {
	return __057_add_last_used_workspace_to_codebase_sql, nil
}

var __061_add_duplicate_space_owner_name_sql = []byte(`--- added a duplicate space with the same owner and name than a previous one
INSERT INTO
   spaces (created_at, updated_at, id, version, name, description, owner_id)
VALUES
   (
      now(), now(), '86af5178-9b41-469b-9096-57e5155c3f32', 0, 'test.Space.one', 'Space desc one', '01b291cd-9399-4f1a-8bbc-d1de66d76192'
   )
;
`)

func _061_add_duplicate_space_owner_name_sql() ([]byte, error) {
	return __061_add_duplicate_space_owner_name_sql, nil
}

var __063_workitem_related_changes_sql = []byte(`--
-- comments
--
insert into spaces (id, name) values ('11111111-6262-0000-0000-000000000000', 'test');
insert into work_item_types (id, name, space_id) values ('11111111-6262-0000-0000-000000000000', 'Test WIT','11111111-6262-0000-0000-000000000000');
insert into work_items (id, space_id, type, fields) values (62001, '11111111-6262-0000-0000-000000000000', '11111111-6262-0000-0000-000000000000', '{"system.title":"Work item 1"}'::json);
insert into work_items (id, space_id, type, fields) values (62002, '11111111-6262-0000-0000-000000000000', '11111111-6262-0000-0000-000000000000', '{"system.title":"Work item 2"}'::json);
insert into work_items (id, space_id, type, fields) values (62003, '11111111-6262-0000-0000-000000000000', '11111111-6262-0000-0000-000000000000', '{"system.title":"Work item 3"}'::json);
-- remove previous comments
delete from comments;
-- add comments linked to work items above
insert into comments (id, parent_id, body, created_at) values ( '11111111-6262-0001-0000-000000000000', '62001', 'a comment', '2017-06-13 09:00:00.0000+00');
insert into comments (id, parent_id, body, created_at) values ( '11111111-6262-0003-0000-000000000000', '62003', 'a comment', '2017-06-13 11:00:00.0000+00');
update comments set deleted_at = '2017-06-13 11:15:00.0000+00' where id =  '11111111-6262-0003-0000-000000000000';

--
-- work item links
--
insert into work_items (id, space_id, type, fields) values (62004, '11111111-6262-0000-0000-000000000000', '11111111-6262-0000-0000-000000000000', '{"system.title":"Work item 3"}'::json);
insert into work_items (id, space_id, type, fields) values (62005, '11111111-6262-0000-0000-000000000000', '11111111-6262-0000-0000-000000000000', '{"system.title":"Work item 3"}'::json);
insert into work_items (id, space_id, type, fields) values (62006, '11111111-6262-0000-0000-000000000000', '11111111-6262-0000-0000-000000000000', '{"system.title":"Work item 3"}'::json);
insert into work_items (id, space_id, type, fields) values (62007, '11111111-6262-0000-0000-000000000000', '11111111-6262-0000-0000-000000000000', '{"system.title":"Work item 3"}'::json);
insert into work_items (id, space_id, type, fields) values (62008, '11111111-6262-0000-0000-000000000000', '11111111-6262-0000-0000-000000000000', '{"system.title":"Work item 3"}'::json);
insert into work_items (id, space_id, type, fields) values (62009, '11111111-6262-0000-0000-000000000000', '11111111-6262-0000-0000-000000000000', '{"system.title":"Work item 3"}'::json);
delete from work_item_links;
insert into work_item_links (id, version, source_id, target_id, created_at) values ('11111111-6262-0001-0000-000000000000', 1, 62004, 62005, '2017-06-13 09:00:00.0000+00');
insert into work_item_links (id, version, source_id, target_id, deleted_at) values ('11111111-6262-0003-0000-000000000000', 1, 62008, 62009, '2017-06-13 11:00:00.0000+00');
update work_item_links set deleted_at = '2017-06-13 11:15:00.0000+00' where id = '11111111-6262-0003-0000-000000000000';


`)

func _063_workitem_related_changes_sql() ([]byte, error) {
	return __063_workitem_related_changes_sql, nil
}

var __065_workitem_id_unique_per_space_sql = []byte(`-- create spaces 1 and 2
insert into spaces (id, name) values ('11111111-0000-0000-0000-000000000000', 'test space 1');
insert into spaces (id, name) values ('22222222-0000-0000-0000-000000000000', 'test space 2');
-- create work item types for spaces 1 and 2
insert into work_item_types (id, name, space_id) values ('11111111-0000-0000-0000-000000000000', 'test type 1', '11111111-0000-0000-0000-000000000000');
insert into work_item_types (id, name, space_id) values ('22222222-0000-0000-0000-000000000000', 'test type 2', '22222222-0000-0000-0000-000000000000');
-- create work item link types for spaces 1 and 2
insert into work_item_link_types (id, name, topology, forward_name, reverse_name, space_id) 
    values ('11111111-0000-0000-0000-000000000000', 'foo', 'dependency', 'foo', 'foo', '11111111-0000-0000-0000-000000000000');
insert into work_item_link_types (id, name, topology, forward_name, reverse_name, space_id) 
    values ('22222222-0000-0000-0000-000000000000', 'bar', 'dependency', 'bar', 'bar', '22222222-0000-0000-0000-000000000000');
-- create identity (for revisions)
insert into identities (id, username) values ('cafebabe-0000-0000-0000-000000000000', 'foo');
-- inserting work items, their revisions and comments in space '1'
insert into work_items (id, type, space_id) values (12347, '11111111-0000-0000-0000-000000000000', '11111111-0000-0000-0000-000000000000');
insert into work_item_revisions (revision_type, modifier_id, work_item_id) values (1, 'cafebabe-0000-0000-0000-000000000000', 12347);
insert into work_item_revisions (revision_type, modifier_id, work_item_id) values (2, 'cafebabe-0000-0000-0000-000000000000', 12347);
insert into comments (parent_id, body) values ('12347', 'blabla');
insert into work_items (id, type, space_id) values (12348, '11111111-0000-0000-0000-000000000000', '11111111-0000-0000-0000-000000000000');
insert into work_item_revisions (revision_type, modifier_id, work_item_id) values (1, 'cafebabe-0000-0000-0000-000000000000', 12348);
insert into work_item_revisions (revision_type, modifier_id, work_item_id) values (2, 'cafebabe-0000-0000-0000-000000000000', 12348);
insert into comments (parent_id, body) values ('12348', 'blabla');
-- inserting work items, their revisions and comments in space '2'
insert into work_items (id, type, space_id) values (12349, '22222222-0000-0000-0000-000000000000', '22222222-0000-0000-0000-000000000000');
insert into work_item_revisions (revision_type, modifier_id, work_item_id) values (1, 'cafebabe-0000-0000-0000-000000000000', 12349);
insert into work_item_revisions (revision_type, modifier_id, work_item_id) values (2, 'cafebabe-0000-0000-0000-000000000000', 12349);
insert into comments (parent_id, body) values ('12349', 'blabla');
insert into work_items (id, type, space_id) values (12350, '22222222-0000-0000-0000-000000000000', '22222222-0000-0000-0000-000000000000');
insert into work_item_revisions (revision_type, modifier_id, work_item_id) values (1, 'cafebabe-0000-0000-0000-000000000000', 12350);
insert into work_item_revisions (revision_type, modifier_id, work_item_id) values (2, 'cafebabe-0000-0000-0000-000000000000', 12350);
insert into comments (parent_id, body) values ('12350', 'blabla');
-- insert links between work items
insert into work_item_links (id, link_type_id, source_id, target_id) values ('11111111-0000-0000-0000-000000000000', '11111111-0000-0000-0000-000000000000', 12347, 12348);
insert into work_item_link_revisions (revision_type, modifier_id, work_item_link_id, work_item_link_version, work_item_link_source_id, work_item_link_target_id, work_item_link_type_id)
  values (1, 'cafebabe-0000-0000-0000-000000000000', '11111111-0000-0000-0000-000000000000',0,12347,12348,'11111111-0000-0000-0000-000000000000');
insert into work_item_links (id, link_type_id, source_id, target_id) values ('22222222-0000-0000-0000-000000000000', '22222222-0000-0000-0000-000000000000', 12349, 12350);
insert into work_item_link_revisions (revision_type, modifier_id, work_item_link_id, work_item_link_version, work_item_link_source_id, work_item_link_target_id, work_item_link_type_id)
  values (1, 'cafebabe-0000-0000-0000-000000000000', '22222222-0000-0000-0000-000000000000',0,12349,12350,'22222222-0000-0000-0000-000000000000');
`)

func _065_workitem_id_unique_per_space_sql() ([]byte, error) {
	return __065_workitem_id_unique_per_space_sql, nil
}

var __066_work_item_links_data_integrity_sql = []byte(`-- prepare data
insert into spaces (id, name) values ('00000066-0000-0000-0000-000000000000', 'test space 1');
insert into work_item_types (id, name, space_id) values ('00000066-0000-0000-0000-000000000000', 'test type 1', '00000066-0000-0000-0000-000000000000');
insert into work_item_link_types (id, name, topology, forward_name, reverse_name, space_id) 
    values ('00000066-0000-0000-0000-000000000000', 'foo', 'dependency', 'foo', 'foo', '00000066-0000-0000-0000-000000000000');
insert into work_items (id, type, space_id) values ('00000066-0000-0000-0000-000000000001', '00000066-0000-0000-0000-000000000000', '00000066-0000-0000-0000-000000000000');
insert into work_items (id, type, space_id) values ('00000066-0000-0000-0000-000000000002', '00000066-0000-0000-0000-000000000000', '00000066-0000-0000-0000-000000000000');
-- insert valid and invalid links
insert into work_item_links (id, link_type_id, source_id, target_id) values ('00000066-0000-0000-0000-000000000001', '00000066-0000-0000-0000-000000000000', '00000066-0000-0000-0000-000000000001', '00000066-0000-0000-0000-000000000002');
insert into work_item_links (id, link_type_id, source_id, target_id) values ('00000066-0000-0000-0000-000000000002', NULL, '00000066-0000-0000-0000-000000000001', '00000066-0000-0000-0000-000000000002');
insert into work_item_links (id, link_type_id, source_id, target_id) values ('00000066-0000-0000-0000-000000000003', '00000066-0000-0000-0000-000000000000', NULL, '00000066-0000-0000-0000-000000000002');
insert into work_item_links (id, link_type_id, source_id, target_id) values ('00000066-0000-0000-0000-000000000004', '00000066-0000-0000-0000-000000000000', '00000066-0000-0000-0000-000000000001', NULL);
`)

func _066_work_item_links_data_integrity_sql() ([]byte, error) {
	return __066_work_item_links_data_integrity_sql, nil
}

var __067_comment_parentid_uuid_sql = []byte(`-- need some work items to migrate the comment_revisions table, too
insert into spaces (id, name) values ('00000067-0000-0000-0000-000000000000', 'test space 67');
insert into work_item_types (id, name, space_id) values ('00000067-0000-0000-0000-000000000000', 'test type 1', '00000067-0000-0000-0000-000000000000');
insert into work_items (id, number, type, space_id) values ('00000067-0000-0000-0000-000000000000', 1, '00000067-0000-0000-0000-000000000000', '00000067-0000-0000-0000-000000000000');
insert into comments (id, parent_id, body) values ('00000067-0000-0000-0000-000000000000', '00000067-0000-0000-0000-000000000000', 'a foo comment');
insert into comment_revisions (id, revision_type, modifier_id, comment_id, comment_parent_id, comment_body) 
    values ('00000067-0000-0000-0000-000000000000', 1, 'cafebabe-0000-0000-0000-000000000000', '00000067-0000-0000-0000-000000000000',  1, 'a foo comment');`)

func _067_comment_parentid_uuid_sql() ([]byte, error) {
	return __067_comment_parentid_uuid_sql, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() ([]byte, error){
	"044-insert-test-data.sql":                    _044_insert_test_data_sql,
	"045-update-work-items.sql":                   _045_update_work_items_sql,
	"046-insert-oauth-states.sql":                 _046_insert_oauth_states_sql,
	"047-insert-codebases.sql":                    _047_insert_codebases_sql,
	"048-unique-idx-failed-insert.sql":            _048_unique_idx_failed_insert_sql,
	"050-users-add-column-company.sql":            _050_users_add_column_company_sql,
	"053-edit-username.sql":                       _053_edit_username_sql,
	"054-add-stackid-to-codebase.sql":             _054_add_stackid_to_codebase_sql,
	"055-assign-root-area-if-missing.sql":         _055_assign_root_area_if_missing_sql,
	"056-assign-root-iteration-if-missing.sql":    _056_assign_root_iteration_if_missing_sql,
	"057-add-last-used-workspace-to-codebase.sql": _057_add_last_used_workspace_to_codebase_sql,
	"061-add-duplicate-space-owner-name.sql":      _061_add_duplicate_space_owner_name_sql,
	"063-workitem-related-changes.sql":            _063_workitem_related_changes_sql,
	"065-workitem-id-unique-per-space.sql":        _065_workitem_id_unique_per_space_sql,
	"066-work_item_links_data_integrity.sql":      _066_work_item_links_data_integrity_sql,
	"067-comment-parentid-uuid.sql":               _067_comment_parentid_uuid_sql,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func     func() ([]byte, error)
	Children map[string]*_bintree_t
}

var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"044-insert-test-data.sql":                    {_044_insert_test_data_sql, map[string]*_bintree_t{}},
	"045-update-work-items.sql":                   {_045_update_work_items_sql, map[string]*_bintree_t{}},
	"046-insert-oauth-states.sql":                 {_046_insert_oauth_states_sql, map[string]*_bintree_t{}},
	"047-insert-codebases.sql":                    {_047_insert_codebases_sql, map[string]*_bintree_t{}},
	"048-unique-idx-failed-insert.sql":            {_048_unique_idx_failed_insert_sql, map[string]*_bintree_t{}},
	"050-users-add-column-company.sql":            {_050_users_add_column_company_sql, map[string]*_bintree_t{}},
	"053-edit-username.sql":                       {_053_edit_username_sql, map[string]*_bintree_t{}},
	"054-add-stackid-to-codebase.sql":             {_054_add_stackid_to_codebase_sql, map[string]*_bintree_t{}},
	"055-assign-root-area-if-missing.sql":         {_055_assign_root_area_if_missing_sql, map[string]*_bintree_t{}},
	"056-assign-root-iteration-if-missing.sql":    {_056_assign_root_iteration_if_missing_sql, map[string]*_bintree_t{}},
	"057-add-last-used-workspace-to-codebase.sql": {_057_add_last_used_workspace_to_codebase_sql, map[string]*_bintree_t{}},
	"061-add-duplicate-space-owner-name.sql":      {_061_add_duplicate_space_owner_name_sql, map[string]*_bintree_t{}},
	"063-workitem-related-changes.sql":            {_063_workitem_related_changes_sql, map[string]*_bintree_t{}},
	"065-workitem-id-unique-per-space.sql":        {_065_workitem_id_unique_per_space_sql, map[string]*_bintree_t{}},
	"066-work_item_links_data_integrity.sql":      {_066_work_item_links_data_integrity_sql, map[string]*_bintree_t{}},
	"067-comment-parentid-uuid.sql":               {_067_comment_parentid_uuid_sql, map[string]*_bintree_t{}},
}}
