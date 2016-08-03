package userRole

import (
	"fmt"
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
