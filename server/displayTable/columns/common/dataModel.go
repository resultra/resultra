package common

import (
	"fmt"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/databaseWrapper"
)

func SaveNewTableColumn(columnType string, parentTableID string, columnID string, properties interface{}) error {

	encodedProps, encodeErr := generic.EncodeJSONString(properties)
	if encodeErr != nil {
		return fmt.Errorf("SaveNewTableViewColumn: Unable to save %v: error = %v", columnType, encodeErr)
	}

	if _, insertErr := databaseWrapper.DBHandle().Exec(
		`INSERT INTO table_view_columns (table_id,column_id,type,properties) VALUES ($1,$2,$3,$4)`,
		parentTableID, columnID, columnType, encodedProps); insertErr != nil {
		return fmt.Errorf("SaveNewTableViewColumn: Can't save %v: error = %v", columnType, insertErr)
	}

	return nil
}

func GetTableColumn(columnType string, parentTableID string, columnID string, properties interface{}) error {

	encodedProps := ""
	getErr := databaseWrapper.DBHandle().QueryRow(`SELECT properties FROM table_view_columns
		 WHERE table_id=$1 AND column_id=$2 AND type=$3 LIMIT 1`,
		parentTableID, columnID, columnType).Scan(&encodedProps)
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

func GetTableColumns(columnType string, parentTableID string, addColumnFunc addColumnCallbackFunc) error {

	rows, queryErr := databaseWrapper.DBHandle().Query(`SELECT column_id,properties
			FROM table_view_columns 
			WHERE table_id=$1 AND type=$2`,
		parentTableID, columnType)
	if queryErr != nil {
		return fmt.Errorf("GetTableViewColumns: Failure querying database: %v", queryErr)
	}

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

func UpdateTableColumn(columnType string, parentTableID string, columnID string, properties interface{}) error {

	encodedProps, encodeErr := generic.EncodeJSONString(properties)
	if encodeErr != nil {
		return fmt.Errorf("UpdateTableViewColumn: failure encoding properties: error = %v", encodeErr)
	}

	if _, updateErr := databaseWrapper.DBHandle().Exec(`UPDATE table_view_columns 
				SET properties=$1
				WHERE table_id=$2 AND column_id=$3`,
		encodedProps, parentTableID, columnID); updateErr != nil {
		return fmt.Errorf("UpdateTableViewColumn: Can't update form component %v: error = %v",
			columnType, updateErr)
	}

	return nil

}

func DeleteTableColumn(parentTableID string, columnID string) error {
	if _, deleteErr := databaseWrapper.DBHandle().Exec(`DELETE FROM table_view_columns 
				WHERE table_id=$1 AND column_id=$2`, parentTableID, columnID); deleteErr != nil {
		return fmt.Errorf("DeleteTableViewColumn: Can't delete form component %v: error = %v",
			columnID, deleteErr)
	}
	return nil
}
