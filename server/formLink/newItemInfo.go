package formLink

import (
	"fmt"

	"net/http"

	"resultra/datasheet/server/form"
	"resultra/datasheet/server/generic/userAuth"
	"resultra/datasheet/server/userRole"
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

		formLink, getFormLinkErr := GetFormLink(formLinkID)
		if getFormLinkErr != nil {
			return nil, getFormLinkErr
		}

		formInfo, getErr := form.GetForm(formLink.FormID)
		if getErr != nil {
			return nil, getErr
		}

		hasPrivs, privsErr := userRole.CurrentUserHasNewItemLinkPrivs(r, formInfo.ParentDatabaseID, formLinkID)
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
