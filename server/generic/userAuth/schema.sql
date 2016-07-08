CREATE TABLE users ( 
	user_id text PRIMARY KEY,
	user_name text NOT NULL,
	first_name text NOT NULL, -- TODO cannot be empty
	last_name text NOT NULL, -- TODO cannot be empty
	email_addr text NOT NULL, -- TODO must be non-empty, unique (case-insensitive)
	password_hash text NOT NULL
);
