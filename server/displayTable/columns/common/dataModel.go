package common

import (
	"database/sql"
	"fmt"
	"resultra/datasheet/server/generic"
)

func SaveNewTableColumn(destDBHandle *sql.DB, columnType string, parentTableID string, columnID string, properties interface{}) error {

	encodedProps, encodeErr := generic.EncodeJSONString(properties)
	if encodeErr != nil {
		return fmt.Errorf("SaveNewTableViewColumn: Unable to save %v: error = %v", columnType, encodeErr)
	}

	if _, insertErr := destDBHandle.Exec(
		`INSERT INTO table_view_columns (table_id,column_id,type,properties) VALUES ($1,$2,$3,$4)`,
		parentTableID, columnID, columnType, encodedProps); insertErr != nil {
		return fmt.Errorf("SaveNewTableViewColumn: Can't save %v: error = %v", columnType, insertErr)
	}

	return nil
}

func GetTableColumn(trackerDBHandle *sql.DB, columnType string, parentTableID string, columnID string, properties interface{}) error {

	encodedProps := ""
	getErr := trackerDBHandle.QueryRow(`SELECT properties FROM table_view_columns
		 WHERE table_id=$1 AND column_id=$2 AND type=$3 LIMIT 1`,
		parentTableID, columnID, columnType).Scan(&encodedProps)
	if getErr != nil {
		return fmt.Errorf("GetTableColumn: Unabled to get table column %v: id = %v: datastore err=%v",
			columnType, columnID, getErr)
	}

	if decodeErr := generic.DecodeJSONString(encodedProps, properties); decodeErr != nil {
		return fmt.Errorf("GetTableColumn: Unabled to decode properties: encoded properties = %v: datastore err=%v",
			encodedProps, decodeErr)
	}

	return nil
}

func GetTableColumnTableID(trackerDBHandle *sql.DB, columnID string) (string, error) {

	tableID := ""
	getErr := trackerDBHandle.QueryRow(`SELECT table_id FROM table_view_columns
		 WHERE column_id=$1 LIMIT 1`,
		columnID).Scan(&tableID)
	if getErr != nil {
		return "", fmt.Errorf("GetTableColumnTable: Unabled to get table id for column: id = %v: datastore err=%v",
			columnID, getErr)
	}
	return tableID, nil

}

func GetTableColumnAndTable(trackerDBHandle *sql.DB, columnType string, parentTableID string, columnID string, properties interface{}) error {

	encodedProps := ""
	tableID := ""
	getErr := trackerDBHandle.QueryRow(`SELECT table_id,properties FROM table_view_columns
		 WHERE column_id=$1 AND type=$2 LIMIT 1`,
		columnID, columnType).Scan(&tableID, &encodedProps)
	if getErr != nil {
		return fmt.Errorf("GetTableViewColumn: Unabled to get table column %v: id = %v: datastore err=%v",
			columnType, columnID, getErr)
	}

	if decodeErr := generic.DecodeJSONString(encodedProps, properties); decodeErr != nil {
		return fmt.Errorf("GetTableViewColumn: Unabled to decode properties: encoded properties = %v: datastore err=%v",
			encodedProps, decodeErr)
	}

	return nil
}

type addColumnCallbackFunc func(string, string) error

func GetTableColumns(srcDBHandle *sql.DB, columnType string, parentTableID string, addColumnFunc addColumnCallbackFunc) error {

	rows, queryErr := srcDBHandle.Query(`SELECT column_id,properties
			FROM table_view_columns 
			WHERE table_id=$1 AND type=$2`,
		parentTableID, columnType)
	if queryErr != nil {
		return fmt.Errorf("GetTableViewColumns: Failure querying database: %v", queryErr)
	}
	defer rows.Close()

	for rows.Next() {
		currColumnID := ""
		encodedProps := ""
		if scanErr := rows.Scan(&currColumnID, &encodedProps); scanErr != nil {
			return fmt.Errorf("GetTableViewColumns: Failure querying database: %v", scanErr)
		}
		if err := addColumnFunc(currColumnID, encodedProps); err != nil {
			return err
		}
	}

	return nil
}

func UpdateTableColumn(trackerDBHandle *sql.DB, columnType string, parentTableID string, columnID string, properties interface{}) error {

	encodedProps, encodeErr := generic.EncodeJSONString(properties)
	if encodeErr != nil {
		return fmt.Errorf("UpdateTableViewColumn: failure encoding properties: error = %v", encodeErr)
	}

	if _, updateErr := trackerDBHandle.Exec(`UPDATE table_view_columns 
				SET properties=$1
				WHERE table_id=$2 AND column_id=$3`,
		encodedProps, parentTableID, columnID); updateErr != nil {
		return fmt.Errorf("UpdateTableViewColumn: Can't update form component %v: error = %v",
			columnType, updateErr)
	}

	return nil

}

func DeleteTableColumn(trackerDBHandle *sql.DB, parentTableID string, columnID string) error {
	if _, deleteErr := trackerDBHandle.Exec(`DELETE FROM table_view_columns 
				WHERE table_id=$1 AND column_id=$2`, parentTableID, columnID); deleteErr != nil {
		return fmt.Errorf("DeleteTableViewColumn: Can't delete table view column %v: error = %v",
			columnID, deleteErr)
	}
	return nil
}

type ColumnInfo struct {
	ColumnID string
	TableID  string
	ColType  string
}

func GetTableColumnInfo(trackerDBHandle *sql.DB, columnID string) (*ColumnInfo, error) {
	colType := ""
	tableID := ""
	getErr := trackerDBHandle.QueryRow(`SELECT table_view_columns.type,table_views.table_id FROM table_view_columns,table_views
		 WHERE table_view_columns.column_id=$1 AND table_view_columns.table_id=table_views.table_id LIMIT 1`,
		columnID).Scan(&colType, &tableID)
	if getErr != nil {
		return nil, fmt.Errorf("GetTableColumnInfo: Unabled to get table column info: id = %v: datastore err=%v", columnID, getErr)
	}

	colInfo := ColumnInfo{
		ColumnID: columnID,
		ColType:  colType,
		TableID:  tableID}

	return &colInfo, nil

}
