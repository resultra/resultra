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

func getDashboardData(w http.ResponseWriter, r *http.Request) {

	var dashboardParams dashboard.GetDashboardDataParams
	if err := decodeJSONRequest(r, &dashboardParams); err != nil {
		WriteErrorResponse(w, err)
		return
	}

	appEngCntxt := appengine.NewContext(r)
	if dashboardData, err := dashboard.GetDashboardData(appEngCntxt, dashboardParams); err != nil {
		WriteErrorResponse(w, err)
	} else {
		writeJSONResponse(w, *dashboardData)
	}

}

func updateBarChartProps(w http.ResponseWriter, r *http.Request) {

	var barChartRef dashboard.BarChartRef
	if err := decodeJSONRequest(r, &barChartRef); err != nil {
		WriteErrorResponse(w, err)
		return
	}

	appEngCntxt := appengine.NewContext(r)
	if barChartRef, err := dashboard.UpdateBarChartProps(appEngCntxt, barChartRef); err != nil {
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

	appEngCntxt := appengine.NewContext(r)
	if barChartRef, err := dashboard.SetBarChartTitle(appEngCntxt, titleParams); err != nil {
		WriteErrorResponse(w, err)
	} else {
		writeJSONResponse(w, barChartRef)
	}

}
