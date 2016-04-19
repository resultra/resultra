package dashboard

import (
	"appengine"
	"net/http"
	"resultra/datasheet/server/generic/api"
)

func newDashboard(w http.ResponseWriter, r *http.Request) {

	var dashboardParam map[string]string
	if err := api.DecodeJSONRequest(r, &dashboardParam); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	appEngCntxt := appengine.NewContext(r)
	if dashboardRef, err := NewDashboard(appEngCntxt, dashboardParam["name"]); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, dashboardRef)
	}

}

func getDashboardData(w http.ResponseWriter, r *http.Request) {

	var dashboardParams GetDashboardDataParams
	if err := api.DecodeJSONRequest(r, &dashboardParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	appEngCntxt := appengine.NewContext(r)
	if dashboardData, err := GetDashboardData(appEngCntxt, dashboardParams); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *dashboardData)
	}

}
