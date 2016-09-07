package global

import (
	"fmt"
	"resultra/datasheet/server/generic"
	"time"
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

func (params SetTextGlobalValueParams) valueType() string { return GlobalTypeText }

func (params SetTextGlobalValueParams) generateValue() (string, error) {

	val := TextValue{Val: params.Value}

	return generic.EncodeJSONString(val)
}

type SetBoolGlobalValueParams struct {
	GlobalValUpdateHeader
	Value bool `json:"value"`
}

func (params SetBoolGlobalValueParams) valueType() string { return GlobalTypeBool }

func (params SetBoolGlobalValueParams) generateValue() (string, error) {

	val := BoolValue{Val: params.Value}

	return generic.EncodeJSONString(val)
}

type SetTimeGlobalValueParams struct {
	GlobalValUpdateHeader
	Value time.Time `json:"value"`
}

func (params SetTimeGlobalValueParams) valueType() string { return GlobalTypeTime }

func (params SetTimeGlobalValueParams) generateValue() (string, error) {

	val := TimeValue{Val: params.Value}

	return generic.EncodeJSONString(val)
}

type SetNumberGlobalValueParams struct {
	GlobalValUpdateHeader
	Value float64 `json:"value"`
}

func (params SetNumberGlobalValueParams) valueType() string { return GlobalTypeNumber }

func (params SetNumberGlobalValueParams) generateValue() (string, error) {

	val := NumberValue{Val: params.Value}

	return generic.EncodeJSONString(val)
}

type SetImageGlobalValueParams struct {
	GlobalValUpdateHeader
	CloudFileName string `json:"cloudFileName"`
	OrigFileName  string `json:"origFileName"`
}

func (params SetImageGlobalValueParams) valueType() string { return GlobalTypeImage }

func (params SetImageGlobalValueParams) generateValue() (string, error) {

	val := ImageValue{CloudFileName: params.CloudFileName,
		OrigFileName: params.OrigFileName}

	return generic.EncodeJSONString(val)
}
