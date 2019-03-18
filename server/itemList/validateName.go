package itemList

import (
	"database/sql"
	"fmt"
	"resultra/tracker/server/generic/stringValidation"
)

func validateUniqueItemListName(trackerDBHandle *sql.DB, databaseID string, listID string, listName string) error {
	// Query to validate the name is unique:
	// 1. Select all the forms in the same database
	// 2. Include forms with the same name.
	// 3. Exclude forms with the same form ID. In other words
	//    the name is considered valid if it is the same as its
	//    existing name.
	rows, queryErr := trackerDBHandle.Query(
		`SELECT list_id,name 
			FROM item_lists
			WHERE database_id=$1 AND name=$2 AND list_id<>$3`,
		databaseID, listName, listID)
	if queryErr != nil {
		return fmt.Errorf("System error validating list name (%v)", queryErr)
	}
	defer rows.Close()

	existingListNameUsedByAnotherList := rows.Next()
	if existingListNameUsedByAnotherList {
		return fmt.Errorf("Invalid list name - names must be unique")
	}

	return nil

}

func validateItemListName(trackerDBHandle *sql.DB, listID string, listName string) error {

	if !stringValidation.WellFormedItemName(listName) {
		return fmt.Errorf("Invalid list name")
	}

	databaseID, err := GetItemListDatabaseID(trackerDBHandle, listID)
	if err != nil {
		return fmt.Errorf("System error validating form name (%v)", err)
	}

	if uniqueErr := validateUniqueItemListName(trackerDBHandle, databaseID, listID, listName); uniqueErr != nil {
		return uniqueErr
	}

	return nil
}

func validateNewItemListName(trackerDBHandle *sql.DB, databaseID string, listName string) error {

	if !stringValidation.WellFormedItemName(listName) {
		return fmt.Errorf("Invalid list name")
	}

	// No form will have an empty formID, so this will cause test for unique
	// form names to return true if any form already has the given formName.
	listID := ""
	if uniqueErr := validateUniqueItemListName(trackerDBHandle, databaseID, listID, listName); uniqueErr != nil {
		return uniqueErr
	}

	return nil
}
