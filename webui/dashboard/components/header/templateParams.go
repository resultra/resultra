package header

import (
	"resultra/datasheet/webui/generic/propertiesSidebar"
)

type HeaderDesignTemplateParams struct {
	ElemPrefix        string
	TitlePanelParams  propertiesSidebar.PanelTemplateParams
	FormatPanelParams propertiesSidebar.PanelTemplateParams
}

var DesignTemplateParams HeaderDesignTemplateParams

func init() {

	elemPrefix := "header_"

	DesignTemplateParams = HeaderDesignTemplateParams{
		ElemPrefix:        elemPrefix,
		TitlePanelParams:  propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Title", PanelID: "headerTitle"},
		FormatPanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Format", PanelID: "headerFormat"}}

}
