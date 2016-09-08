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