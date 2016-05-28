package checkBox

import (
	"resultra/datasheet/webui/generic/propertiesSidebar"
)

type CheckboxTemplateParams struct {
	ElemPrefix        string
	FormatPanelParams propertiesSidebar.PanelTemplateParams
}

var TemplateParams CheckboxTemplateParams

func init() {
	TemplateParams = CheckboxTemplateParams{
		ElemPrefix:        "checkbox_",
		FormatPanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Format", PanelID: "checkboxFormat"}}
}
