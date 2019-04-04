// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package selection

import (
	"github.com/gorilla/mux"
	"github.com/resultra/resultra/server/common/databaseWrapper"
	"github.com/resultra/resultra/server/generic/api"
	"net/http"
)

func init() {
	selectionRouter := mux.NewRouter()

	selectionRouter.HandleFunc("/api/frm/selection/new", newSelection)
	selectionRouter.HandleFunc("/api/frm/selection/resize", resizeSelection)
	selectionRouter.HandleFunc("/api/frm/selection/setLabelFormat", setLabelFormat)
	selectionRouter.HandleFunc("/api/frm/selection/setValueList", setValueList)

	selectionRouter.HandleFunc("/api/frm/selection/setClearValueSupported", setClearValueSupported)

	selectionRouter.HandleFunc("/api/frm/selection/setVisibility", setVisibility)
	selectionRouter.HandleFunc("/api/frm/selection/setPermissions", setPermissions)

	http.Handle("/api/frm/selection/", selectionRouter)
}

func newSelection(w http.ResponseWriter, r *http.Request) {

	selectionParams := NewSelectionParams{}
	if err := api.DecodeJSONRequest(r, &selectionParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}
	if selectionRef, err := saveNewSelection(trackerDBHandle, selectionParams); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *selectionRef)
	}

}

func processSelectionPropUpdate(w http.ResponseWriter, r *http.Request, propUpdater SelectionPropUpdater) {
	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}
	if selectionRef, err := updateSelectionProps(trackerDBHandle, propUpdater); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, selectionRef)
	}
}

func resizeSelection(w http.ResponseWriter, r *http.Request) {
	var resizeParams SelectionResizeParams
	if err := api.DecodeJSONRequest(r, &resizeParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processSelectionPropUpdate(w, r, resizeParams)
}

func setLabelFormat(w http.ResponseWriter, r *http.Request) {
	var params SelectionLabelFormatParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processSelectionPropUpdate(w, r, params)

}

func setVisibility(w http.ResponseWriter, r *http.Request) {
	var params SelectionVisibilityParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processSelectionPropUpdate(w, r, params)

}

func setPermissions(w http.ResponseWriter, r *http.Request) {
	var params SelectionPermissionParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processSelectionPropUpdate(w, r, params)

}

func setClearValueSupported(w http.ResponseWriter, r *http.Request) {
	var params SelectionClearValueSupportedParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processSelectionPropUpdate(w, r, params)
}

func setValueList(w http.ResponseWriter, r *http.Request) {
	var params SelectionValueListParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processSelectionPropUpdate(w, r, params)

}
