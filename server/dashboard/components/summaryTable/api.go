package summaryTable

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/tracker/server/common/databaseWrapper"
	"resultra/tracker/server/generic/api"
)

func init() {

	summaryTableRouter := mux.NewRouter()

	summaryTableRouter.HandleFunc("/api/dashboard/summaryTable/new", newSummaryTableAPI)

	summaryTableRouter.HandleFunc("/api/dashboard/summaryTable/setTitle", setSummaryTableTitle)
	summaryTableRouter.HandleFunc("/api/dashboard/summaryTable/setDimensions", setSummaryTableDimensions)
	summaryTableRouter.HandleFunc("/api/dashboard/summaryTable/setColumns", setSummaryTableColumns)

	summaryTableRouter.HandleFunc("/api/dashboard/summaryTable/setDefaultFilterRules", setDefaultFilterRules)
	summaryTableRouter.HandleFunc("/api/dashboard/summaryTable/setPreFilterRules", setPreFilterRules)

	summaryTableRouter.HandleFunc("/api/dashboard/summaryTable/setHelpPopupMsg", setHelpPopupMsg)

	summaryTableRouter.HandleFunc("/api/dashboard/summaryTable/setRowValueGrouping", setRowValueGrouping)

	http.Handle("/api/dashboard/summaryTable/", summaryTableRouter)
}

func newSummaryTableAPI(w http.ResponseWriter, r *http.Request) {

	var params NewSummaryTableParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if summaryTableRef, err := newSummaryTable(trackerDBHandle, params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, summaryTableRef)
	}

}

func processSummaryTablePropUpdate(w http.ResponseWriter, r *http.Request, propUpdater SummaryTablePropertyUpdater) {

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if summaryTableRef, err := updateSummaryTableProps(trackerDBHandle, propUpdater); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, summaryTableRef)
	}
}

func setSummaryTableTitle(w http.ResponseWriter, r *http.Request) {
	var titleParams SetSummaryTableTitleParams
	if err := api.DecodeJSONRequest(r, &titleParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processSummaryTablePropUpdate(w, r, titleParams)
}

func setSummaryTableDimensions(w http.ResponseWriter, r *http.Request) {

	var params SetSummaryTableDimensionsParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processSummaryTablePropUpdate(w, r, params)
}

func setSummaryTableColumns(w http.ResponseWriter, r *http.Request) {
	var params SetSummaryTableSummaryColumns
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processSummaryTablePropUpdate(w, r, params)
}

func setRowValueGrouping(w http.ResponseWriter, r *http.Request) {
	var params SetRowGroupingParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processSummaryTablePropUpdate(w, r, params)
}

func setDefaultFilterRules(w http.ResponseWriter, r *http.Request) {
	var params SetSummaryTableDefaultFilterRulesParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processSummaryTablePropUpdate(w, r, params)
}

func setPreFilterRules(w http.ResponseWriter, r *http.Request) {
	var params SetSummaryTablePreFilterRulesParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processSummaryTablePropUpdate(w, r, params)
}

func setHelpPopupMsg(w http.ResponseWriter, r *http.Request) {
	var params SetHelpPopupMsgParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processSummaryTablePropUpdate(w, r, params)
}
