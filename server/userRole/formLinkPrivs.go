package userRole

import (
	"fmt"
	"resultra/datasheet/server/generic/databaseWrapper"
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
