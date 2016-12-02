CREATE TABLE item_lists ( 
	list_id text PRIMARY KEY, 
	table_id text REFERENCES data_tables(table_id), 
	name text NOT NULL,
	properties text NOT NULL
); 
