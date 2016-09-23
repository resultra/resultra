package summaryTable

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/generic/api"
)

func init() {

	summaryTableRouter := mux.NewRouter()

	summaryTableRouter.HandleFunc("/api/dashboard/summaryTable/new", newSummaryTableAPI)

	summaryTableRouter.HandleFunc("/api/dashboard/summaryTable/setTitle", setSummaryTableTitle)
	summaryTableRouter.HandleFunc("/api/dashboard/summaryTable/setDimensions", setSummaryTableDimensions)
	summaryTableRouter.HandleFunc("/api/dashboard/summaryTable/setAvailableFilters", setSummaryTableAvailableFilters)
	summaryTableRouter.HandleFunc("/api/dashboard/summaryTable/setDefaultFilters", setSummaryTableDefaultFilters)
	summaryTableRouter.HandleFunc("/api/dashboard/summaryTable/setColumns", setSummaryTableColumns)
	/*

		summaryTableRouter.HandleFunc("/api/dashboard/summaryTable/setRowValueGrouping", setXAxisValueGrouping)

	*/

	http.Handle("/api/dashboard/summaryTable/", summaryTableRouter)
}

func newSummaryTableAPI(w http.ResponseWriter, r *http.Request) {

	var params NewSummaryTableParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if summaryTableRef, err := newSummaryTable(params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, summaryTableRef)
	}

}

func processSummaryTablePropUpdate(w http.ResponseWriter, r *http.Request, propUpdater SummaryTablePropertyUpdater) {

	if summaryTableRef, err := updateSummaryTableProps(propUpdater); err != nil {
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

func setSummaryTableAvailableFilters(w http.ResponseWriter, r *http.Request) {
	var params SetSummaryTableAvailableFilterParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processSummaryTablePropUpdate(w, r, params)
}

func setSummaryTableDefaultFilters(w http.ResponseWriter, r *http.Request) {
	var params SetSummaryTableDefaultFilterParams
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
