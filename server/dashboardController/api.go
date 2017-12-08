package dashboardController

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/common/databaseWrapper"
	"resultra/datasheet/server/generic/api"
	"resultra/datasheet/server/common/userAuth"
)

type DummyStructForInclude struct{ Val int64 }

func init() {

	dashboardControllerRouter := mux.NewRouter()

	dashboardControllerRouter.HandleFunc("/api/dashboardController/getDefaultData", getDefaultDashboardDataAPI)

	dashboardControllerRouter.HandleFunc("/api/dashboardController/getBarChartData", getBarChartDataAPI)
	dashboardControllerRouter.HandleFunc("/api/dashboardController/getSummaryTableData", getSummaryTableDataAPI)
	dashboardControllerRouter.HandleFunc("/api/dashboardController/getGaugeData", getGaugeDataAPI)
	dashboardControllerRouter.HandleFunc("/api/dashboardController/getSummaryValData", getSummaryValDataAPI)

	http.Handle("/api/dashboardController/", dashboardControllerRouter)
}

func getDefaultDashboardDataAPI(w http.ResponseWriter, r *http.Request) {

	var dashboardParams GetDashboardDataParams
	if err := api.DecodeJSONRequest(r, &dashboardParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	currUserID, userErr := userAuth.GetCurrentUserID(r)
	if userErr != nil {
		api.WriteJSONResponse(w, fmt.Errorf("Can't verify user authentication"))
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if dashboardData, err := getDefaultDashboardData(trackerDBHandle, currUserID, dashboardParams); err != nil {
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

	currUserID, userErr := userAuth.GetCurrentUserID(r)
	if userErr != nil {
		api.WriteJSONResponse(w, fmt.Errorf("Can't verify user authentication"))
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if summaryTableData, err := getSummaryTableData(trackerDBHandle, currUserID, params); err != nil {
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

	currUserID, userErr := userAuth.GetCurrentUserID(r)
	if userErr != nil {
		api.WriteJSONResponse(w, fmt.Errorf("Can't verify user authentication"))
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if barChartData, err := getBarChartData(trackerDBHandle, currUserID, barChartParams); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, barChartData)
	}

}

func getGaugeDataAPI(w http.ResponseWriter, r *http.Request) {

	var params GetGaugeDataParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	currUserID, userErr := userAuth.GetCurrentUserID(r)
	if userErr != nil {
		api.WriteJSONResponse(w, fmt.Errorf("Can't verify user authentication"))
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if gaugeData, err := getGaugeData(trackerDBHandle, currUserID, params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, gaugeData)
	}

}

func getSummaryValDataAPI(w http.ResponseWriter, r *http.Request) {

	var params GetSummaryValDataParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	currUserID, userErr := userAuth.GetCurrentUserID(r)
	if userErr != nil {
		api.WriteJSONResponse(w, fmt.Errorf("Can't verify user authentication"))
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if summaryValData, err := getSummaryValData(trackerDBHandle, currUserID, params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, summaryValData)
	}

}
