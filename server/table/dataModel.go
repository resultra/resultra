package table

import (
	"appengine"
	"fmt"
	"github.com/gocql/gocql"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/cassandraWrapper"
)

const TableEntityKind string = "Table"

type Table struct {
	ParentDatabaseID gocql.UUID `json:"parentDatabaseID"`
	TableID          gocql.UUID `json:"tableID"`
	Name             string     `json:"name"`
}

const tableParentIDFieldName string = "ParentDatabaseID"

type NewTableParams struct {
	DatabaseID gocql.UUID `json:'databaseID'`
	Name       string     `json:"name"`
}

func saveNewTable(appEngContext appengine.Context, params NewTableParams) (*Table, error) {

	sanitizedTableName, sanitizeErr := generic.SanitizeName(params.Name)
	if sanitizeErr != nil {
		return nil, sanitizeErr
	}

	// TODO - Validate name is unique

	dbSession, sessionErr := cassandraWrapper.CreateSession()
	if sessionErr != nil {
		return nil, fmt.Errorf("saveNewTable: Can't create database: unable to create database session: error = %v", sessionErr)
	}
	defer dbSession.Close()

	tableID := gocql.TimeUUID()
	if insertErr := dbSession.Query(`INSERT INTO dataTable (databaseID, tableID, name) VALUES (?, ?, ?)`,
		params.DatabaseID, tableID, sanitizedTableName).Exec(); insertErr != nil {
		fmt.Errorf("saveNewTable: Can't create table: unable to create database: error = %v", insertErr)
	}

	newTable := Table{ParentDatabaseID: params.DatabaseID, TableID: tableID, Name: sanitizedTableName}

	return &newTable, nil
}

type GetTableListParams struct {
	DatabaseID gocql.UUID `json:"databaseID"` // parent database
}

func getTableList(appEngContext appengine.Context, params GetTableListParams) ([]Table, error) {

	dbSession, sessionErr := cassandraWrapper.CreateSession()
	if sessionErr != nil {
		return nil, fmt.Errorf("getTableList: Unable to create database session: error = %v", sessionErr)
	}
	defer dbSession.Close()

	tableIter := dbSession.Query(`SELECT databaseID,tableID,name FROM dataTable WHERE databaseID = ?`,
		params.DatabaseID).Iter()

	var currTable Table
	tables := []Table{}
	for tableIter.Scan(&currTable.ParentDatabaseID, &currTable.TableID, &currTable.Name) {
		tables = append(tables, currTable)
	}
	if closeErr := tableIter.Close(); closeErr != nil {
		fmt.Errorf("getTableList: Failure querying database: %v", closeErr)
	}

	return tables, nil

}
