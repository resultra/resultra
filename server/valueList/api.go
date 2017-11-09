package valueList

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/common/databaseWrapper"
	"resultra/datasheet/server/generic/api"
)

type DummyStructForInclude struct{ Val int64 }

func init() {
	valueListRouter := mux.NewRouter()

	valueListRouter.HandleFunc("/api/valueList/new", newValueListAPI)
	valueListRouter.HandleFunc("/api/valueList/get", getValueListAPI)
	valueListRouter.HandleFunc("/api/valueList/getList", getValueListsAPI)

	valueListRouter.HandleFunc("/api/valueList/setName", setName)
	valueListRouter.HandleFunc("/api/valueList/setValues", setValues)

	http.Handle("/api/valueList/", valueListRouter)
}

func newValueListAPI(w http.ResponseWriter, r *http.Request) {

	params := NewValueListParams{}
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	newPreset, err := newValueList(trackerDBHandle, params)
	if err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *newPreset)
	}

}

func getValueListAPI(w http.ResponseWriter, r *http.Request) {

	params := GetValueListParams{}
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	presets, err := GetValueList(trackerDBHandle, params.ValueListID)
	if err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, presets)
	}

}

func getValueListsAPI(w http.ResponseWriter, r *http.Request) {

	params := GetValueListsParams{}
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	presets, err := getAllValueLists(trackerDBHandle, params.ParentDatabaseID)
	if err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, presets)
	}

}

func processValueListPropUpdate(w http.ResponseWriter, r *http.Request, propUpdater ValueListPropUpdater) {

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if headerRef, err := updateValueListProps(trackerDBHandle, propUpdater); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, headerRef)
	}
}

func setName(w http.ResponseWriter, r *http.Request) {
	var params ValueListNameParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processValueListPropUpdate(w, r, params)
}

func setValues(w http.ResponseWriter, r *http.Request) {
	var params ValueListValuesParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processValueListPropUpdate(w, r, params)
}
