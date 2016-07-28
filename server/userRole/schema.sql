CREATE TABLE  database_admins ( 
	database_id text REFERENCES databases(database_id), 
	user_id text REFERENCES users(user_id)
);

CREATE TABLE database_roles (
	database_id text REFERENCES databases(database_id), 
	role_id text PRIMARY KEY,
	name text NOT NULL
);

/* Records in user_roles define which database roles a given users belongs to */
CREATE TABLE user_roles (
	user_id text REFERENCES users(user_id), 
	role_id text REFERENCES database_roles(role_id)
);