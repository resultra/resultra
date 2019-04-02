// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package adminController

import (
	"database/sql"
	"fmt"
	"github.com/resultra/resultra/server/common/userAuth"
	"github.com/resultra/resultra/server/userRole"
)

type UserRoleInfo struct {
	UserInfo       userAuth.UserInfo           `json:"userInfo"`
	IsAdmin        bool                        `json:"isAdmin"`
	CollaboratorID string                      `json:"collaboratorID"`
	CustomRoles    []userRole.DatabaseRoleInfo `json:"customRoles"`
	RoleMembership map[string]bool             `json:"roleMembership"`
}

type GetUserRolesParams struct {
	DatabaseID string `json:"databaseID"`
}

// get a list of users and their roles.
func getUserRolesInfo(trackerDBHandle *sql.DB, params GetUserRolesParams) ([]UserRoleInfo, error) {

	userRoles := []UserRoleInfo{}
	usersRoleInfo, err := userRole.GetAllUsersRoleInfo(trackerDBHandle, params.DatabaseID)
	if err != nil {
		return nil, fmt.Errorf("getUserRolesInfo: %v", err)
	}
	for _, currRoleInfo := range usersRoleInfo {
		currInfo := UserRoleInfo{
			UserInfo:       currRoleInfo.UserInfo,
			IsAdmin:        currRoleInfo.IsAdmin,
			CollaboratorID: currRoleInfo.CollaboratorID,
			CustomRoles:    currRoleInfo.RoleInfo}
		userRoles = append(userRoles, currInfo)
	}

	return userRoles, nil
}

type RoleInfo struct {
	AdminUsers  []userAuth.UserInfo       `json:"adminUsers"`
	CustomRoles []userRole.CustomRoleInfo `json:"customRoles"`
}

type GetRoleInfoParams GetUserRolesParams

func getRoleInfo(trackerDBHandle *sql.DB, params GetRoleInfoParams) (*RoleInfo, error) {

	adminUserInfo, err := userRole.GetDatabaseAdminUserInfo(trackerDBHandle, params.DatabaseID)
	if err != nil {
		return nil, fmt.Errorf("getUserRolesInfo: %v", err)
	}

	customRoleInfo, err := userRole.GetCustomRoleInfo(trackerDBHandle, params.DatabaseID)
	if err != nil {
		return nil, fmt.Errorf("getUserRolesInfo: %v", err)
	}

	roleInfo := RoleInfo{
		AdminUsers:  adminUserInfo,
		CustomRoles: customRoleInfo}

	return &roleInfo, nil
}
