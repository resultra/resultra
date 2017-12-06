package workspace

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/common/databaseWrapper"
	"resultra/datasheet/server/generic/api"
)

type DummyStructForInclude struct {
	Val int64
}

func init() {
	workspaceRouter := mux.NewRouter()

	workspaceRouter.HandleFunc("/api/workspace/setName", setNameAPI)

	http.Handle("/api/workspace/", workspaceRouter)
}

type SetNameParams struct {
	NewName string `json:"newName"`
}

func setNameAPI(w http.ResponseWriter, r *http.Request) {

	var params SetNameParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if err := setWorkspaceName(trackerDBHandle, params.NewName); err != nil {
		api.WriteJSONResponse(w, fmt.Sprintf("%v", err))
		return
	}

	response := true
	api.WriteJSONResponse(w, response)

}
