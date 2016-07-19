CREATE TABLE  data_tables ( 
	table_id text PRIMARY KEY,
	database_id text REFERENCES databases (database_id), 
	name text NOT NULL
);
