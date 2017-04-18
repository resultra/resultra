CREATE TABLE value_lists (
	value_list_id text PRIMARY KEY,
	database_id text REFERENCES databases(database_id), 
	name text NOT NULL,
	properties text NOT NULL
); 
