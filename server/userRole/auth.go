package userRole

import (
	"fmt"
	"net/http"
	"resultra/datasheet/server/generic/databaseWrapper"
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
			databaseID, currUserID, getErr)
	}

	return nil
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
