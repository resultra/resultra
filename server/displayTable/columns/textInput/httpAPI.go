// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package textInput

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/resultra/resultra/server/common/databaseWrapper"
	"github.com/resultra/resultra/server/generic/api"
)

func init() {
	textInputRouter := mux.NewRouter()

	textInputRouter.HandleFunc("/api/tableView/textInput/new", newTextInput)

	textInputRouter.HandleFunc("/api/tableView/textInput/get", getTextInputAPI)

	textInputRouter.HandleFunc("/api/tableView/textInput/setLabelFormat", setLabelFormat)
	textInputRouter.HandleFunc("/api/tableView/textInput/setPermissions", setPermissions)
	textInputRouter.HandleFunc("/api/tableView/textInput/setValueList", setValueList)
	textInputRouter.HandleFunc("/api/tableView/textInput/setValidation", setValidation)
	textInputRouter.HandleFunc("/api/tableView/textInput/setClearValueSupported", setClearValueSupported)
	textInputRouter.HandleFunc("/api/tableView/textInput/setHelpPopupMsg", setHelpPopupMsg)

	textInputRouter.HandleFunc("/api/tableView/textInput/validateInput", validateInputAPI)

	http.Handle("/api/tableView/textInput/", textInputRouter)
}

func newTextInput(w http.ResponseWriter, r *http.Request) {

	textInputParams := NewTextInputParams{}
	if err := api.DecodeJSONRequest(r, &textInputParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if textInputRef, err := saveNewTextInput(trackerDBHandle, textInputParams); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *textInputRef)
	}

}

type GetTextInputParams struct {
	ParentTableID string `json:"parentTableID"`
	TextInputID   string `json:"textInputID"`
}

func getTextInputAPI(w http.ResponseWriter, r *http.Request) {

	var params GetTextInputParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	textInput, err := getTextInput(trackerDBHandle, params.ParentTableID, params.TextInputID)
	if err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	api.WriteJSONResponse(w, *textInput)
}

func validateInputAPI(w http.ResponseWriter, r *http.Request) {

	var params TextInputValidateInputParams
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

func processTextInputPropUpdate(w http.ResponseWriter, r *http.Request, propUpdater TextInputPropUpdater) {

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if textInputRef, err := updateTextInputProps(trackerDBHandle, propUpdater); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, textInputRef)
	}
}

func setLabelFormat(w http.ResponseWriter, r *http.Request) {
	var params TextInputLabelFormatParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processTextInputPropUpdate(w, r, params)
}

func setPermissions(w http.ResponseWriter, r *http.Request) {
	var params TextInputPermissionParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processTextInputPropUpdate(w, r, params)
}

func setValueList(w http.ResponseWriter, r *http.Request) {
	var params TextInputValueListParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processTextInputPropUpdate(w, r, params)
}

func setValidation(w http.ResponseWriter, r *http.Request) {
	var params TextInputValidationParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processTextInputPropUpdate(w, r, params)
}

func setClearValueSupported(w http.ResponseWriter, r *http.Request) {
	var params TextInputClearValueSupportedParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processTextInputPropUpdate(w, r, params)
}

func setHelpPopupMsg(w http.ResponseWriter, r *http.Request) {
	var params HelpPopupMsgParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processTextInputPropUpdate(w, r, params)
}
