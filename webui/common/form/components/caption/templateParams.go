package caption

import (
	"resultra/datasheet/webui/common/form/components/common/visibility"
	"resultra/datasheet/webui/generic/propertiesSidebar"
)

type CaptionTemplateParams struct {
	ElemPrefix            string
	FormatPanelParams     propertiesSidebar.PanelTemplateParams
	CaptionPanelParams    propertiesSidebar.PanelTemplateParams
	VisibilityPanelParams visibility.VisibilityPropertyTemplateParams
}

var TemplateParams CaptionTemplateParams

func init() {

	elemPrefix := "caption_"

	TemplateParams = CaptionTemplateParams{
		ElemPrefix:            elemPrefix,
		VisibilityPanelParams: visibility.NewComponentVisibilityTempalteParams(elemPrefix, "captionVisibility"),
		FormatPanelParams:     propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Format", PanelID: "captionFormat"},
		CaptionPanelParams:    propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Caption", PanelID: "headerCaption"}}
}
