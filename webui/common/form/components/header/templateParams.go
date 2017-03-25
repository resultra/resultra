package header

import (
	"resultra/datasheet/webui/common/form/components/common/delete"
	"resultra/datasheet/webui/common/form/components/common/visibility"
	"resultra/datasheet/webui/generic/propertiesSidebar"
)

type HeaderTemplateParams struct {
	ElemPrefix            string
	FormatPanelParams     propertiesSidebar.PanelTemplateParams
	LabelPanelParams      propertiesSidebar.PanelTemplateParams
	VisibilityPanelParams visibility.VisibilityPropertyTemplateParams
	DeletePanelParams     delete.DeletePropertyPanelTemplateParams
}

var TemplateParams HeaderTemplateParams

func init() {

	elemPrefix := "header_"

	TemplateParams = HeaderTemplateParams{
		ElemPrefix:            elemPrefix,
		FormatPanelParams:     propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Format", PanelID: "headerFormat"},
		DeletePanelParams:     delete.NewDeletePropertyPanelTemplateParams(elemPrefix, "headerDelete", "Delete Header"),
		VisibilityPanelParams: visibility.NewComponentVisibilityTempalteParams(elemPrefix, "headerVisibility"),
		LabelPanelParams:      propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Header Text", PanelID: "headerLabel"}}

}
