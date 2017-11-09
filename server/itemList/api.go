package itemList

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/common/databaseWrapper"
	"resultra/datasheet/server/generic/api"
	"resultra/datasheet/server/userRole"
)

func init() {

	itemListRouter := mux.NewRouter()

	itemListRouter.HandleFunc("/api/itemList/new", newListAPI)
	itemListRouter.HandleFunc("/api/itemList/get", getListAPI)

	itemListRouter.HandleFunc("/api/itemList/list", getItemListListAPI)
	itemListRouter.HandleFunc("/api/itemList/getUserItemListList", getUserItemListListAPI)

	itemListRouter.HandleFunc("/api/itemList/setName", setItemListName)
	itemListRouter.HandleFunc("/api/itemList/setDefaultSortRules", setDefaultSortRules)
	itemListRouter.HandleFunc("/api/itemList/setDefaultSortFields", setDefaultSortFields)

	itemListRouter.HandleFunc("/api/itemList/setDefaultFilterRules", setDefaultFilterRules)
	itemListRouter.HandleFunc("/api/itemList/setDefaultFilterFields", setDefaultFilterFields)

	itemListRouter.HandleFunc("/api/itemList/setPreFilterRules", setPreFilterRules)

	itemListRouter.HandleFunc("/api/itemList/setAlternateViews", setAlternateViews)

	itemListRouter.HandleFunc("/api/itemList/setDefaultView", setDefaultView)

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

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if verifyErr := userRole.VerifyCurrUserIsDatabaseAdmin(trackerDBHandle,
		r, params.ParentDatabaseID); verifyErr != nil {
		api.WriteErrorResponse(w, verifyErr)
		return
	}

	if formRef, err := newItemList(trackerDBHandle, params); err != nil {
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

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if theList, err := GetItemList(trackerDBHandle, params.ListID); err != nil {
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

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if listList, err := GetAllSortedItemLists(trackerDBHandle, params.DatabaseID); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, listList)
	}

}

func getUserItemListListAPI(w http.ResponseWriter, r *http.Request) {

	var params GetItemListLIstParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if listList, err := GetAllUserSortedItemLists(r, params.DatabaseID); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, listList)
	}

}

func processItemListPropUpdate(w http.ResponseWriter, r *http.Request, propUpdater ItemListPropUpdater) {

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if updatedList, err := updateItemListProps(trackerDBHandle, propUpdater); err != nil {
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

func setPreFilterRules(w http.ResponseWriter, r *http.Request) {
	var params SetPreFilterRulesParams
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

func setAlternateViews(w http.ResponseWriter, r *http.Request) {
	var params SetAlternateViewsParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processItemListPropUpdate(w, r, params)
}

func setDefaultView(w http.ResponseWriter, r *http.Request) {
	var params SetDefaultViewParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processItemListPropUpdate(w, r, params)
}

func setDefaultFilterFields(w http.ResponseWriter, r *http.Request) {
	var params SetDefaultFilterFieldsParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processItemListPropUpdate(w, r, params)
}

func setDefaultSortFields(w http.ResponseWriter, r *http.Request) {
	var params SetDefaultSortFieldsParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processItemListPropUpdate(w, r, params)
}

func validateListNameAPI(w http.ResponseWriter, r *http.Request) {

	listName := r.FormValue("listName")
	listID := r.FormValue("listID")

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if err := validateItemListName(trackerDBHandle, listID, listName); err != nil {
		api.WriteJSONResponse(w, fmt.Sprintf("%v", err))
		return
	}

	response := true
	api.WriteJSONResponse(w, response)

}

func validateNewItemListNameAPI(w http.ResponseWriter, r *http.Request) {

	listName := r.FormValue("listName")
	databaseID := r.FormValue("databaseID")

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if err := validateNewItemListName(trackerDBHandle, databaseID, listName); err != nil {
		api.WriteJSONResponse(w, fmt.Sprintf("%v", err))
		return
	}

	response := true
	api.WriteJSONResponse(w, response)

}
