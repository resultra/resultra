package databaseController

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/common/databaseWrapper"
	"resultra/datasheet/server/generic/api"
	"resultra/datasheet/server/trackerDatabase"
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

	if templateList, err := getCurrentUserTemplateTrackers(trackerDBHandle, r); err != nil {
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

	if dbList, err := getCurrentUserTrackingDatabases(trackerDBHandle, r); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, dbList)
	}

}
