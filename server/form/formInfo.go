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
	TextBoxes   []textBox.TextBox       `json:"textBoxes"`
	CheckBoxes  []checkBox.CheckBox     `json:"checkBoxes"`
	DatePickers []datePicker.DatePicker `json:"datePickers"`
	HtmlEditors []htmlEditor.HtmlEditor `json:"htmlEditors"`
	Images      []image.Image           `json:"images"`
}

type GetFormInfoParams struct {
	FormID string `json:"formID"`
}

func getFormInfo(appEngContext appengine.Context, params GetFormInfoParams) (*FormInfo, error) {

	textBoxes, err := textBox.GetTextBoxes(appEngContext, params.FormID)
	if err != nil {
		return nil, err
	}

	checkBoxes, err := checkBox.GetCheckBoxes(appEngContext, params.FormID)
	if err != nil {
		return nil, err
	}

	datePickers, err := datePicker.GetDatePickers(appEngContext, params.FormID)
	if err != nil {
		return nil, err
	}

	htmlEditors, err := htmlEditor.GetHtmlEditors(appEngContext, params.FormID)
	if err != nil {
		return nil, err
	}

	images, err := image.GetImages(appEngContext, params.FormID)
	if err != nil {
		return nil, err
	}

	formInfo := FormInfo{
		TextBoxes:   textBoxes,
		CheckBoxes:  checkBoxes,
		DatePickers: datePickers,
		HtmlEditors: htmlEditors,
		Images:      images}

	return &formInfo, nil
}
