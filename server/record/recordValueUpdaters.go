package record

import (
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/generic"
	"time"
)

// Update a text field value

type SetRecordTextValueParams struct {
	RecordUpdateHeader
	ValueFormat CellUpdateValueFormat `json:"valueFormat"`
	Value       *string               `json:"value"`
}

func (setValParams SetRecordTextValueParams) fieldType() string { return field.FieldTypeText }

func (setValParams SetRecordTextValueParams) doCollapseRecentValues() bool { return true }

func (valParams SetRecordTextValueParams) generateCellValue() (string, error) {

	cellVal := TextCellValue{Val: valParams.Value}

	return generic.EncodeJSONString(cellVal)
}

func (valParams SetRecordTextValueParams) getUpdateProperties() CellUpdateProperties {
	props := CellUpdateProperties{valParams.ValueFormat}

	return props
}

// Update a user field value

type SetRecordUserValueParams struct {
	RecordUpdateHeader
	ValueFormat CellUpdateValueFormat `json:"valueFormat"`
	UserID      *string               `json:"userID"`
}

func (setValParams SetRecordUserValueParams) fieldType() string { return field.FieldTypeUser }

func (setValParams SetRecordUserValueParams) doCollapseRecentValues() bool { return true }

func (valParams SetRecordUserValueParams) generateCellValue() (string, error) {

	cellVal := UserCellValue{UserID: valParams.UserID}

	return generic.EncodeJSONString(cellVal)
}

func (valParams SetRecordUserValueParams) getUpdateProperties() CellUpdateProperties {
	props := CellUpdateProperties{valParams.ValueFormat}

	return props
}

// Update a long text field value

type SetRecordLongTextValueParams struct {
	RecordUpdateHeader
	ValueFormat CellUpdateValueFormat `json:"valueFormat"`
	Value       *string               `json:"value"`
}

func (setValParams SetRecordLongTextValueParams) fieldType() string { return field.FieldTypeLongText }

func (setValParams SetRecordLongTextValueParams) doCollapseRecentValues() bool { return true }

func (valParams SetRecordLongTextValueParams) generateCellValue() (string, error) {

	cellVal := TextCellValue{Val: valParams.Value}

	return generic.EncodeJSONString(cellVal)
}

func (valParams SetRecordLongTextValueParams) getUpdateProperties() CellUpdateProperties {
	props := CellUpdateProperties{valParams.ValueFormat}

	return props
}

// Update a comment field value

type SetRecordCommentValueParams struct {
	RecordUpdateHeader
	ValueFormat CellUpdateValueFormat `json:"valueFormat"`
	CommentText string                `json:"commentText"`
	Attachments []string              `json:"attachments"`
}

func (setValParams SetRecordCommentValueParams) fieldType() string { return field.FieldTypeComment }

func (setValParams SetRecordCommentValueParams) doCollapseRecentValues() bool { return false }

func (valParams SetRecordCommentValueParams) generateCellValue() (string, error) {

	cellVal := CommentCellValue{
		CommentText: valParams.CommentText,
		Attachments: valParams.Attachments}

	return generic.EncodeJSONString(cellVal)
}

func (valParams SetRecordCommentValueParams) getUpdateProperties() CellUpdateProperties {
	props := CellUpdateProperties{valParams.ValueFormat}

	return props
}

// Update a number field value

type SetRecordNumberValueParams struct {
	RecordUpdateHeader
	ValueFormat CellUpdateValueFormat `json:"valueFormat"`
	Value       *float64              `json:"value"`
}

func (setValParams SetRecordNumberValueParams) fieldType() string { return field.FieldTypeNumber }

func (setValParams SetRecordNumberValueParams) doCollapseRecentValues() bool { return true }

func (valParams SetRecordNumberValueParams) generateCellValue() (string, error) {

	cellVal := NumberCellValue{Val: valParams.Value}

	return generic.EncodeJSONString(cellVal)
}

func (valParams SetRecordNumberValueParams) getUpdateProperties() CellUpdateProperties {

	props := CellUpdateProperties{valParams.ValueFormat}

	return props
}

// Update a number field value

type SetRecordBoolValueParams struct {
	RecordUpdateHeader
	ValueFormat CellUpdateValueFormat `json:"valueFormat"`
	Value       *bool                 `json:"value"`
}

func (setValParams SetRecordBoolValueParams) fieldType() string { return field.FieldTypeBool }

func (setValParams SetRecordBoolValueParams) doCollapseRecentValues() bool { return true }

func (valParams SetRecordBoolValueParams) generateCellValue() (string, error) {

	cellVal := BoolCellValue{Val: valParams.Value}

	return generic.EncodeJSONString(cellVal)
}

func (valParams SetRecordBoolValueParams) getUpdateProperties() CellUpdateProperties {

	props := CellUpdateProperties{valParams.ValueFormat}

	return props
}

type SetRecordTimeValueParams struct {
	RecordUpdateHeader
	ValueFormat CellUpdateValueFormat `json:"valueFormat"`
	Value       *time.Time            `json:"value"`
}

func (setValParams SetRecordTimeValueParams) fieldType() string { return field.FieldTypeTime }

func (setValParams SetRecordTimeValueParams) doCollapseRecentValues() bool { return true }

func (valParams SetRecordTimeValueParams) generateCellValue() (string, error) {

	cellVal := TimeCellValue{Val: valParams.Value}

	return generic.EncodeJSONString(cellVal)
}

func (valParams SetRecordTimeValueParams) getUpdateProperties() CellUpdateProperties {
	props := CellUpdateProperties{valParams.ValueFormat}

	return props
}

type SetRecordFileValueParams struct {
	RecordUpdateHeader
	ValueFormat CellUpdateValueFormat `json:"valueFormat"`
	Attachments []string              `json:"attachments"`
}

func (setValParams SetRecordFileValueParams) fieldType() string { return field.FieldTypeFile }

func (setValParams SetRecordFileValueParams) doCollapseRecentValues() bool { return true }

func (valParams SetRecordFileValueParams) generateCellValue() (string, error) {

	cellValue := FileCellValue{Attachments: valParams.Attachments}

	return generic.EncodeJSONString(cellValue)
}

func (valParams SetRecordFileValueParams) getUpdateProperties() CellUpdateProperties {
	props := CellUpdateProperties{valParams.ValueFormat}

	return props
}
