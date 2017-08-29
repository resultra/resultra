package dashboard

import (
	"fmt"
	"resultra/datasheet/server/dashboard/components/barChart"
	"resultra/datasheet/server/dashboard/components/gauge"
	"resultra/datasheet/server/dashboard/components/header"
	"resultra/datasheet/server/dashboard/components/summaryTable"
	"resultra/datasheet/server/dashboard/components/summaryValue"
	"resultra/datasheet/server/generic/uniqueID"
)

func cloneDashboardComponents(remappedIDs uniqueID.UniqueIDRemapper, srcParentDashboard string) error {

	if err := barChart.CloneBarCharts(remappedIDs, srcParentDashboard); err != nil {
		return fmt.Errorf("cloneDashboardComponents: %v", err)
	}

	if err := summaryTable.CloneSummaryTables(remappedIDs, srcParentDashboard); err != nil {
		return fmt.Errorf("cloneDashboardComponents: %v", err)
	}

	if err := header.CloneHeaders(remappedIDs, srcParentDashboard); err != nil {
		return fmt.Errorf("cloneDashboardComponents: %v", err)
	}

	if err := gauge.CloneGauges(remappedIDs, srcParentDashboard); err != nil {
		return fmt.Errorf("cloneDashboardComponents: %v", err)
	}

	if err := summaryValue.CloneSummaryVals(remappedIDs, srcParentDashboard); err != nil {
		return fmt.Errorf("cloneDashboardComponents: %v", err)
	}

	return nil
}
