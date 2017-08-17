CREATE TABLE  alerts ( 
	database_id text REFERENCES databases(database_id), 
	alert_id text PRIMARY KEY, 
	name text NOT NULL,
	properties text NOT NULL
); 