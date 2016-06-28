create table datasheet.users ( 
	user_id uuid,
	user_name text,
	first_name text,
	last_name text, 
	email_addr text,
	password_hash text,
	PRIMARY KEY (user_id,email_addr) 
);
CREATE INDEX user_email on datasheet.users(email_addr);
