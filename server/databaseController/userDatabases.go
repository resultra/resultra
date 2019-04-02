// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package databaseController

import (
	"fmt"
	"net/http"

	"database/sql"
	"github.com/resultra/resultra/server/common/userAuth"
)

type UserTrackingDatabaseInfo struct {
	DatabaseID   string `json:"databaseID"`
	DatabaseName string `json:"databaseName"`
	IsAdmin      bool   `json:"isAdmin"`
	IsActive     bool   `json:"isActive"`
}

type GetTrackerListParams struct {
	IncludeInactive bool `json:"includeInactive"`
}

func getCurrentUserTrackingDatabases(params GetTrackerListParams,
	trackerDBHandle *sql.DB, req *http.Request) ([]UserTrackingDatabaseInfo, error) {

	currUserID, userErr := userAuth.GetCurrentUserID(req)
	if userErr != nil {
		return nil, fmt.Errorf("getCurrentUserTrackingDatabases: can't get current user: %v", userErr)
	}

	rows, queryErr := trackerDBHandle.Query(
		`SELECT databases.database_id, databases.name,databases.is_active FROM collaborators,databases WHERE 
			collaborators.user_id=$1 AND 
			collaborators.database_id = databases.database_id`, currUserID)
	if queryErr != nil {
		return nil, fmt.Errorf("getCurrentUserTrackingDatabases: Failure querying database: %v", queryErr)
	}
	defer rows.Close()

	trackingInfoByDatabase := map[string]UserTrackingDatabaseInfo{}
	for rows.Next() {
		var currTrackingDBInfo UserTrackingDatabaseInfo
		currTrackingDBInfo.IsAdmin = true
		if scanErr := rows.Scan(&currTrackingDBInfo.DatabaseID, &currTrackingDBInfo.DatabaseName,
			&currTrackingDBInfo.IsActive); scanErr != nil {
			return nil, fmt.Errorf("getDatabaseDashboardsInfo: Failure querying database: %v", scanErr)
		}
		trackingInfoByDatabase[currTrackingDBInfo.DatabaseID] = currTrackingDBInfo
	}

	collabRows, collabQueryErr := trackerDBHandle.Query(
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

		if currTrackingInfo.IsActive {
			userTrackingDBInfo = append(userTrackingDBInfo, currTrackingInfo)
		} else {
			// Only include inactive trackers if the current user is also the admin
			if params.IncludeInactive && currTrackingInfo.IsAdmin {
				userTrackingDBInfo = append(userTrackingDBInfo, currTrackingInfo)
			}
		}
	}

	return userTrackingDBInfo, nil

}
