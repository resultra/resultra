package dashboardController

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/generic/api"
)

type DummyStructForInclude struct{ Val int64 }

func init() {

	dashboardControllerRouter := mux.NewRouter()

	dashboardControllerRouter.HandleFunc("/api/dashboardController/getDefaultData", getDefaultDashboardDataAPI)

	dashboardControllerRouter.HandleFunc("/api/dashboardController/getBarChartData", getBarChartDataAPI)
	dashboardControllerRouter.HandleFunc("/api/dashboardController/getSummaryTableData", getSummaryTableDataAPI)

	http.Handle("/api/dashboardController/", dashboardControllerRouter)
}

func getDefaultDashboardDataAPI(w http.ResponseWriter, r *http.Request) {

	var dashboardParams GetDashboardDataParams
	if err := api.DecodeJSONRequest(r, &dashboardParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if dashboardData, err := getDefaultDashboardData(dashboardParams); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *dashboardData)
	}

}

func getSummaryTableDataAPI(w http.ResponseWriter, r *http.Request) {

	var params GetSummaryTableDataParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if summaryTableData, err := getSummaryTableData(params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, summaryTableData)
	}

}

func getBarChartDataAPI(w http.ResponseWriter, r *http.Request) {

	var barChartParams GetBarChartDataParams
	if err := api.DecodeJSONRequest(r, &barChartParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if barChartData, err := getBarChartData(barChartParams); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, barChartData)
	}

}
