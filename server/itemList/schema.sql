CREATE TABLE item_lists ( 
	list_id text PRIMARY KEY, 
	database_id text REFERENCES databases(database_id), 
	name text NOT NULL,
	properties text NOT NULL
); 
