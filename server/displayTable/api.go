package displayTable

import (
	"fmt"
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

	tableRouter.HandleFunc("/api/tableView/new", newTableAPI)
	tableRouter.HandleFunc("/api/tableView/setName", setTableName)

	tableRouter.HandleFunc("/api/tableView/list", listTableAPI)
	tableRouter.HandleFunc("/api/tableView/get", getTableAPI)

	tableRouter.HandleFunc("/api/tableView/validateTableName", validateTableNameAPI)
	tableRouter.HandleFunc("/api/tableView/validateNewTableName", validateNewTableNameAPI)

	http.Handle("/api/tableView/", tableRouter)
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

type ListTableParams struct {
	ParentDatabaseID string `json:"parentDatabaseID"`
}

func listTableAPI(w http.ResponseWriter, r *http.Request) {

	var params ListTableParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if verifyErr := userRole.VerifyCurrUserIsDatabaseAdmin(
		r, params.ParentDatabaseID); verifyErr != nil {
		api.WriteErrorResponse(w, verifyErr)
		return
	}

	if tableRefs, err := getAllTables(params.ParentDatabaseID); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, tableRefs)
	}

}

type GetTableParams struct {
	TableID string `json:"tableID"`
}

func getTableAPI(w http.ResponseWriter, r *http.Request) {

	var params GetTableParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	tableRef, err := GetTable(params.TableID)
	if err != nil {
		api.WriteErrorResponse(w, err)
	}

	if verifyErr := userRole.VerifyCurrUserIsDatabaseAdmin(
		r, tableRef.ParentDatabaseID); verifyErr != nil {
		api.WriteErrorResponse(w, verifyErr)
		return
	}

	api.WriteJSONResponse(w, tableRef)

}

func validateTableNameAPI(w http.ResponseWriter, r *http.Request) {

	tableName := r.FormValue("tableName")
	formID := r.FormValue("formID")

	if err := validateTableName(formID, tableName); err != nil {
		api.WriteJSONResponse(w, fmt.Sprintf("%v", err))
		return
	}

	response := true
	api.WriteJSONResponse(w, response)

}

func validateNewTableNameAPI(w http.ResponseWriter, r *http.Request) {

	tableName := r.FormValue("tableName")
	databaseID := r.FormValue("databaseID")

	if err := validateNewTableName(databaseID, tableName); err != nil {
		api.WriteJSONResponse(w, fmt.Sprintf("%v", err))
		return
	}

	response := true
	api.WriteJSONResponse(w, response)

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
