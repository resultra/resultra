package table

import (
	"fmt"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/databaseWrapper"
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

	sanitizedTableName, sanitizeErr := generic.SanitizeName(params.Name)
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
