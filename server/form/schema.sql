CREATE TABLE forms ( 
	database_id text REFERENCES databases(database_id), 
	form_id text PRIMARY KEY, 
	name text NOT NULL,
	properties text NOT NULL
); 
