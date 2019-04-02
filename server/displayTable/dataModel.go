// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package displayTable

import (
	"database/sql"
	"fmt"
	"github.com/resultra/resultra/server/generic"
	"github.com/resultra/resultra/server/generic/stringValidation"
	"github.com/resultra/resultra/server/generic/uniqueID"
	"github.com/resultra/resultra/server/trackerDatabase"
)

type NewTableParams struct {
	ParentDatabaseID string `json:"parentDatabaseID"`
	Name             string `json:"name"`
}

type DisplayTable struct {
	ParentDatabaseID string                 `json:"parentDatabaseID"`
	Name             string                 `json:"name"`
	TableID          string                 `json:"tableID"`
	Properties       DisplayTableProperties `json:"properties"`
}

func saveTable(destDBHandle *sql.DB, newTable DisplayTable) error {

	encodedTableProps, encodeErr := generic.EncodeJSONString(newTable.Properties)
	if encodeErr != nil {
		return fmt.Errorf("saveTable: failure encoding properties: error = %v", encodeErr)
	}

	if _, insertErr := destDBHandle.Exec(`INSERT INTO table_views
			 	(database_id,table_id,name,properties) VALUES ($1,$2,$3,$4)`,
		newTable.ParentDatabaseID, newTable.TableID, newTable.Name, encodedTableProps); insertErr != nil {
		return fmt.Errorf("saveTable: Can't create display table: error = %v", insertErr)
	}
	return nil

}

func newTable(trackerDBHandle *sql.DB, params NewTableParams) (*DisplayTable, error) {

	sanitizedName, sanitizeErr := stringValidation.SanitizeName(params.Name)
	if sanitizeErr != nil {
		return nil, sanitizeErr
	}

	newTable := DisplayTable{
		ParentDatabaseID: params.ParentDatabaseID,
		TableID:          uniqueID.GenerateUniqueID(),
		Name:             sanitizedName,
		Properties:       newDefaultDisplayTableProperties()}

	if err := saveTable(trackerDBHandle, newTable); err != nil {
		return nil, fmt.Errorf("newTable: error saving table: %v", err)
	}

	return &newTable, nil
}

func GetTable(trackerDBHandle *sql.DB, tableID string) (*DisplayTable, error) {

	tableName := ""
	encodedProps := ""
	databaseID := ""
	getErr := trackerDBHandle.QueryRow(`SELECT database_id,name,properties FROM table_views
		 WHERE table_id=$1 LIMIT 1`, tableID).Scan(&databaseID, &tableName, &encodedProps)
	if getErr != nil {
		return nil, fmt.Errorf("GetTable: Unabled to get table: table ID = %v: datastore err=%v",
			tableID, getErr)
	}

	tableProps := newDefaultDisplayTableProperties()
	if decodeErr := generic.DecodeJSONString(encodedProps, &tableProps); decodeErr != nil {
		return nil, fmt.Errorf("GetTable: can't decode properties: %v", encodedProps)
	}

	getTable := DisplayTable{
		ParentDatabaseID: databaseID,
		TableID:          tableID,
		Name:             tableName,
		Properties:       tableProps}

	return &getTable, nil
}

type GetTableListParams struct {
	ParentDatabaseID string `json:"parentDatabaseID"`
}

func getAllTablesFromSrc(srcDBHandle *sql.DB, parentDatabaseID string) ([]DisplayTable, error) {

	rows, queryErr := srcDBHandle.Query(
		`SELECT database_id,table_id,name,properties FROM table_views WHERE database_id = $1`,
		parentDatabaseID)
	if queryErr != nil {
		return nil, fmt.Errorf("GetAllTables: Failure querying database: %v", queryErr)
	}
	defer rows.Close()

	tables := []DisplayTable{}
	for rows.Next() {
		var currTable DisplayTable
		encodedProps := ""

		if scanErr := rows.Scan(&currTable.ParentDatabaseID, &currTable.TableID, &currTable.Name, &encodedProps); scanErr != nil {
			return nil, fmt.Errorf("GetAllTables: Failure querying database: %v", scanErr)
		}

		tableProps := newDefaultDisplayTableProperties()
		if decodeErr := generic.DecodeJSONString(encodedProps, &tableProps); decodeErr != nil {
			return nil, fmt.Errorf("GetAllTables: can't decode properties: %v", encodedProps)
		}
		currTable.Properties = tableProps

		tables = append(tables, currTable)
	}

	return tables, nil

}

func getAllTables(trackerDBHandle *sql.DB, parentDatabaseID string) ([]DisplayTable, error) {
	return getAllTablesFromSrc(trackerDBHandle, parentDatabaseID)
}

func CloneTables(cloneParams *trackerDatabase.CloneDatabaseParams) error {

	remappedDatabaseID, err := cloneParams.IDRemapper.GetExistingRemappedID(cloneParams.SourceDatabaseID)
	if err != nil {
		return fmt.Errorf("CloneTableTables: Error getting remapped table ID: %v", err)
	}

	tables, err := getAllTablesFromSrc(cloneParams.SrcDBHandle, cloneParams.SourceDatabaseID)
	if err != nil {
		return fmt.Errorf("CloneTables: Error getting tables for parent database ID = %v: %v",
			cloneParams.SourceDatabaseID, err)
	}

	for _, currTable := range tables {

		destTable := currTable
		destTable.ParentDatabaseID = remappedDatabaseID

		destTableID, err := cloneParams.IDRemapper.AllocNewRemappedID(currTable.TableID)
		if err != nil {
			return fmt.Errorf("CloneTableTables: %v", err)
		}
		destTable.TableID = destTableID

		destProps, err := currTable.Properties.Clone(cloneParams.IDRemapper)
		if err != nil {
			return fmt.Errorf("CloneTableTables: %v", err)
		}
		destTable.Properties = *destProps

		if err := saveTable(cloneParams.DestDBHandle, destTable); err != nil {
			return fmt.Errorf("CloneTableTables: %v", err)
		}

		if err := cloneTableCols(cloneParams, currTable.TableID); err != nil {
			return fmt.Errorf("Clone tables: %v", err)
		}

	}

	return nil

}

func updateExistingTable(trackerDBHandle *sql.DB, tableID string, updatedTable *DisplayTable) (*DisplayTable, error) {

	encodedProps, encodeErr := generic.EncodeJSONString(updatedTable.Properties)
	if encodeErr != nil {
		return nil, fmt.Errorf("updateExistingTable: failure encoding properties: error = %v", encodeErr)
	}

	if _, updateErr := trackerDBHandle.Exec(`UPDATE table_views 
				SET properties=$1, name=$2
				WHERE table_id=$3`,
		encodedProps, updatedTable.Name, tableID); updateErr != nil {
		return nil, fmt.Errorf("updateExistingTable: Can't update table properties %v: error = %v",
			tableID, updateErr)
	}

	return updatedTable, nil

}

func getTableDatabaseID(trackerDBHandle *sql.DB, tableID string) (string, error) {

	theTable, err := GetTable(trackerDBHandle, tableID)
	if err != nil {
		return "", nil
	}
	return theTable.ParentDatabaseID, nil
}

type TableNameValidationInfo struct {
	Name string
	ID   string
}

func validateUniqueTableName(trackerDBHandle *sql.DB, databaseID string, tableID string, tableName string) error {
	// Query to validate the name is unique:
	// 1. Select all the tables in the same database
	// 2. Include tables with the same name.
	// 3. Exclude tables with the same table ID. In other words
	//    the name is considered valid if it is the same as its
	//    existing name.
	rows, queryErr := trackerDBHandle.Query(
		`SELECT table_views.table_id,table_views.name 
			FROM table_views,databases
			WHERE databases.database_id=$1 AND
			table_views.database_id=databases.database_id AND
				table_views.name=$2 AND table_views.table_id<>$3`,
		databaseID, tableName, tableID)
	if queryErr != nil {
		return fmt.Errorf("System error validating table name (%v)", queryErr)
	}
	defer rows.Close()

	existingTableNameUsedByAnotherTable := rows.Next()
	if existingTableNameUsedByAnotherTable {
		return fmt.Errorf("Invalid table name - names must be unique")
	}

	return nil

}

func validateTableName(trackerDBHandle *sql.DB, tableID string, tableName string) error {

	if !stringValidation.WellFormedItemName(tableName) {
		return fmt.Errorf("Invalid table name")
	}

	databaseID, err := getTableDatabaseID(trackerDBHandle, tableID)
	if err != nil {
		return fmt.Errorf("System error validating table name (%v)", err)
	}

	if uniqueErr := validateUniqueTableName(trackerDBHandle, databaseID, tableID, tableName); uniqueErr != nil {
		return uniqueErr
	}

	return nil
}

func validateNewTableName(trackerDBHandle *sql.DB, databaseID string, tableName string) error {

	if !stringValidation.WellFormedItemName(tableName) {
		return fmt.Errorf("Invalid table name")
	}

	// No table will have an empty tableID, so this will cause test for unique
	// table names to return true if any table already has the given tableName.
	tableID := ""
	if uniqueErr := validateUniqueTableName(trackerDBHandle, databaseID, tableID, tableName); uniqueErr != nil {
		return uniqueErr
	}

	return nil
}
