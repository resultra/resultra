// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package barChart

import (
	"resultra/tracker/webui/admin/common/inputProperties"
	"resultra/tracker/webui/common/recordFilter"
	"resultra/tracker/webui/dashboard/components/common/delete"
	"resultra/tracker/webui/dashboard/components/common/newComponentDialog"
	"resultra/tracker/webui/dashboard/components/common/valueSummary"
	"resultra/tracker/webui/generic/propertiesSidebar"
)

type BarChartDesignTemplateParams struct {
	ElemPrefix               string
	ValueGroupingParams      newComponentDialog.ValueGroupingTemplateParams
	NewValueGroupingParams   newComponentDialog.ValueGroupingTemplateParams
	ValueSummaryParams       valueSummary.ValueSummaryTemplateParams
	TitlePanelParams         propertiesSidebar.PanelTemplateParams
	XAxisPanelParams         propertiesSidebar.PanelTemplateParams
	YAxisPanelParams         propertiesSidebar.PanelTemplateParams
	FilteringPanelParams     propertiesSidebar.PanelTemplateParams
	FilterPropPanelParams    recordFilter.FilterPanelTemplateParams
	PreFilteringPanelParams  propertiesSidebar.PanelTemplateParams
	PreFilterPropPanelParams recordFilter.FilterPanelTemplateParams
	HelpPopupParams          inputProperties.HelpPopupPropertyTemplateParams
	DeletePanelParams        delete.DeletePropertyPanelTemplateParams
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
		"Field or time increment for X axis' values", "Group Values By"}

	newBarChartElemPrefix := "newBarChart_"
	newValueGroupingParams := newComponentDialog.ValueGroupingTemplateParams{
		newBarChartElemPrefix, "Configure which values are shown along the X axis and how these values are grouped",
		"Field or time increment for X axis' values", "Group Values By"}

	valueSummaryParams := valueSummary.ValueSummaryTemplateParams{
		elemPrefix, "Configure how values are summarized along the Y axis.",
		"Field to summarize with", "Summarize values by"}

	DesignTemplateParams = BarChartDesignTemplateParams{
		ElemPrefix:               elemPrefix,
		ValueGroupingParams:      valueGroupingParams,
		NewValueGroupingParams:   newValueGroupingParams,
		ValueSummaryParams:       valueSummaryParams,
		TitlePanelParams:         propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Title", PanelID: "barChartTitle"},
		HelpPopupParams:          inputProperties.NewHelpPopupTemplateParams(elemPrefix, "barChartHelp"),
		XAxisPanelParams:         propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "X Axis", PanelID: "barChartXAxis"},
		YAxisPanelParams:         propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Y Axis", PanelID: "barChartYAxis"},
		DeletePanelParams:        delete.NewDeletePropertyPanelTemplateParams(elemPrefix, "barChartDelete", "Delete Bar Chart"),
		FilteringPanelParams:     propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Default Filtering", PanelID: "barChartFiltering"},
		FilterPropPanelParams:    recordFilter.NewFilterPanelTemplateParams(elemPrefix),
		PreFilteringPanelParams:  propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Pre-Filtering", PanelID: "barChartPreFiltering"},
		PreFilterPropPanelParams: recordFilter.NewFilterPanelTemplateParams(preFilterElemPrefix)}

	ViewTemplateParams = BarChartViewTemplateParams{
		ElemPrefix:              elemPrefix,
		FilteringPanelParams:    propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Filtering", PanelID: "barChartFiltering"},
		FilterConfigPanelParams: recordFilter.NewFilterPanelTemplateParams(elemPrefix)}

}
