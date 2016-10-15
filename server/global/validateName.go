package global

import (
	"fmt"
	"resultra/datasheet/server/generic/databaseWrapper"
	"resultra/datasheet/server/generic/stringValidation"
)

func validateUniqueGlobalName(databaseID string, globalID string, globalName string) error {
	// Query to validate the name is unique:
	// 1. Select all the globals in the same database
	// 2. Include globals with the same name.
	// 3. Exclude globals with the same globals ID. In other words
	//    the name is considered valid if it is the same as its
	//    existing name.
	rows, queryErr := databaseWrapper.DBHandle().Query(
		`SELECT globals.global_id,globals.name 
			FROM globals,databases
			WHERE databases.database_id=$1 AND
				globals.database_id=databases.database_id AND
				globals.name=$2 AND globals.global_id<>$3`,
		databaseID, globalName, globalID)
	if queryErr != nil {
		return fmt.Errorf("System error validating global name (%v)", queryErr)
	}

	existingGlobalNameUsedByAnotherForm := rows.Next()
	if existingGlobalNameUsedByAnotherForm {
		return fmt.Errorf("Invalid global name - names must be unique")
	}

	return nil

}

func validateGlobalName(globalID string, globalName string) error {

	if !stringValidation.WellFormedItemName(globalName) {
		return fmt.Errorf("Invalid name")
	}

	databaseID, err := getGlobalDatabaseID(globalID)
	if err != nil {
		return fmt.Errorf("System error validating global name (%v)", err)
	}

	if uniqueErr := validateUniqueGlobalName(databaseID, globalID, globalName); uniqueErr != nil {
		return uniqueErr
	}

	return nil
}

func validateNewGlobalName(databaseID string, globalName string) error {

	if !stringValidation.WellFormedItemName(globalName) {
		return fmt.Errorf("Invalid global name")
	}

	// No global will have an empty formID, so this will cause test for unique
	// global names to return true if any global already has the given globalName.
	globalID := ""
	if uniqueErr := validateUniqueGlobalName(databaseID, globalID, globalName); uniqueErr != nil {
		return uniqueErr
	}

	return nil
}
