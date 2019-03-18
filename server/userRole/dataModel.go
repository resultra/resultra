package userRole

import (
	"database/sql"
	"fmt"
	"log"
	"resultra/tracker/server/common/userAuth"
	"resultra/tracker/server/generic/stringValidation"
	"resultra/tracker/server/generic/uniqueID"
	"resultra/tracker/server/trackerDatabase"
)

func AddDatabaseAdmin(trackerDBHandle *sql.DB, databaseID string, userID string) error {

	// TODO verify the current user has permissions to add the user as an admin.
	// TODO add verification to only allow a single database_id,user_id pair

	_, err := AddCollaborator(trackerDBHandle, databaseID, userID, true)
	if err != nil {
		return fmt.Errorf("addDatabaseAdmin: Can't add database admin user ID = %v to database with ID = %v: error = %v",
			userID, databaseID, err)
	}

	return nil

}

func AddCollaborator(trackerDBHandle *sql.DB, databaseID string, userID string, isAdmin bool) (*CollaboratorInfo, error) {

	if existingCollab, err := GetCollaborator(trackerDBHandle, databaseID, userID); err == nil {
		return existingCollab, nil
	}

	collabID := uniqueID.GenerateUniqueID()

	if _, insertErr := trackerDBHandle.Exec(
		`INSERT INTO collaborators (collaborator_id,database_id,user_id,is_admin) VALUES ($1,$2,$3,$4)`,
		collabID, databaseID, userID, isAdmin); insertErr != nil {
		return nil, fmt.Errorf("AddCollaborator: Can't add collaborator with user ID = %v to database with ID = %v: error = %v",
			userID, databaseID, insertErr)
	}

	newCollabInfo := &CollaboratorInfo{
		CollaboratorID: collabID,
		UserID:         userID,
		DatabaseID:     databaseID,
		IsAdmin:        isAdmin}

	return newCollabInfo, nil

}

func AddNonAdminCollaborator(trackerDBHandle *sql.DB, databaseID string, userID string) (*CollaboratorInfo, error) {
	return AddCollaborator(trackerDBHandle, databaseID, userID, false)
}

type CollaboratorInfo struct {
	CollaboratorID string
	UserID         string
	DatabaseID     string
	IsAdmin        bool
}

func GetCollaborator(trackerDBHandle *sql.DB, databaseID string, userID string) (*CollaboratorInfo, error) {

	collabID := ""
	isAdmin := false
	getErr := trackerDBHandle.QueryRow(`SELECT collaborator_id,is_admin FROM collaborators
		 WHERE user_id=$1 AND database_id=$2 LIMIT 1`, databaseID, userID).Scan(&collabID, &isAdmin)
	if getErr != nil {
		return nil, fmt.Errorf("getCollaborator: Unabled to get collaborator with databaseID = %v, userID = %v: datastore err=%v",
			databaseID, userID, getErr)
	}

	collabInfo := CollaboratorInfo{
		CollaboratorID: collabID,
		UserID:         userID,
		DatabaseID:     databaseID,
		IsAdmin:        isAdmin}

	return &collabInfo, nil

}

func getAllCollaborators(trackerDBHandle *sql.DB, databaseID string) ([]CollaboratorInfo, error) {

	rows, queryErr := trackerDBHandle.Query(
		`SELECT collaborator_id, user_id,is_admin FROM collaborators
				WHERE database_id=$1`, databaseID)
	if queryErr != nil {
		return nil, fmt.Errorf("GetDatabaseRoles: Failure querying database: %v", queryErr)
	}
	defer rows.Close()

	collabs := []CollaboratorInfo{}
	for rows.Next() {

		currCollabInfo := CollaboratorInfo{}
		isAdmin := false
		if scanErr := rows.Scan(&currCollabInfo.CollaboratorID, &currCollabInfo.UserID, &isAdmin); scanErr != nil {
			return nil, fmt.Errorf("GetDatabaseRoles: Failure querying database: %v", scanErr)
		}
		currCollabInfo.DatabaseID = databaseID
		currCollabInfo.IsAdmin = isAdmin
		collabs = append(collabs, currCollabInfo)
	}

	return collabs, nil

}

type GetAllCollaborUserInfoParams struct {
	DatabaseID string `json:"databaseID"`
}

func GetAllCollaboratorUserInfo(trackerDBHandle *sql.DB, params GetAllCollaborUserInfoParams) ([]userAuth.UserInfo, error) {

	collabInfo, err := getAllCollaborators(trackerDBHandle, params.DatabaseID)
	if err != nil {
		return nil, fmt.Errorf("getAllCollaboratorUserInfo: %v", err)
	}

	collabUserInfo := []userAuth.UserInfo{}
	for _, currCollab := range collabInfo {

		// TODo - Put a cache behind userAuth.GetUserInfoByID
		userInfo, err := userAuth.GetUserInfoByID(trackerDBHandle, currCollab.UserID)
		if err != nil {
			return nil, fmt.Errorf("GetAllUsersRoleInfo: Failure getting user info: userID=%v, error = %v", currCollab.UserID, err)
		}
		collabUserInfo = append(collabUserInfo, *userInfo)
	}
	return collabUserInfo, nil

}

func GetCollaboratorByID(trackerDBHandle *sql.DB, collaboratorID string) (*CollaboratorInfo, error) {

	userID := ""
	databaseID := ""
	isAdmin := false
	getErr := trackerDBHandle.QueryRow(`SELECT user_id,database_id,is_admin FROM collaborators
		 WHERE collaborator_id=$1 LIMIT 1`, collaboratorID).Scan(&userID, &databaseID, &isAdmin)
	if getErr != nil {
		return nil, fmt.Errorf("getCollaboratorByID: Unabled to get collaborator with ID = %v: datastore err=%v",
			collaboratorID, getErr)
	}

	collabInfo := CollaboratorInfo{
		CollaboratorID: collaboratorID,
		UserID:         userID,
		DatabaseID:     databaseID,
		IsAdmin:        isAdmin}

	return &collabInfo, nil

}

type DatabaseRole struct {
	DatabaseID string `json:"databaseID"`
	RoleID     string `json:"roleID"`
	RoleName   string `json:"roleName"`
}

func saveDatabaseRole(destDBHandle *sql.DB, role DatabaseRole) (*DatabaseRole, error) {

	if _, insertErr := destDBHandle.Exec(
		`INSERT INTO database_roles (database_id,role_id,name) VALUES ($1,$2,$3)`,
		role.DatabaseID, role.RoleID, role.RoleName); insertErr != nil {
		return nil, fmt.Errorf("addDatabaseRole: Can't add database role to database with ID = %v: error = %v",
			role.DatabaseID, insertErr)
	}
	return &role, nil

}

func addDatabaseRole(destDBHandle *sql.DB, databaseID string, roleName string) (*DatabaseRole, error) {

	// TODO verify the current user has admin permissions to modify roles

	sanitizedRoleName, sanitizeErr := stringValidation.SanitizeName(roleName)
	if sanitizeErr != nil {
		return nil, sanitizeErr
	}

	roleID := uniqueID.GenerateUniqueID()

	dbRole := DatabaseRole{
		DatabaseID: databaseID,
		RoleID:     roleID,
		RoleName:   sanitizedRoleName}

	return saveDatabaseRole(destDBHandle, dbRole)

}

type DatabaseRoleInfo struct {
	ParentDatabaseID string `json:"parentDatabaseID"`
	RoleID           string `json:"roleID"`
	RoleName         string `json:"roleName"`
}

type DatabaseRolesParams struct {
	DatabaseID string `json:"databaseID"`
}

func getDatabaseRolesFromSrc(srcDBHandle *sql.DB, databaseID string) ([]DatabaseRoleInfo, error) {

	rows, queryErr := srcDBHandle.Query(
		`SELECT role_id, name FROM database_roles
				WHERE database_id=$1`, databaseID)
	if queryErr != nil {
		return nil, fmt.Errorf("GetDatabaseRoles: Failure querying database: %v", queryErr)
	}
	defer rows.Close()

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

func GetDatabaseRoles(trackerDBHandle *sql.DB, databaseID string) ([]DatabaseRoleInfo, error) {
	return getDatabaseRolesFromSrc(trackerDBHandle, databaseID)
}

func CloneRoles(cloneParams *trackerDatabase.CloneDatabaseParams) error {

	destDatabaseID, err := cloneParams.IDRemapper.GetExistingRemappedID(cloneParams.SourceDatabaseID)
	if err != nil {
		return fmt.Errorf("CloneGlobals: Unable to get mapped ID for source database: %v", err)
	}

	roles, err := getDatabaseRolesFromSrc(cloneParams.SrcDBHandle, cloneParams.SourceDatabaseID)
	if err != nil {
		return fmt.Errorf("CloneRoles: Unable to retrieve roles: databaseID=%v, error=%v ",
			cloneParams.SourceDatabaseID, err)
	}
	for _, currRole := range roles {

		remappedID := uniqueID.GenerateUniqueID()
		cloneParams.IDRemapper[currRole.RoleID] = remappedID

		destRole := DatabaseRole{
			DatabaseID: destDatabaseID,
			RoleID:     remappedID,
			RoleName:   currRole.RoleName}

		if _, err := saveDatabaseRole(cloneParams.DestDBHandle, destRole); err != nil {
			return fmt.Errorf("CloneRoles: Can't clone role: error = %v", err)
		}
	}

	return nil

}

func GetUserRole(trackerDBHandle *sql.DB, roleID string) (*DatabaseRoleInfo, error) {
	roleInfo := DatabaseRoleInfo{}
	getErr := trackerDBHandle.QueryRow(`SELECT database_id,role_id, name FROM database_roles
		 WHERE role_id=$1 LIMIT 1`, roleID).Scan(&roleInfo.ParentDatabaseID,
		&roleInfo.RoleID, &roleInfo.RoleName)
	if getErr != nil {
		return nil, fmt.Errorf("GetUserRole: Unabled to get role: role ID = %v: datastore err=%v",
			roleID, getErr)
	}

	return &roleInfo, nil

}

func updateExistingRole(trackerDBHandle *sql.DB, roleID string, updatedRole *DatabaseRoleInfo) (*DatabaseRoleInfo, error) {

	if _, updateErr := trackerDBHandle.Exec(`UPDATE database_roles 
				SET name=$1
				WHERE role_id=$2`, updatedRole.RoleName, roleID); updateErr != nil {
		return nil, fmt.Errorf("updateExistingRole: Can't update role properties %v: error = %v",
			roleID, updateErr)
	}

	return updatedRole, nil

}

type NewDatabaseRoleParams struct {
	DatabaseID string `json:"databaseID"`
	RoleName   string `json:"roleName"`
}

func NewDatabaseRole(trackerDBHandle *sql.DB, params NewDatabaseRoleParams) (*DatabaseRole, error) {
	log.Printf("NewDatabaseRole: %+v", params)

	// TODO Wrap all the database writes from this function into a transaction.

	newRole, newRoleErr := addDatabaseRole(trackerDBHandle, params.DatabaseID, params.RoleName)
	if newRoleErr != nil {
		return nil, fmt.Errorf("NewDatabaseRole: %v", newRoleErr)
	}

	return newRole, nil
}

func AddCollaboratorRole(trackerDBHandle *sql.DB, roleID string, collaboratorID string) error {
	// TODO verify the current user has permissions to add the user as an admin.

	if _, insertErr := trackerDBHandle.Exec(
		`INSERT INTO collaborator_roles (role_id,collaborator_id) VALUES ($1,$2)`,
		roleID, collaboratorID); insertErr != nil {
		return fmt.Errorf("AddCollaboratorRole: Can't add collaborator with ID = %v to role with ID = %v: error = %v",
			collaboratorID, roleID, insertErr)
	}

	return nil

}

func GetDatabaseAdminUserInfo(trackerDBHandle *sql.DB, databaseID string) ([]userAuth.UserInfo, error) {

	rows, queryErr := trackerDBHandle.Query(
		`SELECT users.user_id,users.user_name,users.first_name,users.last_name 
				FROM collaborators,users
				WHERE collaborators.database_id=$1
				   AND collaborators.user_id=users.user_id
				   AND collaborators.is_admin='1'`, databaseID)
	if queryErr != nil {
		return nil, fmt.Errorf("GetDatabaseAdminUserInfo: Failure querying database: %v", queryErr)
	}
	defer rows.Close()

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

func GetCustomRoleListInfo(trackerDBHandle *sql.DB, databaseID string) ([]CustomListRoleInfo, error) {

	rows, queryErr := trackerDBHandle.Query(
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
	defer rows.Close()

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

func GetCustomRoleDashboardInfo(trackerDBHandle *sql.DB, databaseID string) ([]CustomRoleDashboardInfo, error) {

	rows, queryErr := trackerDBHandle.Query(
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
	defer rows.Close()

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

func GetCustomRoleInfo(trackerDBHandle *sql.DB, databaseID string) ([]CustomRoleInfo, error) {

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

	// Get a complete list of all roles for the database. This will
	// include roles with no users, no dashboard priviliges
	// or no list privileges.
	allRoles, allRolesErr := GetDatabaseRoles(trackerDBHandle, databaseID)
	if allRolesErr != nil {
		return nil, fmt.Errorf("GetCustomRoleInfo: %v", allRolesErr)
	}
	for _, roleInfo := range allRoles {
		getOrAllocRoleInfo(roleInfo.RoleID, roleInfo.RoleName)
	}

	customListInfo, listErr := GetCustomRoleListInfo(trackerDBHandle, databaseID)
	if listErr != nil {
		return nil, fmt.Errorf("GetCustomRoleInfo: %v", listErr)
	}
	for _, listInfo := range customListInfo {

		roleInfo := getOrAllocRoleInfo(listInfo.RoleID, listInfo.RoleName)
		roleInfo.ListPrivs = listInfo.ListPrivs
	}

	customDashboardInfo, dashErr := GetCustomRoleDashboardInfo(trackerDBHandle, databaseID)
	if dashErr != nil {
		return nil, fmt.Errorf("GetCustomRoleInfo: %v", dashErr)
	}
	for _, dashInfo := range customDashboardInfo {

		roleInfo := getOrAllocRoleInfo(dashInfo.RoleID, dashInfo.RoleName)
		roleInfo.DashboardPrivs = dashInfo.DashboardPrivs

	}

	usersRoleInfo, err := GetAllUsersRoleInfo(trackerDBHandle, databaseID)
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
	UserInfo       userAuth.UserInfo  `json:"userInfo"`
	CollaboratorID string             `json:"collaboratorID"`
	IsAdmin        bool               `json:"isAdmin"`
	RoleInfo       []DatabaseRoleInfo `json:"roleInfo"`
}

func GetCollaboratorRoleInfo(trackerDBHandle *sql.DB, databaseID string, collaboratorID string) (*UserRoleInfo, error) {

	collabInfo, collabErr := GetCollaboratorByID(trackerDBHandle, collaboratorID)
	if collabErr != nil {
		return nil, fmt.Errorf("GetCollaboratorRoleInfo: Failure querying database: %v", collabErr)
	}
	userInfo, err := userAuth.GetUserInfoByID(trackerDBHandle, collabInfo.UserID)
	if err != nil {
		return nil, fmt.Errorf("GetCollaboratorRoleInfo: Failure getting user info: userID=%v, error = %v", collabInfo.UserID, err)
	}
	userRoleInfo := UserRoleInfo{UserInfo: *userInfo,
		CollaboratorID: collabInfo.CollaboratorID,
		RoleInfo:       []DatabaseRoleInfo{},
		IsAdmin:        collabInfo.IsAdmin}

	rows, queryErr := trackerDBHandle.Query(
		`SELECT collaborators.user_id, database_roles.role_id,database_roles.name
				FROM collaborators,database_roles,collaborator_roles
				WHERE database_roles.database_id=$1
				   AND database_roles.role_id=collaborator_roles.role_id
				   AND collaborator_roles.collaborator_id=$2
				   AND collaborators.collaborator_id=collaborator_roles.collaborator_id`, databaseID, collaboratorID)
	if queryErr != nil {
		return nil, fmt.Errorf("GetCollaboratorRoleInfo: Failure querying database: %v", queryErr)
	}
	defer rows.Close()

	for rows.Next() {

		var currUserID string
		var currRoleInfo DatabaseRoleInfo

		if scanErr := rows.Scan(&currUserID, &currRoleInfo.RoleID, &currRoleInfo.RoleName); scanErr != nil {
			return nil, fmt.Errorf("GetCollaboratorRoleInfo: Failure querying database: %v", scanErr)
		}
		currRoleInfo.ParentDatabaseID = databaseID
		userRoleInfo.RoleInfo = append(userRoleInfo.RoleInfo, currRoleInfo)
	}

	return &userRoleInfo, nil

}

type GetCollaboratorRoleInfoParams struct {
	CollaboratorID string `json:"collaboratorID"`
	DatabaseID     string `json:"databaseID"`
}

func GetCollaboratorRoleInfoAPI(trackerDBHandle *sql.DB, params GetCollaboratorRoleInfoParams) (*UserRoleInfo, error) {
	return GetCollaboratorRoleInfo(trackerDBHandle, params.DatabaseID, params.CollaboratorID)
}

type SetCollaboratorRoleInfoParams struct {
	CollaboratorID string `json:"collaboratorID"`
	UserID         string `json:"userID"`
	DatabaseID     string `json:"databaseID"`
	RoleID         string `json:"roleID"`
	MemberOfRole   bool   `json:"memberOfRole"`
}

func SetCollaboratorRoleInfo(trackerDBHandle *sql.DB, params SetCollaboratorRoleInfoParams) error {
	if _, deleteErr := trackerDBHandle.Exec(`DELETE FROM collaborator_roles 
				WHERE role_id=$1 AND collaborator_id = $2`, params.RoleID, params.CollaboratorID); deleteErr != nil {
		return fmt.Errorf("SetUserRoleInfo: Can't update role properties %+v: error = %v",
			params, deleteErr)
	}
	if params.MemberOfRole {
		return AddCollaboratorRole(trackerDBHandle, params.RoleID, params.CollaboratorID)
	}
	return nil

}

// Aggregate the role information by user.
func GetAllUsersRoleInfo(trackerDBHandle *sql.DB, databaseID string) ([]UserRoleInfo, error) {

	collabs, collabErr := getAllCollaborators(trackerDBHandle, databaseID)
	if collabErr != nil {
		return nil, fmt.Errorf("GetAllUsersRoleInfo: Failure querying database: %v", collabErr)
	}

	roleInfoByUserID := map[string]*UserRoleInfo{}
	for _, currCollab := range collabs {
		userInfo, err := userAuth.GetUserInfoByID(trackerDBHandle, currCollab.UserID)
		if err != nil {
			return nil, fmt.Errorf("GetAllUsersRoleInfo: Failure getting user info: userID=%v, error = %v", currCollab.UserID, err)
		}

		userRoleInfo := &UserRoleInfo{UserInfo: *userInfo,
			CollaboratorID: currCollab.CollaboratorID,
			RoleInfo:       []DatabaseRoleInfo{},
			IsAdmin:        currCollab.IsAdmin}
		roleInfoByUserID[currCollab.UserID] = userRoleInfo
	}

	rows, queryErr := trackerDBHandle.Query(
		`SELECT collaborators.user_id, database_roles.role_id,database_roles.name
				FROM collaborators,database_roles,collaborator_roles
				WHERE collaborators.database_id=$1
				   AND database_roles.role_id=collaborator_roles.role_id 
				   AND collaborator_roles.collaborator_id = collaborators.collaborator_id`, databaseID)
	if queryErr != nil {
		return nil, fmt.Errorf("GetAllUsersRoleInfo: Failure querying database: %v", queryErr)
	}
	defer rows.Close()

	for rows.Next() {

		var currUserID string
		var currRoleInfo DatabaseRoleInfo

		if scanErr := rows.Scan(&currUserID, &currRoleInfo.RoleID, &currRoleInfo.RoleName); scanErr != nil {
			return nil, fmt.Errorf("GetAllUsersRoleInfo: Failure querying database: %v", scanErr)
		}

		var userRoleInfo *UserRoleInfo
		userRoleInfo, foundInfo := roleInfoByUserID[currUserID]
		if !foundInfo {
			return nil, fmt.Errorf("GetAllUsersRoleInfo: Failure getting collaborator info: userID=%v", currUserID)
		}
		userRoleInfo.RoleInfo = append(userRoleInfo.RoleInfo, currRoleInfo)
	}

	usersRoleInfo := []UserRoleInfo{}
	for _, currRoleInfo := range roleInfoByUserID {
		usersRoleInfo = append(usersRoleInfo, *currRoleInfo)
	}

	return usersRoleInfo, nil
}

type RoleCollaboratorInfo struct {
	UserInfo       userAuth.UserInfo `json:"userInfo"`
	CollaboratorID string            `json:"collaboratorID"`
	IsMemberOfRole bool              `json:"isMemberOfRole"`
}

type GetRoleCollaboratorsParams struct {
	DatabaseID string `json:"databaseID"`
	RoleID     string `json:"roleID"`
}

func GetRoleCollaborators(trackerDBHandle *sql.DB, params GetRoleCollaboratorsParams) ([]RoleCollaboratorInfo, error) {
	allUserRoleInfo, err := GetAllUsersRoleInfo(trackerDBHandle, params.DatabaseID)
	if err != nil {
		return nil, fmt.Errorf("GetRoleCollaborators: %v", err)
	}

	roleCollabInfo := []RoleCollaboratorInfo{}
	for _, currUserRoleInfo := range allUserRoleInfo {
		memberOfRole := func() bool {
			for _, currRoleInfo := range currUserRoleInfo.RoleInfo {
				if currRoleInfo.RoleID == params.RoleID {
					return true
				}
			}
			return false
		}
		currUserRoleCollabInfo := RoleCollaboratorInfo{
			UserInfo:       currUserRoleInfo.UserInfo,
			CollaboratorID: currUserRoleInfo.CollaboratorID,
			IsMemberOfRole: memberOfRole()}
		roleCollabInfo = append(roleCollabInfo, currUserRoleCollabInfo)
	}
	return roleCollabInfo, nil
}

type RoleUserInfo struct {
	RoleID    string              `json:"roleID"`
	RoleName  string              `json:"roleName"`
	RoleUsers []userAuth.UserInfo `json:"roleUsers"`
}

func GetRoleUserInfoByRoleID(trackerDBHandle *sql.DB, databaseID string) (map[string]RoleUserInfo, error) {

	roleInfo, roleInfoErr := GetCustomRoleInfo(trackerDBHandle, databaseID)
	if roleInfoErr != nil {
		return nil, fmt.Errorf("GetRoleUserInfo: Failure querying database: %v", roleInfoErr)
	}

	roleUserInfoByRoleID := map[string]RoleUserInfo{}

	for _, currRoleInfo := range roleInfo {
		roleUserInfo := RoleUserInfo{
			RoleID:    currRoleInfo.RoleID,
			RoleName:  currRoleInfo.RoleName,
			RoleUsers: currRoleInfo.RoleUsers}
		roleUserInfoByRoleID[currRoleInfo.RoleID] = roleUserInfo
	}

	return roleUserInfoByRoleID, nil

}

type DeleteCollaboratorParams struct {
	DatabaseID     string `json:"databaseID"`
	CollaboratorID string `json:"collaboratorID"`
}

func DeleteCollaborator(trackerDBHandle *sql.DB, params DeleteCollaboratorParams) error {
	if _, deleteErr := trackerDBHandle.Exec(`DELETE FROM collaborators 
				WHERE database_id=$1 AND collaborator_id=$2`,
		params.DatabaseID, params.CollaboratorID); deleteErr != nil {
		return fmt.Errorf("DeleteCollaborator: Can't update collaborator %+v: error = %v",
			params, deleteErr)
	}
	return nil
}
