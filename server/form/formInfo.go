package form

import (
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

func getFormInfo(params GetFormInfoParams) (*FormInfo, error) {

	textBoxes, err := textBox.GetTextBoxes(params.FormID)
	if err != nil {
		return nil, err
	}

	checkBoxes, err := checkBox.GetCheckBoxes(params.FormID)
	if err != nil {
		return nil, err
	}

	datePickers, err := datePicker.GetDatePickers(params.FormID)
	if err != nil {
		return nil, err
	}

	htmlEditors, err := htmlEditor.GetHtmlEditors(params.FormID)
	if err != nil {
		return nil, err
	}

	images, err := image.GetImages(params.FormID)
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
