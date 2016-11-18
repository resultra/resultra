package record

import (
	"fmt"
	"log"
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
		log.Printf("GetNumberFieldValue: value not found for fieldID = %v", fieldID)
		return 0.0, false
	}

	if theNum, validType := rawVal.(float64); validType {
		log.Printf("GetNumberFieldValue: got value for fieldID = %v, value = %v", fieldID, theNum)
		return theNum, true
	} else {
		log.Printf("GetNumberFieldValue: invalid type for value for fieldID = %v, value = %v", fieldID, rawVal)
		return 0.0, false
	}
}

// TODO (Important) - Save time values as time.Time. When saved then restored from an interface{} value, time values get restored as
// strings rather than dates. This makes it necessary to decode the strings after the fact. A more type safe way to store the
// values would be to have a different map of values for each type; i.e. bool, time, etc.
func (recFieldVals RecFieldValues) GetTimeFieldValue(fieldID string) (time.Time, bool) {
	// Time fields are stored as strings when serialized using the RecFieldValues
	// To return the actual date, the string needs to be deserialized into a time.Time type.
	timeVal := time.Time{}

	rawVal, foundVal := recFieldVals[fieldID]
	if !foundVal {
		log.Printf("GetTimeFieldValue: rawVal not found for field = %v", fieldID)
		return timeVal, false
	}

	timeStr, foundStrVal := rawVal.(string)
	if !foundStrVal {
		log.Printf("GetTimeFieldValue: string not found")
		return timeVal, false
	}

	timeVal, parseErr := time.Parse(time.RFC3339, timeStr)
	if parseErr != nil {
		log.Printf("GetTimeFieldValue: parse failed")
		return timeVal, false
	}

	return timeVal, true

}
