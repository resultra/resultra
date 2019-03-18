// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package databaseWrapper

import (
	"database/sql"
	"fmt"
)

var trackerDatabaseSchema = `

CREATE TABLE IF NOT EXISTS databases (
   database_id text PRIMARY KEY,
   name text NOT NULL,
   properties text NOT NULL,
   description text NOT NULL,
   is_template boolean NOT NULL,
   is_active boolean NOT NULL DEFAULT '1',
   created_by_user_id text REFERENCES users (user_id)
);


CREATE TABLE IF NOT EXISTS users ( 
	user_id text PRIMARY KEY,
	user_name text NOT NULL,
	first_name text NOT NULL, -- TODO cannot be empty
	last_name text NOT NULL, -- TODO cannot be empty
	email_addr text NOT NULL, -- TODO must be non-empty, unique (case-insensitive)
	password_hash text NOT NULL,
	properties text NOT NULL DEFAULT '{}',
    is_active boolean NOT NULL DEFAULT '1',
	is_workspace_admin bool NOT NULL DEFAULT '0'
);

CREATE UNIQUE INDEX email_unique_index on users (LOWER(email_addr));
CREATE UNIQUE INDEX username_unique_index on users (LOWER(user_name));

CREATE TABLE IF NOT EXISTS password_reset_links (
	reset_id text PRIMARY KEY,
	user_id text REFERENCES users (user_id),
	reset_timestamp_utc timestamp NOT NULL
);

CREATE TABLE IF NOT EXISTS user_invites (
	invite_id text PRIMARY KEY,
	from_user_id text REFERENCES users (user_id),
	invite_timestamp_utc timestamp NOT NULL,
	invitee_email_addr text NOT NULL,
	invite_msg text NOT NULL
);


CREATE TABLE IF NOT EXISTS workspace_info (
	single_row_id int PRIMARY KEY DEFAULT '1',
	schema_version int NOT NULL,
	name text NOT NULL,
	properties text NOT NULL
);

-- Create the default workspace information
INSERT INTO workspace_info (schema_version,name,properties) VALUES ('1','Trackers','{}');

CREATE TABLE IF NOT EXISTS fields ( 
	field_id text PRIMARY KEY, 
	database_id text REFERENCES databases (database_id), 
	name text NOT NULL, 
	type text NOT NULL, 
	ref_name text NOT NULL, 
	calc_field_eqn text, 
	is_calc_field boolean NOT NULL, 
    is_active boolean NOT NULL DEFAULT '1',
	preprocessed_formula_text text
); 

CREATE UNIQUE INDEX field_ref_name_unique_index on fields (LOWER(ref_name),database_id);

CREATE TABLE IF NOT EXISTS globals ( 
	global_id text PRIMARY KEY,
	database_id text REFERENCES databases (database_id), 
	name text NOT NULL, 
	ref_name text NOT NULL, 
    is_active boolean NOT NULL DEFAULT '1',
	type text NOT NULL
);

CREATE UNIQUE INDEX global_ref_name_unique_index on globals (LOWER(ref_name),database_id);

CREATE TABLE IF NOT EXISTS global_updates (
	update_id text PRIMARY KEY,
	global_id text REFERENCES globals (global_id),
	update_timestamp_utc timestamp NOT NULL,
	value text NOT NULL -- value encoded as JSON
);

CREATE TABLE IF NOT EXISTS attachments (
	attachment_id text PRIMARY KEY,
	database_id text REFERENCES databases (database_id),
	user_id text REFERENCES users(user_id),
	create_timestamp_utc timestamp NOT NULL,
	orig_file_name text NOT NULL,
	type text NOT NULL,
	cloud_file_name text NOT NULL,
	title text,
	caption text
);

CREATE TABLE IF NOT EXISTS records ( 
	database_id text REFERENCES databases (database_id), 
	record_id text PRIMARY KEY,
	is_draft_record boolean NOT NULL,
	create_timestamp_utc timestamp NOT NULL,
    is_active boolean NOT NULL DEFAULT '1',
	sequence_num int NOT NULL
);

CREATE TABLE IF NOT EXISTS cell_updates (
	update_id text PRIMARY KEY,
	database_id text REFERENCES databases (database_id) ON DELETE CASCADE,
	record_id text REFERENCES records (record_id) ON DELETE CASCADE,
	field_id text REFERENCES fields (field_id) ON DELETE CASCADE,
	change_set_id text, -- used to segregate uncommitted changes made in modal dialogs, set to NULL for baseline changes.
	update_timestamp_utc timestamp NOT NULL,
	user_id text REFERENCES users(user_id),
	value text NOT NULL -- value encoded as JSON
);

CREATE TABLE IF NOT EXISTS  alerts ( 
	database_id text REFERENCES databases(database_id), 
	alert_id text PRIMARY KEY, 
	name text NOT NULL,
    is_active boolean NOT NULL DEFAULT '1',
	properties text NOT NULL
);

CREATE TABLE IF NOT EXISTS alert_notification_times (
	database_id text REFERENCES databases(database_id), 
	user_id text REFERENCES users(user_id),
	latest_alert_timestamp_utc timestamp NOT NULL
);

CREATE UNIQUE INDEX role_name_unique_index on alerts (LOWER(name),database_id);

CREATE TABLE IF NOT EXISTS  dashboards ( 
	database_id text REFERENCES databases(database_id), 
	dashboard_id text PRIMARY KEY, 
	name text NOT NULL,
    is_active boolean NOT NULL DEFAULT '1',
	properties text NOT NULL DEFAULT '{}'
);

CREATE UNIQUE INDEX dashboard_name_unique_index on dashboards (LOWER(name),database_id);


CREATE TABLE IF NOT EXISTS table_views ( 
	database_id text REFERENCES databases(database_id), 
	table_id text PRIMARY KEY, 
	name text NOT NULL,
    is_active boolean NOT NULL DEFAULT '1',
	properties text NOT NULL DEFAULT '{}'
); 

CREATE TABLE IF NOT EXISTS forms ( 
	database_id text REFERENCES databases(database_id), 
	form_id text PRIMARY KEY, 
	name text NOT NULL,
    is_active boolean NOT NULL DEFAULT '1',
	properties text NOT NULL DEFAULT '{}'
); 

CREATE UNIQUE INDEX form_name_unique_index on forms (LOWER(name),database_id);


CREATE TABLE IF NOT EXISTS item_lists ( 
	list_id text PRIMARY KEY, 
	database_id text REFERENCES databases(database_id), 
	name text NOT NULL,
    is_active boolean NOT NULL DEFAULT '1',
	properties text NOT NULL DEFAULT '{}'
); 

CREATE UNIQUE INDEX list_name_unique_index on item_lists (LOWER(name),database_id);

CREATE TABLE IF NOT EXISTS form_links (
	link_id text PRIMARY KEY,
	form_id text REFERENCES forms(form_id), 
	name text NOT NULL,
	include_in_sidebar boolean NOT NULL,
	shared_link_enabled boolean,
	shared_link_id text,
    is_active boolean NOT NULL DEFAULT '1',
	properties text NOT NULL
); 

CREATE TABLE IF NOT EXISTS value_lists (
	value_list_id text PRIMARY KEY,
	database_id text REFERENCES databases(database_id), 
	name text NOT NULL,
    is_active boolean NOT NULL DEFAULT '1',
	properties text NOT NULL DEFAULT '{}'
); 

CREATE UNIQUE INDEX value_list_name_unique_index on value_lists (LOWER(name),database_id);


CREATE TABLE IF NOT EXISTS dashboard_components (
	dashboard_id text REFERENCES dashboards(dashboard_id), 
	component_id text PRIMARY KEY,
	properties text NOT NULL DEFAULT '{}',
	type text NOT NULL
);

CREATE TABLE IF NOT EXISTS form_components (
	form_id text REFERENCES forms(form_id), 
	component_id text PRIMARY KEY,
	properties text NOT NULL DEFAULT '{}',
	type text NOT NULL
);

CREATE TABLE IF NOT EXISTS table_view_columns (
	table_id text REFERENCES table_views(table_id), 
	column_id text PRIMARY KEY,
	properties text NOT NULL DEFAULT '{}',
	type text NOT NULL
);

CREATE TABLE IF NOT EXISTS field_comments (
	comment_id text NOT NULL,
	user_id text REFERENCES users(user_id),
	record_id text  REFERENCES records (record_id) ON DELETE CASCADE, 
	field_id text REFERENCES fields(field_id) ON DELETE CASCADE,
	create_timestamp_utc timestamp NOT NULL,
	update_timestamp_utc timestamp NOT NULL,
	comment text NOT NULL
);


CREATE TABLE IF NOT EXISTS database_roles (
	database_id text REFERENCES databases(database_id), 
	role_id text PRIMARY KEY,
	name text NOT NULL
);

CREATE TABLE IF NOT EXISTS collaborators (
	collaborator_id text PRIMARY KEY,
	user_id text REFERENCES users(user_id),
	is_admin boolean NOT NULL DEFAULT '0',
	database_id text REFERENCES databases(database_id),
	UNIQUE(user_id,database_id)
);

CREATE TABLE IF NOT EXISTS collaborator_roles (
	collaborator_id text REFERENCES collaborators(collaborator_id) ON DELETE CASCADE, 
	role_id text REFERENCES database_roles(role_id) ON DELETE CASCADE
);


CREATE TABLE IF NOT EXISTS list_role_privs (
	list_id text REFERENCES item_lists(list_id) ON DELETE CASCADE, 
	role_id text REFERENCES database_roles(role_id) ON DELETE CASCADE,
	privs text NOT NULL
);


CREATE TABLE IF NOT EXISTS dashboard_role_privs (
	dashboard_id text REFERENCES dashboards(dashboard_id) ON DELETE CASCADE, 
	role_id text REFERENCES database_roles(role_id) ON DELETE CASCADE,
	privs text NOT NULL
);

CREATE TABLE IF NOT EXISTS new_item_form_link_role_privs (
	link_id text REFERENCES form_links(link_id) ON DELETE CASCADE, 
	role_id text REFERENCES database_roles(role_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS alert_role_privs (
	alert_id text REFERENCES alerts(alert_id) ON DELETE CASCADE, 
	role_id text REFERENCES database_roles(role_id) ON DELETE CASCADE
);

` // END of SCHEMA

func InitNewTrackerDatabaseToDest(trackerDBHandle *sql.DB) error {
	if _, createErr := trackerDBHandle.Exec(trackerDatabaseSchema); createErr != nil {
		return fmt.Errorf("can't initialize database: %v", createErr)
	}
	return nil

}
