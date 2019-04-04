// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package valueList

import (
	"github.com/gorilla/mux"
	"github.com/resultra/resultra/server/common/databaseWrapper"
	"github.com/resultra/resultra/server/generic/api"
	"net/http"
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
