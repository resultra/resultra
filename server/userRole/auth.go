// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package userRole

import (
	"database/sql"
	"fmt"
	"github.com/resultra/resultra/server/common/databaseWrapper"
	"github.com/resultra/resultra/server/common/userAuth"
	"net/http"
)

func VerifyCurrUserIsDatabaseAdmin(trackerDBHandle *sql.DB, req *http.Request, databaseID string) error {

	currUserID, userErr := userAuth.GetCurrentUserID(req)
	if userErr != nil {
		return fmt.Errorf("verifyCurrUserIsDatabaseAdmin: can't verify user: %v", userErr)
	}

	queryUserID := ""
	getErr := trackerDBHandle.QueryRow(
		`SELECT user_id 
			FROM collaborators 
			WHERE database_id=$1 AND user_id=$2 AND is_admin='1' LIMIT 1`,
		databaseID, currUserID).Scan(&queryUserID)
	if getErr != nil {
		return fmt.Errorf(
			"verifyCurrUserIsDatabaseAdmin: can't validate user = %v as admin for database = %v: err=%v",
			currUserID, databaseID, getErr)
	}

	return nil
}

func CurrUserIsDatabaseAdmin(req *http.Request, databaseID string) bool {

	// TODO - process error instead of just returning false.
	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(req)
	if dbErr != nil {
		return false
	}

	verifyCurrUserAdminErr := VerifyCurrUserIsDatabaseAdmin(trackerDBHandle, req, databaseID)
	if verifyCurrUserAdminErr != nil {
		return false
	} else {
		return true
	}
}

func getFormDatabaseID(trackerDBHandle *sql.DB, formID string) (string, error) {

	databaseID := ""
	getErr := trackerDBHandle.QueryRow(
		`SELECT database_id 
			FROM forms 
			WHERE forms.form_id=$1 LIMIT 1`,
		formID).Scan(&databaseID)
	if getErr != nil {
		return "", fmt.Errorf(
			"getFormDatabaseID: can't get database for form = %v: err=%v",
			formID, getErr)
	}

	return databaseID, nil

}

func VerifyCurrUserIsDatabaseAdminForForm(req *http.Request, formID string) error {

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(req)
	if dbErr != nil {
		return dbErr
	}

	databaseID, err := getFormDatabaseID(trackerDBHandle, formID)
	if err != nil {
		return fmt.Errorf("VerifyCurrUserIsDatabaseAdminForForm: %v", err)
	}

	if err := VerifyCurrUserIsDatabaseAdmin(trackerDBHandle, req, databaseID); err != nil {
		return fmt.Errorf("VerifyCurrUserIsDatabaseAdminForTable: %v", err)
	}

	return nil
}

func getDashboardDatabaseID(trackerDBHandle *sql.DB, dashboardID string) (string, error) {

	databaseID := ""
	getErr := trackerDBHandle.QueryRow(
		`SELECT database_id 
			FROM dashboards 
			WHERE dashboards.dashboard_id=$1 LIMIT 1`,
		dashboardID).Scan(&databaseID)
	if getErr != nil {
		return "", fmt.Errorf(
			"getDashboardDatabaseID: can't get database for dashboard = %v: err=%v",
			dashboardID, getErr)
	}

	return databaseID, nil

}

func VerifyCurrUserIsDatabaseAdminForDashboard(req *http.Request, dashboardID string) error {

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(req)
	if dbErr != nil {
		return dbErr
	}

	databaseID, err := getDashboardDatabaseID(trackerDBHandle, dashboardID)
	if err != nil {
		return fmt.Errorf("VerifyCurrUserIsDatabaseAdminForDashboard: %v", err)
	}

	if err := VerifyCurrUserIsDatabaseAdmin(trackerDBHandle, req, databaseID); err != nil {
		return fmt.Errorf("VerifyCurrUserIsDatabaseAdminForTable: %v", err)
	}

	return nil
}

func getItemListDatabaseID(trackerDBHandle *sql.DB, listID string) (string, error) {

	databaseID := ""
	getErr := trackerDBHandle.QueryRow(
		`SELECT database_id 
			FROM item_lists 
			WHERE item_lists.list_id=$1 LIMIT 1`,
		listID).Scan(&databaseID)
	if getErr != nil {
		return "", fmt.Errorf(
			"getListDatabaseID: can't get database for item list = %v: err=%v",
			listID, getErr)
	}

	return databaseID, nil

}

func VerifyCurrUserIsDatabaseAdminForItemList(req *http.Request, listID string) error {

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(req)
	if dbErr != nil {
		return dbErr
	}

	databaseID, err := getItemListDatabaseID(trackerDBHandle, listID)
	if err != nil {
		return fmt.Errorf("VerifyCurrUserIsDatabaseAdminForItemList: %v", err)
	}

	if err := VerifyCurrUserIsDatabaseAdmin(trackerDBHandle, req, databaseID); err != nil {
		return fmt.Errorf("VerifyCurrUserIsDatabaseAdminForItemList: %v", err)
	}

	return nil
}

func getNewItemLinkDatabaseID(trackerDBHandle *sql.DB, linkID string) (string, error) {

	databaseID := ""
	getErr := trackerDBHandle.QueryRow(
		`SELECT forms.database_id 
			FROM forms,form_links 
			WHERE form_links.link_id=$1 
				AND form_links.form_id=forms.form_id
			LIMIT 1`,
		linkID).Scan(&databaseID)
	if getErr != nil {
		return "", fmt.Errorf(
			"getNewItemLinkDatabaseID: can't get database for new item form link = %v: err=%v",
			linkID, getErr)
	}

	return databaseID, nil

}

func VerifyCurrUserIsDatabaseAdminForNewItemLink(req *http.Request, linkID string) error {

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(req)
	if dbErr != nil {
		return dbErr
	}

	databaseID, err := getNewItemLinkDatabaseID(trackerDBHandle, linkID)
	if err != nil {
		return fmt.Errorf("VerifyCurrUserIsDatabaseAdminForItemList: %v", err)
	}

	if err := VerifyCurrUserIsDatabaseAdmin(trackerDBHandle, req, databaseID); err != nil {
		return fmt.Errorf("VerifyCurrUserIsDatabaseAdminForItemList: %v", err)
	}

	return nil
}

func getAlertDatabaseID(trackerDBHandle *sql.DB, alertID string) (string, error) {

	databaseID := ""
	getErr := trackerDBHandle.QueryRow(
		`SELECT alerts.database_id 
			FROM alerts 
			WHERE alerts.alert_id=$1 
			LIMIT 1`,
		alertID).Scan(&databaseID)
	if getErr != nil {
		return "", fmt.Errorf(
			"getAlertDatabaseID: can't get database for alert = %v: err=%v",
			alertID, getErr)
	}

	return databaseID, nil

}

func VerifyCurrUserIsDatabaseAdminForAlert(req *http.Request, alertID string) error {

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(req)
	if dbErr != nil {
		return dbErr
	}

	databaseID, err := getAlertDatabaseID(trackerDBHandle, alertID)
	if err != nil {
		return fmt.Errorf("VerifyCurrUserIsDatabaseAdminForAlert: %v", err)
	}

	if err := VerifyCurrUserIsDatabaseAdmin(trackerDBHandle, req, databaseID); err != nil {
		return fmt.Errorf("VerifyCurrUserIsDatabaseAdminForAlert: %v", err)
	}

	return nil
}

func GetUserRoleDatabaseID(trackerDBHandle *sql.DB, roleID string) (string, error) {
	databaseID := ""
	getErr := trackerDBHandle.QueryRow(
		`SELECT database_id 
			FROM database_roles 
			WHERE database_roles.role_id=$1 LIMIT 1`,
		roleID).Scan(&databaseID)
	if getErr != nil {
		return "", fmt.Errorf(
			"getUserRoleDatabaseID: can't get database for user role = %v: err=%v",
			roleID, getErr)
	}

	return databaseID, nil
}

func VerifyCurrUserIsDatabaseAdminForUserRole(req *http.Request, roleID string) error {

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(req)
	if dbErr != nil {
		return dbErr
	}

	databaseID, err := GetUserRoleDatabaseID(trackerDBHandle, roleID)
	if err != nil {
		return fmt.Errorf("VerifyCurrUserIsDatabaseAdminForUserRole: %v", err)
	}

	if err := VerifyCurrUserIsDatabaseAdmin(trackerDBHandle, req, databaseID); err != nil {
		return fmt.Errorf("VerifyCurrUserIsDatabaseAdminForItemList: %v", err)
	}

	return nil
}
