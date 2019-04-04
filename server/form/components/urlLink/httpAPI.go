// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package urlLink

import (
	"github.com/gorilla/mux"
	"github.com/resultra/resultra/server/common/databaseWrapper"
	"github.com/resultra/resultra/server/generic/api"
	"net/http"
)

func init() {
	urlLinkRouter := mux.NewRouter()

	urlLinkRouter.HandleFunc("/api/frm/urlLink/new", newUrlLink)
	urlLinkRouter.HandleFunc("/api/frm/urlLink/resize", resizeUrlLink)
	urlLinkRouter.HandleFunc("/api/frm/urlLink/setLabelFormat", setLabelFormat)
	urlLinkRouter.HandleFunc("/api/frm/urlLink/setVisibility", setVisibility)
	urlLinkRouter.HandleFunc("/api/frm/urlLink/setPermissions", setPermissions)
	urlLinkRouter.HandleFunc("/api/frm/urlLink/setValidation", setValidation)
	urlLinkRouter.HandleFunc("/api/frm/urlLink/setClearValueSupported", setClearValueSupported)
	urlLinkRouter.HandleFunc("/api/frm/urlLink/setHelpPopupMsg", setHelpPopupMsg)

	urlLinkRouter.HandleFunc("/api/frm/urlLink/validateInput", validateInputAPI)

	http.Handle("/api/frm/urlLink/", urlLinkRouter)
}

func newUrlLink(w http.ResponseWriter, r *http.Request) {

	urlLinkParams := NewUrlLinkParams{}
	if err := api.DecodeJSONRequest(r, &urlLinkParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if urlLinkRef, err := saveNewUrlLink(trackerDBHandle, urlLinkParams); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *urlLinkRef)
	}

}

func validateInputAPI(w http.ResponseWriter, r *http.Request) {

	var params UrlLinkValidateInputParams
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

func processUrlLinkPropUpdate(w http.ResponseWriter, r *http.Request, propUpdater UrlLinkPropUpdater) {
	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}
	if urlLinkRef, err := updateUrlLinkProps(trackerDBHandle, propUpdater); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, urlLinkRef)
	}
}

func resizeUrlLink(w http.ResponseWriter, r *http.Request) {
	var resizeParams UrlLinkResizeParams
	if err := api.DecodeJSONRequest(r, &resizeParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processUrlLinkPropUpdate(w, r, resizeParams)
}

func setLabelFormat(w http.ResponseWriter, r *http.Request) {
	var params UrlLinkLabelFormatParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processUrlLinkPropUpdate(w, r, params)
}

func setVisibility(w http.ResponseWriter, r *http.Request) {
	var params UrlLinkVisibilityParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processUrlLinkPropUpdate(w, r, params)
}

func setPermissions(w http.ResponseWriter, r *http.Request) {
	var params UrlLinkPermissionParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processUrlLinkPropUpdate(w, r, params)
}

func setValidation(w http.ResponseWriter, r *http.Request) {
	var params UrlLinkValidationParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processUrlLinkPropUpdate(w, r, params)
}

func setClearValueSupported(w http.ResponseWriter, r *http.Request) {
	var params UrlLinkClearValueSupportedParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processUrlLinkPropUpdate(w, r, params)
}

func setHelpPopupMsg(w http.ResponseWriter, r *http.Request) {
	var params HelpPopupMsgParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processUrlLinkPropUpdate(w, r, params)
}
