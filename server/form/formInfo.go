package form

import (
	"appengine"
	"resultra/datasheet/server/form/components/checkBox"
	"resultra/datasheet/server/form/components/datePicker"
	"resultra/datasheet/server/form/components/htmlEditor"
	"resultra/datasheet/server/form/components/image"
	"resultra/datasheet/server/form/components/textBox"
)

type FormInfo struct {
	TextBoxes   []textBox.TextBoxRef       `json:"textBoxes"`
	CheckBoxes  []checkBox.CheckBoxRef     `json:"checkBoxes"`
	DatePickers []datePicker.DatePickerRef `json:"datePickers"`
	HtmlEditors []htmlEditor.HtmlEditorRef `json:"htmlEditors"`
	Images      []image.ImageRef           `json:"images"`
}

type GetFormInfoParams struct {
	FormID string `json:"formID"`
}

func getFormInfo(appEngContext appengine.Context, params GetFormInfoParams) (*FormInfo, error) {

	textBoxRefs, err := textBox.GetTextBoxes(appEngContext, params.FormID)
	if err != nil {
		return nil, err
	}

	checkBoxRefs, err := checkBox.GetCheckBoxes(appEngContext, params.FormID)
	if err != nil {
		return nil, err
	}

	datePickerRefs, err := datePicker.GetDatePickers(appEngContext, params.FormID)
	if err != nil {
		return nil, err
	}

	htmlEditorRefs, err := htmlEditor.GetHtmlEditors(appEngContext, params.FormID)
	if err != nil {
		return nil, err
	}

	imageRefs, err := image.GetImages(appEngContext, params.FormID)
	if err != nil {
		return nil, err
	}

	formInfo := FormInfo{
		TextBoxes:   textBoxRefs,
		CheckBoxes:  checkBoxRefs,
		DatePickers: datePickerRefs,
		HtmlEditors: htmlEditorRefs,
		Images:      imageRefs}

	return &formInfo, nil
}
