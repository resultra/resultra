package summaryValue

import (
	"fmt"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/dashboard/components/common"
	"resultra/datasheet/server/dashboard/values"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/uniqueID"
)

const summaryValEntityKind string = "summaryVal"

// DashboardBarChart is the datastore object for dashboard bar charts.
type SummaryVal struct {
	ParentDashboardID string `json:"parentDashboardID"`

	SummaryValID string `json:"summaryValID"`

	// DataSrcTable is the table the bar chart gets its data from
	Properties SummaryValProps `json:"properties"`
}

type NewSummaryValParams struct {
	ParentDashboardID string `json:"parentDashboardID"`

	ValSummary values.NewValSummaryParams `json:"valSummary"`

	Geometry componentLayout.LayoutGeometry `json:"geometry"`
}

func saveSummaryVal(newSummaryVal SummaryVal) error {

	if saveErr := common.SaveNewDashboardComponent(summaryValEntityKind,
		newSummaryVal.ParentDashboardID, newSummaryVal.SummaryValID, newSummaryVal.Properties); saveErr != nil {
		return fmt.Errorf("newSummaryVal: Unable to save summaryVal component: error = %v", saveErr)
	}
	return nil

}

func newSummaryVal(params NewSummaryValParams) (*SummaryVal, error) {

	if len(params.ParentDashboardID) <= 0 {
		return nil, fmt.Errorf("newSummaryVal: Error creating summary table: missing parent dashboard ID")
	}

	if !componentLayout.ValidGeometry(params.Geometry) {
		return nil, fmt.Errorf("newSummaryVal: Invalid geometry for bar chart: %+v", params.Geometry)
	}

	valSummary, valSummaryErr := values.NewValSummary(params.ValSummary)
	if valSummaryErr != nil {
		return nil, fmt.Errorf("NewBarChart: Error creating summary values for bar chart: error = %v", valSummaryErr)
	}

	summaryValProps := newDefaultSummaryValProps()
	summaryValProps.Geometry = params.Geometry
	summaryValProps.ValSummary = *valSummary

	newSummaryVal := SummaryVal{
		ParentDashboardID: params.ParentDashboardID,
		SummaryValID:      uniqueID.GenerateSnowflakeID(),
		Properties:        summaryValProps}

	if saveErr := saveSummaryVal(newSummaryVal); saveErr != nil {
		return nil, fmt.Errorf("newSummaryVal: Unable to save summary component with params=%+v: error = %v", params, saveErr)
	}

	return &newSummaryVal, nil
}

func GetSummaryVal(parentDashboardID string, summaryValID string) (*SummaryVal, error) {

	summaryValProps := newDefaultSummaryValProps()
	if getErr := common.GetDashboardComponent(summaryValEntityKind, parentDashboardID, summaryValID, &summaryValProps); getErr != nil {
		return nil, fmt.Errorf("getBarChart: Unable to retrieve bar chart component: %v", getErr)
	}

	summaryVal := SummaryVal{
		ParentDashboardID: parentDashboardID,
		SummaryValID:      summaryValID,
		Properties:        summaryValProps}

	return &summaryVal, nil

}

func GetSummaryVals(parentDashboardID string) ([]SummaryVal, error) {

	summaryVals := []SummaryVal{}
	addSummaryVal := func(summaryValID string, encodedProps string) error {

		summaryValProps := newDefaultSummaryValProps()
		if decodeErr := generic.DecodeJSONString(encodedProps, &summaryValProps); decodeErr != nil {
			return fmt.Errorf("GetSummaryVals: can't decode properties: %v", encodedProps)
		}

		currSummaryVal := SummaryVal{
			ParentDashboardID: parentDashboardID,
			SummaryValID:      summaryValID,
			Properties:        summaryValProps}

		summaryVals = append(summaryVals, currSummaryVal)

		return nil
	}
	if getErr := common.GetDashboardComponents(summaryValEntityKind, parentDashboardID, addSummaryVal); getErr != nil {
		return nil, fmt.Errorf("getBarCharts: Can't get bar chart components: %v")
	}

	return summaryVals, nil
}

func CloneSummaryVals(remappedIDs uniqueID.UniqueIDRemapper, srcParentDashboardID string) error {

	remappedDashboardID, err := remappedIDs.GetExistingRemappedID(srcParentDashboardID)
	if err != nil {
		return fmt.Errorf("CloneSummaryVals: %v", err)
	}

	summaryVals, err := GetSummaryVals(srcParentDashboardID)
	if err != nil {
		return fmt.Errorf("CloneSummaryVals: %v", err)
	}

	for _, srcSummaryVal := range summaryVals {

		remappedSummaryValID, err := remappedIDs.AllocNewRemappedID(srcSummaryVal.SummaryValID)
		if err != nil {
			return fmt.Errorf("CloneSummaryVals: %v", err)
		}

		clonedProps, err := srcSummaryVal.Properties.Clone(remappedIDs)
		if err != nil {
			return fmt.Errorf("CloneSummaryVals: %v", err)
		}

		destSummaryVal := SummaryVal{
			ParentDashboardID: remappedDashboardID,
			SummaryValID:      remappedSummaryValID,
			Properties:        *clonedProps}

		if err := saveSummaryVal(destSummaryVal); err != nil {
			return fmt.Errorf("CloneSummaryVals: %v", err)
		}
	}

	return nil
}

func updateExistingSummaryVal(updatedSummaryVal *SummaryVal) (*SummaryVal, error) {

	if updateErr := common.UpdateDashboardComponent(summaryValEntityKind, updatedSummaryVal.ParentDashboardID,
		updatedSummaryVal.SummaryValID, updatedSummaryVal.Properties); updateErr != nil {
		return nil, fmt.Errorf("Error updating summary table %+v: %v", updatedSummaryVal, updateErr)
	}

	return updatedSummaryVal, nil

}
