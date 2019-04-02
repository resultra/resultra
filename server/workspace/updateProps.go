// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package workspace

import (
	"fmt"
	"net/http"
	"github.com/resultra/resultra/server/common/databaseWrapper"
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
