package adminController

import (
	"fmt"
	"resultra/datasheet/server/generic/userAuth"
	"resultra/datasheet/server/userRole"
	//	"resultra/datasheet/server/generic/databaseWrapper"
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
func getUserRolesInfo(params GetUserRolesParams) ([]UserRoleInfo, error) {

	adminUserInfo, err := userRole.GetDatabaseAdminUserInfo(params.DatabaseID)
	if err != nil {
		return nil, fmt.Errorf("getUserRolesInfo: %v", err)
	}

	userRoles := []UserRoleInfo{}
	for _, currUserInfo := range adminUserInfo {
		currAdminInfo := UserRoleInfo{UserInfo: currUserInfo, IsAdmin: true}
		userRoles = append(userRoles, currAdminInfo)
	}

	usersRoleInfo, err := userRole.GetAllUsersRoleInfo(params.DatabaseID)
	if err != nil {
		return nil, fmt.Errorf("getUserRolesInfo: %v", err)
	}
	for _, currRoleInfo := range usersRoleInfo {
		currInfo := UserRoleInfo{
			UserInfo:       currRoleInfo.UserInfo,
			IsAdmin:        false,
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

func getRoleInfo(params GetRoleInfoParams) (*RoleInfo, error) {
	adminUserInfo, err := userRole.GetDatabaseAdminUserInfo(params.DatabaseID)
	if err != nil {
		return nil, fmt.Errorf("getUserRolesInfo: %v", err)
	}

	customRoleInfo, err := userRole.GetCustomRoleInfo(params.DatabaseID)
	if err != nil {
		return nil, fmt.Errorf("getUserRolesInfo: %v", err)
	}

	roleInfo := RoleInfo{
		AdminUsers:  adminUserInfo,
		CustomRoles: customRoleInfo}

	return &roleInfo, nil
}
