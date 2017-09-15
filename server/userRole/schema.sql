CREATE TABLE  database_admins ( 
	database_id text REFERENCES databases(database_id), 
	user_id text REFERENCES users(user_id)
);

CREATE TABLE database_roles (
	database_id text REFERENCES databases(database_id), 
	role_id text PRIMARY KEY,
	name text NOT NULL
);

CREATE TABLE collaborators (
	collaborator_id text PRIMARY KEY,
	user_id text REFERENCES users(user_id), 
	database_id text REFERENCES databases(database_id),
	UNIQUE(user_id,database_id)
);

/* Records in user_roles define which database roles a given users belongs to */
CREATE TABLE collaborator_roles (
	collaborator_id text REFERENCES collaborators(collaborator_id) ON DELETE CASCADE, 
	role_id text REFERENCES database_roles(role_id) ON DELETE CASCADE
);




CREATE TABLE list_role_privs (
	list_id text REFERENCES item_lists(list_id) ON DELETE CASCADE, 
	role_id text REFERENCES database_roles(role_id) ON DELETE CASCADE,
	privs text NOT NULL
);


CREATE TABLE dashboard_role_privs (
	dashboard_id text REFERENCES dashboards(dashboard_id) ON DELETE CASCADE, 
	role_id text REFERENCES database_roles(role_id) ON DELETE CASCADE,
	privs text NOT NULL
);

CREATE TABLE new_item_form_link_role_privs (
	link_id text REFERENCES form_links(link_id) ON DELETE CASCADE, 
	role_id text REFERENCES database_roles(role_id) ON DELETE CASCADE
);

CREATE TABLE alert_role_privs (
	alert_id text REFERENCES alerts(alert_id) ON DELETE CASCADE, 
	role_id text REFERENCES database_roles(role_id) ON DELETE CASCADE
);

