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

	headerRouter.HandleFunc("/api/frm/header/new", newHeader)
	headerRouter.HandleFunc("/api/frm/header/resize", resizeHeader)
	headerRouter.HandleFunc("/api/frm/header/setLabel", setHeaderLabel)
	headerRouter.HandleFunc("/api/frm/header/setUnderlined", setUnderlined)
	headerRouter.HandleFunc("/api/frm/header/setSize", setSize)
	headerRouter.HandleFunc("/api/frm/header/setVisibility", setVisibility)

	http.Handle("/api/frm/header/", headerRouter)
}

func newHeader(w http.ResponseWriter, r *http.Request) {

	params := NewHeaderParams{}
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if headerRef, err := saveNewHeader(trackerDBHandle, params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *headerRef)
	}

}

func processHeaderPropUpdate(w http.ResponseWriter, r *http.Request, propUpdater HeaderPropUpdater) {
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

func resizeHeader(w http.ResponseWriter, r *http.Request) {
	var params HeaderResizeParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processHeaderPropUpdate(w, r, params)
}

func setHeaderLabel(w http.ResponseWriter, r *http.Request) {
	var params HeaderLabelParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processHeaderPropUpdate(w, r, params)
}

func setUnderlined(w http.ResponseWriter, r *http.Request) {
	var params HeaderUnderlinedParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processHeaderPropUpdate(w, r, params)
}

func setSize(w http.ResponseWriter, r *http.Request) {
	var params HeaderSizeParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processHeaderPropUpdate(w, r, params)
}

func setVisibility(w http.ResponseWriter, r *http.Request) {
	var params HeaderVisibilityParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processHeaderPropUpdate(w, r, params)
}
