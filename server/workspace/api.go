// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package workspace

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/resultra/resultra/server/common/databaseWrapper"
	"github.com/resultra/resultra/server/generic/api"
	"net/http"
)

type DummyStructForInclude struct {
	Val int64
}

func init() {
	workspaceRouter := mux.NewRouter()

	workspaceRouter.HandleFunc("/api/workspace/setName", setNameAPI)
	workspaceRouter.HandleFunc("/api/workspace/setAllowUserRegistration", setAllowUserRegistrationAPI)
	workspaceRouter.HandleFunc("/api/workspace/getInfo", getInfoAPI)

	http.Handle("/api/workspace/", workspaceRouter)
}

func getInfoAPI(w http.ResponseWriter, r *http.Request) {
	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	workspaceInfo, err := GetWorkspaceInfo(trackerDBHandle)
	if err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	api.WriteJSONResponse(w, workspaceInfo)

}

type SetNameParams struct {
	NewName string `json:"newName"`
}

func setNameAPI(w http.ResponseWriter, r *http.Request) {

	var params SetNameParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if err := setWorkspaceName(trackerDBHandle, params.NewName); err != nil {
		api.WriteJSONResponse(w, fmt.Sprintf("%v", err))
		return
	}

	response := true
	api.WriteJSONResponse(w, response)

}

func setAllowUserRegistrationAPI(w http.ResponseWriter, r *http.Request) {

	var params AllowRegistrationParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if updatedWorkspaceInfo, err := updateWorkspaceProps(r, params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *updatedWorkspaceInfo)
	}

}
