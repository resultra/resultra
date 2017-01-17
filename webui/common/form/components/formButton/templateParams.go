package formButton

import (
	"resultra/datasheet/webui/generic/propertiesSidebar"
)

type ButtonTemplateParams struct {
	ElemPrefix            string
	FormatPanelParams     propertiesSidebar.PanelTemplateParams
	LinkedFormPanelParams propertiesSidebar.PanelTemplateParams
}

var TemplateParams ButtonTemplateParams

func init() {
	TemplateParams = ButtonTemplateParams{
		ElemPrefix:            "button_",
		FormatPanelParams:     propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Format", PanelID: "buttonFormat"},
		LinkedFormPanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Linked Form", PanelID: "buttonForm"}}
}
