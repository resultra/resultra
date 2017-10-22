package itemList

import (
	"fmt"
	"net/http"
	"resultra/datasheet/server/itemList"
	"resultra/datasheet/server/userRole"
)

type ViewListInfo struct {
	ListID         string
	ListName       string
	ListPrivileges string
}

func getViewListInfo(r *http.Request, listID string) (*ViewListInfo, error) {

	listInfo, err := itemList.GetItemList(listID)
	if err != nil {
		return nil, err
	}

	listPrivs, privsErr := userRole.GetCurrentUserItemListPrivs(r, listInfo.ParentDatabaseID, listID)
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
