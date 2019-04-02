// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package userSelection

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/resultra/resultra/server/common/databaseWrapper"
	"github.com/resultra/resultra/server/generic/api"
)

func init() {
	userSelectionRouter := mux.NewRouter()

	userSelectionRouter.HandleFunc("/api/tableView/userSelection/new", newUserSelection)

	userSelectionRouter.HandleFunc("/api/tableView/userSelection/get", getUserSelectionAPI)

	userSelectionRouter.HandleFunc("/api/tableView/userSelection/setLabelFormat", setLabelFormat)
	userSelectionRouter.HandleFunc("/api/tableView/userSelection/setClearValueSupported", setClearValueSupported)
	userSelectionRouter.HandleFunc("/api/tableView/userSelection/setHelpPopupMsg", setHelpPopupMsg)

	userSelectionRouter.HandleFunc("/api/tableView/userSelection/setSelectableRoles", setSelectableRoles)
	userSelectionRouter.HandleFunc("/api/tableView/userSelection/setCurrUserSelectable", setCurrUserSelectable)

	userSelectionRouter.HandleFunc("/api/tableView/userSelection/setPermissions", setPermissions)
	userSelectionRouter.HandleFunc("/api/tableView/userSelection/setValidation", setValidation)
	userSelectionRouter.HandleFunc("/api/tableView/userSelection/validateInput", validateInputAPI)

	http.Handle("/api/tableView/userSelection/", userSelectionRouter)
}

func newUserSelection(w http.ResponseWriter, r *http.Request) {

	userSelectionParams := NewUserSelectionParams{}
	if err := api.DecodeJSONRequest(r, &userSelectionParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if userSelectionRef, err := saveNewUserSelection(trackerDBHandle, userSelectionParams); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *userSelectionRef)
	}

}

type GetUserSelectionParams struct {
	ParentTableID   string `json:"parentTableID"`
	UserSelectionID string `json:"userSelectionID"`
}

func getUserSelectionAPI(w http.ResponseWriter, r *http.Request) {

	var params GetUserSelectionParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	textInput, err := getUserSelection(trackerDBHandle, params.ParentTableID, params.UserSelectionID)
	if err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	api.WriteJSONResponse(w, *textInput)
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

func processUserSelectionPropUpdate(w http.ResponseWriter, r *http.Request, propUpdater UserSelectionPropUpdater) {

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if userSelectionRef, err := updateUserSelectionProps(trackerDBHandle, propUpdater); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, userSelectionRef)
	}
}

func setLabelFormat(w http.ResponseWriter, r *http.Request) {
	var params UserSelectionLabelFormatParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processUserSelectionPropUpdate(w, r, params)
}

func setPermissions(w http.ResponseWriter, r *http.Request) {
	var params UserSelectionPermissionParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processUserSelectionPropUpdate(w, r, params)
}

func setValidation(w http.ResponseWriter, r *http.Request) {
	var params UserSelectionValidationParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processUserSelectionPropUpdate(w, r, params)
}

func setClearValueSupported(w http.ResponseWriter, r *http.Request) {
	var params UserSelectionClearValueSupportedParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processUserSelectionPropUpdate(w, r, params)
}

func setHelpPopupMsg(w http.ResponseWriter, r *http.Request) {
	var params HelpPopupMsgParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processUserSelectionPropUpdate(w, r, params)
}

func setSelectableRoles(w http.ResponseWriter, r *http.Request) {
	var params SelectableRoleParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processUserSelectionPropUpdate(w, r, params)
}

func setCurrUserSelectable(w http.ResponseWriter, r *http.Request) {
	var params CurrUserSelectableParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processUserSelectionPropUpdate(w, r, params)
}
