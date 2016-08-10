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

CREATE TABLE form_role_privs (
	form_id text REFERENCES forms(form_id), 
	role_id text REFERENCES database_roles(role_id),
	privs text NOT NULL
);

CREATE TABLE dashboard_role_privs (
	dashboard_id text REFERENCES dashboards(dashboard_id), 
	role_id text REFERENCES database_roles(role_id),
	privs text NOT NULL
);