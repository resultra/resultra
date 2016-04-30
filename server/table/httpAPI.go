package table

import (
	"appengine"
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/generic/api"
)

func init() {
	tableRouter := mux.NewRouter()

	tableRouter.HandleFunc("/api/table/new", newTable)

	http.Handle("/api/table/", tableRouter)
}

func newTable(w http.ResponseWriter, r *http.Request) {

	var tableParams NewTableParams
	if err := api.DecodeJSONRequest(r, &tableParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	appEngCntxt := appengine.NewContext(r)
	if tableRef, err := saveNewTable(appEngCntxt, tableParams); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *tableRef)
	}

}
