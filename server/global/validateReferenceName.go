package global

import (
	"fmt"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/databaseWrapper"
)

func validateUniqueReferenceName(databaseID string, globalID string, referenceName string) error {
	// Query to validate the reference name is unique:
	// 1. Select all the globals in the same database
	// 2. Include globals with the same reference name.
	// 3. Exclude globals with the same globals ID. In other words
	//    the name is considered valid if it is the same as its
	//    existing name.
	rows, queryErr := databaseWrapper.DBHandle().Query(
		`SELECT globals.global_id,globals.ref_name 
			FROM globals,databases
			WHERE databases.database_id=$1 AND
				globals.database_id=databases.database_id AND
				globals.ref_name=$2 AND globals.global_id<>$3`,
		databaseID, referenceName, globalID)
	if queryErr != nil {
		return fmt.Errorf("System error validating global name (%v)", queryErr)
	}

	existingNameUsedByAnotherGlobal := rows.Next()
	if existingNameUsedByAnotherGlobal {
		return fmt.Errorf("Invalid formula reference name - names must be unique")
	}

	return nil

}

func validateReferenceName(globalID string, referenceName string) error {

	if !generic.WellFormedFormulaReferenceName(referenceName) {
		return fmt.Errorf("Invalid formula reference name: '%v' Cannot be empty and must only contain letters, numbers and underscores",
			referenceName)

	}

	databaseID, err := getGlobalDatabaseID(globalID)
	if err != nil {
		return fmt.Errorf("System error validating global name (%v)", err)
	}

	if uniqueErr := validateUniqueReferenceName(databaseID, globalID, referenceName); uniqueErr != nil {
		return uniqueErr
	}

	return nil
}

func validateNewReferenceName(databaseID string, referenceName string) error {

	if !generic.WellFormedFormulaReferenceName(referenceName) {
		return fmt.Errorf("Invalid formula reference name: '%v' Cannot be empty and must only contain letters, numbers and underscores",
			referenceName)
	}

	// No global will have an empty global ID, so this will cause test for unique
	// global names to return true if any global already has the given referenceName.
	globalID := ""
	if uniqueErr := validateUniqueReferenceName(databaseID, globalID, referenceName); uniqueErr != nil {
		return uniqueErr
	}

	return nil
}
