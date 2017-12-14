package header

import (
	"resultra/datasheet/webui/dashboard/components/common/delete"
	"resultra/datasheet/webui/generic/propertiesSidebar"
)

type HeaderDesignTemplateParams struct {
	ElemPrefix        string
	TitlePanelParams  propertiesSidebar.PanelTemplateParams
	FormatPanelParams propertiesSidebar.PanelTemplateParams
	DeletePanelParams delete.DeletePropertyPanelTemplateParams
}

var DesignTemplateParams HeaderDesignTemplateParams

func init() {

	elemPrefix := "header_"

	DesignTemplateParams = HeaderDesignTemplateParams{
		ElemPrefix:        elemPrefix,
		TitlePanelParams:  propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Title", PanelID: "headerTitle"},
		DeletePanelParams: delete.NewDeletePropertyPanelTemplateParams(elemPrefix, "headerDelete", "Delete Header"),
		FormatPanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Format", PanelID: "headerFormat"}}

}
