package htmlEditor

import (
	"resultra/datasheet/webui/generic/propertiesSidebar"
)

type HTMLEditorTemplateParams struct {
	ElemPrefix        string
	FormatPanelParams propertiesSidebar.PanelTemplateParams
}

var TemplateParams HTMLEditorTemplateParams

func init() {
	TemplateParams = HTMLEditorTemplateParams{
		ElemPrefix:        "htmlEditor_",
		FormatPanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Format", PanelID: "htmlEditorFormat"}}
}
