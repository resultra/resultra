package userRole

import (
	"fmt"
	"log"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/databaseWrapper"
	"resultra/datasheet/server/generic/uniqueID"
	"resultra/datasheet/server/generic/userAuth"
)

func AddDatabaseAdmin(databaseID string, userID string) error {
	// TODO verify the current user has permissions to add the user as an admin.

	if _, insertErr := databaseWrapper.DBHandle().Exec(
		`INSERT INTO database_admins (database_id,user_id) VALUES ($1,$2)`,
		databaseID, userID); insertErr != nil {
		return fmt.Errorf("addDatabaseAdmin: Can't add database admin user ID = %v to database with ID = %v: error = %v",
			userID, databaseID, insertErr)
	}

	return nil

}

type DatabaseRole struct {
	DatabaseID string `json:"databaseID"`
	RoleID     string `json:"roleID"`
	RoleName   string `json:"roleName"`
}

func addDatabaseRole(databaseID string, roleName string) (*DatabaseRole, error) {

	// TODO verify the current user has admin permissions to modify roles

	sanitizedRoleName, sanitizeErr := generic.SanitizeName(roleName)
	if sanitizeErr != nil {
		return nil, sanitizeErr
	}

	roleID := uniqueID.GenerateSnowflakeID()

	if _, insertErr := databaseWrapper.DBHandle().Exec(
		`INSERT INTO database_roles (database_id,role_id,name) VALUES ($1,$2,$3)`,
		databaseID, roleID, sanitizedRoleName); insertErr != nil {
		return nil, fmt.Errorf("addDatabaseRole: Can't add database role to database with ID = %v: error = %v",
			databaseID, insertErr)
	}

	dbRole := DatabaseRole{
		DatabaseID: databaseID,
		RoleID:     roleID,
		RoleName:   sanitizedRoleName}

	return &dbRole, nil

}

type NewDatabaseRoleWithPrivsParams struct {
	DatabaseID     string            `json:"databaseID"`
	RoleName       string            `json:"roleName"`
	FormPrivs      map[string]string `json:"formPrivs"`      // Map of form ID to privilege
	DashboardPrivs map[string]string `json:"dashboardPrivs"` // Map of dashboard ID to privilege
}

func newDatabaseRoleWithPrivs(params NewDatabaseRoleWithPrivsParams) error {
	log.Printf("newDatabaseRoleWithPrivs: %+v", params)

	// TODO Wrap all the database writes from this function into a transaction.

	newRole, newRoleErr := addDatabaseRole(params.DatabaseID, params.RoleName)
	if newRoleErr != nil {
		return fmt.Errorf("newDatabaseRoleWithPrivs: %v", newRoleErr)
	}

	for formID, priv := range params.FormPrivs {

		params := SetFormRolePrivsParams{
			FormID: formID,
			RoleID: newRole.RoleID,
			Privs:  priv}
		if formPrivErr := setFormRolePrivs(params); formPrivErr != nil {
			return fmt.Errorf("newDatabaseRoleWithPrivs: %v", formPrivErr)
		}
	}

	for dashboardID, priv := range params.DashboardPrivs {
		if dashboardPrivErr := setDashboardRolePrivs(newRole.RoleID, dashboardID, priv); dashboardPrivErr != nil {
			return fmt.Errorf("newDatabaseRoleWithPrivs: %v", dashboardPrivErr)
		}
	}

	return nil
}

func addUserRole(roleID string, userID string) error {
	// TODO verify the current user has permissions to add the user as an admin.

	if _, insertErr := databaseWrapper.DBHandle().Exec(
		`INSERT INTO user_roles (role_id,user_id) VALUES ($1,$2)`,
		roleID, userID); insertErr != nil {
		return fmt.Errorf("addUserRole: Can't add user with ID = %v to role with ID = %v: error = %v",
			userID, roleID, insertErr)
	}

	return nil

}

func GetDatabaseAdminUserInfo(databaseID string) ([]userAuth.UserInfo, error) {

	rows, queryErr := databaseWrapper.DBHandle().Query(
		`SELECT users.user_id,users.user_name,users.first_name,users.last_name 
				FROM database_admins,users
				WHERE database_admins.database_id=$1
				   AND database_admins.user_id=users.user_id`, databaseID)
	if queryErr != nil {
		return nil, fmt.Errorf("GetDatabaseAdminUserInfo: Failure querying database: %v", queryErr)
	}

	adminsInfo := []userAuth.UserInfo{}

	for rows.Next() {

		currAdmin := userAuth.UserInfo{}
		if scanErr := rows.Scan(&currAdmin.UserID, &currAdmin.UserName,
			&currAdmin.FirstName, &currAdmin.LastName); scanErr != nil {
			return nil, fmt.Errorf("GetDatabaseAdminUserInfo: Failure querying database: %v", scanErr)
		}
		adminsInfo = append(adminsInfo, currAdmin)
	}

	return adminsInfo, nil

}

type FormPrivInfo struct {
	FormID   string `json:"formID"`
	FormName string `json:"formName"`
	Privs    string `json:"privs"`
}

type CustomFormRoleInfo struct {
	RoleID    string         `json:"roleID"`
	RoleName  string         `json:"roleName"`
	FormPrivs []FormPrivInfo `json:"formPrivs"`
}

func GetCustomRoleFormInfo(databaseID string) ([]CustomFormRoleInfo, error) {

	rows, queryErr := databaseWrapper.DBHandle().Query(
		`SELECT database_roles.role_id,database_roles.name,
					forms.form_id,forms.name,form_role_privs.privs
				FROM form_role_privs,database_roles,forms
				WHERE database_roles.database_id=$1
				   AND database_roles.role_id=form_role_privs.role_id
				   AND form_role_privs.form_id=forms.form_id 
				ORDER BY database_roles.role_id`, databaseID)
	if queryErr != nil {
		return nil, fmt.Errorf("GetCustomRoleInfo: Failure querying database: %v", queryErr)
	}

	roleInfoMap := map[string]*CustomFormRoleInfo{}

	for rows.Next() {

		currFormPrivInfo := FormPrivInfo{}
		currRoleName := ""
		currRoleID := ""

		if scanErr := rows.Scan(&currRoleID, &currRoleName,
			&currFormPrivInfo.FormID, &currFormPrivInfo.FormName, &currFormPrivInfo.Privs); scanErr != nil {
			return nil, fmt.Errorf("GetCustomRoleInfo: Failure querying database: %v", scanErr)
		}

		var roleInfo *CustomFormRoleInfo
		if currRoleInfo, roleInfoFound := roleInfoMap[currRoleID]; !roleInfoFound {
			roleInfo = &CustomFormRoleInfo{
				RoleID:    currRoleID,
				RoleName:  currRoleName,
				FormPrivs: []FormPrivInfo{}}
			roleInfoMap[currRoleID] = roleInfo
		} else {
			roleInfo = currRoleInfo
		}

		roleInfo.FormPrivs = append(roleInfo.FormPrivs, currFormPrivInfo)

	}

	customRoleInfo := []CustomFormRoleInfo{}
	for _, currRoleInfo := range roleInfoMap {
		customRoleInfo = append(customRoleInfo, *currRoleInfo)
	}

	return customRoleInfo, nil

}

type DashboardPrivInfo struct {
	DashboardID   string `json:"dashboardID"`
	DashboardName string `json:"dashboardName"`
	Privs         string `json:"privs"`
}

type CustomRoleDashboardInfo struct {
	RoleID         string              `json:"roleID"`
	RoleName       string              `json:"roleName"`
	DashboardPrivs []DashboardPrivInfo `json:"dashboardPrivs"`
}

func GetCustomRoleDashboardInfo(databaseID string) ([]CustomRoleDashboardInfo, error) {

	rows, queryErr := databaseWrapper.DBHandle().Query(
		`SELECT database_roles.role_id,database_roles.name,
					dashboards.dashboard_id,dashboards.name,dashboard_role_privs.privs
				FROM dashboard_role_privs,database_roles,dashboards
				WHERE database_roles.database_id=$1
				   AND database_roles.role_id=dashboard_role_privs.role_id
				   AND dashboard_role_privs.dashboard_id=dashboards.dashboard_id 
				ORDER BY database_roles.role_id`, databaseID)
	if queryErr != nil {
		return nil, fmt.Errorf("GetCustomRoleInfo: Failure querying database: %v", queryErr)
	}

	roleInfoMap := map[string]*CustomRoleDashboardInfo{}

	for rows.Next() {

		currDashPrivInfo := DashboardPrivInfo{}
		currRoleName := ""
		currRoleID := ""

		if scanErr := rows.Scan(&currRoleID, &currRoleName,
			&currDashPrivInfo.DashboardID, &currDashPrivInfo.DashboardName, &currDashPrivInfo.Privs); scanErr != nil {
			return nil, fmt.Errorf("GetCustomRoleDashboardInfo: Failure querying database: %v", scanErr)
		}

		var roleInfo *CustomRoleDashboardInfo
		if currRoleInfo, roleInfoFound := roleInfoMap[currRoleID]; !roleInfoFound {
			roleInfo = &CustomRoleDashboardInfo{
				RoleID:         currRoleID,
				RoleName:       currRoleName,
				DashboardPrivs: []DashboardPrivInfo{}}
			roleInfoMap[currRoleID] = roleInfo
		} else {
			roleInfo = currRoleInfo
		}

		roleInfo.DashboardPrivs = append(roleInfo.DashboardPrivs, currDashPrivInfo)

	}

	customRoleInfo := []CustomRoleDashboardInfo{}
	for _, currRoleInfo := range roleInfoMap {
		customRoleInfo = append(customRoleInfo, *currRoleInfo)
	}

	return customRoleInfo, nil

}

type CustomRoleInfo struct {
	RoleID         string              `json:"roleID"`
	RoleName       string              `json:"roleName"`
	FormPrivs      []FormPrivInfo      `json:"formPrivs"`
	DashboardPrivs []DashboardPrivInfo `json:"dashboardPrivs"`
}

func GetCustomRoleInfo(databaseID string) ([]CustomRoleInfo, error) {

	roleInfoMap := map[string]*CustomRoleInfo{}

	customFormInfo, formErr := GetCustomRoleFormInfo(databaseID)
	if formErr != nil {
		return nil, fmt.Errorf("GetCustomRoleInfo: %v", formErr)
	}
	for _, formInfo := range customFormInfo {

		var roleInfo *CustomRoleInfo
		if currRoleInfo, roleInfoFound := roleInfoMap[formInfo.RoleID]; roleInfoFound {
			roleInfo = currRoleInfo
		} else {
			roleInfo = &CustomRoleInfo{
				RoleID:         formInfo.RoleID,
				RoleName:       formInfo.RoleName,
				FormPrivs:      []FormPrivInfo{},
				DashboardPrivs: []DashboardPrivInfo{}}
			roleInfoMap[formInfo.RoleID] = roleInfo
		}
		roleInfo.FormPrivs = formInfo.FormPrivs
	}

	customDashboardInfo, dashErr := GetCustomRoleDashboardInfo(databaseID)
	if dashErr != nil {
		return nil, fmt.Errorf("GetCustomRoleInfo: %v", dashErr)
	}
	for _, dashInfo := range customDashboardInfo {

		var roleInfo *CustomRoleInfo
		if currRoleInfo, roleInfoFound := roleInfoMap[dashInfo.RoleID]; roleInfoFound {
			roleInfo = currRoleInfo
		} else {
			roleInfo = &CustomRoleInfo{
				RoleID:         dashInfo.RoleID,
				RoleName:       dashInfo.RoleName,
				FormPrivs:      []FormPrivInfo{},
				DashboardPrivs: []DashboardPrivInfo{}}
			roleInfoMap[dashInfo.RoleID] = roleInfo
		}
		roleInfo.DashboardPrivs = dashInfo.DashboardPrivs

	}

	customRoleInfo := []CustomRoleInfo{}
	for _, roleInfo := range roleInfoMap {
		customRoleInfo = append(customRoleInfo, *roleInfo)
	}

	return customRoleInfo, nil

}
