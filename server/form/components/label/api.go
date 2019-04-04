// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package label

import (
	"github.com/gorilla/mux"
	"github.com/resultra/resultra/server/common/databaseWrapper"
	"github.com/resultra/resultra/server/generic/api"
	"net/http"
)

func init() {
	labelRouter := mux.NewRouter()

	labelRouter.HandleFunc("/api/frm/label/new", newLabel)
	labelRouter.HandleFunc("/api/frm/label/resize", resizeLabel)
	labelRouter.HandleFunc("/api/frm/label/setLabelFormat", setLabelFormat)
	labelRouter.HandleFunc("/api/frm/label/setVisibility", setVisibility)
	labelRouter.HandleFunc("/api/frm/label/setPermissions", setPermissions)
	labelRouter.HandleFunc("/api/frm/label/setHelpPopupMsg", setHelpPopupMsg)

	labelRouter.HandleFunc("/api/frm/label/setValidation", setValidation)
	labelRouter.HandleFunc("/api/frm/label/validateInput", validateInputAPI)

	http.Handle("/api/frm/label/", labelRouter)
}

func newLabel(w http.ResponseWriter, r *http.Request) {

	labelParams := NewLabelParams{}
	if err := api.DecodeJSONRequest(r, &labelParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if labelRef, err := saveNewLabel(trackerDBHandle, labelParams); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *labelRef)
	}

}

func validateInputAPI(w http.ResponseWriter, r *http.Request) {

	var params ValidateInputParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	validationResp := validateInput(trackerDBHandle, params)
	api.WriteJSONResponse(w, validationResp)
}

func processLabelPropUpdate(w http.ResponseWriter, r *http.Request, propUpdater LabelPropUpdater) {
	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if labelRef, err := updateLabelProps(trackerDBHandle, propUpdater); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, labelRef)
	}
}

func resizeLabel(w http.ResponseWriter, r *http.Request) {
	var resizeParams LabelResizeParams
	if err := api.DecodeJSONRequest(r, &resizeParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processLabelPropUpdate(w, r, resizeParams)
}

func setLabelFormat(w http.ResponseWriter, r *http.Request) {
	var params LabelLabelFormatParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processLabelPropUpdate(w, r, params)
}

func setVisibility(w http.ResponseWriter, r *http.Request) {
	var params LabelVisibilityParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processLabelPropUpdate(w, r, params)
}

func setPermissions(w http.ResponseWriter, r *http.Request) {
	var params LabelPermissionParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processLabelPropUpdate(w, r, params)
}

func setValidation(w http.ResponseWriter, r *http.Request) {
	var params LabelValidationParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processLabelPropUpdate(w, r, params)
}

func setHelpPopupMsg(w http.ResponseWriter, r *http.Request) {
	var params HelpPopupMsgParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processLabelPropUpdate(w, r, params)
}
