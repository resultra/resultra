// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package textSelection

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/resultra/resultra/server/common/databaseWrapper"
	"github.com/resultra/resultra/server/generic/api"
)

func init() {
	textSelectionRouter := mux.NewRouter()

	textSelectionRouter.HandleFunc("/api/tableView/textSelection/new", newTextSelection)

	textSelectionRouter.HandleFunc("/api/tableView/textSelection/get", getTextSelectionAPI)

	textSelectionRouter.HandleFunc("/api/tableView/textSelection/setLabelFormat", setLabelFormat)
	textSelectionRouter.HandleFunc("/api/tableView/textSelection/setPermissions", setPermissions)
	textSelectionRouter.HandleFunc("/api/tableView/textSelection/setValueList", setValueList)
	textSelectionRouter.HandleFunc("/api/tableView/textSelection/setValidation", setValidation)
	textSelectionRouter.HandleFunc("/api/tableView/textSelection/setClearValueSupported", setClearValueSupported)
	textSelectionRouter.HandleFunc("/api/tableView/textSelection/setHelpPopupMsg", setHelpPopupMsg)

	textSelectionRouter.HandleFunc("/api/tableView/textSelection/validateInput", validateInputAPI)

	http.Handle("/api/tableView/textSelection/", textSelectionRouter)
}

func newTextSelection(w http.ResponseWriter, r *http.Request) {

	textSelectionParams := NewTextSelectionParams{}
	if err := api.DecodeJSONRequest(r, &textSelectionParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if textSelectionRef, err := saveNewTextSelection(trackerDBHandle, textSelectionParams); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *textSelectionRef)
	}

}

type GetTextSelectionParams struct {
	ParentTableID string `json:"parentTableID"`
	SelectionID   string `json:"selectionID"`
}

func getTextSelectionAPI(w http.ResponseWriter, r *http.Request) {

	var params GetTextSelectionParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	textSelection, err := getTextSelection(trackerDBHandle, params.ParentTableID, params.SelectionID)
	if err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	api.WriteJSONResponse(w, *textSelection)
}

func validateInputAPI(w http.ResponseWriter, r *http.Request) {

	var params TextSelectionValidateInputParams
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

func processTextSelectionPropUpdate(w http.ResponseWriter, r *http.Request, propUpdater TextSelectionPropUpdater) {

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if textInputRef, err := updateTextSelectionProps(trackerDBHandle, propUpdater); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, textInputRef)
	}
}

func setLabelFormat(w http.ResponseWriter, r *http.Request) {
	var params TextSelectionLabelFormatParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processTextSelectionPropUpdate(w, r, params)
}

func setPermissions(w http.ResponseWriter, r *http.Request) {
	var params TextSelectionPermissionParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processTextSelectionPropUpdate(w, r, params)
}

func setValueList(w http.ResponseWriter, r *http.Request) {
	var params TextSelectionValueListParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processTextSelectionPropUpdate(w, r, params)
}

func setValidation(w http.ResponseWriter, r *http.Request) {
	var params TextSelectionValidationParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processTextSelectionPropUpdate(w, r, params)
}

func setClearValueSupported(w http.ResponseWriter, r *http.Request) {
	var params TextSelectionClearValueSupportedParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processTextSelectionPropUpdate(w, r, params)
}

func setHelpPopupMsg(w http.ResponseWriter, r *http.Request) {
	var params HelpPopupMsgParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processTextSelectionPropUpdate(w, r, params)
}
