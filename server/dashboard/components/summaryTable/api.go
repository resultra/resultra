package summaryTable

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/generic/api"
)

func init() {

	summaryTableRouter := mux.NewRouter()

	summaryTableRouter.HandleFunc("/api/dashboard/summaryTable/new", newSummaryTableAPI)
	summaryTableRouter.HandleFunc("/api/dashboard/summaryTable/getData", getSummaryTableDataAPI)

	/*

		summaryTableRouter.HandleFunc("/api/dashboard/summaryTable/setTitle", setBarChartTitle)
		summaryTableRouter.HandleFunc("/api/dashboard/summaryTable/setDimensions", setBarChartDimensions)
		summaryTableRouter.HandleFunc("/api/dashboard/summaryTable/setRowValueGrouping", setXAxisValueGrouping)

		summaryTableRouter.HandleFunc("/api/dashboard/summaryTable/setAvailableFilters", setBarChartAvailableFilters)
		summaryTableRouter.HandleFunc("/api/dashboard/summaryTable/setDefaultFilters", setBarChartDefaultFilters)
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

func getSummaryTableDataAPI(w http.ResponseWriter, r *http.Request) {

	var params GetSummaryTableDataParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if summaryTableData, err := GetSummaryTableData(params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, summaryTableData)
	}

}
