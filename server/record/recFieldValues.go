package record

import (
	"time"
)

// RecFieldValues is the low-level/base type for storing field values of different types,
// mapped by field ID. Users of this type are responsible to know what type of object is
// stored in the map and make the appropriate function calls.
type RecFieldValues map[string]interface{}

func (recFieldVals RecFieldValues) ValueIsSet(fieldID string) bool {
	theVal, valueExists := recFieldVals[fieldID]
	if valueExists {
		// The value will be set to nil if it is cleared within the UI. This is the same as it not being set
		// in the first place.
		if theVal == nil {
			return false
		} else {
			return true
		}
	} else {
		return false
	}
}

func (recFieldVals RecFieldValues) GetTextFieldValue(fieldID string) (string, bool) {
	rawVal, foundVal := recFieldVals[fieldID]

	if !foundVal {
		return "", false
	}

	if theStr, validType := rawVal.(string); validType {
		return theStr, true
	} else {
		return "", false
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
