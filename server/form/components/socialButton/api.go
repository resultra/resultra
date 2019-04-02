// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package socialButton

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/resultra/resultra/server/common/databaseWrapper"
	"github.com/resultra/resultra/server/generic/api"
)

func init() {
	socialButtonRouter := mux.NewRouter()

	socialButtonRouter.HandleFunc("/api/frm/socialButton/new", newSocialButton)
	socialButtonRouter.HandleFunc("/api/frm/socialButton/resize", resizeSocialButton)
	socialButtonRouter.HandleFunc("/api/frm/socialButton/setTooltips", setTooltips)
	socialButtonRouter.HandleFunc("/api/frm/socialButton/setIcon", setIcon)
	socialButtonRouter.HandleFunc("/api/frm/socialButton/setLabelFormat", setLabelFormat)
	socialButtonRouter.HandleFunc("/api/frm/socialButton/setHelpPopupMsg", setHelpPopupMsg)

	socialButtonRouter.HandleFunc("/api/frm/socialButton/setVisibility", setVisibility)

	http.Handle("/api/frm/socialButton/", socialButtonRouter)
}

func newSocialButton(w http.ResponseWriter, r *http.Request) {

	params := NewSocialButtonParams{}
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if socialButtonRef, err := saveNewSocialButton(trackerDBHandle, params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *socialButtonRef)
	}

}

func processSocialButtonPropUpdate(w http.ResponseWriter, r *http.Request, propUpdater SocialButtonPropUpdater) {
	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}
	if socialButtonRef, err := updateSocialButtonProps(trackerDBHandle, propUpdater); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, socialButtonRef)
	}
}

func resizeSocialButton(w http.ResponseWriter, r *http.Request) {
	var resizeParams SocialButtonResizeParams
	if err := api.DecodeJSONRequest(r, &resizeParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processSocialButtonPropUpdate(w, r, resizeParams)
}

func setTooltips(w http.ResponseWriter, r *http.Request) {
	var tooltipParams SocialButtonTooltipParams
	if err := api.DecodeJSONRequest(r, &tooltipParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processSocialButtonPropUpdate(w, r, tooltipParams)
}

func setIcon(w http.ResponseWriter, r *http.Request) {
	var params SocialButtonIconParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processSocialButtonPropUpdate(w, r, params)
}

func setLabelFormat(w http.ResponseWriter, r *http.Request) {
	var params SocialButtonLabelFormatParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processSocialButtonPropUpdate(w, r, params)
}

func setVisibility(w http.ResponseWriter, r *http.Request) {
	var params SocialButtonVisibilityParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processSocialButtonPropUpdate(w, r, params)
}

func setPermissions(w http.ResponseWriter, r *http.Request) {
	var params SocialButtonPermissionParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processSocialButtonPropUpdate(w, r, params)
}

func setHelpPopupMsg(w http.ResponseWriter, r *http.Request) {
	var params HelpPopupMsgParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processSocialButtonPropUpdate(w, r, params)
}
