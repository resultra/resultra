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
