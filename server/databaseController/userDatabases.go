package databaseController

import (
	"fmt"
	"net/http"

	"resultra/datasheet/server/common/databaseWrapper"
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

	trackingInfoByDatabase := map[string]UserTrackingDatabaseInfo{}
	for rows.Next() {
		var currTrackingDBInfo UserTrackingDatabaseInfo
		currTrackingDBInfo.IsAdmin = true
		if scanErr := rows.Scan(&currTrackingDBInfo.DatabaseID, &currTrackingDBInfo.DatabaseName); scanErr != nil {
			return nil, fmt.Errorf("getDatabaseDashboardsInfo: Failure querying database: %v", scanErr)
		}
		trackingInfoByDatabase[currTrackingDBInfo.DatabaseID] = currTrackingDBInfo
	}

	collabRows, collabQueryErr := databaseWrapper.DBHandle().Query(
		`SELECT databases.database_id, databases.name 
			FROM databases,collaborators  
			WHERE databases.database_id = collaborators.database_id
			AND collaborators.user_id = $1`, currUserID)
	if collabQueryErr != nil {
		return nil, fmt.Errorf("getCurrentUserTrackingDatabases: Failure querying database: %v", collabQueryErr)
	}
	for collabRows.Next() {
		var currTrackingDBInfo UserTrackingDatabaseInfo
		currTrackingDBInfo.IsAdmin = false
		if scanErr := collabRows.Scan(&currTrackingDBInfo.DatabaseID, &currTrackingDBInfo.DatabaseName); scanErr != nil {
			return nil, fmt.Errorf("getDatabaseDashboardsInfo: Failure querying database: %v", scanErr)
		}

		// Don't overwrite the tracking information if its already populated as an administrator
		if _, foundDBInfo := trackingInfoByDatabase[currTrackingDBInfo.DatabaseID]; foundDBInfo == false {
			trackingInfoByDatabase[currTrackingDBInfo.DatabaseID] = currTrackingDBInfo
		}
	}

	userTrackingDBInfo := []UserTrackingDatabaseInfo{}
	for _, currTrackingInfo := range trackingInfoByDatabase {
		userTrackingDBInfo = append(userTrackingDBInfo, currTrackingInfo)
	}

	return userTrackingDBInfo, nil

}
