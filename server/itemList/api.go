package itemList

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/generic/api"
	"resultra/datasheet/server/userRole"
)

func init() {

	itemListRouter := mux.NewRouter()

	itemListRouter.HandleFunc("/api/itemList/new", newListAPI)
	itemListRouter.HandleFunc("/api/itemList/get", getListAPI)
	itemListRouter.HandleFunc("/api/itemList/list", getItemListListAPI)

	itemListRouter.HandleFunc("/api/itemList/setName", setItemListName)
	itemListRouter.HandleFunc("/api/itemList/setDefaultSortRules", setDefaultSortRules)
	itemListRouter.HandleFunc("/api/itemList/setDefaultFilterRules", setDefaultFilterRules)

	itemListRouter.HandleFunc("/api/itemList/validateListName", validateListNameAPI)
	itemListRouter.HandleFunc("/api/itemList/validateNewListName", validateNewItemListNameAPI)

	http.Handle("/api/itemList/", itemListRouter)
}

func newListAPI(w http.ResponseWriter, r *http.Request) {

	var params NewItemListParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if verifyErr := userRole.VerifyCurrUserIsDatabaseAdminForTable(
		r, params.ParentTableID); verifyErr != nil {
		api.WriteErrorResponse(w, verifyErr)
		return
	}

	if formRef, err := newItemList(params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *formRef)
	}

}

type GetListParams struct {
	ListID string `json:"listID"`
}

func getListAPI(w http.ResponseWriter, r *http.Request) {

	var params GetListParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if theList, err := GetItemList(params.ListID); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *theList)
	}

}

type GetItemListLIstParams struct {
	DatabaseID string `json:"databaseID"`
}

func getItemListListAPI(w http.ResponseWriter, r *http.Request) {

	var params GetItemListLIstParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if listList, err := getDatabaseItemLists(params.DatabaseID); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, listList)
	}

}

func processItemListPropUpdate(w http.ResponseWriter, r *http.Request, propUpdater ItemListPropUpdater) {
	if updatedList, err := updateItemListProps(propUpdater); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, updatedList)
	}
}

func setItemListName(w http.ResponseWriter, r *http.Request) {
	var params SetItemListNameParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processItemListPropUpdate(w, r, params)
}

func setDefaultFilterRules(w http.ResponseWriter, r *http.Request) {
	var params SetFilterRulesParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processItemListPropUpdate(w, r, params)
}

func setDefaultSortRules(w http.ResponseWriter, r *http.Request) {
	var params SetDefaultSortRulesParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processItemListPropUpdate(w, r, params)
}

func validateListNameAPI(w http.ResponseWriter, r *http.Request) {

	listName := r.FormValue("listName")
	listID := r.FormValue("listID")

	if err := validateItemListName(listID, listName); err != nil {
		api.WriteJSONResponse(w, fmt.Sprintf("%v", err))
		return
	}

	response := true
	api.WriteJSONResponse(w, response)

}

func validateNewItemListNameAPI(w http.ResponseWriter, r *http.Request) {

	listName := r.FormValue("listName")
	databaseID := r.FormValue("databaseID")

	if err := validateNewItemListName(databaseID, listName); err != nil {
		api.WriteJSONResponse(w, fmt.Sprintf("%v", err))
		return
	}

	response := true
	api.WriteJSONResponse(w, response)

}
