package field

import (
	"fmt"
	"resultra/datasheet/server/generic/databaseWrapper"
	"resultra/datasheet/server/generic/stringValidation"
)

func validateUniqueFieldName(databaseID string, fieldID string, fieldName string) error {
	// Query to validate the name is unique:
	// 1. Select all the fields in the same database
	// 2. Include fields with the same name.
	// 3. Exclude fields with the same form ID. In other words
	//    the name is considered valid if it is the same as its
	//    existing name.
	rows, queryErr := databaseWrapper.DBHandle().Query(
		`SELECT fields.field_id,fields.name 
			FROM fields,databases
			WHERE databases.database_id=$1 AND
			fields.database_id=databases.database_id AND
				fields.name=$2 AND fields.field_id<>$3`,
		databaseID, fieldName, fieldID)
	if queryErr != nil {
		return fmt.Errorf("System error validating form name (%v)", queryErr)
	}

	existingFieldNameUsedByAnotherField := rows.Next()
	if existingFieldNameUsedByAnotherField {
		return fmt.Errorf("Invalid field name - names must be unique")
	}

	return nil

}

func validateExistingFieldName(fieldID string, fieldName string) error {

	if !stringValidation.WellFormedItemName(fieldName) {
		return fmt.Errorf("Invalid field name")
	}

	fieldInfo, err := GetField(fieldID)
	if err != nil {
		return fmt.Errorf("System error validating field name (%v)", err)
	}

	if uniqueErr := validateUniqueFieldName(fieldInfo.ParentDatabaseID, fieldID, fieldName); uniqueErr != nil {
		return uniqueErr
	}

	return nil
}

func validateNewFieldName(databaseID string, fieldName string) error {

	if !stringValidation.WellFormedItemName(fieldName) {
		return fmt.Errorf("Invalid field name")
	}

	// No form will have an empty formID, so this will cause test for unique
	// form names to return true if any form already has the given formName.
	fieldID := ""
	if uniqueErr := validateUniqueFieldName(databaseID, fieldID, fieldName); uniqueErr != nil {
		return uniqueErr
	}

	return nil
}
