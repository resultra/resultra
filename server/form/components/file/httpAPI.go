// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package file

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/resultra/resultra/server/common/databaseWrapper"
	"github.com/resultra/resultra/server/generic/api"
)

func init() {
	fileRouter := mux.NewRouter()

	fileRouter.HandleFunc("/api/frm/file/new", newFile)
	fileRouter.HandleFunc("/api/frm/file/resize", resizeFile)
	fileRouter.HandleFunc("/api/frm/file/setLabelFormat", setLabelFormat)
	fileRouter.HandleFunc("/api/frm/file/setVisibility", setVisibility)
	fileRouter.HandleFunc("/api/frm/file/setPermissions", setPermissions)
	fileRouter.HandleFunc("/api/frm/file/setValidation", setValidation)
	fileRouter.HandleFunc("/api/frm/file/setClearValueSupported", setClearValueSupported)
	fileRouter.HandleFunc("/api/frm/file/setHelpPopupMsg", setHelpPopupMsg)

	fileRouter.HandleFunc("/api/frm/file/validateInput", validateInputAPI)

	http.Handle("/api/frm/file/", fileRouter)
}

func newFile(w http.ResponseWriter, r *http.Request) {

	fileParams := NewFileParams{}
	if err := api.DecodeJSONRequest(r, &fileParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if fileRef, err := saveNewFile(trackerDBHandle, fileParams); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *fileRef)
	}

}

func validateInputAPI(w http.ResponseWriter, r *http.Request) {

	var params FileValidateInputParams
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

func processFilePropUpdate(w http.ResponseWriter, r *http.Request, propUpdater FilePropUpdater) {
	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}
	if fileRef, err := updateFileProps(trackerDBHandle, propUpdater); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, fileRef)
	}
}

func resizeFile(w http.ResponseWriter, r *http.Request) {
	var resizeParams FileResizeParams
	if err := api.DecodeJSONRequest(r, &resizeParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processFilePropUpdate(w, r, resizeParams)
}

func setLabelFormat(w http.ResponseWriter, r *http.Request) {
	var params FileLabelFormatParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processFilePropUpdate(w, r, params)
}

func setVisibility(w http.ResponseWriter, r *http.Request) {
	var params FileVisibilityParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processFilePropUpdate(w, r, params)
}

func setPermissions(w http.ResponseWriter, r *http.Request) {
	var params FilePermissionParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processFilePropUpdate(w, r, params)
}

func setValidation(w http.ResponseWriter, r *http.Request) {
	var params FileValidationParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processFilePropUpdate(w, r, params)
}

func setClearValueSupported(w http.ResponseWriter, r *http.Request) {
	var params FileClearValueSupportedParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processFilePropUpdate(w, r, params)
}

func setHelpPopupMsg(w http.ResponseWriter, r *http.Request) {
	var params HelpPopupMsgParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processFilePropUpdate(w, r, params)
}
