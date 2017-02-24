CREATE TABLE attachments (
	attachment_id text PRIMARY KEY,
	database_id text REFERENCES databases (database_id),
	user_id text NOT NULL,
	create_timestamp_utc timestamp NOT NULL,
	orig_file_name text NOT NULL,
	cloud_file_name text NOT NULL,
	title text,
	caption text
);