// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package file

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/tracker/server/common/databaseWrapper"
	"resultra/tracker/server/generic/api"
)

func init() {
	fileRouter := mux.NewRouter()

	fileRouter.HandleFunc("/api/tableView/file/new", newFile)

	fileRouter.HandleFunc("/api/tableView/file/get", getFileAPI)

	fileRouter.HandleFunc("/api/tableView/file/setLabelFormat", setLabelFormat)
	fileRouter.HandleFunc("/api/tableView/file/setPermissions", setPermissions)
	fileRouter.HandleFunc("/api/tableView/file/setValidation", setValidation)
	fileRouter.HandleFunc("/api/tableView/file/setClearValueSupported", setClearValueSupported)
	fileRouter.HandleFunc("/api/tableView/file/setHelpPopupMsg", setHelpPopupMsg)

	fileRouter.HandleFunc("/api/tableView/file/validateInput", validateInputAPI)

	http.Handle("/api/tableView/file/", fileRouter)
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

type GetFileParams struct {
	ParentTableID string `json:"parentTableID"`
	FileID        string `json:"fileID"`
}

func getFileAPI(w http.ResponseWriter, r *http.Request) {

	var params GetFileParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	file, err := getFile(trackerDBHandle, params.ParentTableID, params.FileID)
	if err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	api.WriteJSONResponse(w, *file)
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

func setLabelFormat(w http.ResponseWriter, r *http.Request) {
	var params FileLabelFormatParams
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
