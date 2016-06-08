create keyspace datasheet with replication = { 'class':'SimpleStrategy','replication_factor':1};

create table datasheet.database(
	databaseID uuid,
	name text,
	PRIMARY KEY (databaseID)
);

create table datasheet.dataTable ( 
	databaseID uuid, 
	tableID uuid,
	name text,
	PRIMARY KEY ((databaseID),tableID) 
);

create table datasheet.field ( 
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
