CREATE TABLE  alerts ( 
	database_id text REFERENCES databases(database_id), 
	alert_id text PRIMARY KEY, 
	name text NOT NULL,
	properties text NOT NULL
); 

CREATE TABLE alert_notifications {
	alert_id REFERENCES alerts(alert_id) ON DELETE CASCADE,
	update_id REFERENCES cell_updates(update_id) ON DELETE CASCADE,
	UNIQUE(alert_id,update_id)
}