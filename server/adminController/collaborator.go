// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package adminController

import (
	"database/sql"
	"fmt"
	"resultra/tracker/server/userRole"
)

type AddCollaboratorParams struct {
	DatabaseID string   `json:"databaseID"`
	UserID     string   `json:"userID"`
	RoleIDs    []string `json:"roleIDs"`
}

func addCollaborator(trackerDBHandle *sql.DB, params AddCollaboratorParams) (*UserRoleInfo, error) {

	collabInfo, err := userRole.AddNonAdminCollaborator(trackerDBHandle, params.DatabaseID, params.UserID)
	if err != nil {
		return nil, fmt.Errorf("addCollaborator: %v", err)
	}

	for _, currRoleID := range params.RoleIDs {
		if err := userRole.AddCollaboratorRole(trackerDBHandle, currRoleID, collabInfo.CollaboratorID); err != nil {
			return nil, fmt.Errorf("addCollaborator: Error adding role for collaborator: %v", err)
		}
	}

	collabUserRoleInfo, err := userRole.GetCollaboratorRoleInfo(trackerDBHandle, params.DatabaseID, collabInfo.CollaboratorID)
	if err != nil {
		return nil, fmt.Errorf("addCollaborator: Error adding updated role info for collaborator: %v", err)
	}

	userRoleInfo := UserRoleInfo{
		UserInfo:       collabUserRoleInfo.UserInfo,
		IsAdmin:        false,
		CollaboratorID: collabInfo.CollaboratorID,
		CustomRoles:    collabUserRoleInfo.RoleInfo}

	return &userRoleInfo, nil
}
