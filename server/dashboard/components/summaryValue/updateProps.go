package summaryValue

import (
	"fmt"
	"log"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/dashboard/values"
	"resultra/datasheet/server/generic/threshold"
	"resultra/datasheet/server/recordFilter"
)

// The BarChartPropertyUpdater interface along with UpdateBarChartProps() implement a harness for
// property updates. All property updates consiste of: (1) Retrieve the entity from the datastore,
// (2) Do the update on the entity itself, (3) Save the updated entity back to the datastore.
// Steps (1) and (3) can be done in a wrapper function UpdateBarChartProps, while only (2) needs
// be defined for each different property update. The goal is to minimize code bloat of property
// setting code and also make property updating code more uniform and less error prone.
type SummaryValPropertyUpdater interface {
	uniqueSummaryValID() string
	parentDashboardID() string
	updateSummaryValProps(summaryVal *SummaryVal) error
}

type SummaryValUniqueIDSummaryVal struct {
	ParentDashboardID string `json:"parentDashboardID"`
	SummaryValID      string `json:"summaryValID"`
}

func (idSummaryVal SummaryValUniqueIDSummaryVal) parentDashboardID() string {
	return idSummaryVal.ParentDashboardID
}

func (idSummaryVal SummaryValUniqueIDSummaryVal) uniqueSummaryValID() string {
	return idSummaryVal.SummaryValID
}

func updateSummaryValProps(propUpdater SummaryValPropertyUpdater) (*SummaryVal, error) {

	// Retrieve the bar chart from the data store
	summaryValForUpdate, getErr := GetSummaryVal(propUpdater.parentDashboardID(), propUpdater.uniqueSummaryValID())
	if getErr != nil {
		return nil, fmt.Errorf("updateSummaryValProps: Unable to get existing summaryVal: %v", getErr)
	}

	// Do the actual update
	propUpdateErr := propUpdater.updateSummaryValProps(summaryValForUpdate)
	if propUpdateErr != nil {
		return nil, fmt.Errorf("updateSummaryValProps: Unable to update existing summaryVal: %v", propUpdateErr)
	}

	// Save the updated bar chart back to the data store
	updatedSummaryVal, updateErr := updateExistingSummaryVal(summaryValForUpdate)
	if updateErr != nil {
		return nil, fmt.Errorf("updateSummaryValProps: Unable to update existing summaryVal: %v", updateErr)
	}

	return updatedSummaryVal, nil

}

// Title Property

type SetSummaryValTitleParams struct {
	// Embed a common summaryVal to reference the BarChart in the datastore. This summaryVal also supports
	// the niqueBarChartID() method to retrieve the unique ID. So, once decoded, the struct can be passed as an
	// BarChartPropertyUpdater interface to a generic/reusable function to process the property update.
	SummaryValUniqueIDSummaryVal
	NewTitle string `json:"newTitle"`
}

func (titleParam SetSummaryValTitleParams) updateSummaryValProps(summaryVal *SummaryVal) error {

	log.Printf("Updating summaryVal title: %v", titleParam.NewTitle)

	summaryVal.Properties.Title = titleParam.NewTitle

	return nil
}

// Dimensions Property

type SetSummaryValDimensionsParams struct {
	SummaryValUniqueIDSummaryVal
	Geometry componentLayout.LayoutGeometry `json:"geometry"`
}

func (params SetSummaryValDimensionsParams) updateSummaryValProps(summaryVal *SummaryVal) error {

	if !componentLayout.ValidGeometry(params.Geometry) {
		return fmt.Errorf("setBarChartDimensions: Invalid geometry for bar chart: %+v", params.Geometry)
	}

	summaryVal.Properties.Geometry = params.Geometry

	return nil
}

type SetValSummaryParams struct {
	SummaryValUniqueIDSummaryVal
	ValSummary values.ValSummary `json:"valSummary"`
}

func (params SetValSummaryParams) updateSummaryValProps(summaryVal *SummaryVal) error {

	summaryVal.Properties.ValSummary = params.ValSummary

	return nil
}

type SetDefaultFilterRulesParams struct {
	SummaryValUniqueIDSummaryVal
	DefaultFilterRules recordFilter.RecordFilterRuleSet `json:"defaultFilterRules"`
}

func (params SetDefaultFilterRulesParams) updateSummaryValProps(summaryVal *SummaryVal) error {

	summaryVal.Properties.DefaultFilterRules = params.DefaultFilterRules

	return nil
}

type SetPreFilterRulesParams struct {
	SummaryValUniqueIDSummaryVal
	PreFilterRules recordFilter.RecordFilterRuleSet `json:"preFilterRules"`
}

func (params SetPreFilterRulesParams) updateSummaryValProps(summaryVal *SummaryVal) error {

	summaryVal.Properties.PreFilterRules = params.PreFilterRules

	return nil
}

type SetThresholdsParams struct {
	SummaryValUniqueIDSummaryVal
	ThresholdVals []threshold.ThresholdValues `json:"thresholdVals"`
}

func (updateParams SetThresholdsParams) updateSummaryValProps(summaryVal *SummaryVal) error {

	summaryVal.Properties.ThresholdVals = updateParams.ThresholdVals

	return nil
}
