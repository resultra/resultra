// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package formButton

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/resultra/resultra/server/common/databaseWrapper"
	"github.com/resultra/resultra/server/generic/api"
)

func init() {
	buttonRouter := mux.NewRouter()

	buttonRouter.HandleFunc("/api/tableView/formButton/new", newButton)

	buttonRouter.HandleFunc("/api/tableView/formButton/get", getButtonAPI)
	buttonRouter.HandleFunc("/api/tableView/formButton/getFromButtonID", getButtonFromIDAPI)

	buttonRouter.HandleFunc("/api/tableView/formButton/setPopupBehavior", setPopupBehavior)

	buttonRouter.HandleFunc("/api/tableView/formButton/setDefaultVals", setDefaultVals)
	buttonRouter.HandleFunc("/api/tableView/formButton/setSize", setSize)
	buttonRouter.HandleFunc("/api/tableView/formButton/setColorScheme", setColorScheme)
	buttonRouter.HandleFunc("/api/tableView/formButton/setIcon", setIcon)
	buttonRouter.HandleFunc("/api/tableView/formButton/setButtonLabelFormat", setButtonLabelFormat)

	http.Handle("/api/tableView/formButton/", buttonRouter)
}

func newButton(w http.ResponseWriter, r *http.Request) {

	params := NewButtonParams{}
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if headerRef, err := saveNewButton(trackerDBHandle, params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *headerRef)
	}

}

type GetButtonParams struct {
	ParentTableID string `json:"parentTableID"`
	ButtonID      string `json:"buttonID"`
}

func getButtonAPI(w http.ResponseWriter, r *http.Request) {

	var params GetButtonParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	formButton, err := getButton(trackerDBHandle, params.ParentTableID, params.ButtonID)
	if err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	api.WriteJSONResponse(w, *formButton)
}

type GetButtonFromIDParams struct {
	ButtonID string `json:"buttonID"`
}

func getButtonFromIDAPI(w http.ResponseWriter, r *http.Request) {

	var params GetButtonFromIDParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	formButton, err := getButtonFromButtonID(trackerDBHandle, params.ButtonID)
	if err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	api.WriteJSONResponse(w, *formButton)
}

func processButtonPropUpdate(w http.ResponseWriter, r *http.Request, propUpdater ButtonPropUpdater) {

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if headerRef, err := updateButtonProps(trackerDBHandle, propUpdater); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, headerRef)
	}
}

func setPopupBehavior(w http.ResponseWriter, r *http.Request) {
	var params ButtonBehaviorParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processButtonPropUpdate(w, r, params)
}

func setDefaultVals(w http.ResponseWriter, r *http.Request) {
	var params ButtonDefaultValParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processButtonPropUpdate(w, r, params)
}

func setSize(w http.ResponseWriter, r *http.Request) {
	var params ButtonSizeParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processButtonPropUpdate(w, r, params)
}

func setColorScheme(w http.ResponseWriter, r *http.Request) {
	var params ButtonColorSchemeParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processButtonPropUpdate(w, r, params)
}

func setIcon(w http.ResponseWriter, r *http.Request) {
	var params ButtonIconParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processButtonPropUpdate(w, r, params)
}

func setButtonLabelFormat(w http.ResponseWriter, r *http.Request) {
	var params ButtonLabelFormatParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processButtonPropUpdate(w, r, params)
}
