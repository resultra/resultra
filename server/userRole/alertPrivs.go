// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package userRole

import (
	"database/sql"
	"fmt"
	"net/http"
	"github.com/resultra/resultra/server/common/userAuth"
	"github.com/resultra/resultra/server/trackerDatabase"
)

type SetAlertRolePrivsParams struct {
	AlertID      string `json:"alertID"`
	RoleID       string `json:"roleID"`
	AlertEnabled bool   `json:"alertEnabled"`
}

func setAlertRolePrivsToDest(destDBHandle *sql.DB, params SetAlertRolePrivsParams) error {

	if _, deleteErr := destDBHandle.Exec(
		`DELETE FROM alert_role_privs where role_id=$1 and alert_id=$2`,
		params.RoleID, params.AlertID); deleteErr != nil {
		return fmt.Errorf("SetAlertRolePrivs: Can't delete old privs: error = %v", deleteErr)
	}

	if params.AlertEnabled {
		if _, insertErr := destDBHandle.Exec(
			`INSERT INTO alert_role_privs (role_id,alert_id) VALUES ($1,$2)`,
			params.RoleID, params.AlertID); insertErr != nil {
			return fmt.Errorf("SetAlertRolePrivs: Can't set list privileges: error = %v", insertErr)
		}
	}

	return nil

}

func SetAlertRolePrivs(trackerDBHandle *sql.DB, params SetAlertRolePrivsParams) error {
	return setAlertRolePrivsToDest(trackerDBHandle, params)
}

func getAllAlertPrivsFromSrc(srcDBHandle *sql.DB, parentDatabaseID string) ([]SetAlertRolePrivsParams, error) {

	// Retrieve alerts for which the privileges have been explicitely set.
	rows, queryErr := srcDBHandle.Query(
		`SELECT alert_role_privs.alert_id,alert_role_privs.role_id
			FROM alert_role_privs,alerts
			WHERE alerts.database_id=$1 AND
				alert_role_privs.alert_id = alerts.alert_id`, parentDatabaseID)
	if queryErr != nil {
		return nil, fmt.Errorf("getAllAlertPrivsFromSrc: Failure querying database: %v", queryErr)
	}
	defer rows.Close()

	privs := []SetAlertRolePrivsParams{}
	for rows.Next() {
		currPriv := SetAlertRolePrivsParams{}
		currPriv.AlertEnabled = true // presence in the database means the link is enabled

		if scanErr := rows.Scan(&currPriv.AlertID, &currPriv.RoleID); scanErr != nil {
			return nil, fmt.Errorf("GetAlertPrivs: Failure querying database: %v", scanErr)
		}
		privs = append(privs, currPriv)
	}

	return privs, nil
}

func CloneAlertPrivs(cloneParams *trackerDatabase.CloneDatabaseParams) error {

	privs, err := getAllAlertPrivsFromSrc(cloneParams.SrcDBHandle, cloneParams.SourceDatabaseID)
	if err != nil {
		return fmt.Errorf("CloneRoles: Unable to retrieve roles: databaseID=%v, error=%v ",
			cloneParams.SourceDatabaseID, err)
	}
	for _, currPriv := range privs {

		destAlertID, err := cloneParams.IDRemapper.GetExistingRemappedID(currPriv.AlertID)
		if err != nil {
			return fmt.Errorf("CloneListPrivs: Unable to get mapped ID for source database: %v", err)
		}

		destRoleID, err := cloneParams.IDRemapper.GetExistingRemappedID(currPriv.RoleID)
		if err != nil {
			return fmt.Errorf("CloneListPrivs: Unable to get mapped ID for source database: %v", err)
		}

		destPrivs := SetAlertRolePrivsParams{
			AlertID:      destAlertID,
			RoleID:       destRoleID,
			AlertEnabled: true}

		if err := setAlertRolePrivsToDest(cloneParams.DestDBHandle, destPrivs); err != nil {
			return fmt.Errorf("CloneListPrivs: Can't clone list privilege: error = %v", err)
		}
	}

	return nil

}

type GetRoleAlertPrivParams struct {
	RoleID string `json:"roleID"`
}

type RoleAlertPriv struct {
	AlertID      string `json:"alertID"`
	AlertName    string `json:"alertName"`
	AlertEnabled bool   `json:"alertEnabled"`
}

func getDefaultAlertPrivs(trackerDBHandle *sql.DB, databaseID string) ([]RoleAlertPriv, error) {

	rows, queryErr := trackerDBHandle.Query(
		`SELECT alerts.alert_id,alerts.name
			FROM alerts
			WHERE alerts.database_id=$1`, databaseID)
	if queryErr != nil {
		return nil, fmt.Errorf("getDefaultAlertPrivs: Failure querying database: %v", queryErr)
	}
	defer rows.Close()

	roleAlertPrivs := []RoleAlertPriv{}
	for rows.Next() {
		currPrivInfo := RoleAlertPriv{}
		currPrivInfo.AlertEnabled = false // presence in the database means the alert is enabled

		if scanErr := rows.Scan(&currPrivInfo.AlertID, &currPrivInfo.AlertName); scanErr != nil {
			return nil, fmt.Errorf("GetRoleListPrivs: Failure querying database: %v", scanErr)
		}
		roleAlertPrivs = append(roleAlertPrivs, currPrivInfo)
	}
	return roleAlertPrivs, nil
}

// For a given role, get the list of roles with privileges
func GetRoleAlertPrivs(trackerDBHandle *sql.DB, roleID string) ([]RoleAlertPriv, error) {

	roleDatabaseID, roleDBErr := GetUserRoleDatabaseID(trackerDBHandle, roleID)
	if roleDBErr != nil {
		return nil, fmt.Errorf("GetAlertPrivs: Failure querying database: %v", roleDBErr)
	}

	defaultAlerts, getAlertErr := getDefaultAlertPrivs(trackerDBHandle, roleDatabaseID)
	if getAlertErr != nil {
		return nil, fmt.Errorf("GetAlertPrivs: Failure querying database: %v", getAlertErr)
	}

	// Initially set privilieges to false for all alerts
	privsByAlertID := map[string]RoleAlertPriv{}
	for _, defaultPriv := range defaultAlerts {
		privsByAlertID[defaultPriv.AlertID] = defaultPriv
	}

	// Retrieve alerts for which the privileges have been explicitely set.
	rows, queryErr := trackerDBHandle.Query(
		`SELECT alerts.alert_id,alerts.name
			FROM alert_role_privs,alerts
			WHERE alert_role_privs.role_id=$1 AND
				alert_role_privs.alert_id = alerts.alert_id`, roleID)
	if queryErr != nil {
		return nil, fmt.Errorf("GetAlertPrivs: Failure querying database: %v", queryErr)
	}
	defer rows.Close()

	for rows.Next() {
		currPrivInfo := RoleAlertPriv{}
		currPrivInfo.AlertEnabled = true // presence in the database means the link is enabled

		if scanErr := rows.Scan(&currPrivInfo.AlertID, &currPrivInfo.AlertName); scanErr != nil {
			return nil, fmt.Errorf("GetAlertPrivs: Failure querying database: %v", scanErr)
		}
		privsByAlertID[currPrivInfo.AlertID] = currPrivInfo
	}

	// Flatten the privileges back down to a list.
	roleAlertPrivs := []RoleAlertPriv{}
	for _, currPrivInfo := range privsByAlertID {
		roleAlertPrivs = append(roleAlertPrivs, currPrivInfo)
	}

	return roleAlertPrivs, nil
}

type GetAlertRolePrivParams struct {
	AlertID string `json:"alertID"`
}

type AlertRolePriv struct {
	RoleID       string `json:"roleID"`
	RoleName     string `json:"roleName"`
	AlertEnabled bool   `json:"alertEnabled"`
}

func GetAlertRolePrivs(trackerDBHandle *sql.DB, alertID string) ([]AlertRolePriv, error) {

	alertDatabaseID, alertDBErr := getAlertDatabaseID(trackerDBHandle, alertID)
	if alertDBErr != nil {
		return nil, fmt.Errorf("GetAlertPrivs: Failure querying database: %v", alertDBErr)
	}

	allRoles, getRoleErr := GetDatabaseRoles(trackerDBHandle, alertDatabaseID)
	if getRoleErr != nil {
		return nil, fmt.Errorf("GetAlertPrivs: Failure querying database: %v", getRoleErr)
	}

	// Initially set privilieges to false for all alerts
	privsByRoleID := map[string]AlertRolePriv{}
	for _, role := range allRoles {
		defaultPriv := AlertRolePriv{
			RoleID:       role.RoleID,
			RoleName:     role.RoleName,
			AlertEnabled: false}
		privsByRoleID[defaultPriv.RoleID] = defaultPriv
	}

	// Retrieve alerts for which the privileges have been explicitely set.
	rows, queryErr := trackerDBHandle.Query(
		`SELECT database_roles.role_id,database_roles.name
			FROM alert_role_privs,database_roles
			WHERE alert_role_privs.alert_id=$1 AND
				alert_role_privs.role_id = database_roles.role_id`, alertID)
	if queryErr != nil {
		return nil, fmt.Errorf("GetAlertRolePrivs: Failure querying database: %v", queryErr)
	}
	defer rows.Close()

	for rows.Next() {
		currPrivInfo := AlertRolePriv{}
		currPrivInfo.AlertEnabled = true // presence in the database means the link is enabled

		if scanErr := rows.Scan(&currPrivInfo.RoleID, &currPrivInfo.RoleName); scanErr != nil {
			return nil, fmt.Errorf("GetAlertPrivs: Failure querying database: %v", scanErr)
		}
		privsByRoleID[currPrivInfo.RoleID] = currPrivInfo
	}

	// Flatten the privileges back down to a list.
	alertRolePrivs := []AlertRolePriv{}
	for _, currPrivInfo := range privsByRoleID {
		alertRolePrivs = append(alertRolePrivs, currPrivInfo)
	}

	return alertRolePrivs, nil
}

func GetAlertsWithUserPrivs(trackerDBHandle *sql.DB, databaseID string, userID string) (map[string]bool, error) {
	rows, queryErr := trackerDBHandle.Query(
		`SELECT alert_role_privs.alert_id
				FROM alert_role_privs,database_roles,collaborator_roles,collaborators
				WHERE database_roles.database_id=$1
					AND alert_role_privs.role_id=database_roles.role_id
					AND database_roles.role_id=collaborator_roles.role_id
					AND collaborator_roles.collaborator_id=collaborators.collaborator_id
					AND collaborators.user_id=$2`, databaseID, userID)
	if queryErr != nil {
		return nil, fmt.Errorf("GetItemListsWithUserPrivs: Failure querying database: %v", queryErr)
	}
	defer rows.Close()

	visibleAlerts := map[string]bool{}
	for rows.Next() {
		alertID := ""
		if scanErr := rows.Scan(&alertID); scanErr != nil {
			return nil, fmt.Errorf("GetAlertsWithUserPrivs: Failure querying database: %v", scanErr)
		}
		visibleAlerts[alertID] = true
	}

	return visibleAlerts, nil
}

func CurrentUserHasAlertPrivs(trackerDBHandle *sql.DB, req *http.Request,
	databaseID string, alertID string) (bool, error) {

	currUserID, userErr := userAuth.GetCurrentUserID(req)
	if userErr != nil {
		return false, fmt.Errorf("verifyCurrUserIsDatabaseAdmin: can't verify user: %v", userErr)
	}

	rows, queryErr := trackerDBHandle.Query(
		`SELECT alert_role_privs.alert_id
					FROM alert_role_privs,database_roles,collaborator_roles,collaborators
					WHERE database_roles.database_id=$1
					AND alert_role_privs.alert_id=$2
						AND alert_role_privs.role_id=database_roles.role_id
						AND database_roles.role_id=collaborator_roles.role_id
						AND collaborator_roles.collaborator_id=collaborators.collaborator_id
						AND collaborators.user_id=$3`, databaseID, alertID, currUserID)
	if queryErr != nil {
		return false, fmt.Errorf("CurrentUserHasAlertPrivs: Failure querying database: %v", queryErr)
	}
	defer rows.Close()

	// Default to false (no priveleges) unless there's privileges set in the database.
	privs := false

	for rows.Next() {
		alertWithPrivs := ""
		if scanErr := rows.Scan(&alertWithPrivs); scanErr != nil {
			return false, fmt.Errorf("CurrentUserHasAlertPrivs: Failure querying database: %v", scanErr)
		}
		privs = true
	}

	return privs, nil

}
