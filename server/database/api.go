package database

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/generic/api"
)

func init() {
	databaseRouter := mux.NewRouter()

	databaseRouter.HandleFunc("/api/database/new", newDatabase)

	http.Handle("/api/database/", databaseRouter)
}

func newDatabase(w http.ResponseWriter, r *http.Request) {

	var dbParams NewDatabaseParams
	if err := api.DecodeJSONRequest(r, &dbParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if newDB, err := saveNewDatabase(dbParams); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *newDB)
	}

}
