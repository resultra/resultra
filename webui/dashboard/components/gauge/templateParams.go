package gauge

import (
	"resultra/datasheet/webui/common/recordFilter"
	"resultra/datasheet/webui/dashboard/components/common/valueSummary"
	"resultra/datasheet/webui/generic/propertiesSidebar"
)

type GaugeDesignTemplateParams struct {
	ElemPrefix               string
	TitlePanelParams         propertiesSidebar.PanelTemplateParams
	FormatPanelParams        propertiesSidebar.PanelTemplateParams
	RangePanelParams         propertiesSidebar.PanelTemplateParams
	ValueSummaryParams       valueSummary.ValueSummaryTemplateParams
	FilteringPanelParams     propertiesSidebar.PanelTemplateParams
	FilterPropPanelParams    recordFilter.FilterPanelTemplateParams
	PreFilteringPanelParams  propertiesSidebar.PanelTemplateParams
	PreFilterPropPanelParams recordFilter.FilterPanelTemplateParams
}

var DesignTemplateParams GaugeDesignTemplateParams

func init() {

	elemPrefix := "gauge_"
	preFilterElemPrefix := "gaugePreFilter_"

	valueSummaryParams := valueSummary.ValueSummaryTemplateParams{
		elemPrefix, "Configure how values are summarized.",
		"Field to summarize with", "Summarize values by"}

	DesignTemplateParams = GaugeDesignTemplateParams{
		ElemPrefix:               elemPrefix,
		ValueSummaryParams:       valueSummaryParams,
		TitlePanelParams:         propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Title", PanelID: "gaugeTitle"},
		FormatPanelParams:        propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Format", PanelID: "gaugeFormat"},
		RangePanelParams:         propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Value Range", PanelID: "gaugeRange"},
		FilteringPanelParams:     propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Default Filtering", PanelID: "gaugeFiltering"},
		FilterPropPanelParams:    recordFilter.NewFilterPanelTemplateParams(elemPrefix),
		PreFilteringPanelParams:  propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Pre-Filtering", PanelID: "gaugePreFiltering"},
		PreFilterPropPanelParams: recordFilter.NewFilterPanelTemplateParams(preFilterElemPrefix)}

}
