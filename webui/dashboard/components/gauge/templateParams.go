// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package gauge

import (
	"github.com/resultra/resultra/webui/admin/common/inputProperties"
	"github.com/resultra/resultra/webui/common/recordFilter"
	"github.com/resultra/resultra/webui/common/valueThreshold"
	"github.com/resultra/resultra/webui/dashboard/components/common/delete"
	"github.com/resultra/resultra/webui/dashboard/components/common/valueSummary"
	"github.com/resultra/resultra/webui/generic/propertiesSidebar"
)

type GaugeDesignTemplateParams struct {
	ElemPrefix               string
	TitlePanelParams         propertiesSidebar.PanelTemplateParams
	FormatPanelParams        propertiesSidebar.PanelTemplateParams
	RangePanelParams         propertiesSidebar.PanelTemplateParams
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
type GaugeViewTemplateParams struct {
	ElemPrefix              string
	FilteringPanelParams    propertiesSidebar.PanelTemplateParams
	FilterConfigPanelParams recordFilter.FilterPanelTemplateParams
}

var DesignTemplateParams GaugeDesignTemplateParams
var ViewTemplateParams GaugeViewTemplateParams

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
		HelpPopupParams:          inputProperties.NewHelpPopupTemplateParams(elemPrefix, "gaugeHelp"),
		FormatPanelParams:        propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Format", PanelID: "gaugeFormat"},
		RangePanelParams:         propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Value Range", PanelID: "gaugeRange"},
		FilteringPanelParams:     propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Default Filtering", PanelID: "gaugeFiltering"},
		ThresholdPanelParams:     propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Value Thresholds", PanelID: "gaugeThreshold"},
		DeletePanelParams:        delete.NewDeletePropertyPanelTemplateParams(elemPrefix, "gaugeDelete", "Delete Gauge"),
		ThresholdValueParams:     valueThreshold.NewThresholdValuesTemplateParams(elemPrefix),
		ValueSummaryPanelParams:  propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Value Summary", PanelID: "gaugeValSummary"},
		FilterPropPanelParams:    recordFilter.NewFilterPanelTemplateParams(elemPrefix),
		PreFilteringPanelParams:  propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Pre-Filtering", PanelID: "gaugePreFiltering"},
		PreFilterPropPanelParams: recordFilter.NewFilterPanelTemplateParams(preFilterElemPrefix)}

	ViewTemplateParams = GaugeViewTemplateParams{
		ElemPrefix:              elemPrefix,
		FilteringPanelParams:    propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Filtering", PanelID: "dashboardGaugeFiltering"},
		FilterConfigPanelParams: recordFilter.NewFilterPanelTemplateParams(elemPrefix)}

}
