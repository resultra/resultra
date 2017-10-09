CREATE TABLE databases (
   database_id text PRIMARY KEY,
   name text NOT NULL,
   properties text NOT NULL
);

CREATE TABLE users ( 
	user_id text PRIMARY KEY,
	user_name text NOT NULL,
	first_name text NOT NULL, -- TODO cannot be empty
	last_name text NOT NULL, -- TODO cannot be empty
	email_addr text NOT NULL, -- TODO must be non-empty, unique (case-insensitive)
	password_hash text NOT NULL
);

CREATE TABLE fields ( 
	field_id text PRIMARY KEY, 
	database_id text REFERENCES databases (database_id), 
	name text NOT NULL, 
	type text NOT NULL, 
	ref_name text NOT NULL, 
	calc_field_eqn text, 
	is_calc_field boolean NOT NULL, 
	preprocessed_formula_text text
); 

CREATE TABLE globals ( 
	global_id text PRIMARY KEY,
	database_id text REFERENCES databases (database_id), 
	name text NOT NULL, 
	ref_name text NOT NULL, 
	type text NOT NULL
);

CREATE TABLE global_updates (
	update_id text PRIMARY KEY,
	global_id text REFERENCES globals (global_id),
	update_timestamp_utc timestamp NOT NULL,
	value text NOT NULL -- value encoded as JSON
);

CREATE TABLE attachments (
	attachment_id text PRIMARY KEY,
	database_id text REFERENCES databases (database_id),
	user_id text NOT NULL,
	create_timestamp_utc timestamp NOT NULL,
	orig_file_name text NOT NULL,
	type text NOT NULL,
	cloud_file_name text NOT NULL,
	title text,
	caption text
);

CREATE TABLE records ( 
	database_id text REFERENCES databases (database_id), 
	record_id text PRIMARY KEY,
	is_draft_record boolean NOT NULL,
	create_timestamp_utc timestamp NOT NULL,
	sequence_num int NOT NULL
);

CREATE TABLE cell_updates (
	update_id text PRIMARY KEY,
	database_id text REFERENCES databases (database_id) ON DELETE CASCADE,
	record_id text REFERENCES records (record_id) ON DELETE CASCADE,
	field_id text REFERENCES fields (field_id) ON DELETE CASCADE,
	change_set_id text, -- used to segregate uncommitted changes made in modal dialogs, set to NULL for baseline changes.
	update_timestamp_utc timestamp NOT NULL,
	user_id text REFERENCES users(user_id),
	value text NOT NULL -- value encoded as JSON
);

CREATE TABLE  alerts ( 
	database_id text REFERENCES databases(database_id), 
	alert_id text PRIMARY KEY, 
	name text NOT NULL,
	properties text NOT NULL
); 

CREATE TABLE  dashboards ( 
	database_id text REFERENCES databases(database_id), 
	dashboard_id text PRIMARY KEY, 
	name text NOT NULL,
	properties text NOT NULL
); 

CREATE TABLE table_views ( 
	database_id text REFERENCES databases(database_id), 
	table_id text PRIMARY KEY, 
	name text NOT NULL,
	properties text NOT NULL
); 

CREATE TABLE forms ( 
	database_id text REFERENCES databases(database_id), 
	form_id text PRIMARY KEY, 
	name text NOT NULL,
	properties text NOT NULL
); 

CREATE TABLE item_lists ( 
	list_id text PRIMARY KEY, 
	database_id text REFERENCES databases(database_id), 
	name text NOT NULL,
	properties text NOT NULL
); 

CREATE TABLE form_links (
	link_id text PRIMARY KEY,
	form_id text REFERENCES forms(form_id), 
	name text NOT NULL,
	include_in_sidebar boolean NOT NULL,
	shared_link_enabled boolean,
	shared_link_id text,
	properties text NOT NULL
); 

CREATE TABLE value_lists (
	value_list_id text PRIMARY KEY,
	database_id text REFERENCES databases(database_id), 
	name text NOT NULL,
	properties text NOT NULL
); 

CREATE TABLE dashboard_components (
	dashboard_id text REFERENCES dashboards(dashboard_id), 
	component_id text PRIMARY KEY,
	properties text NOT NULL,
	type text NOT NULL
);

CREATE TABLE form_components (
	form_id text REFERENCES forms(form_id), 
	component_id text PRIMARY KEY,
	properties text NOT NULL,
	type text NOT NULL
);

CREATE TABLE table_view_columns (
	table_id text REFERENCES table_views(table_id), 
	column_id text PRIMARY KEY,
	properties text NOT NULL,
	type text NOT NULL
);

CREATE TABLE field_comments (
	comment_id text NOT NULL,
	user_id text REFERENCES users(user_id),
	record_id text  REFERENCES records (record_id) ON DELETE CASCADE, 
	field_id text REFERENCES fields(field_id) ON DELETE CASCADE,
	create_timestamp_utc timestamp NOT NULL,
	update_timestamp_utc timestamp NOT NULL,
	comment text NOT NULL
);

CREATE TABLE  database_admins ( 
	database_id text REFERENCES databases(database_id), 
	user_id text REFERENCES users(user_id)
);

CREATE TABLE database_roles (
	database_id text REFERENCES databases(database_id), 
	role_id text PRIMARY KEY,
	name text NOT NULL
);

CREATE TABLE collaborators (
	collaborator_id text PRIMARY KEY,
	user_id text REFERENCES users(user_id), 
	database_id text REFERENCES databases(database_id),
	UNIQUE(user_id,database_id)
);

CREATE TABLE collaborator_roles (
	collaborator_id text REFERENCES collaborators(collaborator_id) ON DELETE CASCADE, 
	role_id text REFERENCES database_roles(role_id) ON DELETE CASCADE
);


CREATE TABLE list_role_privs (
	list_id text REFERENCES item_lists(list_id) ON DELETE CASCADE, 
	role_id text REFERENCES database_roles(role_id) ON DELETE CASCADE,
	privs text NOT NULL
);


CREATE TABLE dashboard_role_privs (
	dashboard_id text REFERENCES dashboards(dashboard_id) ON DELETE CASCADE, 
	role_id text REFERENCES database_roles(role_id) ON DELETE CASCADE,
	privs text NOT NULL
);

CREATE TABLE new_item_form_link_role_privs (
	link_id text REFERENCES form_links(link_id) ON DELETE CASCADE, 
	role_id text REFERENCES database_roles(role_id) ON DELETE CASCADE
);

CREATE TABLE alert_role_privs (
	alert_id text REFERENCES alerts(alert_id) ON DELETE CASCADE, 
	role_id text REFERENCES database_roles(role_id) ON DELETE CASCADE
);

