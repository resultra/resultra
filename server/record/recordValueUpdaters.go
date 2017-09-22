package record

import (
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/generic"
	"time"
)

// Update a text field value

type SetRecordTextValueParams struct {
	RecordUpdateHeader
	Value *string `json:"value"`
}

func (setValParams SetRecordTextValueParams) fieldType() string { return field.FieldTypeText }

func (setValParams SetRecordTextValueParams) doCollapseRecentValues() bool { return true }

func (valParams SetRecordTextValueParams) generateCellValue() (string, error) {

	cellVal := TextCellValue{Val: valParams.Value}

	return generic.EncodeJSONString(cellVal)
}

// Update a user field value

type SetRecordUserValueParams struct {
	RecordUpdateHeader
	UserIDs []string `json:"userIDs"`
}

func (setValParams SetRecordUserValueParams) fieldType() string { return field.FieldTypeUser }

func (setValParams SetRecordUserValueParams) doCollapseRecentValues() bool { return true }

func (valParams SetRecordUserValueParams) generateCellValue() (string, error) {

	cellVal := UserCellValue{UserIDs: valParams.UserIDs}

	return generic.EncodeJSONString(cellVal)
}

// Update a long text field value

type SetRecordLongTextValueParams struct {
	RecordUpdateHeader
	Value *string `json:"value"`
}

func (setValParams SetRecordLongTextValueParams) fieldType() string { return field.FieldTypeLongText }

func (setValParams SetRecordLongTextValueParams) doCollapseRecentValues() bool { return true }

func (valParams SetRecordLongTextValueParams) generateCellValue() (string, error) {

	cellVal := TextCellValue{Val: valParams.Value}

	return generic.EncodeJSONString(cellVal)
}

// Update a comment field value

type SetRecordCommentValueParams struct {
	RecordUpdateHeader
	CommentText string   `json:"commentText"`
	Attachments []string `json:"attachments"`
}

func (setValParams SetRecordCommentValueParams) fieldType() string { return field.FieldTypeComment }

func (setValParams SetRecordCommentValueParams) doCollapseRecentValues() bool { return false }

func (valParams SetRecordCommentValueParams) generateCellValue() (string, error) {

	cellVal := CommentCellValue{
		CommentText: valParams.CommentText,
		Attachments: valParams.Attachments}

	return generic.EncodeJSONString(cellVal)
}

// Update a number field value

type SetRecordNumberValueParams struct {
	RecordUpdateHeader
	Value *float64 `json:"value"`
}

func (setValParams SetRecordNumberValueParams) fieldType() string { return field.FieldTypeNumber }

func (setValParams SetRecordNumberValueParams) doCollapseRecentValues() bool { return true }

func (valParams SetRecordNumberValueParams) generateCellValue() (string, error) {

	cellVal := NumberCellValue{Val: valParams.Value}

	return generic.EncodeJSONString(cellVal)
}

// Update a number field value

type SetRecordBoolValueParams struct {
	RecordUpdateHeader
	Value *bool `json:"value"`
}

func (setValParams SetRecordBoolValueParams) fieldType() string { return field.FieldTypeBool }

func (setValParams SetRecordBoolValueParams) doCollapseRecentValues() bool { return true }

func (valParams SetRecordBoolValueParams) generateCellValue() (string, error) {

	cellVal := BoolCellValue{Val: valParams.Value}

	return generic.EncodeJSONString(cellVal)
}

type SetRecordTimeValueParams struct {
	RecordUpdateHeader
	Value *time.Time `json:"value"`
}

func (setValParams SetRecordTimeValueParams) fieldType() string { return field.FieldTypeTime }

func (setValParams SetRecordTimeValueParams) doCollapseRecentValues() bool { return true }

func (valParams SetRecordTimeValueParams) generateCellValue() (string, error) {

	cellVal := TimeCellValue{Val: valParams.Value}

	return generic.EncodeJSONString(cellVal)
}

type SetRecordAttachmentValueParams struct {
	RecordUpdateHeader
	Attachments []string `json:"attachments"`
}

func (setValParams SetRecordAttachmentValueParams) fieldType() string {
	return field.FieldTypeAttachment
}

func (setValParams SetRecordAttachmentValueParams) doCollapseRecentValues() bool { return true }

func (valParams SetRecordAttachmentValueParams) generateCellValue() (string, error) {

	cellValue := AttachmentCellValue{Attachments: valParams.Attachments}

	return generic.EncodeJSONString(cellValue)
}

type SetRecordLabelValueParams struct {
	RecordUpdateHeader
	Labels []string `json:"labels"`
}

func (setValParams SetRecordLabelValueParams) fieldType() string { return field.FieldTypeLabel }

func (setValParams SetRecordLabelValueParams) doCollapseRecentValues() bool { return true }

func (valParams SetRecordLabelValueParams) generateCellValue() (string, error) {

	cellValue := LabelCellValue{Labels: valParams.Labels}

	return generic.EncodeJSONString(cellValue)
}

type SetRecordEmailAddrValueParams struct {
	RecordUpdateHeader
	Value *string `json:"value"`
}

func (setValParams SetRecordEmailAddrValueParams) fieldType() string { return field.FieldTypeEmail }

func (setValParams SetRecordEmailAddrValueParams) doCollapseRecentValues() bool { return true }

func (valParams SetRecordEmailAddrValueParams) generateCellValue() (string, error) {

	cellVal := EmailAddrCellValue{Val: valParams.Value}

	return generic.EncodeJSONString(cellVal)
}

type SetRecordUrlValueParams struct {
	RecordUpdateHeader
	Value *string `json:"value"`
}

func (setValParams SetRecordUrlValueParams) fieldType() string { return field.FieldTypeURL }

func (setValParams SetRecordUrlValueParams) doCollapseRecentValues() bool { return true }

func (valParams SetRecordUrlValueParams) generateCellValue() (string, error) {

	cellVal := UrlCellValue{Val: valParams.Value}

	return generic.EncodeJSONString(cellVal)
}

type SetRecordFileAddrValueParams struct {
	RecordUpdateHeader
	Attachment *string `json:"attachment"`
}

func (setValParams SetRecordFileAddrValueParams) fieldType() string { return field.FieldTypeFile }

func (setValParams SetRecordFileAddrValueParams) doCollapseRecentValues() bool { return true }

func (valParams SetRecordFileAddrValueParams) generateCellValue() (string, error) {

	cellVal := FileCellValue{Attachment: valParams.Attachment}

	return generic.EncodeJSONString(cellVal)
}

type SetRecordImageValueParams struct {
	RecordUpdateHeader
	Attachment *string `json:"attachment"`
}

func (setValParams SetRecordImageValueParams) fieldType() string { return field.FieldTypeImage }

func (setValParams SetRecordImageValueParams) doCollapseRecentValues() bool { return true }

func (valParams SetRecordImageValueParams) generateCellValue() (string, error) {

	cellVal := ImageCellValue{Attachment: valParams.Attachment}

	return generic.EncodeJSONString(cellVal)
}
