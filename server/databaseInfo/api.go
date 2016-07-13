package databaseInfo

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/generic/api"
)

type DummyStructForInclude struct {
	Val int64
}

func init() {
	databaseInfoRouter := mux.NewRouter()

	databaseInfoRouter.HandleFunc("/api/databaseInfo/getInfo", getDatabaseInfoAPI)

	http.Handle("/api/databaseInfo/", databaseInfoRouter)
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
