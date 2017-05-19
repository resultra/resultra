CREATE TABLE table_view_columns (
	table_id text REFERENCES table_views(table_id), 
	column_id text PRIMARY KEY,
	properties text NOT NULL,
	type text NOT NULL
);
