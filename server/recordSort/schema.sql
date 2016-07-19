CREATE TABLE form_sort_rules ( 
	form_id text REFERENCES forms(form_id), 
	sort_rules text NOT NULL
);