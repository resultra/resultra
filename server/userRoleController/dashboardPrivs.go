// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package userRoleController

import (
	"database/sql"
	"fmt"
	"github.com/resultra/resultra/server/dashboard"
	"github.com/resultra/resultra/server/userRole"
)

func getRoleDashboardPrivsWithDefaults(trackerDBHandle *sql.DB, roleID string) ([]userRole.RoleDashboardPriv, error) {

	roleDB, err := userRole.GetUserRoleDatabaseID(trackerDBHandle, roleID)
	if err != nil {
		return nil, fmt.Errorf("getRoleDashboardPrivsWithDefaults: %v", err)
	}

	privsByDashboardID := map[string]userRole.RoleDashboardPriv{}

	// Start off with no privileges as the default for all dashboards
	allDashboards, err := dashboard.GetAllDashboards(trackerDBHandle, roleDB)
	if err != nil {
		return nil, fmt.Errorf("getRoleListPrivsWithDefaults: %v", err)
	}
	for _, currDashboard := range allDashboards {
		privsByDashboardID[currDashboard.DashboardID] = userRole.RoleDashboardPriv{
			DashboardID:   currDashboard.DashboardID,
			DashboardName: currDashboard.Name,
			Privs:         userRole.DashboardRolePrivsNone}
	}

	// Update the privileges for those dashboards with an explicit set of privileges set.
	explicitDashboardPrivs, err := userRole.GetRoleDashboardPrivs(trackerDBHandle, roleID)
	if err != nil {
		return nil, fmt.Errorf("getRoleListPrivsWithDefaults: %v", err)
	}
	for _, currPriv := range explicitDashboardPrivs {
		privsByDashboardID[currPriv.DashboardID] = currPriv
	}

	// Flatten the list
	roleDashboardPrivs := []userRole.RoleDashboardPriv{}
	for _, currPriv := range privsByDashboardID {
		roleDashboardPrivs = append(roleDashboardPrivs, currPriv)
	}

	return roleDashboardPrivs, nil

}
