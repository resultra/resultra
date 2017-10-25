package dashboard

import (
	"fmt"
	"resultra/datasheet/server/dashboard/components/barChart"
	"resultra/datasheet/server/dashboard/components/gauge"
	"resultra/datasheet/server/dashboard/components/header"
	"resultra/datasheet/server/dashboard/components/summaryTable"
	"resultra/datasheet/server/dashboard/components/summaryValue"
	"resultra/datasheet/server/trackerDatabase"
)

func cloneDashboardComponents(cloneParams *trackerDatabase.CloneDatabaseParams, srcParentDashboard string) error {

	if err := barChart.CloneBarCharts(cloneParams, srcParentDashboard); err != nil {
		return fmt.Errorf("cloneDashboardComponents: %v", err)
	}

	if err := summaryTable.CloneSummaryTables(cloneParams, srcParentDashboard); err != nil {
		return fmt.Errorf("cloneDashboardComponents: %v", err)
	}

	if err := header.CloneHeaders(cloneParams, srcParentDashboard); err != nil {
		return fmt.Errorf("cloneDashboardComponents: %v", err)
	}

	if err := gauge.CloneGauges(cloneParams, srcParentDashboard); err != nil {
		return fmt.Errorf("cloneDashboardComponents: %v", err)
	}

	if err := summaryValue.CloneSummaryVals(cloneParams, srcParentDashboard); err != nil {
		return fmt.Errorf("cloneDashboardComponents: %v", err)
	}

	return nil
}
