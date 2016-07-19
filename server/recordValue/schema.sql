CREATE TABLE record_val_results ( 
	table_id text REFERENCES data_tables (table_id), 
	record_id text  REFERENCES records (record_id), 
	field_vals text NOT NULL,
	filter_matches text NOT NULL,
	update_timestamp_utc timestamp NOT NULL
);
