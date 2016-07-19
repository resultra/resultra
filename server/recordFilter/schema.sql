CREATE TABLE  filters ( 
	table_id text REFERENCES data_tables(table_id), 
	filter_id text PRIMARY KEY, 
	name text NOT NULL
); 

CREATE TABLE  filter_rules ( 
	filter_id text REFERENCES filters(filter_id), 
	rule_id text PRIMARY KEY,
	field_id text REFERENCES fields(field_id),
	rule_def_id text NOT NULL,
	text_param text,
	number_param double precision
);