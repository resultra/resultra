package global

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/generic/api"
	"resultra/datasheet/server/userRole"
)

type DummyStructForInclude struct{ Val int64 }

func init() {

	globalRouter := mux.NewRouter()

	globalRouter.HandleFunc("/api/global/new", newGlobalAPI)
	globalRouter.HandleFunc("/api/global/getList", getListAPI)

	globalRouter.HandleFunc("/api/global/validateName", validateNameAPI)
	globalRouter.HandleFunc("/api/global/validateNewName", validateNewNameAPI)

	http.Handle("/api/global/", globalRouter)
}

func newGlobalAPI(w http.ResponseWriter, r *http.Request) {

	var params NewGlobalParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if verifyErr := userRole.VerifyCurrUserIsDatabaseAdmin(
		r, params.ParentDatabaseID); verifyErr != nil {
		api.WriteErrorResponse(w, verifyErr)
		return
	}

	if globalRef, err := newGlobal(params); err != nil {
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

	if verifyErr := userRole.VerifyCurrUserIsDatabaseAdmin(
		r, params.ParentDatabaseID); verifyErr != nil {
		api.WriteErrorResponse(w, verifyErr)
		return
	}

	if globals, err := getGlobals(params.ParentDatabaseID); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, globals)
	}

}

func validateNameAPI(w http.ResponseWriter, r *http.Request) {

	globalName := r.FormValue("globalName")
	globalID := r.FormValue("globalID")

	if err := validateGlobalName(globalID, globalName); err != nil {
		api.WriteJSONResponse(w, fmt.Sprintf("%v", err))
		return
	}

	response := true
	api.WriteJSONResponse(w, response)

}

func validateNewNameAPI(w http.ResponseWriter, r *http.Request) {

	globalName := r.FormValue("globalName")
	databaseID := r.FormValue("databaseID")

	if err := validateNewGlobalName(databaseID, globalName); err != nil {
		api.WriteJSONResponse(w, fmt.Sprintf("%v", err))
		return
	}

	response := true
	api.WriteJSONResponse(w, response)

}
