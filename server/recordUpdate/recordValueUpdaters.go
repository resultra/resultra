package recordUpdate

import (
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
	(*rec)[setValParams.FieldID] = setValParams.Value
}

// Update a long text field value

type SetRecordLongTextValueParams struct {
	RecordUpdateHeader
	Value string `json:"value"`
}

func (setValParams SetRecordLongTextValueParams) fieldType() string { return field.FieldTypeLongText }

func (setValParams SetRecordLongTextValueParams) updateRecordValue(rec *record.Record) {
	(*rec)[setValParams.FieldID] = setValParams.Value
}

// Update a number field value

type SetRecordNumberValueParams struct {
	RecordUpdateHeader
	Value float64 `json:"value"`
}

func (setValParams SetRecordNumberValueParams) fieldType() string { return field.FieldTypeNumber }

func (setValParams SetRecordNumberValueParams) updateRecordValue(rec *record.Record) {
	(*rec)[setValParams.FieldID] = setValParams.Value
}

// Update a number field value

type SetRecordBoolValueParams struct {
	RecordUpdateHeader
	Value bool `json:"value"`
}

func (setValParams SetRecordBoolValueParams) fieldType() string { return field.FieldTypeBool }

func (setValParams SetRecordBoolValueParams) updateRecordValue(rec *record.Record) {
	(*rec)[setValParams.FieldID] = setValParams.Value
}

type SetRecordTimeValueParams struct {
	RecordUpdateHeader
	Value time.Time `json:"value"`
}

func (setValParams SetRecordTimeValueParams) fieldType() string { return field.FieldTypeTime }

func (setValParams SetRecordTimeValueParams) updateRecordValue(rec *record.Record) {
	(*rec)[setValParams.FieldID] = setValParams.Value
}
