package textBox

import (
	"resultra/datasheet/webui/generic/propertiesSidebar"
)

type TextboxTemplateParams struct {
	ElemPrefix        string
	FormatPanelParams propertiesSidebar.PanelTemplateParams
}

var TemplateParams TextboxTemplateParams

func init() {
	TemplateParams = TextboxTemplateParams{
		ElemPrefix:        "textBox_",
		FormatPanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Format", PanelID: "textboxFormat"}}
}
