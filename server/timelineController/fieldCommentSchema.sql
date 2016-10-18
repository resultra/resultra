CREATE TABLE field_comments (
	comment_id text NOT NULL,
	user_id text REFERENCES users(user_id),
	record_id text  REFERENCES records (record_id), 
	field_id text REFERENCES fields(field_id),
	create_timestamp_utc timestamp NOT NULL,
	update_timestamp_utc timestamp NOT NULL,
	comment text NOT NULL
);