package userRole

import (
	"fmt"
	"net/http"
	"resultra/datasheet/server/common/databaseWrapper"
	"resultra/datasheet/server/generic/userAuth"
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

type GetRoleDashboardPrivParams struct {
	RoleID string `json:"roleID"`
}

type RoleDashboardPriv struct {
	DashboardID   string `json:"dashboardID"`
	DashboardName string `json:"dashboardName"`
	Privs         string `json:"privs"`
}

func GetRoleDashboardPrivs(roleID string) ([]RoleDashboardPriv, error) {

	rows, queryErr := databaseWrapper.DBHandle().Query(
		`SELECT dashboards.dashboard_id,dashboards.name,dashboard_role_privs.privs
			FROM dashboard_role_privs,dashboards
			WHERE dashboard_role_privs.role_id=$1 AND
				dashboard_role_privs.dashboard_id = dashboards.dashboard_id`, roleID)
	if queryErr != nil {
		return nil, fmt.Errorf("GetRoleDashboardPrivs: Failure querying database: %v", queryErr)
	}

	roleDashboardPrivs := []RoleDashboardPriv{}
	for rows.Next() {
		currPrivInfo := RoleDashboardPriv{}

		if scanErr := rows.Scan(&currPrivInfo.DashboardID, &currPrivInfo.DashboardName, &currPrivInfo.Privs); scanErr != nil {
			return nil, fmt.Errorf("GetRoleDashboardPrivs: Failure querying database: %v", scanErr)
		}
		roleDashboardPrivs = append(roleDashboardPrivs, currPrivInfo)
	}

	return roleDashboardPrivs, nil
}

// Query the database to build a lookup table of dashboardIDs for which the user has
// view permissions on the dashboard.
func GetDashboardsWithUserViewPrivs(databaseID string, userID string) (map[string]bool, error) {
	rows, queryErr := databaseWrapper.DBHandle().Query(
		`SELECT dashboard_role_privs.dashboard_id, dashboard_role_privs.privs
				FROM dashboard_role_privs,database_roles,collaborator_roles,collaborators
				WHERE database_roles.database_id=$1
					AND dashboard_role_privs.role_id=database_roles.role_id
					AND database_roles.role_id=collaborator_roles.role_id
					AND collaborator_roles.collaborator_id=collaborators.collaborator_id
					AND collaborators.user_id=$2`, databaseID, userID)
	if queryErr != nil {
		return nil, fmt.Errorf("GetCustomRoleInfo: Failure querying database: %v", queryErr)
	}

	visibleDashboards := map[string]bool{}
	for rows.Next() {
		dashboardID := ""
		privs := ""
		if scanErr := rows.Scan(&dashboardID, &privs); scanErr != nil {
			return nil, fmt.Errorf("GetCustomRoleDashboardInfo: Failure querying database: %v", scanErr)
		}
		if privs == DashboardRolePrivsView {
			visibleDashboards[dashboardID] = true
		}
	}

	return visibleDashboards, nil
}

func CurrentUserHasDashboardViewPrivs(req *http.Request,
	databaseID string, dashboardID string) (bool, error) {

	privs := false

	if CurrUserIsDatabaseAdmin(req, databaseID) {
		return true, nil
	}

	currUserID, userErr := userAuth.GetCurrentUserID(req)
	if userErr != nil {
		return false, fmt.Errorf("verifyCurrUserIsDatabaseAdmin: can't verify user: %v", userErr)
	}

	rows, queryErr := databaseWrapper.DBHandle().Query(
		`SELECT dashboard_role_privs.privs
					FROM dashboard_role_privs,database_roles,collaborator_roles,collaborators
					WHERE database_roles.database_id=$1
					AND dashboard_role_privs.dashboard_id=$2
						AND dashboard_role_privs.role_id=database_roles.role_id
						AND database_roles.role_id=collaborator_roles.role_id
						AND collaborator_roles.collaborator_id=collaborators.collaborator_id
						AND collaborators.user_id=$3`, databaseID, dashboardID, currUserID)
	if queryErr != nil {
		return false, fmt.Errorf("GetCustomRoleInfo: Failure querying database: %v", queryErr)
	}

	for rows.Next() {
		rolePrivs := ""
		if scanErr := rows.Scan(&rolePrivs); scanErr != nil {
			return false, fmt.Errorf("GetCustomRoleDashboardInfo: Failure querying database: %v", scanErr)
		}
		// If the user's privileges under any role are to view the dashboard, then set the privileges to view
		if rolePrivs == DashboardRolePrivsView {
			privs = true
		}
	}

	return privs, nil

}
