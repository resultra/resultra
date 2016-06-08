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
	tableID uuid, 
	record_id uuid, 
	field_values text, 
	PRIMARY KEY ((tableID),record_id) 
); 