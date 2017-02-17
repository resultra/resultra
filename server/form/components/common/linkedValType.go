package common

import (
	"fmt"
	"resultra/datasheet/server/field"
)

type FieldTypeValidationFunc func(string) bool

func ValidateField(fieldID string, fieldTypeValidationFunc FieldTypeValidationFunc) error {
	field, fieldErr := field.GetField(fieldID)
	if fieldErr != nil {
		return fmt.Errorf("ValidateField: Can't get field with field ID = '%v': datastore error=%v",
			fieldID, fieldErr)
	}

	if !fieldTypeValidationFunc(field.Type) {
		return fmt.Errorf("ValidateField: Invalid field type %v", field.Type)
	}

	return nil

}
