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
