// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package userAuth

import (
	"fmt"
	"net/http"
	"github.com/resultra/resultra/server/common/databaseWrapper"
)

type UserIDInterface interface {
	getUserID() string
}

type UserIDHeader struct {
	UserID string `json:"userID"`
}

func (idHeader UserIDHeader) getUserID() string {
	return idHeader.UserID
}

type UserPropUpdater interface {
	UserIDInterface
	updateProps(userInfo *UserInfo) error
}

func updateUserProps(r *http.Request, propUpdater UserPropUpdater) (*UserInfo, error) {

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		return nil, fmt.Errorf("updateWorkspaceProps: Unable to get existing workspace information: %v", dbErr)
	}

	userInfoForUpdate, getErr := GetUserInfoByID(trackerDBHandle, propUpdater.getUserID())
	if getErr != nil {
		return nil, fmt.Errorf("updateUserProps: Unable to get existing workspace information: %v", getErr)
	}

	if propUpdateErr := propUpdater.updateProps(userInfoForUpdate); propUpdateErr != nil {
		return nil, fmt.Errorf("updateUserProps: Unable to update existing user properties: %v", propUpdateErr)
	}

	updateErr := updateUserProperties(trackerDBHandle, propUpdater.getUserID(), userInfoForUpdate.Properties)
	if updateErr != nil {
		return nil, fmt.Errorf("updateUserProps: Unable to update existing user properties: datastore update error =  %v", updateErr)
	}

	return userInfoForUpdate, nil
}

type ProfileBioParams struct {
	Bio string `json:"bio"`
}

func (updateParams ProfileBioParams) updateProps(userInfo *UserInfo) error {

	//	userInfo.Properties.Bio = updateParams.Bio

	return nil
}
