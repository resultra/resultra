package record

import (
	"fmt"
	"time"
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

func (recFieldVals RecFieldValues) GetNumberFieldValue(fieldID string) (float64, bool) {
	rawVal, foundVal := recFieldVals[fieldID]

	if !foundVal {
		return 0.0, false
	}

	if theNum, validType := rawVal.(float64); validType {
		return theNum, true
	} else {
		return 0.0, false
	}
}

func (recFieldVals RecFieldValues) GetBoolFieldValue(fieldID string) (bool, bool) {

	rawVal, foundVal := recFieldVals[fieldID]

	if !foundVal {
		return false, false
	}

	if theBool, validType := rawVal.(bool); validType {
		return theBool, true
	} else {
		return false, false
	}
}

func (recFieldVals RecFieldValues) GetTimeFieldValue(fieldID string) (time.Time, bool) {
	// Time fields are stored as strings when serialized using the RecFieldValues
	// To return the actual date, the string needs to be deserialized into a time.Time type.
	timeVal := time.Time{}

	rawVal, foundVal := recFieldVals[fieldID]
	if !foundVal {
		return timeVal, false
	}

	valAsTime, timeValFound := rawVal.(time.Time)
	if timeValFound {
		return valAsTime, true
	} else {
		return timeVal, false
	}
}
