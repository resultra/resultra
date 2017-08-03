package userRole

import (
	"fmt"
	"net/http"
	"resultra/datasheet/server/generic/databaseWrapper"
	"resultra/datasheet/server/generic/userAuth"
)

type SetNewItemFormLinkRolePrivsParams struct {
	LinkID      string `json:"linkID"`
	RoleID      string `json:"roleID"`
	LinkEnabled bool   `json:"linkEnabled"`
}

func SetNewItemFormLinkRolePrivs(params SetNewItemFormLinkRolePrivsParams) error {

	if _, deleteErr := databaseWrapper.DBHandle().Exec(
		`DELETE FROM new_item_form_link_role_privs where role_id=$1 and link_id=$2`,
		params.RoleID, params.LinkID); deleteErr != nil {
		return fmt.Errorf("SetNewItemFormLinkRolePrivs: Can't delete old privs: error = %v", deleteErr)
	}

	if params.LinkEnabled {
		if _, insertErr := databaseWrapper.DBHandle().Exec(
			`INSERT INTO new_item_form_link_role_privs (role_id,link_id) VALUES ($1,$2)`,
			params.RoleID, params.LinkID); insertErr != nil {
			return fmt.Errorf("SetNewItemFormLinkRolePrivs: Can't set list privileges: error = %v", insertErr)
		}
	}

	return nil

}

type GetNewItemPrivParams struct {
	RoleID string `json:"roleID"`
}

type RoleNewItemPriv struct {
	LinkID      string `json:"linkID"`
	LinkName    string `json:"linkName"`
	LinkEnabled bool   `json:"linkEnabled"`
}

func getDefaultFormLinks(databaseID string) ([]RoleNewItemPriv, error) {

	rows, queryErr := databaseWrapper.DBHandle().Query(
		`SELECT form_links.link_id,form_links.name
			FROM form_links,forms
			WHERE form_links.form_id=forms.form_id AND
				forms.database_id=$1`, databaseID)
	if queryErr != nil {
		return nil, fmt.Errorf("GetRoleListPrivs: Failure querying database: %v", queryErr)
	}
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

func GetNewItemPrivs(roleID string) ([]RoleNewItemPriv, error) {

	roleDatabaseID, roleDBErr := GetUserRoleDatabaseID(roleID)
	if roleDBErr != nil {
		return nil, fmt.Errorf("GetNewItemPrivs: Failure querying database: %v", roleDBErr)
	}

	defaultFormLinks, getLinkErr := getDefaultFormLinks(roleDatabaseID)
	if getLinkErr != nil {
		return nil, fmt.Errorf("GetNewItemPrivs: Failure querying database: %v", getLinkErr)
	}

	privsByLinkID := map[string]RoleNewItemPriv{}
	for _, defaultPriv := range defaultFormLinks {
		privsByLinkID[defaultPriv.LinkID] = defaultPriv
	}

	rows, queryErr := databaseWrapper.DBHandle().Query(
		`SELECT form_links.link_id,form_links.name
			FROM new_item_form_link_role_privs,form_links
			WHERE new_item_form_link_role_privs.role_id=$1 AND
				new_item_form_link_role_privs.link_id = form_links.link_id`, roleID)
	if queryErr != nil {
		return nil, fmt.Errorf("GetRoleListPrivs: Failure querying database: %v", queryErr)
	}

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

func GetNewItemLinksWithUserPrivs(databaseID string, userID string) (map[string]bool, error) {
	rows, queryErr := databaseWrapper.DBHandle().Query(
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

func CurrentUserHasNewItemLinkPrivs(req *http.Request,
	databaseID string, linkID string) (bool, error) {

	if CurrUserIsDatabaseAdmin(req, databaseID) {
		return true, nil
	}

	currUserID, userErr := userAuth.GetCurrentUserID(req)
	if userErr != nil {
		return false, fmt.Errorf("verifyCurrUserIsDatabaseAdmin: can't verify user: %v", userErr)
	}

	rows, queryErr := databaseWrapper.DBHandle().Query(
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
