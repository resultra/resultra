package global

import (
	"fmt"
	"resultra/datasheet/server/generic"
)

type GlobalValUpdater interface {
	parentDatabaseID() string
	globalID() string
	valueType() string
	generateValue() (string, error)
}

// RecordUpdateHeader is a common header for all record value updates. It also implements
// part of the RecorddUpdater interface. This struct should be embedded in other structs
// used to update values of specific types.
type GlobalValUpdateHeader struct {
	ParentDatabaseID string `json:"parentDatabaseID"`
	GlobalID         string `json:"globalID"`
}

func (header GlobalValUpdateHeader) parentDatabaseID() string {
	return header.ParentDatabaseID
}

func (header GlobalValUpdateHeader) globalID() string {
	return header.GlobalID
}

func updateGlobalValue(updater GlobalValUpdater) (*GlobalValUpdate, error) {

	global, getErr := getGlobal(updater.globalID())
	if getErr != nil {
		return nil, fmt.Errorf("UpdateGlobalValue: Can't retrieve global with given ID: %v", getErr)
	}

	if global.Type != updater.valueType() {
		return nil, fmt.Errorf("UpdateGlobalValue: type mismatch: got %v, expecting %v", updater.valueType(), global.Type)
	}

	encodedValue, encodeErr := updater.generateValue()
	if encodeErr != nil {
		return nil, fmt.Errorf("UpdateGlobalValue: Error encoding value: %v", encodeErr)
	}

	valUpdate, updateErr := saveValUpdate(updater.globalID(), encodedValue)
	if updateErr != nil {
		return nil, fmt.Errorf("UpdateGlobalValue: Error saving value: %v", updateErr)
	}

	return valUpdate, nil

}

type SetTextGlobalValueParams struct {
	GlobalValUpdateHeader
	Value string `json:"value"`
}

type TextValue struct {
	Val string `json:"val"`
}

func (params SetTextGlobalValueParams) valueType() string { return GlobalTypeText }

func (params SetTextGlobalValueParams) generateValue() (string, error) {

	val := TextValue{Val: params.Value}

	return generic.EncodeJSONString(val)
}
