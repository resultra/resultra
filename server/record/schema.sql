CREATE TABLE records ( 
	database_id text REFERENCES databases (database_id), 
	record_id text PRIMARY KEY,
	is_draft_record boolean NOT NULL,
	create_timestamp_utc timestamp NOT NULL
);

CREATE TABLE cell_updates (
	update_id text PRIMARY KEY,
	database_id text REFERENCES databases (database_id),
	record_id text REFERENCES records (record_id),
	field_id text REFERENCES fields (field_id),
	change_set_id text -- used to segregate uncommitted changes made in modal dialogs, set to NULL for baseline changes.
	update_timestamp_utc timestamp NOT NULL,
	user_id text REFERENCES users(user_id),
	value text NOT NULL, -- value encoded as JSON
	properties text NOT NULL -- properties encoded as JSON
);