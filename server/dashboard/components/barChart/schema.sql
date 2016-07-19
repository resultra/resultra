CREATE TABLE bar_charts (
	dashboard_id text REFERENCES dashboards(dashboard_id), 
	barchart_id text PRIMARY KEY,
	properties text NOT NULL
); 
