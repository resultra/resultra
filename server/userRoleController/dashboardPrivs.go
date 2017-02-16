package userRoleController

import (
	"fmt"
	"resultra/datasheet/server/dashboard"
	"resultra/datasheet/server/userRole"
)

func getRoleDashboardPrivsWithDefaults(roleID string) ([]userRole.RoleDashboardPriv, error) {

	roleDB, err := userRole.GetUserRoleDatabaseID(roleID)
	if err != nil {
		return nil, fmt.Errorf("getRoleDashboardPrivsWithDefaults: %v", err)
	}

	privsByDashboardID := map[string]userRole.RoleDashboardPriv{}

	// Start off with no privileges as the default for all dashboards
	allDashboards, err := dashboard.GetAllDashboards(roleDB)
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
	explicitDashboardPrivs, err := userRole.GetRoleDashboardPrivs(roleID)
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
