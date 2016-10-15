package header

import (
	"resultra/datasheet/webui/generic/propertiesSidebar"
)

type HeaderTemplateParams struct {
	ElemPrefix        string
	FormatPanelParams propertiesSidebar.PanelTemplateParams
	LabelPanelParams  propertiesSidebar.PanelTemplateParams
}

var TemplateParams HeaderTemplateParams

func init() {
	TemplateParams = HeaderTemplateParams{
		ElemPrefix:        "header_",
		FormatPanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Format", PanelID: "headerFormat"},
		LabelPanelParams:  propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Header Text", PanelID: "headerLabel"}}
}
