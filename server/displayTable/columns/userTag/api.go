// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package userTag

import (
	"github.com/gorilla/mux"
	"github.com/resultra/resultra/server/common/databaseWrapper"
	"github.com/resultra/resultra/server/generic/api"
	"net/http"
)

func init() {
	userTagRouter := mux.NewRouter()

	userTagRouter.HandleFunc("/api/tableView/userTag/new", newUserTag)

	userTagRouter.HandleFunc("/api/tableView/userTag/get", getUserTagAPI)

	userTagRouter.HandleFunc("/api/tableView/userTag/setLabelFormat", setLabelFormat)
	userTagRouter.HandleFunc("/api/tableView/userTag/setClearValueSupported", setClearValueSupported)
	userTagRouter.HandleFunc("/api/tableView/userTag/setHelpPopupMsg", setHelpPopupMsg)

	userTagRouter.HandleFunc("/api/tableView/userTag/setSelectableRoles", setSelectableRoles)
	userTagRouter.HandleFunc("/api/tableView/userTag/setCurrUserSelectable", setCurrUserSelectable)

	userTagRouter.HandleFunc("/api/tableView/userTag/setPermissions", setPermissions)
	userTagRouter.HandleFunc("/api/tableView/userTag/setValidation", setValidation)
	userTagRouter.HandleFunc("/api/tableView/userTag/validateInput", validateInputAPI)

	http.Handle("/api/tableView/userTag/", userTagRouter)
}

func newUserTag(w http.ResponseWriter, r *http.Request) {

	userTagParams := NewUserTagParams{}
	if err := api.DecodeJSONRequest(r, &userTagParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if userTagRef, err := saveNewUserTag(trackerDBHandle, userTagParams); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *userTagRef)
	}

}

type GetUserTagParams struct {
	ParentTableID string `json:"parentTableID"`
	UserTagID     string `json:"userTagID"`
}

func getUserTagAPI(w http.ResponseWriter, r *http.Request) {

	var params GetUserTagParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	textInput, err := getUserTag(trackerDBHandle, params.ParentTableID, params.UserTagID)
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

func processUserTagPropUpdate(w http.ResponseWriter, r *http.Request, propUpdater UserTagPropUpdater) {

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if userTagRef, err := updateUserTagProps(trackerDBHandle, propUpdater); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, userTagRef)
	}
}

func setLabelFormat(w http.ResponseWriter, r *http.Request) {
	var params UserTagLabelFormatParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processUserTagPropUpdate(w, r, params)
}

func setPermissions(w http.ResponseWriter, r *http.Request) {
	var params UserTagPermissionParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processUserTagPropUpdate(w, r, params)
}

func setValidation(w http.ResponseWriter, r *http.Request) {
	var params UserTagValidationParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processUserTagPropUpdate(w, r, params)
}

func setClearValueSupported(w http.ResponseWriter, r *http.Request) {
	var params UserTagClearValueSupportedParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processUserTagPropUpdate(w, r, params)
}

func setHelpPopupMsg(w http.ResponseWriter, r *http.Request) {
	var params HelpPopupMsgParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processUserTagPropUpdate(w, r, params)
}

func setSelectableRoles(w http.ResponseWriter, r *http.Request) {
	var params SelectableRoleParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processUserTagPropUpdate(w, r, params)
}

func setCurrUserSelectable(w http.ResponseWriter, r *http.Request) {
	var params CurrUserSelectableParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processUserTagPropUpdate(w, r, params)
}
