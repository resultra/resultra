package userRole

import (
	"fmt"
	"resultra/datasheet/server/generic/databaseWrapper"
)

const ListRolePrivsNone string = "none"
const ListRolePrivsView string = "view"
const ListRolePrivsEdit string = "edit"

type ListPriv struct {
	ListID string `json:"listID"`
	Privs  string `json:"privs"`
}

func verifyListRolePrivs(privs string) error {
	if (privs == ListRolePrivsNone) ||
		(privs == ListRolePrivsView) ||
		(privs == ListRolePrivsEdit) {
		return nil
	} else {
		return fmt.Errorf("verifyListRolePrivs: Invalid privileges: %v", privs)
	}
}

type SetListRolePrivsParams struct {
	ListID string `json:"listID"`
	RoleID string `json:"roleID"`
	Privs  string `json:"privs"`
}

func setListRolePrivs(params SetListRolePrivsParams) error {

	if privsErr := verifyListRolePrivs(params.Privs); privsErr != nil {
		return fmt.Errorf("setListRolePrivs: error = %v", privsErr)
	}

	if _, deleteErr := databaseWrapper.DBHandle().Exec(
		`DELETE FROM list_role_privs where role_id=$1 and list_id=$2`,
		params.RoleID, params.ListID); deleteErr != nil {
		return fmt.Errorf("setListRolePrivs: Can't delete old privs: error = %v", deleteErr)
	}

	if _, insertErr := databaseWrapper.DBHandle().Exec(
		`INSERT INTO list_role_privs (role_id,list_id,privs) VALUES ($1,$2,$3)`,
		params.RoleID, params.ListID, params.Privs); insertErr != nil {
		return fmt.Errorf("setListRolePrivs: Can't set list privileges: error = %v", insertErr)
	}

	return nil

}

type ListRolePriv struct {
	RoleID   string `json:"roleID"`
	RoleName string `json:"roleName"`
	Privs    string `json:"privs"`
}

func getListRolePrivs(listID string) ([]ListRolePriv, error) {

	databaseID, databaseErr := getItemListDatabaseID(listID)
	if databaseErr != nil {
		return nil, fmt.Errorf("getListRolePrivs: Error retrieving database info for list: %v", listID)
	}

	roles, rolesErr := getDatabaseRoles(databaseID)
	if rolesErr != nil {
		return nil, fmt.Errorf("getListRolePrivs: Error getting roles for list: %v", listID)
	}

	// Start with a default for all roles
	rolePrivMap := map[string]*ListRolePriv{}
	for _, currRole := range roles {
		defaultRolePriv := ListRolePriv{
			RoleID:   currRole.RoleID,
			RoleName: currRole.RoleName,
			Privs:    ListRolePrivsNone}
		rolePrivMap[currRole.RoleID] = &defaultRolePriv
	}

	// Overwrite the defaults for those roles with an explicit role setting for the list.
	rows, queryErr := databaseWrapper.DBHandle().Query(
		`SELECT database_roles.role_id,database_roles.name,list_role_privs.privs
			FROM list_role_privs,database_roles
			WHERE list_role_privs.list_id=$1
				AND database_roles.role_id=list_role_privs.role_id`, listID)
	if queryErr != nil {
		return nil, fmt.Errorf("getListRolePrivs: Failure querying database: %v", queryErr)
	}

	for rows.Next() {
		currPrivInfo := ListRolePriv{}

		if scanErr := rows.Scan(&currPrivInfo.RoleID, &currPrivInfo.RoleName, &currPrivInfo.Privs); scanErr != nil {
			return nil, fmt.Errorf("getFormRolePrivs: Failure querying database: %v", scanErr)
		}
		rolePrivMap[currPrivInfo.RoleID] = &currPrivInfo
	}

	listRolePrivs := []ListRolePriv{}
	for _, currRolePriv := range rolePrivMap {
		listRolePrivs = append(listRolePrivs, *currRolePriv)
	}

	return listRolePrivs, nil

}

type GetRoleListPrivParams struct {
	RoleID string `json:"roleID"`
}

type RoleListPriv struct {
	ListID   string `json:"listID"`
	ListName string `json:"listName"`
	Privs    string `json:"privs"`
}

func getRoleListPrivs(roleID string) ([]RoleListPriv, error) {
	return nil, fmt.Errorf("Not implemented")
}
