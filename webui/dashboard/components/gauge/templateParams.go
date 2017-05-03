package gauge

import (
	"resultra/datasheet/webui/dashboard/components/common/valueSummary"
	"resultra/datasheet/webui/generic/propertiesSidebar"
)

type GaugeDesignTemplateParams struct {
	ElemPrefix         string
	TitlePanelParams   propertiesSidebar.PanelTemplateParams
	FormatPanelParams  propertiesSidebar.PanelTemplateParams
	ValueSummaryParams valueSummary.ValueSummaryTemplateParams
}

var DesignTemplateParams GaugeDesignTemplateParams

func init() {

	elemPrefix := "gauge_"

	valueSummaryParams := valueSummary.ValueSummaryTemplateParams{
		elemPrefix, "Configure how values are summarized.",
		"Field to summarize with", "Summarize values by"}

	DesignTemplateParams = GaugeDesignTemplateParams{
		ElemPrefix:         elemPrefix,
		ValueSummaryParams: valueSummaryParams,
		TitlePanelParams:   propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Title", PanelID: "gaugeTitle"},
		FormatPanelParams:  propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Format", PanelID: "gaugeFormat"}}

}
