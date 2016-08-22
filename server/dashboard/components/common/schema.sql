CREATE TABLE dashboard_components (
	dashboard_id text REFERENCES dashboards(dashboard_id), 
	component_id text PRIMARY KEY,
	properties text NOT NULL,
	type text NOT NULL
);