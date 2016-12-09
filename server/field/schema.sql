CREATE TABLE fields ( 
	field_id text PRIMARY KEY, 
	database_id text REFERENCES databases (database_id), 
	name text NOT NULL, 
	type text NOT NULL, 
	ref_name text NOT NULL, 
	calc_field_eqn text, 
	is_calc_field boolean NOT NULL, 
	preprocessed_formula_text text
); 