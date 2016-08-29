package global

import (
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
