CREATE TABLE form_components (
	form_id text REFERENCES forms(form_id), 
	component_id text PRIMARY KEY,
	properties text NOT NULL,
	type text NOT NULL
);
