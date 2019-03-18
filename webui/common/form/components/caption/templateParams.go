package caption

import (
	"resultra/tracker/webui/common/form/components/common/delete"
	"resultra/tracker/webui/common/form/components/common/visibility"
	"resultra/tracker/webui/generic/propertiesSidebar"
)

type CaptionTemplateParams struct {
	ElemPrefix            string
	FormatPanelParams     propertiesSidebar.PanelTemplateParams
	CaptionPanelParams    propertiesSidebar.PanelTemplateParams
	VisibilityPanelParams visibility.VisibilityPropertyTemplateParams
	DeletePanelParams     delete.DeletePropertyPanelTemplateParams
}

var TemplateParams CaptionTemplateParams

func init() {

	elemPrefix := "caption_"

	TemplateParams = CaptionTemplateParams{
		ElemPrefix:            elemPrefix,
		VisibilityPanelParams: visibility.NewComponentVisibilityTempalteParams(elemPrefix, "captionVisibility"),
		DeletePanelParams:     delete.NewDeletePropertyPanelTemplateParams(elemPrefix, "captionDelete", "Delete Caption"),
		FormatPanelParams:     propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Format", PanelID: "captionFormat"},
		CaptionPanelParams:    propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Caption", PanelID: "headerCaption"}}
}
