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

func setBarChartTitle(w http.ResponseWriter, r *http.Request) {

	var titleParams dashboard.SetBarChartTitleParams
	if err := decodeJSONRequest(r, &titleParams); err != nil {
		WriteErrorResponse(w, err)
		return
	}

	appEngCntxt := appengine.NewContext(r)
	if barChartRef, err := dashboard.UpdateBarChartProps(
		appEngCntxt, titleParams.UniqueID, titleParams); err != nil {
		WriteErrorResponse(w, err)
	} else {
		writeJSONResponse(w, barChartRef)
	}

}

func setBarChartDimensions(w http.ResponseWriter, r *http.Request) {

	var params dashboard.SetBarChartDimensionsParams
	if err := decodeJSONRequest(r, &params); err != nil {
		WriteErrorResponse(w, err)
		return
	}

	appEngCntxt := appengine.NewContext(r)
	if barChartRef, err := dashboard.UpdateBarChartProps(
		appEngCntxt, params.UniqueID, params); err != nil {
		WriteErrorResponse(w, err)
	} else {
		writeJSONResponse(w, barChartRef)
	}

}
