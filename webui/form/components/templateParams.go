package components

import (
	"resultra/datasheet/webui/form/components/checkBox"
	"resultra/datasheet/webui/form/components/datePicker"
	//	"resultra/datasheet/webui/form/components/htmlEditor"
	"resultra/datasheet/webui/form/components/image"
	"resultra/datasheet/webui/form/components/textBox"
)

type ComponentViewTemplateParams struct {
	CheckBoxParams   checkBox.CheckboxViewTemplateParams
	DatePickerParams datePicker.DatePickerViewTemplateParams
	TextBoxParams    textBox.TextboxViewTemplateParams
	ImageParams      image.ImageViewTemplateParams
}

//var DesignTemplateParams ComponentDesignTemplateParams
var ViewTemplateParams ComponentViewTemplateParams

func init() {

	ViewTemplateParams = ComponentViewTemplateParams{
		CheckBoxParams:   checkBox.ViewTemplateParams,
		DatePickerParams: datePicker.ViewTemplateParams,
		TextBoxParams:    textBox.ViewTemplateParams,
		ImageParams:      image.ViewTemplateParams}

}
