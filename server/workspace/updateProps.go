package workspace

import (
	"fmt"
	"net/http"
	"resultra/tracker/server/common/databaseWrapper"
)

type WorkspacePropUpdater interface {
	updateProps(workspaceInfo *WorkspaceInfo) error
}

func updateWorkspaceProps(r *http.Request, propUpdater WorkspacePropUpdater) (*WorkspaceInfo, error) {

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		return nil, fmt.Errorf("updateWorkspaceProps: Unable to get existing workspace information: %v", dbErr)
	}

	workspaceInfoForUpdate, getErr := GetWorkspaceInfo(trackerDBHandle)
	if getErr != nil {
		return nil, fmt.Errorf("updateWorkspaceProps: Unable to get existing workspace information: %v", getErr)
	}

	if propUpdateErr := propUpdater.updateProps(workspaceInfoForUpdate); propUpdateErr != nil {
		return nil, fmt.Errorf("updateWorkspaceProps: Unable to update existing workspace properties: %v", propUpdateErr)
	}

	updateErr := updateWorkspaceProperties(trackerDBHandle, workspaceInfoForUpdate.Properties)
	if updateErr != nil {
		return nil, fmt.Errorf("updateWorkspaceProps: Unable to update existing gauge indicator properties: datastore update error =  %v", updateErr)
	}

	return workspaceInfoForUpdate, nil
}

type AllowRegistrationParams struct {
	AllowUserRegistration bool `json:"allowUserRegistration"`
}

func (updateParams AllowRegistrationParams) updateProps(workspaceInfo *WorkspaceInfo) error {

	workspaceInfo.Properties.AllowUserRegistration = updateParams.AllowUserRegistration

	return nil
}
