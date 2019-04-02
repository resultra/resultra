// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package components

import (
	"github.com/resultra/resultra/webui/dashboard/components/barChart"
	"github.com/resultra/resultra/webui/dashboard/components/gauge"
	"github.com/resultra/resultra/webui/dashboard/components/header"
	"github.com/resultra/resultra/webui/dashboard/components/summaryTable"
	"github.com/resultra/resultra/webui/dashboard/components/summaryValue"
)

type ComponentDesignTemplateParams struct {
	BarChartParams     barChart.BarChartDesignTemplateParams
	SummaryTableParams summaryTable.SummaryTableDesignTemplateParams
	HeaderParams       header.HeaderDesignTemplateParams
	GaugeParams        gauge.GaugeDesignTemplateParams
	SummaryValParams   summaryValue.SummaryValDesignTemplateParams
}

type ComponentViewTemplateParams struct {
	SummaryTableParams summaryTable.SummaryTableViewTemplateParams
	BarChartParams     barChart.BarChartViewTemplateParams
	GaugeParams        gauge.GaugeViewTemplateParams
	SummaryValParams   summaryValue.SummaryValViewTemplateParams
}

var DesignTemplateParams ComponentDesignTemplateParams
var ViewTemplateParams ComponentViewTemplateParams

func init() {
	DesignTemplateParams = ComponentDesignTemplateParams{
		BarChartParams:     barChart.DesignTemplateParams,
		SummaryTableParams: summaryTable.DesignTemplateParams,
		HeaderParams:       header.DesignTemplateParams,
		GaugeParams:        gauge.DesignTemplateParams,
		SummaryValParams:   summaryValue.DesignTemplateParams}

	ViewTemplateParams = ComponentViewTemplateParams{
		SummaryTableParams: summaryTable.ViewTemplateParams,
		BarChartParams:     barChart.ViewTemplateParams,
		GaugeParams:        gauge.ViewTemplateParams,
		SummaryValParams:   summaryValue.ViewTemplateParams}
}
