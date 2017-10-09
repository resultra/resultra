package field

import (
	"fmt"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/common/databaseWrapper"
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
		return fmt.Errorf("System error validating field name (%v)", queryErr)
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

	// No field will have an empty formID, so this will cause test for unique
	// form names to return true if any form already has the given formName.
	fieldID := ""
	if uniqueErr := validateUniqueFieldName(databaseID, fieldID, fieldName); uniqueErr != nil {
		return uniqueErr
	}

	return nil
}

func validateUniqueFieldRefName(databaseID string, fieldID string, refName string) error {
	// Query to validate the name is unique:
	// 1. Select all the fields in the same database
	// 2. Include fields with the same name.
	// 3. Exclude fields with the same form ID. In other words
	//    the name is considered valid if it is the same as its
	//    existing name.
	rows, queryErr := databaseWrapper.DBHandle().Query(
		`SELECT fields.field_id,fields.ref_name 
			FROM fields,databases
			WHERE databases.database_id=$1 AND
			fields.database_id=databases.database_id AND
				fields.ref_name=$2 AND fields.field_id<>$3`,
		databaseID, refName, fieldID)
	if queryErr != nil {
		return fmt.Errorf("System error validating field reference name (%v)", queryErr)
	}

	existingFieldNameUsedByAnotherField := rows.Next()
	if existingFieldNameUsedByAnotherField {
		return fmt.Errorf("Invalid field reference name - names must be unique")
	}

	return nil

}

func validateExistingFieldRefName(fieldID string, refName string) error {

	if !generic.WellFormedFormulaReferenceName(refName) {
		return fmt.Errorf("Invalid field reference name")
	}

	fieldInfo, err := GetField(fieldID)
	if err != nil {
		return fmt.Errorf("System error validating field name (%v)", err)
	}

	if uniqueErr := validateUniqueFieldName(fieldInfo.ParentDatabaseID, fieldID, refName); uniqueErr != nil {
		return uniqueErr
	}

	return nil
}

func validateNewFieldRefName(databaseID string, refName string) error {

	if !generic.WellFormedFormulaReferenceName(refName) {
		return fmt.Errorf("Invalid field reference name")
	}

	// No field will have an empty fieldID, so this will cause test for unique
	// form names to return true if any form already has the given formName.
	fieldID := ""
	if uniqueErr := validateUniqueFieldName(databaseID, fieldID, refName); uniqueErr != nil {
		return uniqueErr
	}

	return nil
}

type FieldTypeValidationFunc func(string) bool

func ValidateField(fieldID string, fieldTypeValidationFunc FieldTypeValidationFunc) error {
	field, fieldErr := GetField(fieldID)
	if fieldErr != nil {
		return fmt.Errorf("ValidateField: Can't get field with field ID = '%v': datastore error=%v",
			fieldID, fieldErr)
	}

	if !fieldTypeValidationFunc(field.Type) {
		return fmt.Errorf("ValidateField: Invalid field type %v", field.Type)
	}

	return nil

}
