// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package emailAddr

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/tracker/server/common/databaseWrapper"
	"resultra/tracker/server/generic/api"
)

func init() {
	emailAddrRouter := mux.NewRouter()

	emailAddrRouter.HandleFunc("/api/frm/emailAddr/new", newEmailAddr)
	emailAddrRouter.HandleFunc("/api/frm/emailAddr/resize", resizeEmailAddr)
	emailAddrRouter.HandleFunc("/api/frm/emailAddr/setLabelFormat", setLabelFormat)
	emailAddrRouter.HandleFunc("/api/frm/emailAddr/setVisibility", setVisibility)
	emailAddrRouter.HandleFunc("/api/frm/emailAddr/setPermissions", setPermissions)
	emailAddrRouter.HandleFunc("/api/frm/emailAddr/setValidation", setValidation)
	emailAddrRouter.HandleFunc("/api/frm/emailAddr/setClearValueSupported", setClearValueSupported)
	emailAddrRouter.HandleFunc("/api/frm/emailAddr/setHelpPopupMsg", setHelpPopupMsg)

	emailAddrRouter.HandleFunc("/api/frm/emailAddr/validateInput", validateInputAPI)

	http.Handle("/api/frm/emailAddr/", emailAddrRouter)
}

func newEmailAddr(w http.ResponseWriter, r *http.Request) {

	emailAddrParams := NewEmailAddrParams{}
	if err := api.DecodeJSONRequest(r, &emailAddrParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}
	if emailAddrRef, err := saveNewEmailAddr(trackerDBHandle, emailAddrParams); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *emailAddrRef)
	}

}

func validateInputAPI(w http.ResponseWriter, r *http.Request) {

	var params EmailAddrValidateInputParams
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

func processEmailAddrPropUpdate(w http.ResponseWriter, r *http.Request, propUpdater EmailAddrPropUpdater) {
	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}
	if emailAddrRef, err := updateEmailAddrProps(trackerDBHandle, propUpdater); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, emailAddrRef)
	}
}

func resizeEmailAddr(w http.ResponseWriter, r *http.Request) {
	var resizeParams EmailAddrResizeParams
	if err := api.DecodeJSONRequest(r, &resizeParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processEmailAddrPropUpdate(w, r, resizeParams)
}

func setLabelFormat(w http.ResponseWriter, r *http.Request) {
	var params EmailAddrLabelFormatParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processEmailAddrPropUpdate(w, r, params)
}

func setVisibility(w http.ResponseWriter, r *http.Request) {
	var params EmailAddrVisibilityParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processEmailAddrPropUpdate(w, r, params)
}

func setPermissions(w http.ResponseWriter, r *http.Request) {
	var params EmailAddrPermissionParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processEmailAddrPropUpdate(w, r, params)
}

func setValidation(w http.ResponseWriter, r *http.Request) {
	var params EmailAddrValidationParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processEmailAddrPropUpdate(w, r, params)
}

func setClearValueSupported(w http.ResponseWriter, r *http.Request) {
	var params EmailAddrClearValueSupportedParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processEmailAddrPropUpdate(w, r, params)
}

func setHelpPopupMsg(w http.ResponseWriter, r *http.Request) {
	var params HelpPopupMsgParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processEmailAddrPropUpdate(w, r, params)
}
