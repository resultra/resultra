package recordUpdate

import (
	"appengine"
	"fmt"
	"resultra/datasheet/server/field"
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

// Update a long text field value

type SetRecordLongTextValueParams struct {
	RecordUpdateHeader
	Value string `json:"value"`
}

func (setValParams SetRecordLongTextValueParams) fieldType() string { return field.FieldTypeLongText }

func (setValParams SetRecordLongTextValueParams) updateRecordValue(rec *record.Record) {
	(*rec).FieldValues[setValParams.FieldID] = setValParams.Value
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

// Update a number field value

type SetRecordBoolValueParams struct {
	RecordUpdateHeader
	Value bool `json:"value"`
}

func (setValParams SetRecordBoolValueParams) fieldType() string { return field.FieldTypeBool }

func (setValParams SetRecordBoolValueParams) updateRecordValue(rec *record.Record) {
	(*rec).FieldValues[setValParams.FieldID] = setValParams.Value
}

type SetRecordTimeValueParams struct {
	RecordUpdateHeader
	Value time.Time `json:"value"`
}

func (setValParams SetRecordTimeValueParams) fieldType() string { return field.FieldTypeTime }

func (setValParams SetRecordTimeValueParams) updateRecordValue(rec *record.Record) {
	(*rec).FieldValues[setValParams.FieldID] = setValParams.Value
}

type SetRecordFileValueParams struct {
	RecordUpdateHeader
	CloudFileName string `json:"cloudFileName"`
}

func (setValParams SetRecordFileValueParams) fieldType() string { return field.FieldTypeFile }

func (setValParams SetRecordFileValueParams) updateRecordValue(rec *record.Record) {
	(*rec).FieldValues[setValParams.FieldID] = setValParams.CloudFileName
}

// setRecordFileNameFieldValue. Although the parameters for a record update with a filename aren't passed through the http request,
// the standard record updating mechanism can be used to update the field with the filename.
func setRecordFileNameFieldValue(appEngContext appengine.Context,
	recordID string, fieldID string, fileName string) (*record.Record, error) {
	updateRecordHeader := RecordUpdateHeader{
		RecordID: recordID,
		FieldID:  fieldID}
	updateRecordParams := SetRecordFileValueParams{
		RecordUpdateHeader: updateRecordHeader,
		CloudFileName:      fileName}
	updatedRecord, updateErr := UpdateRecordValue(appEngContext, updateRecordParams)
	if updateErr != nil {
		return nil, fmt.Errorf("uploadFile: Unable to update record for newly uploaded file: %v", updateErr)
	}
	return updatedRecord, nil
}
