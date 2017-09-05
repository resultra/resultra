package barChart

import (
	"resultra/datasheet/webui/admin/common/inputProperties"
	"resultra/datasheet/webui/common/recordFilter"
	"resultra/datasheet/webui/dashboard/components/common/newComponentDialog"
	"resultra/datasheet/webui/dashboard/components/common/valueSummary"
	"resultra/datasheet/webui/generic/propertiesSidebar"
)

type BarChartDesignTemplateParams struct {
	ElemPrefix               string
	ValueGroupingParams      newComponentDialog.ValueGroupingTemplateParams
	ValueSummaryParams       valueSummary.ValueSummaryTemplateParams
	TitlePanelParams         propertiesSidebar.PanelTemplateParams
	XAxisPanelParams         propertiesSidebar.PanelTemplateParams
	YAxisPanelParams         propertiesSidebar.PanelTemplateParams
	FilteringPanelParams     propertiesSidebar.PanelTemplateParams
	FilterPropPanelParams    recordFilter.FilterPanelTemplateParams
	PreFilteringPanelParams  propertiesSidebar.PanelTemplateParams
	PreFilterPropPanelParams recordFilter.FilterPanelTemplateParams
	HelpPopupParams          inputProperties.HelpPopupPropertyTemplateParams
}

// Template parameters when the summary table is in view mode
type BarChartViewTemplateParams struct {
	ElemPrefix              string
	FilteringPanelParams    propertiesSidebar.PanelTemplateParams
	FilterConfigPanelParams recordFilter.FilterPanelTemplateParams
}

var DesignTemplateParams BarChartDesignTemplateParams
var ViewTemplateParams BarChartViewTemplateParams

func init() {

	elemPrefix := "barChart_"
	preFilterElemPrefix := "barChartPreFilter_"

	valueGroupingParams := newComponentDialog.ValueGroupingTemplateParams{
		elemPrefix, "Configure which values are shown along the X axis and how these values are grouped",
		"Field for X axis' values", "Group Values By"}

	valueSummaryParams := valueSummary.ValueSummaryTemplateParams{
		elemPrefix, "Configure how values are summarized along the Y axis.",
		"Field to summarize with", "Summarize values by"}

	DesignTemplateParams = BarChartDesignTemplateParams{
		ElemPrefix:               elemPrefix,
		ValueGroupingParams:      valueGroupingParams,
		ValueSummaryParams:       valueSummaryParams,
		TitlePanelParams:         propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Title", PanelID: "barChartTitle"},
		HelpPopupParams:          inputProperties.NewHelpPopupTemplateParams(elemPrefix, "barChartHelp"),
		XAxisPanelParams:         propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "X Axis", PanelID: "barChartXAxis"},
		YAxisPanelParams:         propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Y Axis", PanelID: "barChartYAxis"},
		FilteringPanelParams:     propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Default Filtering", PanelID: "barChartFiltering"},
		FilterPropPanelParams:    recordFilter.NewFilterPanelTemplateParams(elemPrefix),
		PreFilteringPanelParams:  propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Pre-Filtering", PanelID: "barChartPreFiltering"},
		PreFilterPropPanelParams: recordFilter.NewFilterPanelTemplateParams(preFilterElemPrefix)}

	ViewTemplateParams = BarChartViewTemplateParams{
		ElemPrefix:              elemPrefix,
		FilteringPanelParams:    propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Filtering", PanelID: "barChartFiltering"},
		FilterConfigPanelParams: recordFilter.NewFilterPanelTemplateParams(elemPrefix)}

}
