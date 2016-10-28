package adminController

import (
	"fmt"
	"resultra/datasheet/server/userRole"
)

type AddCollaboratorParams struct {
	DatabaseID string   `json:"databaseID"`
	UserID     string   `json:"userID"`
	RoleIDs    []string `json:"roleIDs"`
}

func addCollaborator(params AddCollaboratorParams) (*UserRoleInfo, error) {

	for _, currRoleID := range params.RoleIDs {
		if err := userRole.AddUserRole(currRoleID, params.UserID); err != nil {
			return nil, fmt.Errorf("addCollaborator: Error adding role for collaborator: %v", err)
		}
	}

	collabUserRoleInfo, err := userRole.GetUserRoleInfo(params.DatabaseID, params.UserID)
	if err != nil {
		return nil, fmt.Errorf("addCollaborator: Error adding updated role info for collaborator: %v", err)
	}

	userRoleInfo := UserRoleInfo{
		UserInfo:    collabUserRoleInfo.UserInfo,
		IsAdmin:     false,
		CustomRoles: collabUserRoleInfo.RoleInfo}

	return &userRoleInfo, nil
}
