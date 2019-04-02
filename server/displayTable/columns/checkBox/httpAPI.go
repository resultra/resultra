// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package checkBox

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/resultra/resultra/server/common/databaseWrapper"
	"github.com/resultra/resultra/server/generic/api"
)

func init() {
	checkBoxRouter := mux.NewRouter()

	checkBoxRouter.HandleFunc("/api/tableView/checkBox/new", newCheckBox)

	checkBoxRouter.HandleFunc("/api/tableView/checkBox/get", getCheckBoxAPI)

	checkBoxRouter.HandleFunc("/api/tableView/checkBox/setColorScheme", setColorScheme)
	checkBoxRouter.HandleFunc("/api/tableView/checkBox/setStrikethrough", setStrikethrough)
	checkBoxRouter.HandleFunc("/api/tableView/checkBox/setLabelFormat", setLabelFormat)
	checkBoxRouter.HandleFunc("/api/tableView/checkBox/setPermissions", setPermissions)
	checkBoxRouter.HandleFunc("/api/tableView/checkBox/setClearValueSupported", setClearValueSupported)
	checkBoxRouter.HandleFunc("/api/tableView/checkBox/setHelpPopupMsg", setHelpPopupMsg)

	checkBoxRouter.HandleFunc("/api/tableView/checkBox/setValidation", setValidation)
	checkBoxRouter.HandleFunc("/api/tableView/checkBox/validateInput", validateInputAPI)

	http.Handle("/api/tableView/checkBox/", checkBoxRouter)
}

func newCheckBox(w http.ResponseWriter, r *http.Request) {

	checkBoxParams := NewCheckBoxParams{}
	if err := api.DecodeJSONRequest(r, &checkBoxParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if checkBoxRef, err := saveNewCheckBox(trackerDBHandle, checkBoxParams); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *checkBoxRef)
	}

}

type GetCheckBoxParams struct {
	ParentTableID string `json:"parentTableID"`
	CheckBoxID    string `json:"checkBoxID"`
}

func getCheckBoxAPI(w http.ResponseWriter, r *http.Request) {

	var params GetCheckBoxParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	checkBox, err := getCheckBox(trackerDBHandle, params.ParentTableID, params.CheckBoxID)
	if err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	api.WriteJSONResponse(w, *checkBox)
}

func validateInputAPI(w http.ResponseWriter, r *http.Request) {

	var params CheckBoxValidateInputParams
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

func processCheckBoxPropUpdate(w http.ResponseWriter, r *http.Request, propUpdater CheckBoxPropUpdater) {

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if checkBoxRef, err := updateCheckBoxProps(trackerDBHandle, propUpdater); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, checkBoxRef)
	}
}

func setColorScheme(w http.ResponseWriter, r *http.Request) {
	var params CheckBoxColorSchemeParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processCheckBoxPropUpdate(w, r, params)
}

func setStrikethrough(w http.ResponseWriter, r *http.Request) {
	var params CheckBoxStrikethroughParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processCheckBoxPropUpdate(w, r, params)
}

func setLabelFormat(w http.ResponseWriter, r *http.Request) {
	var params CheckBoxLabelFormatParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processCheckBoxPropUpdate(w, r, params)
}

func setPermissions(w http.ResponseWriter, r *http.Request) {
	var params CheckBoxPermissionParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processCheckBoxPropUpdate(w, r, params)
}

func setValidation(w http.ResponseWriter, r *http.Request) {
	var params CheckBoxValidationParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processCheckBoxPropUpdate(w, r, params)
}

func setClearValueSupported(w http.ResponseWriter, r *http.Request) {
	var params CheckBoxClearValueSupportedParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processCheckBoxPropUpdate(w, r, params)
}

func setHelpPopupMsg(w http.ResponseWriter, r *http.Request) {
	var params HelpPopupMsgParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processCheckBoxPropUpdate(w, r, params)
}
