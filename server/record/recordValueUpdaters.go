package record

import (
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/generic"
	"time"
)

// Update a text field value

type SetRecordTextValueParams struct {
	RecordUpdateHeader
	Value string `json:"value"`
}

func (setValParams SetRecordTextValueParams) fieldType() string { return field.FieldTypeText }

func (setValParams SetRecordTextValueParams) updateRecordValue(rec *Record) {
	(*rec).FieldValues[setValParams.FieldID] = setValParams.Value
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

func (setValParams SetRecordLongTextValueParams) updateRecordValue(rec *Record) {
	(*rec).FieldValues[setValParams.FieldID] = setValParams.Value
}

func (valParams SetRecordLongTextValueParams) generateCellValue() (string, error) {

	cellVal := TextCellValue{Val: valParams.Value}

	return generic.EncodeJSONString(cellVal)
}

// Update a number field value

type SetRecordNumberValueParams struct {
	RecordUpdateHeader
	Value float64 `json:"value"`
}

func (setValParams SetRecordNumberValueParams) fieldType() string { return field.FieldTypeNumber }

func (setValParams SetRecordNumberValueParams) updateRecordValue(rec *Record) {
	(*rec).FieldValues[setValParams.FieldID] = setValParams.Value
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

func (setValParams SetRecordBoolValueParams) fieldType() string { return field.FieldTypeBool }

func (setValParams SetRecordBoolValueParams) updateRecordValue(rec *Record) {
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

func (setValParams SetRecordTimeValueParams) updateRecordValue(rec *Record) {
	(*rec).FieldValues[setValParams.FieldID] = setValParams.Value
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

func (setValParams SetRecordFileValueParams) updateRecordValue(rec *Record) {
	(*rec).FieldValues[setValParams.FieldID] = setValParams.CloudFileName
}

func (valParams SetRecordFileValueParams) generateCellValue() (string, error) {

	cellValue := FileCellValue{
		CloudName: valParams.CloudFileName,
		OrigName:  valParams.OrigFileName}

	return generic.EncodeJSONString(cellValue)
}
