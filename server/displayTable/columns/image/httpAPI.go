// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package image

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/tracker/server/common/databaseWrapper"
	"resultra/tracker/server/generic/api"
)

func init() {
	imageRouter := mux.NewRouter()

	imageRouter.HandleFunc("/api/tableView/image/new", newImage)

	imageRouter.HandleFunc("/api/tableView/image/get", getImageAPI)

	imageRouter.HandleFunc("/api/tableView/image/setLabelFormat", setLabelFormat)
	imageRouter.HandleFunc("/api/tableView/image/setPermissions", setPermissions)
	imageRouter.HandleFunc("/api/tableView/image/setValidation", setValidation)
	imageRouter.HandleFunc("/api/tableView/image/setClearValueSupported", setClearValueSupported)
	imageRouter.HandleFunc("/api/tableView/image/setHelpPopupMsg", setHelpPopupMsg)

	imageRouter.HandleFunc("/api/tableView/image/validateInput", validateInputAPI)

	http.Handle("/api/tableView/image/", imageRouter)
}

func newImage(w http.ResponseWriter, r *http.Request) {

	imageParams := NewImageParams{}
	if err := api.DecodeJSONRequest(r, &imageParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if imageRef, err := saveNewImage(trackerDBHandle, imageParams); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *imageRef)
	}

}

type GetImageParams struct {
	ParentTableID string `json:"parentTableID"`
	ImageID       string `json:"imageID"`
}

func getImageAPI(w http.ResponseWriter, r *http.Request) {

	var params GetImageParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	image, err := getImage(trackerDBHandle, params.ParentTableID, params.ImageID)
	if err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	api.WriteJSONResponse(w, *image)
}

func validateInputAPI(w http.ResponseWriter, r *http.Request) {

	var params ImageValidateInputParams
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

func processImagePropUpdate(w http.ResponseWriter, r *http.Request, propUpdater ImagePropUpdater) {

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if imageRef, err := updateImageProps(trackerDBHandle, propUpdater); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, imageRef)
	}
}

func setLabelFormat(w http.ResponseWriter, r *http.Request) {
	var params ImageLabelFormatParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processImagePropUpdate(w, r, params)
}

func setPermissions(w http.ResponseWriter, r *http.Request) {
	var params ImagePermissionParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processImagePropUpdate(w, r, params)
}

func setValidation(w http.ResponseWriter, r *http.Request) {
	var params ImageValidationParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processImagePropUpdate(w, r, params)
}

func setClearValueSupported(w http.ResponseWriter, r *http.Request) {
	var params ImageClearValueSupportedParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processImagePropUpdate(w, r, params)
}

func setHelpPopupMsg(w http.ResponseWriter, r *http.Request) {
	var params HelpPopupMsgParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processImagePropUpdate(w, r, params)
}
