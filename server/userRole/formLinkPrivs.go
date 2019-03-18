// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package userRole

import (
	"database/sql"
	"fmt"
	"net/http"
	"resultra/tracker/server/common/userAuth"
	"resultra/tracker/server/trackerDatabase"
)

type SetNewItemFormLinkRolePrivsParams struct {
	LinkID      string `json:"linkID"`
	RoleID      string `json:"roleID"`
	LinkEnabled bool   `json:"linkEnabled"`
}

func setNewItemFormLinkRolePrivsToDest(destDBHandle *sql.DB, params SetNewItemFormLinkRolePrivsParams) error {

	if _, deleteErr := destDBHandle.Exec(
		`DELETE FROM new_item_form_link_role_privs where role_id=$1 and link_id=$2`,
		params.RoleID, params.LinkID); deleteErr != nil {
		return fmt.Errorf("SetNewItemFormLinkRolePrivs: Can't delete old privs: error = %v", deleteErr)
	}

	if params.LinkEnabled {
		if _, insertErr := destDBHandle.Exec(
			`INSERT INTO new_item_form_link_role_privs (role_id,link_id) VALUES ($1,$2)`,
			params.RoleID, params.LinkID); insertErr != nil {
			return fmt.Errorf("SetNewItemFormLinkRolePrivs: Can't set list privileges: error = %v", insertErr)
		}
	}

	return nil

}

func getAllItemLinkPrivsFromSrc(srcDBHandle *sql.DB, parentDatabaseID string) ([]SetNewItemFormLinkRolePrivsParams, error) {

	rows, queryErr := srcDBHandle.Query(
		`SELECT new_item_form_link_role_privs.link_id,new_item_form_link_role_privs.role_id
			FROM new_item_form_link_role_privs,form_links,forms
			WHERE forms.database_id=$1
				AND new_item_form_link_role_privs.link_id = form_links.link_id 
				AND form_links.form_id=forms.form_id`, parentDatabaseID)
	if queryErr != nil {
		return nil, fmt.Errorf("GetRoleListPrivs: Failure querying database: %v", queryErr)
	}
	defer rows.Close()

	privs := []SetNewItemFormLinkRolePrivsParams{}

	for rows.Next() {
		currPrivs := SetNewItemFormLinkRolePrivsParams{}
		currPrivs.LinkEnabled = true // presence in the database means the link is enabled

		if scanErr := rows.Scan(&currPrivs.LinkID, &currPrivs.RoleID); scanErr != nil {
			return nil, fmt.Errorf("getAllItemLinkPrivsFromSrc: Failure querying database: %v", scanErr)
		}
		privs = append(privs, currPrivs)
	}
	return privs, nil

}

func CloneNewItemLinkPrivs(cloneParams *trackerDatabase.CloneDatabaseParams) error {

	linkPrivs, err := getAllItemLinkPrivsFromSrc(cloneParams.SrcDBHandle, cloneParams.SourceDatabaseID)
	if err != nil {
		return fmt.Errorf("CloneNewItemLinkPrivs: Unable to retrieve roles: databaseID=%v, error=%v ",
			cloneParams.SourceDatabaseID, err)
	}
	for _, currPriv := range linkPrivs {

		destLinkID, err := cloneParams.IDRemapper.GetExistingRemappedID(currPriv.LinkID)
		if err != nil {
			return fmt.Errorf("CloneNewItemLinkPrivs: Unable to get mapped ID for source database: %v", err)
		}

		destRoleID, err := cloneParams.IDRemapper.GetExistingRemappedID(currPriv.RoleID)
		if err != nil {
			return fmt.Errorf("CloneNewItemLinkPrivs: Unable to get mapped ID for source database: %v", err)
		}

		destPrivs := SetNewItemFormLinkRolePrivsParams{
			LinkID:      destLinkID,
			RoleID:      destRoleID,
			LinkEnabled: true}

		if err := setNewItemFormLinkRolePrivsToDest(cloneParams.DestDBHandle, destPrivs); err != nil {
			return fmt.Errorf("CloneNewItemLinkPrivs: Can't clone list privilege: error = %v", err)
		}
	}

	return nil

}

func SetNewItemFormLinkRolePrivs(trackerDBHandle *sql.DB, params SetNewItemFormLinkRolePrivsParams) error {
	return setNewItemFormLinkRolePrivsToDest(trackerDBHandle, params)
}

type GetNewItemPrivParams struct {
	RoleID string `json:"roleID"`
}

type RoleNewItemPriv struct {
	LinkID      string `json:"linkID"`
	LinkName    string `json:"linkName"`
	LinkEnabled bool   `json:"linkEnabled"`
}

func getDefaultFormLinks(trackerDBHandle *sql.DB, databaseID string) ([]RoleNewItemPriv, error) {

	rows, queryErr := trackerDBHandle.Query(
		`SELECT form_links.link_id,form_links.name
			FROM form_links,forms
			WHERE form_links.form_id=forms.form_id AND
				forms.database_id=$1`, databaseID)
	if queryErr != nil {
		return nil, fmt.Errorf("GetRoleListPrivs: Failure querying database: %v", queryErr)
	}
	defer rows.Close()

	roleNewItemPrivs := []RoleNewItemPriv{}
	for rows.Next() {
		currPrivInfo := RoleNewItemPriv{}
		currPrivInfo.LinkEnabled = false // presence in the database means the link is enabled

		if scanErr := rows.Scan(&currPrivInfo.LinkID, &currPrivInfo.LinkName); scanErr != nil {
			return nil, fmt.Errorf("GetRoleListPrivs: Failure querying database: %v", scanErr)
		}
		roleNewItemPrivs = append(roleNewItemPrivs, currPrivInfo)
	}
	return roleNewItemPrivs, nil
}

func GetRoleNewItemPrivs(trackerDBHandle *sql.DB, roleID string) ([]RoleNewItemPriv, error) {

	roleDatabaseID, roleDBErr := GetUserRoleDatabaseID(trackerDBHandle, roleID)
	if roleDBErr != nil {
		return nil, fmt.Errorf("GetNewItemPrivs: Failure querying database: %v", roleDBErr)
	}

	defaultFormLinks, getLinkErr := getDefaultFormLinks(trackerDBHandle, roleDatabaseID)
	if getLinkErr != nil {
		return nil, fmt.Errorf("GetNewItemPrivs: Failure querying database: %v", getLinkErr)
	}

	privsByLinkID := map[string]RoleNewItemPriv{}
	for _, defaultPriv := range defaultFormLinks {
		privsByLinkID[defaultPriv.LinkID] = defaultPriv
	}

	rows, queryErr := trackerDBHandle.Query(
		`SELECT form_links.link_id,form_links.name
			FROM new_item_form_link_role_privs,form_links
			WHERE new_item_form_link_role_privs.role_id=$1 AND
				new_item_form_link_role_privs.link_id = form_links.link_id`, roleID)
	if queryErr != nil {
		return nil, fmt.Errorf("GetRoleListPrivs: Failure querying database: %v", queryErr)
	}
	defer rows.Close()

	for rows.Next() {
		currPrivInfo := RoleNewItemPriv{}
		currPrivInfo.LinkEnabled = true // presence in the database means the link is enabled

		if scanErr := rows.Scan(&currPrivInfo.LinkID, &currPrivInfo.LinkName); scanErr != nil {
			return nil, fmt.Errorf("GetRoleListPrivs: Failure querying database: %v", scanErr)
		}
		privsByLinkID[currPrivInfo.LinkID] = currPrivInfo
	}

	// Flatten the privileges back down to a list.
	roleNewItemPrivs := []RoleNewItemPriv{}
	for _, currPrivInfo := range privsByLinkID {
		roleNewItemPrivs = append(roleNewItemPrivs, currPrivInfo)
	}

	return roleNewItemPrivs, nil
}

type NewItemRolePriv struct {
	RoleID      string `json:"roleID"`
	RoleName    string `json:"roleName"`
	LinkEnabled bool   `json:"linkEnabled"`
}

type GetNewItemRolePrivParams struct {
	LinkID string `json:"linkID"`
}

func GetNewItemRolePrivs(trackerDBHandle *sql.DB, linkID string) ([]NewItemRolePriv, error) {

	linkDatabaseID, linkDBErr := getNewItemLinkDatabaseID(trackerDBHandle, linkID)
	if linkDBErr != nil {
		return nil, fmt.Errorf("GetNewItemPrivs: Failure querying database: %v", linkDBErr)
	}

	// If there are no explicit privileges for a given role, the default is no privileges. So,
	// to popuplate an array of the dashboard's privileges for all roles, a map must first be populated
	// with a set of defaults.
	privsByRoleID := map[string]NewItemRolePriv{}
	allRoles, rolesErr := getDatabaseRolesFromSrc(trackerDBHandle, linkDatabaseID)
	if rolesErr != nil {
		return nil, fmt.Errorf("getAllDashboardRolesFromSrc: failure querying database: %v", rolesErr)
	}
	for _, currRoleInfo := range allRoles {
		defaultPrivInfo := NewItemRolePriv{
			RoleID:      currRoleInfo.RoleID,
			RoleName:    currRoleInfo.RoleName,
			LinkEnabled: false}
		privsByRoleID[currRoleInfo.RoleID] = defaultPrivInfo
	}

	rows, queryErr := trackerDBHandle.Query(
		`SELECT database_roles.role_id,database_roles.name
			FROM new_item_form_link_role_privs,database_roles
			WHERE new_item_form_link_role_privs.role_id=database_roles.role_id AND
				new_item_form_link_role_privs.link_id = $1`, linkID)
	if queryErr != nil {
		return nil, fmt.Errorf("GetRoleListPrivs: Failure querying database: %v", queryErr)
	}
	defer rows.Close()

	for rows.Next() {
		currPrivInfo := NewItemRolePriv{}
		currPrivInfo.LinkEnabled = true // presence in the database means the link is enabled

		if scanErr := rows.Scan(&currPrivInfo.RoleID, &currPrivInfo.RoleName); scanErr != nil {
			return nil, fmt.Errorf("GetNewItemRolePrivs: Failure querying database: %v", scanErr)
		}
		privsByRoleID[currPrivInfo.RoleID] = currPrivInfo
	}

	// Flatten the privileges back down to a list.
	roleNewItemPrivs := []NewItemRolePriv{}
	for _, currPrivInfo := range privsByRoleID {
		roleNewItemPrivs = append(roleNewItemPrivs, currPrivInfo)
	}

	return roleNewItemPrivs, nil
}

func GetNewItemLinksWithUserPrivs(trackerDBHandle *sql.DB, databaseID string, userID string) (map[string]bool, error) {
	rows, queryErr := trackerDBHandle.Query(
		`SELECT new_item_form_link_role_privs.link_id
				FROM new_item_form_link_role_privs,database_roles,collaborator_roles,collaborators
				WHERE database_roles.database_id=$1
					AND new_item_form_link_role_privs.role_id=database_roles.role_id
					AND database_roles.role_id=collaborator_roles.role_id
					AND collaborator_roles.collaborator_id=collaborators.collaborator_id
					AND collaborators.user_id=$2`, databaseID, userID)
	if queryErr != nil {
		return nil, fmt.Errorf("GetItemListsWithUserPrivs: Failure querying database: %v", queryErr)
	}
	defer rows.Close()

	visibleLinks := map[string]bool{}
	for rows.Next() {
		linkID := ""
		if scanErr := rows.Scan(&linkID); scanErr != nil {
			return nil, fmt.Errorf("GetNewItemLinksWithUserPrivs: Failure querying database: %v", scanErr)
		}
		visibleLinks[linkID] = true
	}

	return visibleLinks, nil
}

func CurrentUserHasNewItemLinkPrivs(trackerDBHandle *sql.DB, req *http.Request,
	databaseID string, linkID string) (bool, error) {

	if CurrUserIsDatabaseAdmin(req, databaseID) {
		return true, nil
	}

	currUserID, userErr := userAuth.GetCurrentUserID(req)
	if userErr != nil {
		return false, fmt.Errorf("verifyCurrUserIsDatabaseAdmin: can't verify user: %v", userErr)
	}

	rows, queryErr := trackerDBHandle.Query(
		`SELECT new_item_form_link_role_privs.link_id
					FROM new_item_form_link_role_privs,database_roles,collaborator_roles,collaborators
					WHERE database_roles.database_id=$1
					AND new_item_form_link_role_privs.link_id=$2
						AND new_item_form_link_role_privs.role_id=database_roles.role_id
						AND database_roles.role_id=collaborator_roles.role_id
						AND collaborator_roles.collaborator_id=collaborators.collaborator_id
						AND collaborators.user_id=$3`, databaseID, linkID, currUserID)
	if queryErr != nil {
		return false, fmt.Errorf("CurrentUserHasNewItemLinkPrivs: Failure querying database: %v", queryErr)
	}
	defer rows.Close()

	// Default to false (no priveleges) unless there's privileges set in the database.
	privs := false

	for rows.Next() {
		linkIDWithPrivs := ""
		if scanErr := rows.Scan(&linkIDWithPrivs); scanErr != nil {
			return false, fmt.Errorf("CurrentUserHasNewItemLinkPrivs: Failure querying database: %v", scanErr)
		}
		privs = true
	}

	return privs, nil

}
