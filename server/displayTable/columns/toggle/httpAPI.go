// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package toggle

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/tracker/server/common/databaseWrapper"
	"resultra/tracker/server/generic/api"
)

func init() {
	toggleRouter := mux.NewRouter()

	toggleRouter.HandleFunc("/api/tableView/toggle/new", newToggle)

	toggleRouter.HandleFunc("/api/tableView/toggle/get", getToggleAPI)

	toggleRouter.HandleFunc("/api/tableView/toggle/setOffColorScheme", setOffColorScheme)
	toggleRouter.HandleFunc("/api/tableView/toggle/setOnColorScheme", setOnColorScheme)
	toggleRouter.HandleFunc("/api/tableView/toggle/setOffLabel", setOffLabel)
	toggleRouter.HandleFunc("/api/tableView/toggle/setOnLabel", setOnLabel)
	toggleRouter.HandleFunc("/api/tableView/toggle/setClearValueSupported", setClearValueSupported)
	toggleRouter.HandleFunc("/api/tableView/toggle/setHelpPopupMsg", setHelpPopupMsg)

	toggleRouter.HandleFunc("/api/tableView/toggle/setLabelFormat", setLabelFormat)
	toggleRouter.HandleFunc("/api/tableView/toggle/setPermissions", setPermissions)

	toggleRouter.HandleFunc("/api/tableView/toggle/setValidation", setValidation)
	toggleRouter.HandleFunc("/api/tableView/toggle/validateInput", validateInputAPI)

	http.Handle("/api/tableView/toggle/", toggleRouter)
}

func newToggle(w http.ResponseWriter, r *http.Request) {

	toggleParams := NewToggleParams{}
	if err := api.DecodeJSONRequest(r, &toggleParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if toggleRef, err := saveNewToggle(trackerDBHandle, toggleParams); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *toggleRef)
	}

}

type GetToggleParams struct {
	ParentTableID string `json:"parentTableID"`
	ToggleID      string `json:"toggleID"`
}

func getToggleAPI(w http.ResponseWriter, r *http.Request) {

	var params GetToggleParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	toggle, err := getToggle(trackerDBHandle, params.ParentTableID, params.ToggleID)
	if err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	api.WriteJSONResponse(w, *toggle)

}

func validateInputAPI(w http.ResponseWriter, r *http.Request) {

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	var params ToggleValidateInputParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	validationResp := validateInput(trackerDBHandle, params)
	api.WriteJSONResponse(w, validationResp)
}

func processTogglePropUpdate(w http.ResponseWriter, r *http.Request, propUpdater TogglePropUpdater) {
	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if toggleRef, err := updateToggleProps(trackerDBHandle, propUpdater); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, toggleRef)
	}
}

func setOffColorScheme(w http.ResponseWriter, r *http.Request) {
	var params ToggleOffColorSchemeParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processTogglePropUpdate(w, r, params)
}

func setOnColorScheme(w http.ResponseWriter, r *http.Request) {
	var params ToggleOnColorSchemeParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processTogglePropUpdate(w, r, params)
}

func setOffLabel(w http.ResponseWriter, r *http.Request) {
	var params ToggleOffLabelParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processTogglePropUpdate(w, r, params)
}

func setOnLabel(w http.ResponseWriter, r *http.Request) {
	var params ToggleOnLabelParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processTogglePropUpdate(w, r, params)
}

func setLabelFormat(w http.ResponseWriter, r *http.Request) {
	var params ToggleLabelFormatParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processTogglePropUpdate(w, r, params)
}

func setPermissions(w http.ResponseWriter, r *http.Request) {
	var params TogglePermissionParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processTogglePropUpdate(w, r, params)
}

func setValidation(w http.ResponseWriter, r *http.Request) {
	var params ToggleValidationParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processTogglePropUpdate(w, r, params)
}

func setClearValueSupported(w http.ResponseWriter, r *http.Request) {
	var params ToggleClearValueSupportedParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processTogglePropUpdate(w, r, params)
}

func setHelpPopupMsg(w http.ResponseWriter, r *http.Request) {
	var params HelpPopupMsgParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processTogglePropUpdate(w, r, params)
}
