package controller

import (
	"appengine"
	"net/http"
	"resultra/datasheet/datamodel/dashboard"
)

func newBarChart(w http.ResponseWriter, r *http.Request) {

	var barChartParams dashboard.NewBarChartParams
	if err := decodeJSONRequest(r, &barChartParams); err != nil {
		WriteErrorResponse(w, err)
		return
	}

	appEngCntxt := appengine.NewContext(r)
	if barChartRef, err := dashboard.NewBarChart(appEngCntxt, barChartParams); err != nil {
		WriteErrorResponse(w, err)
	} else {
		writeJSONResponse(w, barChartRef)
	}

}

func getBarChartData(w http.ResponseWriter, r *http.Request) {

	var barChartParams dashboard.BarChartUniqueID
	if err := decodeJSONRequest(r, &barChartParams); err != nil {
		WriteErrorResponse(w, err)
		return
	}

	appEngCntxt := appengine.NewContext(r)
	if barChartData, err := dashboard.GetBarChartData(appEngCntxt, barChartParams); err != nil {
		WriteErrorResponse(w, err)
	} else {
		writeJSONResponse(w, barChartData)
	}

}

// processBarChartPropUpdate is a helper function to process the property updates in a generic way, reducing
// duplicated code in the individual property updates. There doesn't seem to be a way to generically
// read/decode the parameters, so this is still done explicitely.
func processBarChartPropUpdate(w http.ResponseWriter, r *http.Request, propUpdater dashboard.BarChartPropertyUpdater) {
	appEngCntxt := appengine.NewContext(r)
	if barChartRef, err := dashboard.UpdateBarChartProps(appEngCntxt, propUpdater); err != nil {
		WriteErrorResponse(w, err)
	} else {
		writeJSONResponse(w, barChartRef)
	}
}

func setBarChartTitle(w http.ResponseWriter, r *http.Request) {
	var titleParams dashboard.SetBarChartTitleParams
	if err := decodeJSONRequest(r, &titleParams); err != nil {
		WriteErrorResponse(w, err)
		return
	}
	processBarChartPropUpdate(w, r, titleParams)
}

func setBarChartDimensions(w http.ResponseWriter, r *http.Request) {

	var params dashboard.SetBarChartDimensionsParams
	if err := decodeJSONRequest(r, &params); err != nil {
		WriteErrorResponse(w, err)
		return
	}
	processBarChartPropUpdate(w, r, params)
}
