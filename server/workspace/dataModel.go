package workspace

import (
	"database/sql"
	"fmt"
	"net/http"
	"resultra/datasheet/server/generic/userAuth"
)

type WorkspaceInfo struct {
	Name string `json:"name"`
}

func GetWorkspaceInfo(trackerDBHandle *sql.DB) (*WorkspaceInfo, error) {

	rows, queryErr := trackerDBHandle.Query(`SELECT name FROM workspace_info`)
	if queryErr != nil {
		return nil, fmt.Errorf("GetWorkspaceInfo: Failure querying database: %v", queryErr)
	}
	defer rows.Close()

	existingInfo := rows.Next()
	if existingInfo {

		var workspaceInfo WorkspaceInfo

		if scanErr := rows.Scan(&workspaceInfo.Name); scanErr != nil {
			return nil, fmt.Errorf("GetWorkspaceInfo: failure querying database: %v", scanErr)
		}

		return &workspaceInfo, nil
	}

	return nil, nil

}

const defaultWorkspaceName string = "Trackers"

func GetWorkspaceName(trackerDBHandle *sql.DB) (string, error) {

	workspaceInfo, err := GetWorkspaceInfo(trackerDBHandle)

	if err != nil {
		return "", fmt.Errorf("GetWorkspaceName: %v", err)
	}

	if workspaceInfo == nil {
		return defaultWorkspaceName, nil
	}

	return workspaceInfo.Name, nil

}

func CurrUserIsWorkspaceAdmin(req *http.Request) bool {

	userInfo, userInfoErr := userAuth.GetCurrentUserInfo(req)
	if userInfoErr != nil {
		return false
	}
	return userInfo.IsWorkspaceAdmin
}
