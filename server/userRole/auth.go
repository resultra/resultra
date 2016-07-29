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

func getTableDatabaseID(tableID string) (string, error) {

	databaseID := ""
	getErr := databaseWrapper.DBHandle().QueryRow(
		`SELECT database_id 
			FROM data_tables 
			WHERE table_id=$1 LIMIT 1`,
		tableID).Scan(&databaseID)
	if getErr != nil {
		return "", fmt.Errorf(
			"getTableDatabaseID: can't get database for table = %v: err=%v",
			tableID, getErr)
	}

	return databaseID, nil

}

func VerifyCurrUserIsDatabaseAdminForTable(req *http.Request, tableID string) error {

	databaseID, err := getTableDatabaseID(tableID)
	if err != nil {
		return fmt.Errorf("VerifyCurrUserIsDatabaseAdminForTable: %v", err)
	}

	if err := VerifyCurrUserIsDatabaseAdmin(req, databaseID); err != nil {
		return fmt.Errorf("VerifyCurrUserIsDatabaseAdminForTable: %v", err)
	}

	return nil
}
