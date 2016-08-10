package userRole

import (
	"fmt"
	"resultra/datasheet/server/generic/databaseWrapper"
)

const DashboardRolePrivsNone string = "none"
const DashboardRolePrivsView string = "view"

type DashboardPriv struct {
	DashboardID string `json:"dashboardID"`
	Privs       string `json:"privs"`
}

func verifyDashboardRolePrivs(privs string) error {
	if (privs == DashboardRolePrivsNone) ||
		(privs == DashboardRolePrivsView) {
		return nil
	} else {
		return fmt.Errorf("verifyFormRolePrivs: Invalid privileges: %v", privs)
	}
}

func setDashboardRolePrivs(roleID string, dashboardID string, privs string) error {

	if privsErr := verifyDashboardRolePrivs(privs); privsErr != nil {
		return fmt.Errorf("setDashboardRolePrivs: error = %v", privsErr)
	}

	if _, deleteErr := databaseWrapper.DBHandle().Exec(
		`DELETE FROM dashboard_role_privs where role_id=$1 and dashboard_id=$2`,
		roleID, dashboardID); deleteErr != nil {
		return fmt.Errorf("setDashboardRolePrivs: Can't delete old privs: error = %v", deleteErr)
	}

	if _, insertErr := databaseWrapper.DBHandle().Exec(
		`INSERT INTO dashboard_role_privs (role_id,dashboard_id,privs) VALUES ($1,$2,$3)`,
		roleID, dashboardID, privs); insertErr != nil {
		return fmt.Errorf("setDashboardRolePrivs: Can't set form privileges: error = %v", insertErr)
	}

	return nil

}
