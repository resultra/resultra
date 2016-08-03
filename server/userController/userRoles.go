package userController

import (
	"fmt"
	"resultra/datasheet/server/generic/userAuth"
	"resultra/datasheet/server/userRole"
	//	"resultra/datasheet/server/generic/databaseWrapper"
)

type CustomUserRolesInfo struct {
	RoleID   string `json:"roleID"`
	RoleName string `json:"roleName"`
}

type UserRoleInfo struct {
	UserInfo    userAuth.UserInfo     `json:"userInfo"`
	IsAdmin     bool                  `json:"isAdmin"`
	CustomRoles []CustomUserRolesInfo `json:"customRoles"`
}

type GetUserRolesParams struct {
	DatabaseID string `json:"databaseID"`
}

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

	return userRoles, nil
}
