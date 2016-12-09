CREATE TABLE record_val_results ( 
	database_id text REFERENCES databases (database_id), 
	record_id text  REFERENCES records (record_id), 
	field_vals text NOT NULL,
	update_timestamp_utc timestamp NOT NULL
);
