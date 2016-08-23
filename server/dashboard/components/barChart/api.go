package barChart

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/generic/api"
)

func init() {

	barChartRouter := mux.NewRouter()

	barChartRouter.HandleFunc("/api/dashboard/barChart/new", newBarChart)
	barChartRouter.HandleFunc("/api/dashboard/barChart/getData", getBarChartData)
	barChartRouter.HandleFunc("/api/dashboard/barChart/setTitle", setBarChartTitle)
	barChartRouter.HandleFunc("/api/dashboard/barChart/setDimensions", setBarChartDimensions)

	barChartRouter.HandleFunc("/api/dashboard/barChart/setAvailableFilters", setBarChartAvailableFilters)
	barChartRouter.HandleFunc("/api/dashboard/barChart/setDefaultFilters", setBarChartDefaultFilters)

	http.Handle("/api/dashboard/barChart/", barChartRouter)
}

func newBarChart(w http.ResponseWriter, r *http.Request) {

	var barChartParams NewBarChartParams
	if err := api.DecodeJSONRequest(r, &barChartParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if barChartRef, err := NewBarChart(barChartParams); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, barChartRef)
	}

}

func getBarChartData(w http.ResponseWriter, r *http.Request) {

	var barChartParams BarChartUniqueIDHeader
	if err := api.DecodeJSONRequest(r, &barChartParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if barChartData, err := GetBarChartData(barChartParams.ParentDashboardID, barChartParams.BarChartID); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, barChartData)
	}

}

// processBarChartPropUpdate is a helper function to process the property updates in a generic way, reducing
// duplicated code in the individual property updates. There doesn't seem to be a way to generically
// read/decode the parameters, so this is still done explicitely.
func processBarChartPropUpdate(w http.ResponseWriter, r *http.Request, propUpdater BarChartPropertyUpdater) {

	if barChartRef, err := UpdateBarChartProps(propUpdater); err != nil {
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

func setBarChartAvailableFilters(w http.ResponseWriter, r *http.Request) {
	var params SetBarChartAvailableFilterParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processBarChartPropUpdate(w, r, params)
}

func setBarChartDefaultFilters(w http.ResponseWriter, r *http.Request) {
	var params SetBarChartDefaultFilterParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processBarChartPropUpdate(w, r, params)
}
