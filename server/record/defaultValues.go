package record

import (
	"database/sql"
	"fmt"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/generic/uniqueID"
	"time"
)

const defaultValIDTrue string = "true"
const defaultValIDFalse string = "false"

type DefaultFieldValue struct {
	FieldID        string   `json:"fieldID"`
	DefaultValueID string   `json:"defaultValueID"`
	NumberVal      *float64 `json:"numberVal,omitempty"`
}

func (srcDefaultVal DefaultFieldValue) Clone(remappedIDs uniqueID.UniqueIDRemapper) (*DefaultFieldValue, error) {

	remappedFieldID, err := remappedIDs.GetExistingRemappedID(srcDefaultVal.FieldID)
	if err != nil {
		return nil, fmt.Errorf("DefaultFieldValue.Clone: %v", err)
	}

	destDefaultVal := srcDefaultVal
	destDefaultVal.FieldID = remappedFieldID

	return &destDefaultVal, nil
}

func CloneDefaultFieldValues(remappedIDs uniqueID.UniqueIDRemapper, srcDefaultVals []DefaultFieldValue) ([]DefaultFieldValue, error) {

	destDefaultVals := []DefaultFieldValue{}

	for _, srcDefaultVal := range srcDefaultVals {
		destDefaultVal, err := srcDefaultVal.Clone(remappedIDs)
		if err != nil {
			return nil, fmt.Errorf("CloneDefaultFieldValues: %v", err)
		}
		destDefaultVals = append(destDefaultVals, *destDefaultVal)
	}

	return destDefaultVals, nil
}

type SetDefaultValFunc func(trackerDBHandle *sql.DB, currUserID string, recUpdateHeader RecordUpdateHeader, defaultVal DefaultFieldValue) (*Record, error)
type DefaultValIDDefaultValFuncMap map[string]SetDefaultValFunc

func setBoolTrueDefaultValue(trackerDBHandle *sql.DB, currUserID string, recUpdateHeader RecordUpdateHeader, defaultVal DefaultFieldValue) (*Record, error) {

	trueVal := true

	setValParams := SetRecordBoolValueParams{
		RecordUpdateHeader: recUpdateHeader,
		Value:              &trueVal}

	return UpdateRecordValue(trackerDBHandle, currUserID, setValParams)
}

func setBoolFalseDefaultValue(trackerDBHandle *sql.DB, currUserID string, recUpdateHeader RecordUpdateHeader, defaultVal DefaultFieldValue) (*Record, error) {

	falseVal := false

	setValParams := SetRecordBoolValueParams{
		RecordUpdateHeader: recUpdateHeader,
		Value:              &falseVal}

	return UpdateRecordValue(trackerDBHandle, currUserID, setValParams)
}

var boolFieldDefaultValFuncs = DefaultValIDDefaultValFuncMap{
	defaultValIDTrue:  setBoolTrueDefaultValue,
	defaultValIDFalse: setBoolFalseDefaultValue}

const defaultValIDCurrTime string = "currentTime"

func setCurrTimeDefaultValue(trackerDBHandle *sql.DB, currUserID string, recUpdateHeader RecordUpdateHeader,
	defaultVal DefaultFieldValue) (*Record, error) {

	currTime := time.Now().UTC()

	setValParams := SetRecordTimeValueParams{
		RecordUpdateHeader: recUpdateHeader,
		Value:              &currTime}

	return UpdateRecordValue(trackerDBHandle, currUserID, setValParams)
}

var timeFieldDefaultValFuncs = DefaultValIDDefaultValFuncMap{
	defaultValIDCurrTime: setCurrTimeDefaultValue}

const defaultValIDSpecificVal string = "specificVal"

func setSpecificNumberValDefaultValue(trackerDBHandle *sql.DB, currUserID string, recUpdateHeader RecordUpdateHeader,
	defaultVal DefaultFieldValue) (*Record, error) {

	if defaultVal.NumberVal == nil {
		return nil, fmt.Errorf("setSpecificNumberValDefaultValue: missing default value")
	}

	setValParams := SetRecordNumberValueParams{
		RecordUpdateHeader: recUpdateHeader,
		Value:              defaultVal.NumberVal}

	return UpdateRecordValue(trackerDBHandle, currUserID, setValParams)
}

var numFieldDefaultValFuncs = DefaultValIDDefaultValFuncMap{
	defaultValIDSpecificVal: setSpecificNumberValDefaultValue}

// Get the rule definition based upon the field type
func getDefaultValFuncByFieldType(fieldType string, defaultVal DefaultFieldValue) (SetDefaultValFunc, error) {
	switch fieldType {
	// TODO	case field.FieldTypeText:
	// TODO		case field.FieldTypeNumber:
	case field.FieldTypeBool:
		defaultValFunc, funcFound := boolFieldDefaultValFuncs[defaultVal.DefaultValueID]
		if !funcFound {
			return nil, fmt.Errorf(
				`getRuleDefByFieldType: Failed to retrieve function to set default value function for field type = %v, 
					unrecognized default value ID = %v`,
				fieldType, defaultVal.DefaultValueID)
		} else {
			return defaultValFunc, nil
		}
	case field.FieldTypeTime:
		defaultValFunc, funcFound := timeFieldDefaultValFuncs[defaultVal.DefaultValueID]
		if !funcFound {
			return nil, fmt.Errorf(
				`getRuleDefByFieldType: Failed to retrieve function to set default value function for field type = %v, 
					unrecognized default value ID = %v`,
				fieldType, defaultVal.DefaultValueID)
		} else {
			return defaultValFunc, nil
		}
	case field.FieldTypeNumber:
		defaultValFunc, funcFound := numFieldDefaultValFuncs[defaultVal.DefaultValueID]
		if !funcFound {
			return nil, fmt.Errorf(
				`getRuleDefByFieldType: Failed to retrieve function to set default value function for field type = %v, 
					unrecognized default value ID = %v`,
				fieldType, defaultVal.DefaultValueID)
		} else {
			return defaultValFunc, nil
		}
	default:
		return nil, fmt.Errorf(
			"getDefaultValFuncByFieldType: Failed to retrieve function to set default value: unknown field type = %v",
			fieldType)
	}
}

type SetDefaultValsParams struct {
	ParentDatabaseID string              `json:"parentDatabaseID"`
	RecordID         string              `json:"recordID"`
	ChangeSetID      string              `json:"changeSetID"`
	DefaultVals      []DefaultFieldValue `json:defaultVals`
}

func SetDefaultValues(trackingDBHandle *sql.DB, currUserID string, params SetDefaultValsParams) error {
	for _, defaultVal := range params.DefaultVals {

		updateHeader := RecordUpdateHeader{
			ParentDatabaseID: params.ParentDatabaseID,
			RecordID:         params.RecordID,
			ChangeSetID:      params.ChangeSetID,
			FieldID:          defaultVal.FieldID}

		field, err := field.GetField(trackingDBHandle, defaultVal.FieldID)
		if err != nil {
			return fmt.Errorf("SetDefaultValues: failure retrieving referenced field: %+v", err)
		}

		defaultValFunc, funcErr := getDefaultValFuncByFieldType(field.Type, defaultVal)
		if funcErr != nil {
			return fmt.Errorf("SetDefaultValues: failure retrieving function to set default val: %+v", funcErr)
		}

		_, setErr := defaultValFunc(trackingDBHandle, currUserID, updateHeader, defaultVal)
		if setErr != nil {
			return fmt.Errorf("SetDefaultValues: failure setting default val: %+v", funcErr)
		}
	}
	return nil
}

func ValidateWellFormedDefaultValues(trackingDBHandle *sql.DB, defaultVals []DefaultFieldValue) error {
	for _, defaultVal := range defaultVals {

		field, err := field.GetField(trackingDBHandle, defaultVal.FieldID)
		if err != nil {
			return fmt.Errorf("ValidateWellFormedDefaultValues: failure retrieving referenced field: %+v", err)
		}

		_, funcErr := getDefaultValFuncByFieldType(field.Type, defaultVal)
		if funcErr != nil {
			return fmt.Errorf("ValidateWellFormedDefaultValues: failure retrieving function to set default val: %+v", funcErr)
		}

	}
	return nil
}
