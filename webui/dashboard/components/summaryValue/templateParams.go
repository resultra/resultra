// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package summaryValue

import (
	"resultra/tracker/webui/admin/common/inputProperties"
	"resultra/tracker/webui/common/recordFilter"
	"resultra/tracker/webui/common/valueThreshold"
	"resultra/tracker/webui/dashboard/components/common/delete"
	"resultra/tracker/webui/dashboard/components/common/valueSummary"
	"resultra/tracker/webui/generic/propertiesSidebar"
)

type SummaryValDesignTemplateParams struct {
	ElemPrefix               string
	TitlePanelParams         propertiesSidebar.PanelTemplateParams
	FormatPanelParams        propertiesSidebar.PanelTemplateParams
	ValueSummaryParams       valueSummary.ValueSummaryTemplateParams
	ValueSummaryPanelParams  propertiesSidebar.PanelTemplateParams
	ThresholdPanelParams     propertiesSidebar.PanelTemplateParams
	ThresholdValueParams     valueThreshold.ThresholdValuesPanelTemplateParams
	FilteringPanelParams     propertiesSidebar.PanelTemplateParams
	FilterPropPanelParams    recordFilter.FilterPanelTemplateParams
	PreFilteringPanelParams  propertiesSidebar.PanelTemplateParams
	PreFilterPropPanelParams recordFilter.FilterPanelTemplateParams
	HelpPopupParams          inputProperties.HelpPopupPropertyTemplateParams
	DeletePanelParams        delete.DeletePropertyPanelTemplateParams
}

// Template parameters when in view mode
type SummaryValViewTemplateParams struct {
	ElemPrefix              string
	FilteringPanelParams    propertiesSidebar.PanelTemplateParams
	FilterConfigPanelParams recordFilter.FilterPanelTemplateParams
}

var DesignTemplateParams SummaryValDesignTemplateParams
var ViewTemplateParams SummaryValViewTemplateParams

func init() {

	elemPrefix := "summaryVal_"
	preFilterElemPrefix := "summaryValPreFilter_"

	valueSummaryParams := valueSummary.ValueSummaryTemplateParams{
		elemPrefix, "Configure how values are summarized.",
		"Field to summarize with", "Summarize values by"}

	DesignTemplateParams = SummaryValDesignTemplateParams{
		ElemPrefix:               elemPrefix,
		ValueSummaryParams:       valueSummaryParams,
		TitlePanelParams:         propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Title", PanelID: "summaryValTitle"},
		HelpPopupParams:          inputProperties.NewHelpPopupTemplateParams(elemPrefix, "summaryValHelp"),
		FormatPanelParams:        propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Format", PanelID: "summaryValFormat"},
		FilteringPanelParams:     propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Default Filtering", PanelID: "summaryValFiltering"},
		DeletePanelParams:        delete.NewDeletePropertyPanelTemplateParams(elemPrefix, "summaryValDelete", "Delete Summary Value"),
		ThresholdPanelParams:     propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Value Thresholds", PanelID: "summaryValThreshold"},
		ThresholdValueParams:     valueThreshold.NewThresholdValuesTemplateParams(elemPrefix),
		ValueSummaryPanelParams:  propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Value Summary", PanelID: "summaryValValSummary"},
		FilterPropPanelParams:    recordFilter.NewFilterPanelTemplateParams(elemPrefix),
		PreFilteringPanelParams:  propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Pre-Filtering", PanelID: "summaryValPreFiltering"},
		PreFilterPropPanelParams: recordFilter.NewFilterPanelTemplateParams(preFilterElemPrefix)}

	ViewTemplateParams = SummaryValViewTemplateParams{
		ElemPrefix:              elemPrefix,
		FilteringPanelParams:    propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Filtering", PanelID: "summaryValFiltering"},
		FilterConfigPanelParams: recordFilter.NewFilterPanelTemplateParams(elemPrefix)}

}
