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

func saveNewTable(newTable Table) error {

	if _, insertErr := databaseWrapper.DBHandle().Exec(
		`INSERT INTO data_tables (table_id, database_id, name) VALUES ($1, $2, $3)`,
		newTable.TableID, newTable.ParentDatabaseID, newTable.Name); insertErr != nil {
		return fmt.Errorf("saveNewTable: insert failed: error = %v", insertErr)
	}

	return nil
}

func saveNewEmptyTable(params NewTableParams) (*Table, error) {

	sanitizedTableName, sanitizeErr := stringValidation.SanitizeName(params.Name)
	if sanitizeErr != nil {
		return nil, sanitizeErr
	}

	tableID := uniqueID.GenerateSnowflakeID()

	newTable := Table{ParentDatabaseID: params.DatabaseID, TableID: tableID, Name: sanitizedTableName}

	if err := saveNewTable(newTable); err != nil {
		return nil, err
	}

	return &newTable, nil
}

func CloneTables(remappedIDs uniqueID.UniqueIDRemapper, srcDatabaseID string) error {

	destDatabaseID, err := remappedIDs.GetExistingRemappedID(srcDatabaseID)
	if err != nil {
		return fmt.Errorf("CloneTables: Destination database ID not found for src database = %v", srcDatabaseID)
	}

	tableParams := GetTableListParams{DatabaseID: srcDatabaseID}
	tables, err := GetTableList(tableParams)
	if err != nil {
		return fmt.Errorf("CloneTables: %v", err)
	}
	for _, srcTable := range tables {

		destTableID, err := remappedIDs.AllocNewRemappedID(srcTable.TableID)
		if err != nil {
			return fmt.Errorf("CloneTables: %v", err)
		}

		destTable := srcTable
		destTable.ParentDatabaseID = destDatabaseID
		destTable.TableID = destTableID

		if err := saveNewTable(destTable); err != nil {
			return fmt.Errorf("CloneTables: failure saving cloned table: %v", err)
		}
	}
	return nil
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

func GetTableList(params GetTableListParams) ([]Table, error) {

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
