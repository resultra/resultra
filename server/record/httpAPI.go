// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package record

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/resultra/resultra/server/common/databaseWrapper"
	"github.com/resultra/resultra/server/generic/api"
	"github.com/resultra/resultra/server/generic/uniqueID"
)

func init() {
	recordRouter := mux.NewRouter()

	recordRouter.HandleFunc("/api/record/setDraftStatus", setDraftStatusAPI)

	recordRouter.HandleFunc("/api/record/allocateChangeSetID", allocateChangeSetIDAPI)

	recordRouter.HandleFunc("/api/record/getFieldValChangeInfo", getFieldValChangeInfoAPI)

	http.Handle("/api/record/", recordRouter)
}

func getFieldValChangeInfoAPI(w http.ResponseWriter, r *http.Request) {

	var params GetFieldValChangeInfoParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if valChangeInfo, err := getFieldValChangeInfo(r, params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, valChangeInfo)
	}

}

func setDraftStatusAPI(w http.ResponseWriter, r *http.Request) {

	var params SetDraftStatusParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if err := setDraftStatus(trackerDBHandle, params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, true)
	}
}

type changeSetIDAllocationResponse struct {
	ChangeSetID string `json:"changeSetID"`
}

func allocateChangeSetIDAPI(w http.ResponseWriter, r *http.Request) {

	changeSetID := uniqueID.GenerateUniqueID()

	response := changeSetIDAllocationResponse{changeSetID}

	api.WriteJSONResponse(w, response)
}
