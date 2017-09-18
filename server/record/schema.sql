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