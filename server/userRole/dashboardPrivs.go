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

type SetDashboardRolePrivsParams struct {
	DashboardID string `json:"dashboardID"`
	RoleID      string `json:"roleID"`
	Privs       string `json:"privs"`
}

func SetDashboardRolePrivs(params SetDashboardRolePrivsParams) error {

	if privsErr := verifyDashboardRolePrivs(params.Privs); privsErr != nil {
		return fmt.Errorf("SetDashboardRolePrivs: error = %v", privsErr)
	}

	if _, deleteErr := databaseWrapper.DBHandle().Exec(
		`DELETE FROM dashboard_role_privs where role_id=$1 and dashboard_id=$2`,
		params.RoleID, params.DashboardID); deleteErr != nil {
		return fmt.Errorf("SetDashboardRolePrivs: Can't delete old privs: error = %v", deleteErr)
	}

	if _, insertErr := databaseWrapper.DBHandle().Exec(
		`INSERT INTO dashboard_role_privs (role_id,dashboard_id,privs) VALUES ($1,$2,$3)`,
		params.RoleID, params.DashboardID, params.Privs); insertErr != nil {
		return fmt.Errorf("SetDashboardRolePrivs: Can't set form privileges: error = %v", insertErr)
	}

	return nil

}

type DashboardRolePriv struct {
	RoleID   string `json:"roleID"`
	RoleName string `json:"roleName"`
	Privs    string `json:"privs"`
}

func GetDashboardRolePrivs(dashboardID string) ([]DashboardRolePriv, error) {

	rows, queryErr := databaseWrapper.DBHandle().Query(
		`SELECT database_roles.role_id,database_roles.name,dashboard_role_privs.privs
			FROM dashboard_role_privs,database_roles
			WHERE dashboard_role_privs.dashboard_id=$1
				AND database_roles.role_id=dashboard_role_privs.role_id`, dashboardID)
	if queryErr != nil {
		return nil, fmt.Errorf("getDashboardRolePrivs: Failure querying database: %v", queryErr)
	}

	dashboardRolePrivs := []DashboardRolePriv{}
	for rows.Next() {

		currPrivInfo := DashboardRolePriv{}

		if scanErr := rows.Scan(&currPrivInfo.RoleID, &currPrivInfo.RoleName, &currPrivInfo.Privs); scanErr != nil {
			return nil, fmt.Errorf("getDashboardRolePrivs: Failure querying database: %v", scanErr)
		}

		dashboardRolePrivs = append(dashboardRolePrivs, currPrivInfo)

	}

	return dashboardRolePrivs, nil

}
