package databaseController

import (
	"fmt"
	"net/http"

	"resultra/datasheet/server/generic/databaseWrapper"
	"resultra/datasheet/server/generic/userAuth"
)

type UserTrackingDatabaseInfo struct {
	DatabaseID   string `json:"databaseID"`
	DatabaseName string `json:"databaseName"`
	IsAdmin      bool   `json:"isAdmin"`
}

func getCurrentUserTrackingDatabases(req *http.Request) ([]UserTrackingDatabaseInfo, error) {

	currUserID, userErr := userAuth.GetCurrentUserID(req)
	if userErr != nil {
		return nil, fmt.Errorf("getCurrentUserTrackingDatabases: can't get current user: %v", userErr)
	}

	rows, queryErr := databaseWrapper.DBHandle().Query(
		`SELECT databases.database_id, databases.name FROM database_admins,databases WHERE 
			database_admins.user_id=$1 AND 
			database_admins.database_id = databases.database_id`, currUserID)
	if queryErr != nil {
		return nil, fmt.Errorf("getCurrentUserTrackingDatabases: Failure querying database: %v", queryErr)
	}

	userTrackingDBInfo := []UserTrackingDatabaseInfo{}
	for rows.Next() {
		var currTrackingDBInfo UserTrackingDatabaseInfo
		currTrackingDBInfo.IsAdmin = true
		if scanErr := rows.Scan(&currTrackingDBInfo.DatabaseID, &currTrackingDBInfo.DatabaseName); scanErr != nil {
			return nil, fmt.Errorf("getDatabaseDashboardsInfo: Failure querying database: %v", scanErr)
		}
		userTrackingDBInfo = append(userTrackingDBInfo, currTrackingDBInfo)
	}

	return userTrackingDBInfo, nil

}
