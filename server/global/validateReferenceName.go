// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package global

import (
	"database/sql"
	"fmt"
	"resultra/tracker/server/generic"
)

func validateUniqueReferenceName(trackerDBHandle *sql.DB, databaseID string, globalID string, referenceName string) error {
	// Query to validate the reference name is unique:
	// 1. Select all the globals in the same database
	// 2. Include globals with the same reference name.
	// 3. Exclude globals with the same globals ID. In other words
	//    the name is considered valid if it is the same as its
	//    existing name.
	rows, queryErr := trackerDBHandle.Query(
		`SELECT globals.global_id,globals.ref_name 
			FROM globals,databases
			WHERE databases.database_id=$1 AND
				globals.database_id=databases.database_id AND
				globals.ref_name=$2 AND globals.global_id<>$3`,
		databaseID, referenceName, globalID)
	if queryErr != nil {
		return fmt.Errorf("System error validating global name (%v)", queryErr)
	}
	defer rows.Close()

	existingNameUsedByAnotherGlobal := rows.Next()
	if existingNameUsedByAnotherGlobal {
		return fmt.Errorf("Invalid formula reference name - names must be unique")
	}

	return nil

}

func validateReferenceName(trackerDBHandle *sql.DB, globalID string, referenceName string) error {

	if !generic.WellFormedFormulaReferenceName(referenceName) {
		return fmt.Errorf("Invalid formula reference name: '%v' Cannot be empty and must only contain letters, numbers and underscores",
			referenceName)

	}

	databaseID, err := getGlobalDatabaseID(trackerDBHandle, globalID)
	if err != nil {
		return fmt.Errorf("System error validating global name (%v)", err)
	}

	if uniqueErr := validateUniqueReferenceName(trackerDBHandle, databaseID, globalID, referenceName); uniqueErr != nil {
		return uniqueErr
	}

	return nil
}

func validateNewReferenceName(trackerDBHandle *sql.DB, databaseID string, referenceName string) error {

	if !generic.WellFormedFormulaReferenceName(referenceName) {
		return fmt.Errorf("Invalid formula reference name: '%v' Cannot be empty and must only contain letters, numbers and underscores",
			referenceName)
	}

	// No global will have an empty global ID, so this will cause test for unique
	// global names to return true if any global already has the given referenceName.
	globalID := ""
	if uniqueErr := validateUniqueReferenceName(trackerDBHandle, databaseID, globalID, referenceName); uniqueErr != nil {
		return uniqueErr
	}

	return nil
}
