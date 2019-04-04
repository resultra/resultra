// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package userRole

import (
	"database/sql"
	"fmt"
	"github.com/resultra/resultra/server/common/userAuth"
	"github.com/resultra/resultra/server/trackerDatabase"
	"net/http"
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

// trackerDBHandle

func setListRolePrivsToDest(destDBHandle *sql.DB, params SetListRolePrivsParams) error {

	if privsErr := verifyListRolePrivs(params.Privs); privsErr != nil {
		return fmt.Errorf("SetListRolePrivs: error = %v", privsErr)
	}

	// TODO - Enclose code below in mutex

	if _, deleteErr := destDBHandle.Exec(
		`DELETE FROM list_role_privs where role_id=$1 and list_id=$2`,
		params.RoleID, params.ListID); deleteErr != nil {
		return fmt.Errorf("SetListRolePrivs: Can't delete old privs: error = %v", deleteErr)
	}

	if _, insertErr := destDBHandle.Exec(
		`INSERT INTO list_role_privs (role_id,list_id,privs) VALUES ($1,$2,$3)`,
		params.RoleID, params.ListID, params.Privs); insertErr != nil {
		return fmt.Errorf("SetListRolePrivs: Can't set list privileges: error = %v", insertErr)
	}

	return nil

}

func SetListRolePrivs(trackerDBHandle *sql.DB, params SetListRolePrivsParams) error {
	return setListRolePrivsToDest(trackerDBHandle, params)
}

func getAllListRolePrivsFromSrc(srcDBHandle *sql.DB, parentDatabaseID string) ([]SetListRolePrivsParams, error) {
	// Overwrite the defaults for those roles with an explicit role setting for the list.
	rows, queryErr := srcDBHandle.Query(
		`SELECT list_role_privs.role_id,list_role_privs.list_id,list_role_privs.privs
			FROM list_role_privs,database_roles
			WHERE database_roles.role_id=list_role_privs.role_id
				AND database_roles.database_id=$1`, parentDatabaseID)
	if queryErr != nil {
		return nil, fmt.Errorf("getListRolePrivs: Failure querying database: %v", queryErr)
	}
	defer rows.Close()

	privs := []SetListRolePrivsParams{}
	for rows.Next() {
		currPrivInfo := SetListRolePrivsParams{}

		if scanErr := rows.Scan(&currPrivInfo.RoleID, &currPrivInfo.ListID, &currPrivInfo.Privs); scanErr != nil {
			return nil, fmt.Errorf("getFormRolePrivs: Failure querying database: %v", scanErr)
		}
		privs = append(privs, currPrivInfo)
	}

	return privs, nil

}

func CloneListPrivs(cloneParams *trackerDatabase.CloneDatabaseParams) error {

	listPrivs, err := getAllListRolePrivsFromSrc(cloneParams.SrcDBHandle, cloneParams.SourceDatabaseID)
	if err != nil {
		return fmt.Errorf("CloneRoles: Unable to retrieve roles: databaseID=%v, error=%v ",
			cloneParams.SourceDatabaseID, err)
	}
	for _, currListPriv := range listPrivs {

		destListID, err := cloneParams.IDRemapper.GetExistingRemappedID(currListPriv.ListID)
		if err != nil {
			return fmt.Errorf("CloneListPrivs: Unable to get mapped ID for source database: %v", err)
		}

		destRoleID, err := cloneParams.IDRemapper.GetExistingRemappedID(currListPriv.RoleID)
		if err != nil {
			return fmt.Errorf("CloneListPrivs: Unable to get mapped ID for source database: %v", err)
		}

		destPrivs := SetListRolePrivsParams{
			ListID: destListID,
			RoleID: destRoleID,
			Privs:  currListPriv.Privs}

		if err := setListRolePrivsToDest(cloneParams.DestDBHandle, destPrivs); err != nil {
			return fmt.Errorf("CloneListPrivs: Can't clone list privilege: error = %v", err)
		}
	}

	return nil

}

type ListRolePriv struct {
	RoleID   string `json:"roleID"`
	RoleName string `json:"roleName"`
	Privs    string `json:"privs"`
}

func GetListRolePrivs(trackerDBHandle *sql.DB, listID string) ([]ListRolePriv, error) {

	databaseID, databaseErr := getItemListDatabaseID(trackerDBHandle, listID)
	if databaseErr != nil {
		return nil, fmt.Errorf("getListRolePrivs: Error retrieving database info for list: %v", listID)
	}

	roles, rolesErr := GetDatabaseRoles(trackerDBHandle, databaseID)
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
	rows, queryErr := trackerDBHandle.Query(
		`SELECT database_roles.role_id,database_roles.name,list_role_privs.privs
			FROM list_role_privs,database_roles
			WHERE list_role_privs.list_id=$1
				AND database_roles.role_id=list_role_privs.role_id`, listID)
	if queryErr != nil {
		return nil, fmt.Errorf("getListRolePrivs: Failure querying database: %v", queryErr)
	}
	defer rows.Close()

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

func GetRoleListPrivs(trackerDBHandle *sql.DB, roleID string) ([]RoleListPriv, error) {

	rows, queryErr := trackerDBHandle.Query(
		`SELECT item_lists.list_id,item_lists.name,list_role_privs.privs
			FROM list_role_privs,item_lists
			WHERE list_role_privs.role_id=$1 AND
				item_lists.list_id = list_role_privs.list_id`, roleID)
	if queryErr != nil {
		return nil, fmt.Errorf("GetRoleListPrivs: Failure querying database: %v", queryErr)
	}
	defer rows.Close()

	roleListPrivs := []RoleListPriv{}
	for rows.Next() {
		currPrivInfo := RoleListPriv{}

		if scanErr := rows.Scan(&currPrivInfo.ListID, &currPrivInfo.ListName, &currPrivInfo.Privs); scanErr != nil {
			return nil, fmt.Errorf("GetRoleListPrivs: Failure querying database: %v", scanErr)
		}
		roleListPrivs = append(roleListPrivs, currPrivInfo)
	}

	return roleListPrivs, nil
}

func GetItemListsWithUserPrivs(trackerDBHandle *sql.DB, databaseID string, userID string) (map[string]bool, error) {
	rows, queryErr := trackerDBHandle.Query(
		`SELECT list_role_privs.list_id, list_role_privs.privs
				FROM list_role_privs,database_roles,collaborator_roles,collaborators
				WHERE database_roles.database_id=$1
					AND list_role_privs.role_id=database_roles.role_id
					AND database_roles.role_id=collaborator_roles.role_id
					AND collaborator_roles.collaborator_id=collaborators.collaborator_id
					AND collaborators.user_id=$2`, databaseID, userID)
	if queryErr != nil {
		return nil, fmt.Errorf("GetItemListsWithUserPrivs: Failure querying database: %v", queryErr)
	}
	defer rows.Close()

	visibleLists := map[string]bool{}
	for rows.Next() {
		listID := ""
		privs := ""
		if scanErr := rows.Scan(&listID, &privs); scanErr != nil {
			return nil, fmt.Errorf("GetCustomRoleDashboardInfo: Failure querying database: %v", scanErr)
		}
		if privs != ListRolePrivsNone {
			visibleLists[listID] = true
		}
	}

	return visibleLists, nil
}

func GetCurrentUserItemListPrivs(trackerDBHandle *sql.DB, req *http.Request,
	databaseID string, listID string) (string, error) {

	privs := ListRolePrivsNone // default

	if CurrUserIsDatabaseAdmin(req, databaseID) {
		privs = ListRolePrivsEdit
		return privs, nil
	}

	currUserID, userErr := userAuth.GetCurrentUserID(req)
	if userErr != nil {
		return "", fmt.Errorf("verifyCurrUserIsDatabaseAdmin: can't verify user: %v", userErr)
	}

	rows, queryErr := trackerDBHandle.Query(
		`SELECT list_role_privs.privs
				FROM list_role_privs,database_roles,collaborator_roles,collaborators
				WHERE database_roles.database_id=$1
					AND list_role_privs.list_id=$2
					AND list_role_privs.role_id=database_roles.role_id
					AND database_roles.role_id=collaborator_roles.role_id
					AND collaborator_roles.collaborator_id=collaborators.collaborator_id
					AND collaborators.user_id=$3`, databaseID, listID, currUserID)
	if queryErr != nil {
		return "", fmt.Errorf("GetCurrentUserItemListPrivs: Failure querying database: %v", queryErr)
	}
	defer rows.Close()

	for rows.Next() {
		rolePrivs := ""
		if scanErr := rows.Scan(&rolePrivs); scanErr != nil {
			return "", fmt.Errorf("GetCurrentUserItemListPrivs: Failure querying database: %v", scanErr)
		}
		switch rolePrivs {
		case ListRolePrivsView:
			if privs == ListRolePrivsNone {
				privs = rolePrivs
			}
		case ListRolePrivsEdit:
			privs = ListRolePrivsEdit
		}
	}

	return privs, nil

}
