package table

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/generic/api"
	"resultra/datasheet/server/userRole"
)

func init() {
	tableRouter := mux.NewRouter()

	tableRouter.HandleFunc("/api/table/new", newTable)
	tableRouter.HandleFunc("/api/table/getList", getTableListAPI)

	http.Handle("/api/table/", tableRouter)
}

func newTable(w http.ResponseWriter, r *http.Request) {

	var tableParams NewTableParams
	if err := api.DecodeJSONRequest(r, &tableParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if verifyUserErr := userRole.VerifyCurrUserIsDatabaseAdmin(r, tableParams.DatabaseID); verifyUserErr != nil {
		api.WriteErrorResponse(w, verifyUserErr)
		return
	}

	if newTable, err := saveNewTable(tableParams); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *newTable)
	}

}

func getTableListAPI(w http.ResponseWriter, r *http.Request) {

	var tableParams GetTableListParams
	if err := api.DecodeJSONRequest(r, &tableParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if tableRefs, err := getTableList(tableParams); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, tableRefs)
	}

}
