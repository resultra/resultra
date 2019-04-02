// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package header

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/resultra/resultra/server/common/databaseWrapper"
	"github.com/resultra/resultra/server/generic/api"
)

func init() {

	headerRouter := mux.NewRouter()

	headerRouter.HandleFunc("/api/dashboard/header/new", newHeaderAPI)

	headerRouter.HandleFunc("/api/dashboard/header/setTitle", setHeaderTitle)
	headerRouter.HandleFunc("/api/dashboard/header/setDimensions", setHeaderDimensions)
	headerRouter.HandleFunc("/api/dashboard/header/setSize", setSize)
	headerRouter.HandleFunc("/api/dashboard/header/setUnderlined", setUnderline)

	http.Handle("/api/dashboard/header/", headerRouter)
}

func newHeaderAPI(w http.ResponseWriter, r *http.Request) {

	var params NewHeaderParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if headerRef, err := newHeader(trackerDBHandle, params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, headerRef)
	}

}

func processHeaderPropUpdate(w http.ResponseWriter, r *http.Request, propUpdater HeaderPropertyUpdater) {

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if headerRef, err := updateHeaderProps(trackerDBHandle, propUpdater); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, headerRef)
	}
}

func setHeaderTitle(w http.ResponseWriter, r *http.Request) {
	var titleParams SetHeaderTitleParams
	if err := api.DecodeJSONRequest(r, &titleParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processHeaderPropUpdate(w, r, titleParams)
}

func setHeaderDimensions(w http.ResponseWriter, r *http.Request) {

	var params SetHeaderDimensionsParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processHeaderPropUpdate(w, r, params)
}

func setSize(w http.ResponseWriter, r *http.Request) {

	var params SetHeaderSizeParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processHeaderPropUpdate(w, r, params)
}

func setUnderline(w http.ResponseWriter, r *http.Request) {

	var params SetHeaderUnderlineParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processHeaderPropUpdate(w, r, params)
}
