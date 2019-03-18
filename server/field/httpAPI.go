// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package field

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"resultra/tracker/server/common/databaseWrapper"
	"resultra/tracker/server/generic/api"
	"resultra/tracker/server/userRole"
)

func init() {

	fieldRouter := mux.NewRouter()

	fieldRouter.HandleFunc("/api/field/new", newField)

	fieldRouter.HandleFunc("/api/field/getListByType", getFieldsByType)
	fieldRouter.HandleFunc("/api/field/getSortedListByType", getSortedFieldsByTypeAPI)
	fieldRouter.HandleFunc("/api/field/getAllSortedFields", getAllSortedFieldsAPI)

	fieldRouter.HandleFunc("/api/field/get", getField)

	fieldRouter.HandleFunc("/api/field/validateExistingFieldName", validateExistingFieldNameAPI)
	fieldRouter.HandleFunc("/api/field/validateNewFieldName", validateNewFieldNameAPI)
	fieldRouter.HandleFunc("/api/field/setName", setNameAPI)

	fieldRouter.HandleFunc("/api/field/validateExistingFieldRefName", validateExistingFieldRefNameAPI)
	fieldRouter.HandleFunc("/api/field/validateNewFieldRefName", validateNewFieldRefNameAPI)
	fieldRouter.HandleFunc("/api/field/setRefName", setRefNameAPI)

	http.Handle("/api/field/", fieldRouter)
}

func newField(w http.ResponseWriter, r *http.Request) {

	var newFieldParams NewNonCalcFieldParams
	if err := api.DecodeJSONRequest(r, &newFieldParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if verifyErr := userRole.VerifyCurrUserIsDatabaseAdmin(
		trackerDBHandle, r, newFieldParams.ParentDatabaseID); verifyErr != nil {
		api.WriteErrorResponse(w, verifyErr)
		return
	}

	if newField, err := NewNonCalcField(trackerDBHandle, newFieldParams); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *newField)
	}

}

func getFieldsByType(w http.ResponseWriter, r *http.Request) {

	var fieldListParams GetFieldListParams
	if err := api.DecodeJSONRequest(r, &fieldListParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if fieldsByType, err := GetFieldsByType(trackerDBHandle, fieldListParams); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, fieldsByType)
	}

}

func getSortedFieldsByTypeAPI(w http.ResponseWriter, r *http.Request) {

	var fieldListParams GetSortedFieldListParams
	if err := api.DecodeJSONRequest(r, &fieldListParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if fieldsByType, err := getSortedFieldsByType(trackerDBHandle, fieldListParams); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, fieldsByType)
	}

}

func getAllSortedFieldsAPI(w http.ResponseWriter, r *http.Request) {

	var fieldListParams GetFieldListParams
	if err := api.DecodeJSONRequest(r, &fieldListParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if fieldsByType, err := getAllSortedFields(trackerDBHandle, fieldListParams); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, fieldsByType)
	}

}

type GetFieldParams struct {
	FieldID string `json:"fieldID"`
}

func getField(w http.ResponseWriter, r *http.Request) {

	var params GetFieldParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if fieldInfo, err := GetField(trackerDBHandle, params.FieldID); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, fieldInfo)
	}

}

func validateExistingFieldNameAPI(w http.ResponseWriter, r *http.Request) {

	fieldName := r.FormValue("fieldName")
	fieldID := r.FormValue("fieldID")

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if err := validateExistingFieldName(trackerDBHandle, fieldID, fieldName); err != nil {
		api.WriteJSONResponse(w, fmt.Sprintf("%v", err))
		return
	}

	response := true
	api.WriteJSONResponse(w, response)

}

func validateNewFieldNameAPI(w http.ResponseWriter, r *http.Request) {

	fieldName := r.FormValue("fieldName")
	databaseID := r.FormValue("databaseID")

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if err := validateNewFieldName(trackerDBHandle, databaseID, fieldName); err != nil {
		api.WriteJSONResponse(w, fmt.Sprintf("%v", err))
		return
	}

	response := true
	api.WriteJSONResponse(w, response)

}

func validateExistingFieldRefNameAPI(w http.ResponseWriter, r *http.Request) {

	fieldRefName := r.FormValue("fieldRefName")
	fieldID := r.FormValue("fieldID")

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if err := validateExistingFieldRefName(trackerDBHandle, fieldID, fieldRefName); err != nil {
		api.WriteJSONResponse(w, fmt.Sprintf("%v", err))
		return
	}

	response := true
	api.WriteJSONResponse(w, response)

}

func validateNewFieldRefNameAPI(w http.ResponseWriter, r *http.Request) {

	fieldRefName := r.FormValue("fieldRefName")
	databaseID := r.FormValue("databaseID")

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if err := validateNewFieldRefName(trackerDBHandle, databaseID, fieldRefName); err != nil {
		api.WriteJSONResponse(w, fmt.Sprintf("%v", err))
		return
	}

	response := true
	api.WriteJSONResponse(w, response)

}

func setNameAPI(w http.ResponseWriter, r *http.Request) {
	var params SetFieldNameParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processFieldPropUpdate(w, r, params)
}

func setRefNameAPI(w http.ResponseWriter, r *http.Request) {
	var params SetFieldRefNameParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processFieldPropUpdate(w, r, params)
}
