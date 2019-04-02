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

	buttonRouter.HandleFunc("/api/frm/formButton/new", newButton)

	buttonRouter.HandleFunc("/api/frm/formButton/get", getButtonAPI)

	buttonRouter.HandleFunc("/api/frm/formButton/resize", resizeButton)
	buttonRouter.HandleFunc("/api/frm/formButton/setPopupBehavior", setPopupBehavior)
	buttonRouter.HandleFunc("/api/frm/formButton/setDefaultVals", setDefaultVals)
	buttonRouter.HandleFunc("/api/frm/formButton/setSize", setSize)
	buttonRouter.HandleFunc("/api/frm/formButton/setColorScheme", setColorScheme)
	buttonRouter.HandleFunc("/api/frm/formButton/setIcon", setIcon)
	buttonRouter.HandleFunc("/api/frm/formButton/setVisibility", setVisibility)
	buttonRouter.HandleFunc("/api/frm/formButton/setButtonLabelFormat", setButtonLabelFormat)

	http.Handle("/api/frm/formButton/", buttonRouter)
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
	ButtonID string `json:"buttonID"`
}

func getButtonAPI(w http.ResponseWriter, r *http.Request) {

	params := GetButtonParams{}
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if buttonRef, err := getButtonFromButtonID(trackerDBHandle, params.ButtonID); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *buttonRef)
	}

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

func resizeButton(w http.ResponseWriter, r *http.Request) {
	var params ButtonResizeParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processButtonPropUpdate(w, r, params)
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

func setVisibility(w http.ResponseWriter, r *http.Request) {
	var params ButtonVisibilityParams
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
