package userRole

import (
	"fmt"
	"log"
	"resultra/datasheet/server/generic/databaseWrapper"
	"resultra/datasheet/server/generic/stringValidation"
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

	sanitizedRoleName, sanitizeErr := stringValidation.SanitizeName(roleName)
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

type DatabaseRoleInfo struct {
	ParentDatabaseID string `json:"parentDatabaseID"`
	RoleID           string `json:"roleID"`
	RoleName         string `json:"roleName"`
}

type DatabaseRolesParams struct {
	DatabaseID string `json:"databaseID"`
}

func getDatabaseRoles(databaseID string) ([]DatabaseRoleInfo, error) {

	rows, queryErr := databaseWrapper.DBHandle().Query(
		`SELECT role_id, name FROM database_roles
				WHERE database_id=$1`, databaseID)
	if queryErr != nil {
		return nil, fmt.Errorf("GetDatabaseRoles: Failure querying database: %v", queryErr)
	}

	rolesInfo := []DatabaseRoleInfo{}
	for rows.Next() {

		currRoleInfo := DatabaseRoleInfo{}
		if scanErr := rows.Scan(&currRoleInfo.RoleID, &currRoleInfo.RoleName); scanErr != nil {
			return nil, fmt.Errorf("GetDatabaseRoles: Failure querying database: %v", scanErr)
		}
		currRoleInfo.ParentDatabaseID = databaseID
		rolesInfo = append(rolesInfo, currRoleInfo)
	}

	return rolesInfo, nil

}

func GetUserRole(roleID string) (*DatabaseRoleInfo, error) {
	roleInfo := DatabaseRoleInfo{}
	getErr := databaseWrapper.DBHandle().QueryRow(`SELECT database_id,role_id, name FROM database_roles
		 WHERE role_id=$1 LIMIT 1`, roleID).Scan(&roleInfo.ParentDatabaseID,
		&roleInfo.RoleID, &roleInfo.RoleName)
	if getErr != nil {
		return nil, fmt.Errorf("GetUserRole: Unabled to get role: role ID = %v: datastore err=%v",
			roleID, getErr)
	}

	return &roleInfo, nil

}

func updateExistingRole(roleID string, updatedRole *DatabaseRoleInfo) (*DatabaseRoleInfo, error) {

	if _, updateErr := databaseWrapper.DBHandle().Exec(`UPDATE database_roles 
				SET name=$1
				WHERE role_id=$2`, updatedRole.RoleName, roleID); updateErr != nil {
		return nil, fmt.Errorf("updateExistingRole: Can't update role properties %v: error = %v",
			roleID, updateErr)
	}

	return updatedRole, nil

}

type NewDatabaseRoleWithPrivsParams struct {
	DatabaseID     string            `json:"databaseID"`
	RoleName       string            `json:"roleName"`
	ListPrivs      map[string]string `json:"listPrivs"`      // Map of list ID to privilege
	DashboardPrivs map[string]string `json:"dashboardPrivs"` // Map of dashboard ID to privilege
}

func newDatabaseRoleWithPrivs(params NewDatabaseRoleWithPrivsParams) error {
	log.Printf("newDatabaseRoleWithPrivs: %+v", params)

	// TODO Wrap all the database writes from this function into a transaction.

	newRole, newRoleErr := addDatabaseRole(params.DatabaseID, params.RoleName)
	if newRoleErr != nil {
		return fmt.Errorf("newDatabaseRoleWithPrivs: %v", newRoleErr)
	}

	for listID, priv := range params.ListPrivs {

		params := SetListRolePrivsParams{
			ListID: listID,
			RoleID: newRole.RoleID,
			Privs:  priv}
		if listPrivErr := setListRolePrivs(params); listPrivErr != nil {
			return fmt.Errorf("newDatabaseRoleWithPrivs: %v", listPrivErr)
		}
	}

	for dashboardID, priv := range params.DashboardPrivs {

		setPrivParams := SetDashboardRolePrivsParams{
			DashboardID: dashboardID,
			RoleID:      newRole.RoleID,
			Privs:       priv}

		if dashboardPrivErr := setDashboardRolePrivs(setPrivParams); dashboardPrivErr != nil {
			return fmt.Errorf("newDatabaseRoleWithPrivs: %v", dashboardPrivErr)
		}
	}

	return nil
}

func AddUserRole(roleID string, userID string) error {
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

type ListPrivInfo struct {
	ListID   string `json:"listID"`
	ListName string `json:"listName"`
	Privs    string `json:"privs"`
}

type CustomListRoleInfo struct {
	RoleID    string         `json:"roleID"`
	RoleName  string         `json:"roleName"`
	ListPrivs []ListPrivInfo `json:"listPrivs"`
}

func GetCustomRoleListInfo(databaseID string) ([]CustomListRoleInfo, error) {

	rows, queryErr := databaseWrapper.DBHandle().Query(
		`SELECT database_roles.role_id,database_roles.name,
					item_lists.list_id,item_lists.name,list_role_privs.privs
				FROM list_role_privs,database_roles,item_lists
				WHERE database_roles.database_id=$1
				   AND database_roles.role_id=list_role_privs.role_id
				   AND list_role_privs.list_id=item_lists.list_id 
				ORDER BY database_roles.role_id`, databaseID)
	if queryErr != nil {
		return nil, fmt.Errorf("GetCustomRoleListInfo: Failure querying database: %v", queryErr)
	}

	roleInfoMap := map[string]*CustomListRoleInfo{}

	for rows.Next() {

		currListPrivInfo := ListPrivInfo{}
		currRoleName := ""
		currRoleID := ""

		if scanErr := rows.Scan(&currRoleID, &currRoleName,
			&currListPrivInfo.ListID, &currListPrivInfo.ListName, &currListPrivInfo.Privs); scanErr != nil {
			return nil, fmt.Errorf("GetCustomRoleInfo: Failure querying database: %v", scanErr)
		}

		var roleInfo *CustomListRoleInfo
		if currRoleInfo, roleInfoFound := roleInfoMap[currRoleID]; !roleInfoFound {
			roleInfo = &CustomListRoleInfo{
				RoleID:    currRoleID,
				RoleName:  currRoleName,
				ListPrivs: []ListPrivInfo{}}
			roleInfoMap[currRoleID] = roleInfo
		} else {
			roleInfo = currRoleInfo
		}

		roleInfo.ListPrivs = append(roleInfo.ListPrivs, currListPrivInfo)

	}

	customRoleInfo := []CustomListRoleInfo{}
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
	RoleUsers      []userAuth.UserInfo `json:"roleUsers"`
	ListPrivs      []ListPrivInfo      `json:"listPrivs"`
	DashboardPrivs []DashboardPrivInfo `json:"dashboardPrivs"`
}

func GetCustomRoleInfo(databaseID string) ([]CustomRoleInfo, error) {

	roleInfoMap := map[string]*CustomRoleInfo{}

	getOrAllocRoleInfo := func(roleID string, roleName string) *CustomRoleInfo {
		var roleInfo *CustomRoleInfo
		if currRoleInfo, roleInfoFound := roleInfoMap[roleID]; roleInfoFound {
			roleInfo = currRoleInfo
		} else {
			roleInfo = &CustomRoleInfo{
				RoleID:         roleID,
				RoleName:       roleName,
				RoleUsers:      []userAuth.UserInfo{},
				ListPrivs:      []ListPrivInfo{},
				DashboardPrivs: []DashboardPrivInfo{}}
			roleInfoMap[roleID] = roleInfo
		}
		return roleInfo

	}

	customListInfo, listErr := GetCustomRoleListInfo(databaseID)
	if listErr != nil {
		return nil, fmt.Errorf("GetCustomRoleInfo: %v", listErr)
	}
	for _, listInfo := range customListInfo {

		roleInfo := getOrAllocRoleInfo(listInfo.RoleID, listInfo.RoleName)
		roleInfo.ListPrivs = listInfo.ListPrivs
	}

	customDashboardInfo, dashErr := GetCustomRoleDashboardInfo(databaseID)
	if dashErr != nil {
		return nil, fmt.Errorf("GetCustomRoleInfo: %v", dashErr)
	}
	for _, dashInfo := range customDashboardInfo {

		roleInfo := getOrAllocRoleInfo(dashInfo.RoleID, dashInfo.RoleName)
		roleInfo.DashboardPrivs = dashInfo.DashboardPrivs

	}

	usersRoleInfo, err := GetAllUsersRoleInfo(databaseID)
	if err != nil {
		return nil, fmt.Errorf("GetCustomRoleInfo: %v", err)
	}
	for _, currUserRoleInfo := range usersRoleInfo {
		for _, currRoleInfo := range currUserRoleInfo.RoleInfo {
			roleInfo := getOrAllocRoleInfo(currRoleInfo.RoleID, currRoleInfo.RoleName)
			roleInfo.RoleUsers = append(roleInfo.RoleUsers, currUserRoleInfo.UserInfo)
		}
	}

	// Flatten information from a map to a list
	customRoleInfo := []CustomRoleInfo{}
	for _, roleInfo := range roleInfoMap {
		customRoleInfo = append(customRoleInfo, *roleInfo)
	}

	return customRoleInfo, nil

}

type UserRoleInfo struct {
	UserInfo userAuth.UserInfo  `json:"userInfo"`
	RoleInfo []DatabaseRoleInfo `json:"roleInfo"`
}

func GetUserRoleInfo(databaseID string, userID string) (*UserRoleInfo, error) {

	rows, queryErr := databaseWrapper.DBHandle().Query(
		`SELECT user_roles.user_id, database_roles.role_id,database_roles.name
				FROM database_roles,user_roles
				WHERE database_roles.database_id=$1
				   AND database_roles.role_id=user_roles.role_id
				   AND user_roles.user_id=$2
				   ORDER BY user_roles.user_id`, databaseID, userID)
	if queryErr != nil {
		return nil, fmt.Errorf("GetUserRoleInfo: Failure querying database: %v", queryErr)
	}

	userInfo, err := userAuth.GetUserInfoByID(userID)
	if err != nil {
		return nil, fmt.Errorf("GetUserRoleInfo: Failure getting user info: userID=%v, error = %v", userID, err)
	}

	userRoleInfo := UserRoleInfo{UserInfo: *userInfo, RoleInfo: []DatabaseRoleInfo{}}

	for rows.Next() {

		var currUserID string
		var currRoleInfo DatabaseRoleInfo

		if scanErr := rows.Scan(&currUserID, &currRoleInfo.RoleID, &currRoleInfo.RoleName); scanErr != nil {
			return nil, fmt.Errorf("GetAllUsersRoleInfo: Failure querying database: %v", scanErr)
		}
		userRoleInfo.RoleInfo = append(userRoleInfo.RoleInfo, currRoleInfo)
	}

	return &userRoleInfo, nil

}

// Aggregate the role information by user.
func GetAllUsersRoleInfo(databaseID string) ([]UserRoleInfo, error) {

	rows, queryErr := databaseWrapper.DBHandle().Query(
		`SELECT user_roles.user_id, database_roles.role_id,database_roles.name
				FROM database_roles,user_roles
				WHERE database_roles.database_id=$1
				   AND database_roles.role_id=user_roles.role_id
				   ORDER BY user_roles.user_id`, databaseID)
	if queryErr != nil {
		return nil, fmt.Errorf("GetAllUsersRoleInfo: Failure querying database: %v", queryErr)
	}

	roleInfoByUserID := map[string]*UserRoleInfo{}

	for rows.Next() {

		var currUserID string
		var currRoleInfo DatabaseRoleInfo

		if scanErr := rows.Scan(&currUserID, &currRoleInfo.RoleID, &currRoleInfo.RoleName); scanErr != nil {
			return nil, fmt.Errorf("GetAllUsersRoleInfo: Failure querying database: %v", scanErr)
		}

		var userRoleInfo *UserRoleInfo
		userRoleInfo, foundInfo := roleInfoByUserID[currUserID]
		if !foundInfo {
			userInfo, err := userAuth.GetUserInfoByID(currUserID)
			if err != nil {
				return nil, fmt.Errorf("GetAllUsersRoleInfo: Failure getting user info: userID=%v, error = %v", currUserID, err)
			}

			userRoleInfo = &UserRoleInfo{UserInfo: *userInfo, RoleInfo: []DatabaseRoleInfo{}}
			roleInfoByUserID[currUserID] = userRoleInfo
		}

		userRoleInfo.RoleInfo = append(userRoleInfo.RoleInfo, currRoleInfo)

	}

	usersRoleInfo := []UserRoleInfo{}
	for _, currRoleInfo := range roleInfoByUserID {
		usersRoleInfo = append(usersRoleInfo, *currRoleInfo)
	}

	return usersRoleInfo, nil
}
