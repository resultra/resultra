// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package userRole

import (
	"database/sql"
	"fmt"
	"github.com/resultra/resultra/server/common/databaseWrapper"
	"github.com/resultra/resultra/server/generic/api"
	"github.com/resultra/resultra/server/generic/stringValidation"
	"net/http"
)

type RoleIDInterface interface {
	getRoleID() string
}

type RoleIDHeader struct {
	RoleID string `json:"roleID"`
}

func (idHeader RoleIDHeader) getRoleID() string {
	return idHeader.RoleID
}

type RolePropUpdater interface {
	RoleIDInterface
	updateProps(role *DatabaseRoleInfo) error
}

func UpdateRoleProps(trackerDBHandle *sql.DB, propUpdater RolePropUpdater) (*DatabaseRoleInfo, error) {

	// Retrieve the bar chart from the data store
	roleForUpdate, getErr := GetUserRole(trackerDBHandle, propUpdater.getRoleID())
	if getErr != nil {
		return nil, fmt.Errorf("UpdateRoleProps: Unable to get existing role: %v", getErr)
	}

	if propUpdateErr := propUpdater.updateProps(roleForUpdate); propUpdateErr != nil {
		return nil, fmt.Errorf("UpdateRoleProps: Unable to update existing form properties: %v", propUpdateErr)
	}

	updatedRole, updateErr := updateExistingRole(trackerDBHandle, propUpdater.getRoleID(), roleForUpdate)
	if updateErr != nil {
		return nil, fmt.Errorf("UpdateRoleProps: Unable to update existing role properties: datastore update error =  %v", updateErr)
	}

	return updatedRole, nil
}

func ProcessRolePropUpdate(w http.ResponseWriter, r *http.Request, propUpdater RolePropUpdater) {

	if verifyErr := VerifyCurrUserIsDatabaseAdminForUserRole(r, propUpdater.getRoleID()); verifyErr != nil {
		api.WriteErrorResponse(w, verifyErr)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if updatedRole, err := UpdateRoleProps(trackerDBHandle, propUpdater); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, updatedRole)
	}
}

type SetRoleNameParams struct {
	RoleIDHeader
	NewRoleName string `json:"newRoleName"`
}

func (updateParams SetRoleNameParams) updateProps(role *DatabaseRoleInfo) error {

	if !stringValidation.WellFormedItemName(updateParams.NewRoleName) {
		return fmt.Errorf("update role name: invalid name %v", updateParams.NewRoleName)
	}

	role.RoleName = updateParams.NewRoleName

	return nil
}
