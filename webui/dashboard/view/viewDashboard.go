// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package view

import (
	"fmt"
	"net/http"
	"github.com/resultra/resultra/server/common/databaseWrapper"
	"github.com/resultra/resultra/server/databaseController"
	"github.com/resultra/resultra/server/userRole"
	"github.com/resultra/resultra/webui/dashboard/components"
)

type ViewDashboardInfo struct {
	DatabaseID      string
	DatabaseName    string
	DashboardID     string
	DashboardName   string
	CurrUserIsAdmin bool
	Title           string
	ComponentParams components.ComponentViewTemplateParams
}

func getViewDashboardInfo(r *http.Request, dashboardID string) (*ViewDashboardInfo, error) {

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		return nil, dbErr
	}

	dashboardDbInfo, getErr := databaseController.GetDashboardDatabaseInfo(trackerDBHandle, dashboardID)
	if getErr != nil {
		return nil, getErr
	}

	hasViewPrivs, privsErr := userRole.CurrentUserHasDashboardViewPrivs(trackerDBHandle, r, dashboardDbInfo.DatabaseID, dashboardID)
	if privsErr != nil {
		return nil, privsErr
	}
	if !hasViewPrivs {
		return nil, fmt.Errorf("ERROR: No permissions to view this dashboard")
	}

	isAdmin := userRole.CurrUserIsDatabaseAdmin(r, dashboardDbInfo.DatabaseID)

	dashboardInfo := ViewDashboardInfo{
		DatabaseID:      dashboardDbInfo.DatabaseID,
		DatabaseName:    dashboardDbInfo.DatabaseName,
		DashboardID:     dashboardDbInfo.DashboardID,
		DashboardName:   dashboardDbInfo.DashboardName,
		CurrUserIsAdmin: isAdmin}

	return &dashboardInfo, nil

}
