package record

import (
	"fmt"
)

// RecFieldValues is the low-level/base type for storing field values of different types,
// mapped by field ID. Users of this type are responsible to know what type of object is
// stored in the map and make the appropriate function calls.
type RecFieldValues map[string]interface{}

func (recFieldVals RecFieldValues) ValueIsSet(fieldID string) bool {
	_, valueExists := recFieldVals[fieldID]
	if valueExists {
		return true
	} else {
		return false
	}
}

func (recFieldVals RecFieldValues) GetTextFieldValue(fieldID string) (string, error) {
	rawVal := recFieldVals[fieldID]
	if theStr, validType := rawVal.(string); validType {
		return theStr, nil
	} else {
		return "", fmt.Errorf("Type mismatch retrieving text field value from record: field ID = %v, raw value = %v", fieldID, rawVal)
	}
}
