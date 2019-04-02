// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package databaseController

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/resultra/resultra/server/common/databaseWrapper"
	"github.com/resultra/resultra/server/generic/api"
	"github.com/resultra/resultra/server/trackerDatabase"
)

type DummyStructForInclude struct {
	Val int64
}

func init() {
	databaseRouter := mux.NewRouter()

	databaseRouter.HandleFunc("/api/database/getInfo", getDatabaseInfoAPI)
	databaseRouter.HandleFunc("/api/database/new", newDatabase)

	databaseRouter.HandleFunc("/api/database/getList", getDatabaseListAPI)
	databaseRouter.HandleFunc("/api/database/getTemplateList", getTemplateListAPI)
	databaseRouter.HandleFunc("/api/database/getFactoryTemplateList", getFactoryTemplateListAPI)
	databaseRouter.HandleFunc("/api/database/getUserTemplateList", getUserTemplateListAPI)

	databaseRouter.HandleFunc("/api/database/setName", trackerDatabase.SetNameAPI)
	databaseRouter.HandleFunc("/api/database/setActive", trackerDatabase.SetActiveAPI)
	databaseRouter.HandleFunc("/api/database/setDescription", trackerDatabase.SetDescriptionAPI)
	databaseRouter.HandleFunc("/api/database/setListOrder", trackerDatabase.SetListOrderAPI)
	databaseRouter.HandleFunc("/api/database/setDashboardOrder", trackerDatabase.SetDashboardOrderAPI)
	databaseRouter.HandleFunc("/api/database/setFormLinkOrder", trackerDatabase.SetFormLinkOrderAPI)

	databaseRouter.HandleFunc("/api/database/validateDatabaseName", trackerDatabase.ValidateDatabaseNameAPI)
	databaseRouter.HandleFunc("/api/database/validateNewTrackerName", trackerDatabase.ValidateNewTrackerNameAPI)

	databaseRouter.HandleFunc("/api/database/saveAsTemplate", saveAsTemplate)

	http.Handle("/api/database/", databaseRouter)
}

func getDatabaseInfoAPI(w http.ResponseWriter, r *http.Request) {

	params := DatabaseInfoParams{}
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	dbInfo, err := getDatabaseInfo(trackerDBHandle, params)
	if err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *dbInfo)
	}

}

func newDatabase(w http.ResponseWriter, r *http.Request) {

	var dbParams trackerDatabase.NewDatabaseParams
	if err := api.DecodeJSONRequest(r, &dbParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if newDB, err := createNewDatabase(trackerDBHandle, r, dbParams); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *newDB)
	}

}

func saveAsTemplate(w http.ResponseWriter, r *http.Request) {

	var params SaveAsTemplateParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if templateDB, err := saveExistingDatabaseAsTemplate(r, params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *templateDB)
	}

}

func getTemplateListAPI(w http.ResponseWriter, r *http.Request) {

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	getTemplParams := GetTemplateListParams{
		IncludeInactive: false,
		CurrUserOnly:    false}

	if templateList, err := getCurrentUserTemplateTrackers(getTemplParams, trackerDBHandle, r); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, templateList)
	}

}

func getFactoryTemplateListAPI(w http.ResponseWriter, r *http.Request) {

	if databaseWrapper.FactoryTemplateDatabaseIsConfigured() {
		trackerDBHandle, dbErr := databaseWrapper.GetFactoryTemplateTrackerDatabaseHandle(r)
		if dbErr != nil {
			api.WriteErrorResponse(w, dbErr)
			return
		}
		getTemplParams := GetTemplateListParams{
			IncludeInactive: false,
			CurrUserOnly:    false}

		if templateList, err := getCurrentUserTemplateTrackers(getTemplParams, trackerDBHandle, r); err != nil {
			api.WriteErrorResponse(w, err)
			return
		} else {
			api.WriteJSONResponse(w, templateList)
			return
		}

	} else {
		emptyTemplateList := []UserTemplateTrackerDatabaseInfo{}
		api.WriteJSONResponse(w, emptyTemplateList)
	}

}

type GetUserTemplateParams struct {
	IncludeInactive bool `json:"includeInactive"`
}

func getUserTemplateListAPI(w http.ResponseWriter, r *http.Request) {

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	var params GetUserTemplateParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	getTemplParams := GetTemplateListParams{
		IncludeInactive: params.IncludeInactive,
		CurrUserOnly:    true}

	if templateList, err := getCurrentUserTemplateTrackers(getTemplParams, trackerDBHandle, r); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, templateList)
	}

}

func getDatabaseListAPI(w http.ResponseWriter, r *http.Request) {

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	var params GetTrackerListParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if dbList, err := getCurrentUserTrackingDatabases(params, trackerDBHandle, r); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, dbList)
	}

}
