// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package itemList

import (
	"fmt"
	"github.com/resultra/resultra/server/common/databaseWrapper"
	"github.com/resultra/resultra/server/itemList"
	"github.com/resultra/resultra/server/userRole"
	"net/http"
)

type ViewListInfo struct {
	ListID         string
	ListName       string
	ListPrivileges string
}

func getViewListInfo(r *http.Request, listID string) (*ViewListInfo, error) {

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		return nil, dbErr
	}

	listInfo, err := itemList.GetItemList(trackerDBHandle, listID)
	if err != nil {
		return nil, err
	}

	listPrivs, privsErr := userRole.GetCurrentUserItemListPrivs(trackerDBHandle, r, listInfo.ParentDatabaseID, listID)
	if privsErr != nil {
		return nil, privsErr
	}

	if listPrivs == userRole.ListRolePrivsNone {
		return nil, fmt.Errorf("Invalid permissions loading page. No permissions to view or edit this list.")
	}

	viewListInfo := ViewListInfo{
		ListID:         listID,
		ListName:       listInfo.Name,
		ListPrivileges: listPrivs}

	return &viewListInfo, nil

}
