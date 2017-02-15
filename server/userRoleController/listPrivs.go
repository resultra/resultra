package userRoleController

import (
	"fmt"
	"resultra/datasheet/server/itemList"
	"resultra/datasheet/server/userRole"
)

func getRoleListPrivsWithDefaults(roleID string) ([]userRole.RoleListPriv, error) {

	roleDB, err := userRole.GetUserRoleDatabaseID(roleID)
	if err != nil {
		return nil, fmt.Errorf("getRoleListPrivsWithDefaults: %v", err)
	}

	privsByListID := map[string]userRole.RoleListPriv{}

	// Start off with no privileges as the default for all lists

	allItemLists, err := itemList.GetAllItemLists(roleDB)
	if err != nil {
		return nil, fmt.Errorf("getRoleListPrivsWithDefaults: %v", err)
	}
	for _, currList := range allItemLists {
		privsByListID[currList.ListID] = userRole.RoleListPriv{
			ListID:   currList.ListID,
			ListName: currList.Name,
			Privs:    userRole.ListRolePrivsNone}
	}

	// Update the privileges for those lists with an explicit set of privileges set.
	explicitListPrivs, err := userRole.GetRoleListPrivs(roleID)
	if err != nil {
		return nil, fmt.Errorf("getRoleListPrivsWithDefaults: %v", err)
	}
	for _, currPriv := range explicitListPrivs {
		privsByListID[currPriv.ListID] = currPriv
	}

	// Flatten the list
	roleListPrivs := []userRole.RoleListPriv{}
	for _, currPriv := range privsByListID {
		roleListPrivs = append(roleListPrivs, currPriv)
	}

	return roleListPrivs, nil

}
