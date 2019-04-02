// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package dashboard

import (
	"fmt"
	"github.com/resultra/resultra/server/dashboard/components/barChart"
	"github.com/resultra/resultra/server/dashboard/components/gauge"
	"github.com/resultra/resultra/server/dashboard/components/header"
	"github.com/resultra/resultra/server/dashboard/components/summaryTable"
	"github.com/resultra/resultra/server/dashboard/components/summaryValue"
	"github.com/resultra/resultra/server/trackerDatabase"
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
