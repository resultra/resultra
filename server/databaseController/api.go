package databaseController

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/database"
	"resultra/datasheet/server/generic/api"
)

type DummyStructForInclude struct {
	Val int64
}

func init() {
	databaseRouter := mux.NewRouter()

	databaseRouter.HandleFunc("/api/database/getInfo", getDatabaseInfoAPI)
	databaseRouter.HandleFunc("/api/database/new", newDatabase)

	databaseRouter.HandleFunc("/api/database/getList", getDatabaseListAPI)

	databaseRouter.HandleFunc("/api/database/setName", database.SetNameAPI)
	databaseRouter.HandleFunc("/api/database/validateDatabaseName", database.ValidateDatabaseNameAPI)
	databaseRouter.HandleFunc("/api/database/validateNewTrackerName", database.ValidateNewTrackerNameAPI)

	databaseRouter.HandleFunc("/api/database/saveAsTemplate", saveAsTemplate)

	http.Handle("/api/database/", databaseRouter)
}

func getDatabaseInfoAPI(w http.ResponseWriter, r *http.Request) {

	params := DatabaseInfoParams{}
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	dbInfo, err := getDatabaseInfo(params)
	if err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *dbInfo)
	}

}

func newDatabase(w http.ResponseWriter, r *http.Request) {

	var dbParams database.NewDatabaseParams
	if err := api.DecodeJSONRequest(r, &dbParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if newDB, err := createNewDatabase(r, dbParams); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *newDB)
	}

}

func saveAsTemplate(w http.ResponseWriter, r *http.Request) {

	var params SaveTemplateParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if templateDB, err := saveDatabaseToTemplate(params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *templateDB)
	}

}

func getDatabaseListAPI(w http.ResponseWriter, r *http.Request) {

	if dbList, err := getCurrentUserTrackingDatabases(r); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, dbList)
	}

}
