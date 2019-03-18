// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package formLink

import (
	"fmt"

	"net/http"

	"resultra/tracker/server/common/databaseWrapper"
	"resultra/tracker/server/form"
	"resultra/tracker/server/common/userAuth"
	"resultra/tracker/server/userRole"
)

type NewItemInfo struct {
	FormID          string `json:"formID"`
	FormLinkID      string `json:"formLinkID"`
	LinkName        string `json:"linkName"`
	FormName        string `json:"formName"`
	CurrUserIsAdmin bool   `json:"currUserIsAdmin"`
	DatabaseID      string `json:"databaseID"`
}

func getNewItemInfo(r *http.Request, formLinkID string) (*NewItemInfo, error) {

	_, authErr := userAuth.GetCurrentUserInfo(r)
	if authErr != nil {
		return nil, authErr
	} else {

		trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
		if dbErr != nil {
			return nil, dbErr
		}

		formLink, getFormLinkErr := GetFormLink(trackerDBHandle, formLinkID)
		if getFormLinkErr != nil {
			return nil, getFormLinkErr
		}

		formInfo, getErr := form.GetForm(trackerDBHandle, formLink.FormID)
		if getErr != nil {
			return nil, getErr
		}

		hasPrivs, privsErr := userRole.CurrentUserHasNewItemLinkPrivs(trackerDBHandle, r, formInfo.ParentDatabaseID, formLinkID)
		if privsErr != nil {
			return nil, privsErr
		}
		if !hasPrivs {
			return nil, fmt.Errorf("ERROR: No permissions to add new items with this page")
		}

		isAdmin := userRole.CurrUserIsDatabaseAdmin(r, formInfo.ParentDatabaseID)

		newItemInfo := NewItemInfo{
			FormID:          formLink.FormID,
			FormName:        formInfo.Name,
			DatabaseID:      formInfo.ParentDatabaseID,
			CurrUserIsAdmin: isAdmin,
			FormLinkID:      formLink.LinkID,
			LinkName:        formLink.Name}

		return &newItemInfo, nil
	}

}
