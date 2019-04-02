// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package attachment

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/resultra/resultra/server/common/databaseWrapper"
	"github.com/resultra/resultra/server/generic/api"
)

func init() {
	attachRouter := mux.NewRouter()

	attachRouter.HandleFunc("/api/tableView/attachment/new", newAttachment)

	attachRouter.HandleFunc("/api/tableView/attachment/get", getAttachmentAPI)

	attachRouter.HandleFunc("/api/tableView/attachment/setLabelFormat", setLabelFormat)
	attachRouter.HandleFunc("/api/tableView/attachment/setPermissions", setPermissions)
	attachRouter.HandleFunc("/api/tableView/attachment/setValidation", setValidation)
	attachRouter.HandleFunc("/api/tableView/attachment/setHelpPopupMsg", setHelpPopupMsg)

	attachRouter.HandleFunc("/api/tableView/attachment/validateInput", validateInputAPI)

	http.Handle("/api/tableView/attachment/", attachRouter)
}

func newAttachment(w http.ResponseWriter, r *http.Request) {

	attachmentParams := NewAttachmentParams{}
	if err := api.DecodeJSONRequest(r, &attachmentParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if attachmentRef, err := saveNewAttachment(trackerDBHandle, attachmentParams); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *attachmentRef)
	}

}

type GetAttachmentParams struct {
	ParentTableID string `json:"parentTableID"`
	AttachmentID  string `json:"attachmentID"`
}

func getAttachmentAPI(w http.ResponseWriter, r *http.Request) {

	var params GetAttachmentParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	numberInput, err := getAttachment(trackerDBHandle, params.ParentTableID, params.AttachmentID)
	if err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	api.WriteJSONResponse(w, *numberInput)
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

func processAttachmentPropUpdate(w http.ResponseWriter, r *http.Request, propUpdater AttachmentPropUpdater) {

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if attachmentRef, err := updateAttachmentProps(trackerDBHandle, propUpdater); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, attachmentRef)
	}
}

func setLabelFormat(w http.ResponseWriter, r *http.Request) {
	var resizeParams AttachmentLabelFormatParams
	if err := api.DecodeJSONRequest(r, &resizeParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processAttachmentPropUpdate(w, r, resizeParams)
}

func setPermissions(w http.ResponseWriter, r *http.Request) {
	var resizeParams AttachmentPermissionParams
	if err := api.DecodeJSONRequest(r, &resizeParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processAttachmentPropUpdate(w, r, resizeParams)
}

func setValidation(w http.ResponseWriter, r *http.Request) {
	var resizeParams AttachmentValidationParams
	if err := api.DecodeJSONRequest(r, &resizeParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processAttachmentPropUpdate(w, r, resizeParams)
}

func setHelpPopupMsg(w http.ResponseWriter, r *http.Request) {
	var params HelpPopupMsgParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processAttachmentPropUpdate(w, r, params)
}
