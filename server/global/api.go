// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package global

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"resultra/tracker/server/common/databaseWrapper"
	"resultra/tracker/server/generic/api"
	"resultra/tracker/server/userRole"
)

type DummyStructForInclude struct{ Val int64 }

func init() {

	globalRouter := mux.NewRouter()

	globalRouter.HandleFunc("/api/global/new", newGlobalAPI)
	globalRouter.HandleFunc("/api/global/getList", getListAPI)

	globalRouter.HandleFunc("/api/global/validateName", validateNameAPI)
	globalRouter.HandleFunc("/api/global/validateNewName", validateNewNameAPI)

	globalRouter.HandleFunc("/api/global/validateNewReferenceName", validateNewReferenceNameAPI)

	globalRouter.HandleFunc("/api/global/setTextValue", setTextValue)
	globalRouter.HandleFunc("/api/global/setBoolValue", setBoolValue)
	globalRouter.HandleFunc("/api/global/setTimeValue", setTimeValue)
	globalRouter.HandleFunc("/api/global/setNumberValue", setNumberValue)
	globalRouter.HandleFunc("/api/global/getValues", getValues)

	globalRouter.HandleFunc("/api/global/uploadFileToGlobalValue", uploadFileAPI)
	globalRouter.HandleFunc("/api/global/getGlobalValUrl", getGlobalValUrlAPI)

	http.Handle("/api/global/", globalRouter)
}

func newGlobalAPI(w http.ResponseWriter, r *http.Request) {

	var params NewGlobalParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if verifyErr := userRole.VerifyCurrUserIsDatabaseAdmin(
		trackerDBHandle, r, params.ParentDatabaseID); verifyErr != nil {
		api.WriteErrorResponse(w, verifyErr)
		return
	}

	if globalRef, err := newGlobal(trackerDBHandle, params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *globalRef)
	}

}

func getListAPI(w http.ResponseWriter, r *http.Request) {

	var params GetGlobalsParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if globals, err := GetGlobals(trackerDBHandle, params.ParentDatabaseID); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, globals)
	}

}

func validateNameAPI(w http.ResponseWriter, r *http.Request) {

	globalName := r.FormValue("globalName")
	globalID := r.FormValue("globalID")

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}
	if err := validateGlobalName(trackerDBHandle, globalID, globalName); err != nil {
		api.WriteJSONResponse(w, fmt.Sprintf("%v", err))
		return
	}

	response := true
	api.WriteJSONResponse(w, response)

}

func validateNewNameAPI(w http.ResponseWriter, r *http.Request) {

	globalName := r.FormValue("globalName")
	databaseID := r.FormValue("databaseID")

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}
	if err := validateNewGlobalName(trackerDBHandle, databaseID, globalName); err != nil {
		api.WriteJSONResponse(w, fmt.Sprintf("%v", err))
		return
	}

	response := true
	api.WriteJSONResponse(w, response)

}

func validateNewReferenceNameAPI(w http.ResponseWriter, r *http.Request) {

	refName := r.FormValue("refName")
	databaseID := r.FormValue("databaseID")

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}
	if err := validateNewReferenceName(trackerDBHandle, databaseID, refName); err != nil {
		api.WriteJSONResponse(w, fmt.Sprintf("%v", err))
		return
	}

	response := true
	api.WriteJSONResponse(w, response)

}

func setTextValue(w http.ResponseWriter, r *http.Request) {
	var params SetTextGlobalValueParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}
	globalValUpdate, setErr := updateGlobalValue(trackerDBHandle, params)
	if setErr != nil {
		api.WriteErrorResponse(w, setErr)
		return
	} else {
		api.WriteJSONResponse(w, globalValUpdate)
	}

}

func setTimeValue(w http.ResponseWriter, r *http.Request) {
	var params SetTimeGlobalValueParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}
	globalValUpdate, setErr := updateGlobalValue(trackerDBHandle, params)
	if setErr != nil {
		api.WriteErrorResponse(w, setErr)
		return
	} else {
		api.WriteJSONResponse(w, globalValUpdate)
	}

}

func setBoolValue(w http.ResponseWriter, r *http.Request) {
	var params SetBoolGlobalValueParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}
	globalValUpdate, setErr := updateGlobalValue(trackerDBHandle, params)
	if setErr != nil {
		api.WriteErrorResponse(w, setErr)
		return
	} else {
		api.WriteJSONResponse(w, globalValUpdate)
	}

}

func setNumberValue(w http.ResponseWriter, r *http.Request) {
	var params SetNumberGlobalValueParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}
	globalValUpdate, setErr := updateGlobalValue(trackerDBHandle, params)
	if setErr != nil {
		api.WriteErrorResponse(w, setErr)
		return
	} else {
		api.WriteJSONResponse(w, globalValUpdate)
	}

}

func getValues(w http.ResponseWriter, r *http.Request) {
	var params GetGlobalValuesParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}
	globalVals, getErr := GetGlobalValues(trackerDBHandle, params)
	if getErr != nil {
		api.WriteErrorResponse(w, getErr)
		return
	} else {
		api.WriteJSONResponse(w, globalVals)
	}

}

func uploadFileAPI(w http.ResponseWriter, req *http.Request) {

	if uploadResponse, uploadErr := uploadFile(w, req); uploadErr != nil {
		api.WriteErrorResponse(w, uploadErr)
	} else {
		api.WriteJSONResponse(w, *uploadResponse)
	}

}

func getGlobalValUrlAPI(w http.ResponseWriter, r *http.Request) {

	var params GetGlobalValUrlParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}
	if urlResponse, err := getGlobalValUrl(trackerDBHandle, params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, urlResponse)
	}
}
