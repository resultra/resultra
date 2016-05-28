package datePicker

import (
	"resultra/datasheet/webui/generic/propertiesSidebar"
)

type DatePickerTemplateParams struct {
	ElemPrefix        string
	FormatPanelParams propertiesSidebar.PanelTemplateParams
}

var TemplateParams DatePickerTemplateParams

func init() {
	TemplateParams = DatePickerTemplateParams{
		ElemPrefix:        "datePicker_",
		FormatPanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Format", PanelID: "datePickerFormat"}}
}
