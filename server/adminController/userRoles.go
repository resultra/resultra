package adminController

import (
	"database/sql"
	"fmt"
	"resultra/tracker/server/common/userAuth"
	"resultra/tracker/server/userRole"
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
