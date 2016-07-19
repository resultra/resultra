CREATE TABLE records ( 
	table_id text REFERENCES data_tables (table_id), 
	record_id text PRIMARY KEY
);

CREATE TABLE cell_updates (
	table_id text REFERENCES data_tables (table_id),
	record_id text REFERENCES records (record_id),
	field_id text REFERENCES fields (field_id),
	update_timestamp_utc timestamp NOT NULL,
	value text NOT NULL -- value encoded as JSON
);