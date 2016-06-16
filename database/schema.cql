create keyspace datasheet with replication = { 'class':'SimpleStrategy','replication_factor':1};

// TODO - use timeuuid type instead of uuid

create table datasheet.databases(
	databaseID uuid,
	name text,
	PRIMARY KEY (databaseID)
);

create table datasheet.dataTables ( 
	databaseID uuid, 
	tableID uuid,
	name text,
	PRIMARY KEY ((databaseID),tableID) 
);

create table datasheet.fields ( 
	tableID uuid, 
	fieldID uuid, 
	name text, 
	type text, 
	refName text, 
	calcFieldEqn text, 
	isCalcField boolean, 
	preprocessedFormulaText text, 
	PRIMARY KEY ((tableID),fieldID) 
); 


create table datasheet.forms ( 
	tableID uuid, 
	formID uuid, 
	name text, 
	PRIMARY KEY ((tableID),formID) 
); 


create table datasheet.dashboards ( 
	databaseID uuid, 
	dashboardID uuid, 
	name text, 
	PRIMARY KEY ((databaseID),dashboardID) 
); 


create table datasheet.records ( 
	table_id uuid, 
	record_id uuid, 
	PRIMARY KEY ((table_id),record_id) 
); 

create table datasheet.record_val_results ( 
	table_id uuid, 
	record_id uuid, 
	field_vals text,
	update_timestamp_utc timestamp,
	PRIMARY KEY ((table_id),record_id) 
); 

// In the cell_update table it's necessary to specify
// the clustering keys for all the fields but the value field
// (i.e., record_id,field_id, and update_time)
// This causes Cassandra to any value which happens at a different time
// as unique.
//
// TBD: Should field_id be before record_id in the clustering key order.
//
// TBD (Important) - There is a 2 billion record limit per partition. In this case,
// The limit would be 2 billion cell updates per table. It is unlikely a database
// would exceed this limit, but this table might need to be structured with multiple
// partitions to get around this. 
//
// update_timestamp_utc is ordered in descending order, so the first value will be the
// most recent. This is specified as such to facilitate minimizing repeated updates; i.e.
// if an update by the same user occurs within a small timeframe of another, the value
// can be replaced (TODO - This needs to be implemented in the actual code).
create table datasheet.cell_updates (
	table_id uuid,
	record_id uuid,
	field_id uuid,
	// TODO - Include user ID: user_id uuid,
	update_timestamp_utc timestamp,
	value text, // value encoded as JSON
	PRIMARY KEY ((table_id),record_id,field_id,update_timestamp_utc) 
) WITH CLUSTERING ORDER BY (record_id ASC, field_id ASC, update_timestamp_utc DESC);

create table datasheet.filters ( 
	tableID uuid, 
	filter_id uuid, 
	name text, 
	PRIMARY KEY ((tableID),filter_id) 
); 

// TODO - Should filter_id be the partitioning key for filter_rules?

create table datasheet.filter_rules ( 
	filter_id uuid, 
	rule_id uuid,
	field_id uuid,
	rule_def_id text,
	text_param text,
	number_param double,
	PRIMARY KEY ((filter_id),rule_id) 
);

create table datasheet.bar_charts (
	dashboard_id uuid, 
	barchart_id uuid,
	properties text,
	PRIMARY KEY ((dashboard_id),barchart_id) 
); 

create table datasheet.form_components (
	form_id uuid, 
	component_id uuid,
	properties text,
	type text,
	PRIMARY KEY ((form_id),component_id) 
);
CREATE INDEX component_type on datasheet.form_components(type);
