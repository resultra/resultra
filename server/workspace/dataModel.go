package workspace

import (
	"database/sql"
	"fmt"
	"net/http"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/userAuth"
)

type WorkspaceInfo struct {
	Name       string              `json:"name"`
	Properties WorkspaceProperties `json:"properties"`
}

func GetWorkspaceInfo(trackerDBHandle *sql.DB) (*WorkspaceInfo, error) {

	rows, queryErr := trackerDBHandle.Query(`SELECT name,properties FROM workspace_info`)
	if queryErr != nil {
		return nil, fmt.Errorf("GetWorkspaceInfo: Failure querying database: %v", queryErr)
	}
	defer rows.Close()

	existingInfo := rows.Next()
	if existingInfo {

		var workspaceInfo WorkspaceInfo

		encodedProps := ""

		if scanErr := rows.Scan(&workspaceInfo.Name, &encodedProps); scanErr != nil {
			return nil, fmt.Errorf("GetWorkspaceInfo: failure querying database: %v", scanErr)
		}

		workspaceProps := newDefaultWorkspaceProperties()
		if decodeErr := generic.DecodeJSONString(encodedProps, &workspaceProps); decodeErr != nil {
			return nil, fmt.Errorf("getAlert: can't decode properties: %v", encodedProps)
		}
		workspaceInfo.Properties = workspaceProps

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

func setWorkspaceName(trackerDBHandle *sql.DB, newName string) error {
	if _, updateErr := trackerDBHandle.Exec(`UPDATE workspace_info SET name=$1`, newName); updateErr != nil {
		return fmt.Errorf("setWorkspaceName: Error updating name: error = %v", updateErr)
	}

	return nil
}

func updateWorkspaceProperties(trackerDBHandle *sql.DB, props WorkspaceProperties) error {

	encodedProps, encodeErr := generic.EncodeJSONString(props)
	if encodeErr != nil {
		return fmt.Errorf("updateWorkspaceProperties: failure encoding properties: error = %v", encodeErr)
	}

	if _, updateErr := trackerDBHandle.Exec(`UPDATE workspace_info SET properties=$1`, encodedProps); updateErr != nil {
		return fmt.Errorf("updateWorkspaceProperties: Error updating properties: error = %v", updateErr)
	}

	return nil
}

func CurrUserIsWorkspaceAdmin(req *http.Request) bool {

	userInfo, userInfoErr := userAuth.GetCurrentUserInfo(req)
	if userInfoErr != nil {
		return false
	}
	return userInfo.IsWorkspaceAdmin
}
