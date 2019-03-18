package itemList

import (
	"fmt"
	"net/http"
	"resultra/tracker/server/common/databaseWrapper"
	"resultra/tracker/server/itemList"
	"resultra/tracker/server/userRole"
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
