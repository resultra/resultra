package userRole

import (
	"fmt"
	"net/http"
	"resultra/datasheet/server/common/databaseWrapper"
	"resultra/datasheet/server/generic/userAuth"
)

func VerifyCurrUserIsDatabaseAdmin(req *http.Request, databaseID string) error {

	currUserID, userErr := userAuth.GetCurrentUserID(req)
	if userErr != nil {
		return fmt.Errorf("verifyCurrUserIsDatabaseAdmin: can't verify user: %v", userErr)
	}

	queryUserID := ""
	getErr := databaseWrapper.DBHandle().QueryRow(
		`SELECT user_id 
			FROM database_admins 
			WHERE database_id=$1 AND user_id=$2 LIMIT 1`,
		databaseID, currUserID).Scan(&queryUserID)
	if getErr != nil {
		return fmt.Errorf(
			"verifyCurrUserIsDatabaseAdmin: can't validate user = %v as admin for database = %v: err=%v",
			currUserID, databaseID, getErr)
	}

	return nil
}

func CurrUserIsDatabaseAdmin(req *http.Request, databaseID string) bool {
	verifyCurrUserAdminErr := VerifyCurrUserIsDatabaseAdmin(req, databaseID)
	if verifyCurrUserAdminErr != nil {
		return false
	} else {
		return true
	}
}

func getFormDatabaseID(formID string) (string, error) {

	databaseID := ""
	getErr := databaseWrapper.DBHandle().QueryRow(
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

	databaseID, err := getFormDatabaseID(formID)
	if err != nil {
		return fmt.Errorf("VerifyCurrUserIsDatabaseAdminForForm: %v", err)
	}

	if err := VerifyCurrUserIsDatabaseAdmin(req, databaseID); err != nil {
		return fmt.Errorf("VerifyCurrUserIsDatabaseAdminForTable: %v", err)
	}

	return nil
}

func getDashboardDatabaseID(dashboardID string) (string, error) {

	databaseID := ""
	getErr := databaseWrapper.DBHandle().QueryRow(
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

	databaseID, err := getDashboardDatabaseID(dashboardID)
	if err != nil {
		return fmt.Errorf("VerifyCurrUserIsDatabaseAdminForDashboard: %v", err)
	}

	if err := VerifyCurrUserIsDatabaseAdmin(req, databaseID); err != nil {
		return fmt.Errorf("VerifyCurrUserIsDatabaseAdminForTable: %v", err)
	}

	return nil
}

func getItemListDatabaseID(listID string) (string, error) {

	databaseID := ""
	getErr := databaseWrapper.DBHandle().QueryRow(
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

	databaseID, err := getItemListDatabaseID(listID)
	if err != nil {
		return fmt.Errorf("VerifyCurrUserIsDatabaseAdminForItemList: %v", err)
	}

	if err := VerifyCurrUserIsDatabaseAdmin(req, databaseID); err != nil {
		return fmt.Errorf("VerifyCurrUserIsDatabaseAdminForItemList: %v", err)
	}

	return nil
}

func getNewItemLinkDatabaseID(linkID string) (string, error) {

	databaseID := ""
	getErr := databaseWrapper.DBHandle().QueryRow(
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

	databaseID, err := getNewItemLinkDatabaseID(linkID)
	if err != nil {
		return fmt.Errorf("VerifyCurrUserIsDatabaseAdminForItemList: %v", err)
	}

	if err := VerifyCurrUserIsDatabaseAdmin(req, databaseID); err != nil {
		return fmt.Errorf("VerifyCurrUserIsDatabaseAdminForItemList: %v", err)
	}

	return nil
}

func getAlertDatabaseID(alertID string) (string, error) {

	databaseID := ""
	getErr := databaseWrapper.DBHandle().QueryRow(
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

	databaseID, err := getAlertDatabaseID(alertID)
	if err != nil {
		return fmt.Errorf("VerifyCurrUserIsDatabaseAdminForAlert: %v", err)
	}

	if err := VerifyCurrUserIsDatabaseAdmin(req, databaseID); err != nil {
		return fmt.Errorf("VerifyCurrUserIsDatabaseAdminForAlert: %v", err)
	}

	return nil
}

func GetUserRoleDatabaseID(roleID string) (string, error) {
	databaseID := ""
	getErr := databaseWrapper.DBHandle().QueryRow(
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

	databaseID, err := GetUserRoleDatabaseID(roleID)
	if err != nil {
		return fmt.Errorf("VerifyCurrUserIsDatabaseAdminForUserRole: %v", err)
	}

	if err := VerifyCurrUserIsDatabaseAdmin(req, databaseID); err != nil {
		return fmt.Errorf("VerifyCurrUserIsDatabaseAdminForItemList: %v", err)
	}

	return nil
}
