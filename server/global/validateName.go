// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package global

import (
	"database/sql"
	"fmt"
	"resultra/tracker/server/generic/stringValidation"
)

func validateUniqueGlobalName(trackerDBHandle *sql.DB, databaseID string, globalID string, globalName string) error {
	// Query to validate the name is unique:
	// 1. Select all the globals in the same database
	// 2. Include globals with the same name.
	// 3. Exclude globals with the same globals ID. In other words
	//    the name is considered valid if it is the same as its
	//    existing name.
	rows, queryErr := trackerDBHandle.Query(
		`SELECT globals.global_id,globals.name 
			FROM globals,databases
			WHERE databases.database_id=$1 AND
				globals.database_id=databases.database_id AND
				globals.name=$2 AND globals.global_id<>$3`,
		databaseID, globalName, globalID)
	if queryErr != nil {
		return fmt.Errorf("System error validating global name (%v)", queryErr)
	}
	defer rows.Close()

	existingGlobalNameUsedByAnotherForm := rows.Next()
	if existingGlobalNameUsedByAnotherForm {
		return fmt.Errorf("Invalid global name - names must be unique")
	}

	return nil

}

func validateGlobalName(trackerDBHandle *sql.DB, globalID string, globalName string) error {

	if !stringValidation.WellFormedItemName(globalName) {
		return fmt.Errorf("Invalid name")
	}

	databaseID, err := getGlobalDatabaseID(trackerDBHandle, globalID)
	if err != nil {
		return fmt.Errorf("System error validating global name (%v)", err)
	}

	if uniqueErr := validateUniqueGlobalName(trackerDBHandle, databaseID, globalID, globalName); uniqueErr != nil {
		return uniqueErr
	}

	return nil
}

func validateNewGlobalName(trackerDBHandle *sql.DB, databaseID string, globalName string) error {

	if !stringValidation.WellFormedItemName(globalName) {
		return fmt.Errorf("Invalid global name")
	}

	// No global will have an empty formID, so this will cause test for unique
	// global names to return true if any global already has the given globalName.
	globalID := ""
	if uniqueErr := validateUniqueGlobalName(trackerDBHandle, databaseID, globalID, globalName); uniqueErr != nil {
		return uniqueErr
	}

	return nil
}
