CREATE TABLE  dashboards ( 
	database_id text REFERENCES databases(database_id), 
	dashboard_id text PRIMARY KEY, 
	name text NOT NULL
); 
