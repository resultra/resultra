CREATE TABLE table_views ( 
	database_id text REFERENCES databases(database_id), 
	table_id text PRIMARY KEY, 
	name text NOT NULL,
	properties text NOT NULL
); 