package controller

import (
	"appengine"
	"net/http"
	"resultra/datasheet/datamodel/dashboard"
)

func newDashboard(w http.ResponseWriter, r *http.Request) {

	var dashboardParam map[string]string
	if err := decodeJSONRequest(r, &dashboardParam); err != nil {
		WriteErrorResponse(w, err)
		return
	}

	appEngCntxt := appengine.NewContext(r)
	if dashboardRef, err := dashboard.NewDashboard(appEngCntxt, dashboardParam["name"]); err != nil {
		WriteErrorResponse(w, err)
	} else {
		writeJSONResponse(w, dashboardRef)
	}

}

func newBarChart(w http.ResponseWriter, r *http.Request) {

	var barChartParams dashboard.NewBarChartParams
	if err := decodeJSONRequest(r, &barChartParams); err != nil {
		WriteErrorResponse(w, err)
		return
	}

	appEngCntxt := appengine.NewContext(r)
	if dashboardRef, err := dashboard.NewBarChart(appEngCntxt, barChartParams); err != nil {
		WriteErrorResponse(w, err)
	} else {
		writeJSONResponse(w, dashboardRef)
	}

}

func getBarChartData(w http.ResponseWriter, r *http.Request) {

	var barChartParams dashboard.GetBarChartParams
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
