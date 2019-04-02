// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package userRoleController

import (
	"database/sql"
	"fmt"
	"github.com/resultra/resultra/server/itemList"
	"github.com/resultra/resultra/server/userRole"
)

func getRoleListPrivsWithDefaults(trackerDBHandle *sql.DB, roleID string) ([]userRole.RoleListPriv, error) {

	roleDB, err := userRole.GetUserRoleDatabaseID(trackerDBHandle, roleID)
	if err != nil {
		return nil, fmt.Errorf("getRoleListPrivsWithDefaults: %v", err)
	}

	privsByListID := map[string]userRole.RoleListPriv{}

	// Start off with no privileges as the default for all lists

	allItemLists, err := itemList.GetAllItemLists(trackerDBHandle, roleDB)
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
	explicitListPrivs, err := userRole.GetRoleListPrivs(trackerDBHandle, roleID)
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
