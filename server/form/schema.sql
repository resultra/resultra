CREATE TABLE forms ( 
	table_id text REFERENCES data_tables(table_id), 
	form_id text PRIMARY KEY, 
	name text NOT NULL,
	properties text NOT NULL
); 
