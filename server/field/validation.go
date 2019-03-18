// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package field

import (
	"database/sql"
	"fmt"
	"resultra/tracker/server/generic"
	"resultra/tracker/server/generic/stringValidation"
)

func validateUniqueFieldName(trackerDBHandle *sql.DB, databaseID string, fieldID string, fieldName string) error {
	// Query to validate the name is unique:
	// 1. Select all the fields in the same database
	// 2. Include fields with the same name.
	// 3. Exclude fields with the same form ID. In other words
	//    the name is considered valid if it is the same as its
	//    existing name.
	rows, queryErr := trackerDBHandle.Query(
		`SELECT fields.field_id,fields.name 
			FROM fields,databases
			WHERE databases.database_id=$1 AND
			fields.database_id=databases.database_id AND
				fields.name=$2 AND fields.field_id<>$3`,
		databaseID, fieldName, fieldID)
	if queryErr != nil {
		return fmt.Errorf("System error validating field name (%v)", queryErr)
	}
	defer rows.Close()

	existingFieldNameUsedByAnotherField := rows.Next()
	if existingFieldNameUsedByAnotherField {
		return fmt.Errorf("invalid field name - names must be unique: found existing field with name = %v", fieldName)
	}

	return nil

}

func validateExistingFieldName(trackerDBHandle *sql.DB, fieldID string, fieldName string) error {

	if !stringValidation.WellFormedItemName(fieldName) {
		return fmt.Errorf("Invalid field name")
	}

	fieldInfo, err := GetField(trackerDBHandle, fieldID)
	if err != nil {
		return fmt.Errorf("System error validating field name (%v)", err)
	}

	if uniqueErr := validateUniqueFieldName(trackerDBHandle, fieldInfo.ParentDatabaseID, fieldID, fieldName); uniqueErr != nil {
		return uniqueErr
	}

	return nil
}

func validateNewFieldName(trackerDBHandle *sql.DB, databaseID string, fieldName string) error {

	if !stringValidation.WellFormedItemName(fieldName) {
		return fmt.Errorf("mal-formed field name: %v", fieldName)
	}

	// No field will have an empty formID, so this will cause test for unique
	// form names to return true if any form already has the given formName.
	fieldID := ""
	if uniqueErr := validateUniqueFieldName(trackerDBHandle, databaseID, fieldID, fieldName); uniqueErr != nil {
		return uniqueErr
	}

	return nil
}

func validateUniqueFieldRefName(trackerDBHandle *sql.DB, databaseID string, fieldID string, refName string) error {
	// Query to validate the name is unique:
	// 1. Select all the fields in the same database
	// 2. Include fields with the same name.
	// 3. Exclude fields with the same form ID. In other words
	//    the name is considered valid if it is the same as its
	//    existing name.
	rows, queryErr := trackerDBHandle.Query(
		`SELECT fields.field_id,fields.ref_name 
			FROM fields,databases
			WHERE databases.database_id=$1 AND
			fields.database_id=databases.database_id AND
				fields.ref_name=$2 AND fields.field_id<>$3`,
		databaseID, refName, fieldID)
	if queryErr != nil {
		return fmt.Errorf("System error validating field reference name (%v)", queryErr)
	}
	defer rows.Close()

	existingFieldNameUsedByAnotherField := rows.Next()
	if existingFieldNameUsedByAnotherField {
		return fmt.Errorf("invalid field reference name - names must be unique: found existing field with reference name = %v", refName)
	}

	return nil

}

func validateExistingFieldRefName(trackerDBHandle *sql.DB, fieldID string, refName string) error {

	if !generic.WellFormedFormulaReferenceName(refName) {
		return fmt.Errorf("Invalid field reference name")
	}

	fieldInfo, err := GetField(trackerDBHandle, fieldID)
	if err != nil {
		return fmt.Errorf("System error validating field name (%v)", err)
	}

	if uniqueErr := validateUniqueFieldName(trackerDBHandle, fieldInfo.ParentDatabaseID, fieldID, refName); uniqueErr != nil {
		return uniqueErr
	}

	return nil
}

func validateNewFieldRefName(trackerDBHandle *sql.DB, databaseID string, refName string) error {

	if !generic.WellFormedFormulaReferenceName(refName) {
		return fmt.Errorf("Invalid field reference name")
	}

	// No field will have an empty fieldID, so this will cause test for unique
	// form names to return true if any form already has the given formName.
	fieldID := ""
	if uniqueErr := validateUniqueFieldRefName(trackerDBHandle, databaseID, fieldID, refName); uniqueErr != nil {
		return uniqueErr
	}

	return nil
}

type FieldTypeValidationFunc func(string) bool

func ValidateField(trackerDBHandle *sql.DB, fieldID string, fieldTypeValidationFunc FieldTypeValidationFunc) error {
	field, fieldErr := GetField(trackerDBHandle, fieldID)
	if fieldErr != nil {
		return fmt.Errorf("ValidateField: Can't get field with field ID = '%v': datastore error=%v",
			fieldID, fieldErr)
	}

	if !fieldTypeValidationFunc(field.Type) {
		return fmt.Errorf("ValidateField: Invalid field type %v", field.Type)
	}

	return nil

}
