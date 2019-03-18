package barChart

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/tracker/server/common/databaseWrapper"
	"resultra/tracker/server/generic/api"
)

func init() {

	barChartRouter := mux.NewRouter()

	barChartRouter.HandleFunc("/api/dashboard/barChart/new", newBarChart)
	barChartRouter.HandleFunc("/api/dashboard/barChart/setTitle", setBarChartTitle)
	barChartRouter.HandleFunc("/api/dashboard/barChart/setDimensions", setBarChartDimensions)
	barChartRouter.HandleFunc("/api/dashboard/barChart/setXAxisValueGrouping", setXAxisValueGrouping)
	barChartRouter.HandleFunc("/api/dashboard/barChart/setYAxisSummaryVals", setYAxisSummaryVals)
	barChartRouter.HandleFunc("/api/dashboard/barChart/setHelpPopupMsg", setHelpPopupMsg)

	barChartRouter.HandleFunc("/api/dashboard/barChart/setDefaultFilterRules", setDefaultFilterRules)
	barChartRouter.HandleFunc("/api/dashboard/barChart/setPreFilterRules", setPreFilterRules)

	http.Handle("/api/dashboard/barChart/", barChartRouter)
}

func newBarChart(w http.ResponseWriter, r *http.Request) {

	var barChartParams NewBarChartParams
	if err := api.DecodeJSONRequest(r, &barChartParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if barChartRef, err := NewBarChart(trackerDBHandle, barChartParams); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, barChartRef)
	}

}

// processBarChartPropUpdate is a helper function to process the property updates in a generic way, reducing
// duplicated code in the individual property updates. There doesn't seem to be a way to generically
// read/decode the parameters, so this is still done explicitely.
func processBarChartPropUpdate(w http.ResponseWriter, r *http.Request, propUpdater BarChartPropertyUpdater) {

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if barChartRef, err := UpdateBarChartProps(trackerDBHandle, propUpdater); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, barChartRef)
	}
}

func setBarChartTitle(w http.ResponseWriter, r *http.Request) {
	var titleParams SetBarChartTitleParams
	if err := api.DecodeJSONRequest(r, &titleParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processBarChartPropUpdate(w, r, titleParams)
}

func setBarChartDimensions(w http.ResponseWriter, r *http.Request) {

	var params SetBarChartDimensionsParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processBarChartPropUpdate(w, r, params)
}

func setXAxisValueGrouping(w http.ResponseWriter, r *http.Request) {
	var params SetXAxisValuesParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processBarChartPropUpdate(w, r, params)

}

func setYAxisSummaryVals(w http.ResponseWriter, r *http.Request) {
	var params SetYAxisSummaryParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processBarChartPropUpdate(w, r, params)

}

func setDefaultFilterRules(w http.ResponseWriter, r *http.Request) {
	var params SetDefaultFilterRulesParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processBarChartPropUpdate(w, r, params)

}

func setPreFilterRules(w http.ResponseWriter, r *http.Request) {
	var params SetPreFilterRulesParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processBarChartPropUpdate(w, r, params)

}

func setHelpPopupMsg(w http.ResponseWriter, r *http.Request) {
	var params SetHelpPopupMsgParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processBarChartPropUpdate(w, r, params)

}
