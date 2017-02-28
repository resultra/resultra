package caption

import (
	"resultra/datasheet/webui/generic/propertiesSidebar"
)

type CaptionTemplateParams struct {
	ElemPrefix         string
	FormatPanelParams  propertiesSidebar.PanelTemplateParams
	CaptionPanelParams propertiesSidebar.PanelTemplateParams
}

var TemplateParams CaptionTemplateParams

func init() {
	TemplateParams = CaptionTemplateParams{
		ElemPrefix:         "caption_",
		FormatPanelParams:  propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Format", PanelID: "captionFormat"},
		CaptionPanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Caption", PanelID: "headerCaption"}}
}
