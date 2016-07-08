/* The SQL syntax is deliberately "plain vanilla" SQL. By using plain SQL syntax (not PostgreSQL specific),
   there is a possibility to migrate to other databases as needed. For example, a lightweight test environment could be
   setup in SQLite. */


CREATE TABLE databases (
   database_id text PRIMARY_KEY,
   name text NOT NULL,
   PRIMARY KEY(database_id)
);

-- TODO - Support cascade deletes

CREATE TABLE  data_tables ( 
	table_id text PRIMARY KEY,
	database_id text REFERENCES databases (database_id), 
	name text NOT NULL
);


CREATE TABLE fields ( 
	field_id text PRIMARY KEY, 
	table_id text REFERENCES data_tables (table_id), 
	name text NOT NULL, 
	type text NOT NULL, 
	ref_name text NOT NULL, 
	calc_field_eqn text, 
	is_calc_field boolean NOT NULL, 
	preprocessed_formula_text text
); 

CREATE TABLE records ( 
	table_id text REFERENCES data_tables (table_id), 
	record_id text PRIMARY KEY
);

CREATE TABLE cell_updates (
	table_id text REFERENCES data_tables (table_id),
	record_id text REFERENCES records (record_id),
	field_id text REFERENCES fields (field_id),
	update_timestamp_utc timestamp NOT NULL,
	value text NOT NULL -- value encoded as JSON
);

CREATE TABLE record_val_results ( 
	table_id text REFERENCES data_tables (table_id), 
	record_id text  REFERENCES records (record_id), 
	field_vals text NOT NULL,
	filter_matches text NOT NULL,
	update_timestamp_utc timestamp NOT NULL
);

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

CREATE TABLE  dashboards ( 
	database_id text REFERENCES databases(database_id), 
	dashboard_id text PRIMARY KEY, 
	name text NOT NULL
); 


CREATE TABLE forms ( 
	table_id text REFERENCES data_tables(table_id), 
	form_id text PRIMARY KEY, 
	name text NOT NULL
); 


CREATE TABLE form_sort_rules ( 
	form_id text REFERENCES forms(form_id), 
	sort_rules text NOT NULL
);


CREATE TABLE bar_charts (
	dashboard_id text REFERENCES dashboards(dashboard_id), 
	barchart_id text PRIMARY KEY,
	properties text NOT NULL
); 

CREATE TABLE form_components (
	form_id text REFERENCES forms(form_id), 
	component_id text PRIMARY KEY,
	properties text NOT NULL,
	type text NOT NULL
);
