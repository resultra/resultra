CREATE TABLE account_info (
	account_id text PRIMARY KEY,
	owner_first text NOT NULL,
	owner_last text NOT NULL,
	owner_email text not NULL
); 

CREATE TABLE host_mappings (
	host_name text PRIMARY KEY,
	account_id text REFERENCES account_info(account_id),
	UNIQUE(host_name)
); 

-- In development mode use this to give privileges to the devuser
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO devuser;
ALTER USER devuser CREATEDB;

