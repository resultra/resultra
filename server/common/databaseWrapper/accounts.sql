
-- Inside pgAdmin:
-- (1) Create the database resultra_accounts


-- (2) Open the query tool (right-click on database in TOC), then run the following SQL

CREATE TABLE account_info (
	account_id text PRIMARY KEY,
	owner_first text NOT NULL,
	owner_last text NOT NULL,
	owner_email text NOT NULL,
	db_host_name text NOT NULL
); 

CREATE TABLE host_mappings (
	host_name text PRIMARY KEY,
	account_id text REFERENCES account_info(account_id),
	UNIQUE(host_name)
); 

-- (3) In development mode use this to give privileges to the devuser
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO devuser;
ALTER USER devuser CREATEDB;

