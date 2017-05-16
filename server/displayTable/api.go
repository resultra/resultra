package displayTable

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/generic/api"
	"resultra/datasheet/server/userRole"
)

type DummyStructForInclude struct {
	Val int64
}

func init() {

	tableRouter := mux.NewRouter()

	tableRouter.HandleFunc("/api/displayTable/new", newTableAPI)
	tableRouter.HandleFunc("/api/frm/setName", setTableName)

	http.Handle("/api/displayTable/", tableRouter)
}

func newTableAPI(w http.ResponseWriter, r *http.Request) {

	var params NewTableParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if verifyErr := userRole.VerifyCurrUserIsDatabaseAdmin(
		r, params.ParentDatabaseID); verifyErr != nil {
		api.WriteErrorResponse(w, verifyErr)
		return
	}

	if tableRef, err := newTable(params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *tableRef)
	}

}

func processTablePropUpdate(w http.ResponseWriter, r *http.Request, propUpdater TablePropUpdater) {
	if updatedTable, err := updateTableProps(propUpdater); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, updatedTable)
	}
}

func setTableName(w http.ResponseWriter, r *http.Request) {
	var params SetTableNameParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processTablePropUpdate(w, r, params)
}
