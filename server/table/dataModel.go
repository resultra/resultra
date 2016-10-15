package table

import (
	"fmt"
	"resultra/datasheet/server/generic/databaseWrapper"
	"resultra/datasheet/server/generic/stringValidation"
	"resultra/datasheet/server/generic/uniqueID"
)

const TableEntityKind string = "Table"

type Table struct {
	ParentDatabaseID string `json:"parentDatabaseID"`
	TableID          string `json:"tableID"`
	Name             string `json:"name"`
}

const tableParentIDFieldName string = "ParentDatabaseID"

type NewTableParams struct {
	DatabaseID string `json:'databaseID'`
	Name       string `json:"name"`
}

func saveNewTable(params NewTableParams) (*Table, error) {

	sanitizedTableName, sanitizeErr := stringValidation.SanitizeName(params.Name)
	if sanitizeErr != nil {
		return nil, sanitizeErr
	}

	tableID := uniqueID.GenerateSnowflakeID()

	if _, insertErr := databaseWrapper.DBHandle().Exec(
		`INSERT INTO data_tables (table_id, database_id, name) VALUES ($1, $2, $3)`,
		tableID, params.DatabaseID, sanitizedTableName); insertErr != nil {
		return nil, fmt.Errorf("saveNewTable: insert failed: error = %v", insertErr)
	}

	newTable := Table{ParentDatabaseID: params.DatabaseID, TableID: tableID, Name: sanitizedTableName}

	return &newTable, nil
}

func GetTableDatabaseID(tableID string) (string, error) {

	databaseID := ""
	getErr := databaseWrapper.DBHandle().QueryRow(
		`SELECT database_id 
			FROM data_tables 
			WHERE table_id=$1 LIMIT 1`,
		tableID).Scan(&databaseID)
	if getErr != nil {
		return "", fmt.Errorf(
			"getTableDatabaseID: can't get database for table = %v: err=%v",
			tableID, getErr)
	}

	return databaseID, nil

}

type GetTableListParams struct {
	DatabaseID string `json:"databaseID"` // parent database
}

func getTableList(params GetTableListParams) ([]Table, error) {

	rows, queryErr := databaseWrapper.DBHandle().Query(
		`SELECT database_id,table_id,name FROM data_tables WHERE database_id = $1`, params.DatabaseID)
	if queryErr != nil {
		return nil, fmt.Errorf("getTableList: Failure querying database: %v", queryErr)
	}
	tables := []Table{}
	for rows.Next() {
		var currTable Table
		if scanErr := rows.Scan(&currTable.ParentDatabaseID, &currTable.TableID, &currTable.Name); scanErr != nil {
			return nil, fmt.Errorf("getTableList: Failure querying database: %v", scanErr)
		}
		tables = append(tables, currTable)
	}

	return tables, nil

}
