// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package displayTable

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/resultra/resultra/server/common/databaseWrapper"
	"github.com/resultra/resultra/server/displayTable/columns/common"
	"github.com/resultra/resultra/server/generic/api"
	"github.com/resultra/resultra/server/userRole"
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

	tableRouter.HandleFunc("/api/tableView/setOrderedCols", setOrderedCols)

	tableRouter.HandleFunc("/api/tableView/deleteColumn", deleteColumnAPI)

	tableRouter.HandleFunc("/api/tableView/getColumns", getTableColsAPI)
	tableRouter.HandleFunc("/api/tableView/getTableDisplayInfo", getTableDisplayInfoAPI)

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

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if verifyErr := userRole.VerifyCurrUserIsDatabaseAdmin(trackerDBHandle,
		r, params.ParentDatabaseID); verifyErr != nil {
		api.WriteErrorResponse(w, verifyErr)
		return
	}

	if tableRef, err := newTable(trackerDBHandle, params); err != nil {
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

	/* TODO - Restore authentication for or refactor to get a list to only include table views shown for a given item list
	if verifyErr := userRole.VerifyCurrUserIsDatabaseAdmin(
		r, params.ParentDatabaseID); verifyErr != nil {
		api.WriteErrorResponse(w, verifyErr)
		return
	}
	*/
	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if tableRefs, err := getAllTables(trackerDBHandle, params.ParentDatabaseID); err != nil {
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

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	tableRef, err := GetTable(trackerDBHandle, params.TableID)
	if err != nil {
		api.WriteErrorResponse(w, err)
	}

	if verifyErr := userRole.VerifyCurrUserIsDatabaseAdmin(trackerDBHandle,
		r, tableRef.ParentDatabaseID); verifyErr != nil {
		api.WriteErrorResponse(w, verifyErr)
		return
	}

	api.WriteJSONResponse(w, tableRef)

}

func getTableColsAPI(w http.ResponseWriter, r *http.Request) {

	var params GetTableParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	tableRef, err := GetTable(trackerDBHandle, params.TableID)
	if err != nil {
		api.WriteErrorResponse(w, err)
	}

	if verifyErr := userRole.VerifyCurrUserIsDatabaseAdmin(trackerDBHandle,
		r, tableRef.ParentDatabaseID); verifyErr != nil {
		api.WriteErrorResponse(w, verifyErr)
		return
	}

	tableColInfo, err := getTableDisplayInfo(trackerDBHandle, params.TableID)
	if err != nil {
		api.WriteErrorResponse(w, err)
	}

	api.WriteJSONResponse(w, tableColInfo.Cols)

}

type DeleteTableColParams struct {
	ParentTableID string `json:"parentTableID"`
	ColumnID      string `json:"columnID"`
}

func deleteColumnAPI(w http.ResponseWriter, r *http.Request) {

	var params DeleteTableColParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	tableRef, err := GetTable(trackerDBHandle, params.ParentTableID)
	if err != nil {
		api.WriteErrorResponse(w, err)
	}

	if verifyErr := userRole.VerifyCurrUserIsDatabaseAdmin(trackerDBHandle,
		r, tableRef.ParentDatabaseID); verifyErr != nil {
		api.WriteErrorResponse(w, verifyErr)
		return
	}

	if err := common.DeleteTableColumn(trackerDBHandle, params.ParentTableID, params.ColumnID); err != nil {
		api.WriteErrorResponse(w, err)
	}

	api.WriteJSONResponse(w, true)

}

func getTableDisplayInfoAPI(w http.ResponseWriter, r *http.Request) {

	var params GetTableParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	/* TODO - Restore authentication for or refactor to get a list to only include table views shown for a given item list
	tableRef, err := GetTable(params.TableID)
	if err != nil {
		api.WriteErrorResponse(w, err)
	}

	if verifyErr := userRole.VerifyCurrUserIsDatabaseAdmin(
		r, tableRef.ParentDatabaseID); verifyErr != nil {
		api.WriteErrorResponse(w, verifyErr)
		return
	} */

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	displayInfo, err := getTableDisplayInfo(trackerDBHandle, params.TableID)
	if err != nil {
		api.WriteErrorResponse(w, err)
	}

	api.WriteJSONResponse(w, *displayInfo)

}

func validateTableNameAPI(w http.ResponseWriter, r *http.Request) {

	tableName := r.FormValue("tableName")
	formID := r.FormValue("formID")

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if err := validateTableName(trackerDBHandle, formID, tableName); err != nil {
		api.WriteJSONResponse(w, fmt.Sprintf("%v", err))
		return
	}

	response := true
	api.WriteJSONResponse(w, response)

}

func validateNewTableNameAPI(w http.ResponseWriter, r *http.Request) {

	tableName := r.FormValue("tableName")
	databaseID := r.FormValue("databaseID")

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if err := validateNewTableName(trackerDBHandle, databaseID, tableName); err != nil {
		api.WriteJSONResponse(w, fmt.Sprintf("%v", err))
		return
	}

	response := true
	api.WriteJSONResponse(w, response)

}

func processTablePropUpdate(w http.ResponseWriter, r *http.Request, propUpdater TablePropUpdater) {

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if updatedTable, err := updateTableProps(trackerDBHandle, propUpdater); err != nil {
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

func setOrderedCols(w http.ResponseWriter, r *http.Request) {
	var params SetOrderedColParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processTablePropUpdate(w, r, params)
}
