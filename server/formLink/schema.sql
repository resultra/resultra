CREATE TABLE form_links (
	link_id text PRIMARY KEY,
	form_id text REFERENCES forms(form_id), 
	name text NOT NULL,
	include_in_sidebar boolean NOT NULL,
	properties text NOT NULL
); 
