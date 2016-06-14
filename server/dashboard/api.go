package dashboard

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/generic/api"
)

func init() {

	dashboardRouter := mux.NewRouter()

	dashboardRouter.HandleFunc("/api/dashboard/new", newDashboard)
	dashboardRouter.HandleFunc("/api/dashboard/getData", getDashboardData)

	http.Handle("/api/dashboard/", dashboardRouter)
}

func newDashboard(w http.ResponseWriter, r *http.Request) {

	var dashboardParams NewDashboardParams
	if err := api.DecodeJSONRequest(r, &dashboardParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if dashboardRef, err := NewDashboard(dashboardParams); err != nil {
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

	if dashboardData, err := GetDashboardData(dashboardParams); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *dashboardData)
	}

}
