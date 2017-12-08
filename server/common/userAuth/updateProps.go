package userAuth

import (
	"fmt"
	"net/http"
	"resultra/datasheet/server/common/databaseWrapper"
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
