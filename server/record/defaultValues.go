package record

import (
	"fmt"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/generic/uniqueID"
)

const defaultValIDTrue string = "true"
const defaultValIDFalse string = "false"

type DefaultFieldValue struct {
	FieldID        string `json:"fieldID"`
	DefaultValueID string `json:"defaultValueID"`
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

type SetDefaultValFunc func(currUserID string, recUpdateHeader RecordUpdateHeader, defaultVal DefaultFieldValue) (*Record, error)
type DefaultValIDDefaultValFuncMap map[string]SetDefaultValFunc

var defaultValCellUpdateValFormat CellUpdateValueFormat

func init() {

	defaultValCellUpdateValFormat = CellUpdateValueFormat{"defaultVal", "general"}

}

func setBoolTrueDefaultValue(currUserID string, recUpdateHeader RecordUpdateHeader, defaultVal DefaultFieldValue) (*Record, error) {

	trueVal := true

	setValParams := SetRecordBoolValueParams{
		RecordUpdateHeader: recUpdateHeader,
		Value:              &trueVal,
		ValueFormat:        defaultValCellUpdateValFormat}

	return UpdateRecordValue(currUserID, setValParams)
}

func setBoolFalseDefaultValue(currUserID string, recUpdateHeader RecordUpdateHeader, defaultVal DefaultFieldValue) (*Record, error) {

	falseVal := false

	setValParams := SetRecordBoolValueParams{
		RecordUpdateHeader: recUpdateHeader,
		Value:              &falseVal,
		ValueFormat:        defaultValCellUpdateValFormat}

	return UpdateRecordValue(currUserID, setValParams)
}

var boolFieldDefaultValFuncs = DefaultValIDDefaultValFuncMap{
	defaultValIDTrue:  setBoolTrueDefaultValue,
	defaultValIDFalse: setBoolFalseDefaultValue}

// Get the rule definition based upon the field type
func getDefaultValFuncByFieldType(fieldType string, defaultVal DefaultFieldValue) (SetDefaultValFunc, error) {
	switch fieldType {
	// TODO	case field.FieldTypeText:
	// TODO		case field.FieldTypeNumber:
	// TODO		case field.FieldTypeTime:
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

func SetDefaultValues(currUserID string, params SetDefaultValsParams) error {
	for _, defaultVal := range params.DefaultVals {

		updateHeader := RecordUpdateHeader{
			ParentDatabaseID: params.ParentDatabaseID,
			RecordID:         params.RecordID,
			ChangeSetID:      params.ChangeSetID,
			FieldID:          defaultVal.FieldID}

		field, err := field.GetField(defaultVal.FieldID)
		if err != nil {
			return fmt.Errorf("SetDefaultValues: failure retrieving referenced field: %+v", err)
		}

		defaultValFunc, funcErr := getDefaultValFuncByFieldType(field.Type, defaultVal)
		if funcErr != nil {
			return fmt.Errorf("SetDefaultValues: failure retrieving function to set default val: %+v", funcErr)
		}

		_, setErr := defaultValFunc(currUserID, updateHeader, defaultVal)
		if setErr != nil {
			return fmt.Errorf("SetDefaultValues: failure setting default val: %+v", funcErr)
		}
	}
	return nil
}

func ValidateWellFormedDefaultValues(defaultVals []DefaultFieldValue) error {
	for _, defaultVal := range defaultVals {

		field, err := field.GetField(defaultVal.FieldID)
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
