package recordUpdate

import (
	"appengine"
	"fmt"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/record"
	"time"
)

// Update a text field value

type SetRecordTextValueParams struct {
	RecordUpdateHeader
	Value string `json:"value"`
}

func (setValParams SetRecordTextValueParams) fieldType() string { return field.FieldTypeText }

func (setValParams SetRecordTextValueParams) updateRecordValue(rec *record.Record) {
	(*rec).FieldValues[setValParams.FieldID] = setValParams.Value
}

type TextCellValue struct {
	Val string `json:"val"`
}

func (valParams SetRecordTextValueParams) generateCellValue() (string, error) {

	cellVal := TextCellValue{Val: valParams.Value}

	return generic.EncodeJSONString(cellVal)
}

// Update a long text field value

type SetRecordLongTextValueParams struct {
	RecordUpdateHeader
	Value string `json:"value"`
}

func (setValParams SetRecordLongTextValueParams) fieldType() string { return field.FieldTypeLongText }

func (setValParams SetRecordLongTextValueParams) updateRecordValue(rec *record.Record) {
	(*rec).FieldValues[setValParams.FieldID] = setValParams.Value
}

func (valParams SetRecordLongTextValueParams) generateCellValue() (string, error) {
	return generic.EncodeJSONString(valParams.Value)
}

// Update a number field value

type SetRecordNumberValueParams struct {
	RecordUpdateHeader
	Value float64 `json:"value"`
}

func (setValParams SetRecordNumberValueParams) fieldType() string { return field.FieldTypeNumber }

func (setValParams SetRecordNumberValueParams) updateRecordValue(rec *record.Record) {
	(*rec).FieldValues[setValParams.FieldID] = setValParams.Value
}

type NumberCellValue struct {
	Val float64 `json:'val'`
}

func (valParams SetRecordNumberValueParams) generateCellValue() (string, error) {

	cellVal := NumberCellValue{Val: valParams.Value}

	return generic.EncodeJSONString(cellVal)
}

// Update a number field value

type SetRecordBoolValueParams struct {
	RecordUpdateHeader
	Value bool `json:"value"`
}

type BoolCellValue struct {
	Val bool `json:"val"`
}

func (setValParams SetRecordBoolValueParams) fieldType() string { return field.FieldTypeBool }

func (setValParams SetRecordBoolValueParams) updateRecordValue(rec *record.Record) {
	(*rec).FieldValues[setValParams.FieldID] = setValParams.Value
}

func (valParams SetRecordBoolValueParams) generateCellValue() (string, error) {

	cellVal := BoolCellValue{Val: valParams.Value}

	return generic.EncodeJSONString(cellVal)
}

type SetRecordTimeValueParams struct {
	RecordUpdateHeader
	Value time.Time `json:"value"`
}

func (setValParams SetRecordTimeValueParams) fieldType() string { return field.FieldTypeTime }

func (setValParams SetRecordTimeValueParams) updateRecordValue(rec *record.Record) {
	(*rec).FieldValues[setValParams.FieldID] = setValParams.Value
}

type TimeCellValue struct {
	Val time.Time `json:"val"`
}

func (valParams SetRecordTimeValueParams) generateCellValue() (string, error) {

	cellVal := TimeCellValue{Val: valParams.Value}

	return generic.EncodeJSONString(cellVal)
}

type SetRecordFileValueParams struct {
	RecordUpdateHeader
	CloudFileName string `json:"cloudFileName"`
	OrigFileName  string `json:"origFileName"`
}

func (setValParams SetRecordFileValueParams) fieldType() string { return field.FieldTypeFile }

func (setValParams SetRecordFileValueParams) updateRecordValue(rec *record.Record) {
	(*rec).FieldValues[setValParams.FieldID] = setValParams.CloudFileName
}

type FileCellValue struct {
	CloudName string `json:"cloudName"`
	OrigName  string `json:"origName"`
}

func (valParams SetRecordFileValueParams) generateCellValue() (string, error) {

	cellValue := FileCellValue{
		CloudName: valParams.CloudFileName,
		OrigName:  valParams.OrigFileName}

	return generic.EncodeJSONString(cellValue)
}

// setRecordFileNameFieldValue. Although the parameters for a record update with a filename aren't passed through the http request,
// the standard record updating mechanism can be used to update the field with the filename.
func setRecordFileNameFieldValue(appEngContext appengine.Context,
	parentTableID string,
	recordID string, fieldID string, origFileName string, cloudFileName string) (*record.Record, error) {
	updateRecordHeader := RecordUpdateHeader{
		ParentTableID: parentTableID,
		RecordID:      recordID,
		FieldID:       fieldID}
	updateRecordParams := SetRecordFileValueParams{
		RecordUpdateHeader: updateRecordHeader,
		OrigFileName:       origFileName,
		CloudFileName:      cloudFileName}
	updatedRecord, updateErr := UpdateRecordValue(appEngContext, updateRecordParams)
	if updateErr != nil {
		return nil, fmt.Errorf("uploadFile: Unable to update record for newly uploaded file: %v", updateErr)
	}
	return updatedRecord, nil
}
