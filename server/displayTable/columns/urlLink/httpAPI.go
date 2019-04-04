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

	urlLinkRouter.HandleFunc("/api/tableView/urlLink/new", newUrlLink)

	urlLinkRouter.HandleFunc("/api/tableView/urlLink/get", getUrlLinkAPI)

	urlLinkRouter.HandleFunc("/api/tableView/urlLink/setLabelFormat", setLabelFormat)
	urlLinkRouter.HandleFunc("/api/tableView/urlLink/setPermissions", setPermissions)
	urlLinkRouter.HandleFunc("/api/tableView/urlLink/setValidation", setValidation)
	urlLinkRouter.HandleFunc("/api/tableView/urlLink/setClearValueSupported", setClearValueSupported)
	urlLinkRouter.HandleFunc("/api/tableView/urlLink/setHelpPopupMsg", setHelpPopupMsg)

	urlLinkRouter.HandleFunc("/api/tableView/urlLink/validateInput", validateInputAPI)

	http.Handle("/api/tableView/urlLink/", urlLinkRouter)
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

type GetUrlLinkParams struct {
	ParentTableID string `json:"parentTableID"`
	UrlLinkID     string `json:"urlLinkID"`
}

func getUrlLinkAPI(w http.ResponseWriter, r *http.Request) {

	var params GetUrlLinkParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	urlLink, err := getUrlLink(trackerDBHandle, params.ParentTableID, params.UrlLinkID)
	if err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	api.WriteJSONResponse(w, *urlLink)
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

func setLabelFormat(w http.ResponseWriter, r *http.Request) {
	var params UrlLinkLabelFormatParams
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
