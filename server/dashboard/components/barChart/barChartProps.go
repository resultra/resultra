// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package barChart

import (
	"database/sql"
	"fmt"
	"github.com/resultra/resultra/server/common/componentLayout"
	"github.com/resultra/resultra/server/dashboard/values"
	"github.com/resultra/resultra/server/recordFilter"
	"log"
)

// The BarChartPropertyUpdater interface along with UpdateBarChartProps() implement a harness for
// property updates. All property updates consiste of: (1) Retrieve the entity from the datastore,
// (2) Do the update on the entity itself, (3) Save the updated entity back to the datastore.
// Steps (1) and (3) can be done in a wrapper function UpdateBarChartProps, while only (2) needs
// be defined for each different property update. The goal is to minimize code bloat of property
// setting code and also make property updating code more uniform and less error prone.
type BarChartPropertyUpdater interface {
	uniqueBarChartID() string
	parentDashboardID() string
	updateBarChartProps(barChart *BarChart) error
}

type BarChartUniqueIDHeader struct {
	ParentDashboardID string `json:"parentDashboardID"`
	BarChartID        string `json:"barChartID"`
}

func (idHeader BarChartUniqueIDHeader) parentDashboardID() string {
	return idHeader.ParentDashboardID
}

func (idHeader BarChartUniqueIDHeader) uniqueBarChartID() string {
	return idHeader.BarChartID
}

func UpdateBarChartProps(trackerDBHandle *sql.DB, propUpdater BarChartPropertyUpdater) (*BarChart, error) {

	// Retrieve the bar chart from the data store
	barChartForUpdate, getBarChartErr := GetBarChart(trackerDBHandle, propUpdater.parentDashboardID(), propUpdater.uniqueBarChartID())
	if getBarChartErr != nil {
		return nil, fmt.Errorf("updateBarChartProps: Unable to get existing bar chart: %v", getBarChartErr)
	}

	// Do the actual update
	propUpdateErr := propUpdater.updateBarChartProps(barChartForUpdate)
	if propUpdateErr != nil {
		return nil, fmt.Errorf("updateBarChartProps: Unable to update existing bar chart: %v", propUpdateErr)
	}

	// Save the updated bar chart back to the data store
	updatedBarChart, updateErr := updateExistingBarChart(trackerDBHandle, barChartForUpdate)
	if updateErr != nil {
		return nil, fmt.Errorf("updateBarChartProps: Unable to update existing bar chart: %v", updateErr)
	}

	return updatedBarChart, nil

}

// Title Property

type SetBarChartTitleParams struct {
	// Embed a common header to reference the BarChart in the datastore. This header also supports
	// the niqueBarChartID() method to retrieve the unique ID. So, once decoded, the struct can be passed as an
	// BarChartPropertyUpdater interface to a generic/reusable function to process the property update.
	BarChartUniqueIDHeader
	NewTitle string `json:"newTitle"`
}

func (titleParam SetBarChartTitleParams) updateBarChartProps(barChart *BarChart) error {

	log.Printf("Updating bar chart title: %v", titleParam.NewTitle)

	barChart.Properties.Title = titleParam.NewTitle

	return nil
}

// Dimensions Property

type SetBarChartDimensionsParams struct {
	BarChartUniqueIDHeader
	Geometry componentLayout.LayoutGeometry `json:"geometry"`
}

func (params SetBarChartDimensionsParams) updateBarChartProps(barChart *BarChart) error {

	if !componentLayout.ValidGeometry(params.Geometry) {
		return fmt.Errorf("setBarChartDimensions: Invalid geometry for bar chart: %+v", params.Geometry)
	}

	barChart.Properties.Geometry = params.Geometry

	return nil
}

type SetXAxisValuesParams struct {
	BarChartUniqueIDHeader
	XAxisValueGrouping values.ValGrouping `json:"xAxisValueGrouping"`
}

func (params SetXAxisValuesParams) updateBarChartProps(barChart *BarChart) error {

	barChart.Properties.XAxisVals = params.XAxisValueGrouping

	return nil
}

type SetYAxisSummaryParams struct {
	BarChartUniqueIDHeader
	YAxisValSummary values.ValSummary `json:"yAxisValSummary"`
}

func (params SetYAxisSummaryParams) updateBarChartProps(barChart *BarChart) error {

	barChart.Properties.YAxisVals = params.YAxisValSummary

	return nil
}

type SetDefaultFilterRulesParams struct {
	BarChartUniqueIDHeader
	DefaultFilterRules recordFilter.RecordFilterRuleSet `json:"defaultFilterRules"`
}

func (params SetDefaultFilterRulesParams) updateBarChartProps(barChart *BarChart) error {

	barChart.Properties.DefaultFilterRules = params.DefaultFilterRules

	return nil
}

type SetPreFilterRulesParams struct {
	BarChartUniqueIDHeader
	PreFilterRules recordFilter.RecordFilterRuleSet `json:"preFilterRules"`
}

func (params SetPreFilterRulesParams) updateBarChartProps(barChart *BarChart) error {

	barChart.Properties.PreFilterRules = params.PreFilterRules

	return nil
}

type SetHelpPopupMsgParams struct {
	BarChartUniqueIDHeader
	PopupMsg string `json:"popupMsg"`
}

func (params SetHelpPopupMsgParams) updateBarChartProps(barChart *BarChart) error {

	barChart.Properties.HelpPopupMsg = params.PopupMsg

	return nil
}
