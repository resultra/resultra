// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package tag

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/resultra/resultra/server/common/databaseWrapper"
	"github.com/resultra/resultra/server/generic/api"
)

func init() {
	tagRouter := mux.NewRouter()

	tagRouter.HandleFunc("/api/tableView/tag/new", newTag)

	tagRouter.HandleFunc("/api/tableView/tag/get", getTagAPI)

	tagRouter.HandleFunc("/api/tableView/tag/setLabelFormat", setLabelFormat)
	tagRouter.HandleFunc("/api/tableView/tag/setHelpPopupMsg", setHelpPopupMsg)

	tagRouter.HandleFunc("/api/tableView/tag/setPermissions", setPermissions)
	tagRouter.HandleFunc("/api/tableView/tag/setValidation", setValidation)
	tagRouter.HandleFunc("/api/tableView/tag/validateInput", validateInputAPI)

	http.Handle("/api/tableView/tag/", tagRouter)
}

func newTag(w http.ResponseWriter, r *http.Request) {

	tagParams := NewTagParams{}
	if err := api.DecodeJSONRequest(r, &tagParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if tagRef, err := saveNewTag(trackerDBHandle, tagParams); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *tagRef)
	}

}

type GetTagParams struct {
	ParentTableID string `json:"parentTableID"`
	TagID         string `json:"tagID"`
}

func getTagAPI(w http.ResponseWriter, r *http.Request) {

	var params GetTagParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	textInput, err := getTag(trackerDBHandle, params.ParentTableID, params.TagID)
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

func processTagPropUpdate(w http.ResponseWriter, r *http.Request, propUpdater TagPropUpdater) {

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if tagRef, err := updateTagProps(trackerDBHandle, propUpdater); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, tagRef)
	}
}

func setLabelFormat(w http.ResponseWriter, r *http.Request) {
	var params TagLabelFormatParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processTagPropUpdate(w, r, params)
}

func setPermissions(w http.ResponseWriter, r *http.Request) {
	var params TagPermissionParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processTagPropUpdate(w, r, params)
}

func setValidation(w http.ResponseWriter, r *http.Request) {
	var params TagValidationParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processTagPropUpdate(w, r, params)
}

func setHelpPopupMsg(w http.ResponseWriter, r *http.Request) {
	var params HelpPopupMsgParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processTagPropUpdate(w, r, params)
}
