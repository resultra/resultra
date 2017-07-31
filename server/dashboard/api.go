package dashboard

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/generic/api"
	"resultra/datasheet/server/userRole"
)

func init() {

	dashboardRouter := mux.NewRouter()

	dashboardRouter.HandleFunc("/api/dashboard/new", newDashboard)

	dashboardRouter.HandleFunc("/api/dashboard/getProperties", getDashboardPropsAPI)
	dashboardRouter.HandleFunc("/api/dashboard/setName", setNameAPI)
	dashboardRouter.HandleFunc("/api/dashboard/setLayout", setLayoutAPI)
	dashboardRouter.HandleFunc("/api/dashboard/getUserDashboardList", getUserDashboardListAPI)

	dashboardRouter.HandleFunc("/api/dashboard/validateNewDashboardName", validateNewDashboardNameAPI)
	dashboardRouter.HandleFunc("/api/dashboard/validateDashboardName", validateDashboardNameAPI)
	dashboardRouter.HandleFunc("/api/dashboard/validateComponentTitle", validateComponentTitleAPI)

	http.Handle("/api/dashboard/", dashboardRouter)
}

func newDashboard(w http.ResponseWriter, r *http.Request) {

	var dashboardParams NewDashboardParams
	if err := api.DecodeJSONRequest(r, &dashboardParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if verifyUserErr := userRole.VerifyCurrUserIsDatabaseAdmin(
		r, dashboardParams.DatabaseID); verifyUserErr != nil {
		api.WriteErrorResponse(w, verifyUserErr)
		return
	}

	if dashboardRef, err := NewDashboard(dashboardParams); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, dashboardRef)
	}

}

func getDashboardPropsAPI(w http.ResponseWriter, r *http.Request) {

	var params GetDashboardParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if dashboardProps, err := GetDashboard(params.DashboardID); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *dashboardProps)
	}

}

func validateNewDashboardNameAPI(w http.ResponseWriter, r *http.Request) {

	dashboardName := r.FormValue("dashboardName")
	databaseID := r.FormValue("databaseID")

	if err := validateNewDashboardName(databaseID, dashboardName); err != nil {
		api.WriteJSONResponse(w, fmt.Sprintf("%v", err))
		return
	}

	response := true
	api.WriteJSONResponse(w, response)

}

func validateDashboardNameAPI(w http.ResponseWriter, r *http.Request) {
	dashboardName := r.FormValue("dashboardName")
	dashboardID := r.FormValue("dashboardID")

	if err := validateDashboardName(dashboardID, dashboardName); err != nil {
		api.WriteJSONResponse(w, fmt.Sprintf("%v", err))
		return
	}

	response := true
	api.WriteJSONResponse(w, response)

}

func validateComponentTitleAPI(w http.ResponseWriter, r *http.Request) {

	title := r.FormValue(`title`)

	if err := validateComponentTitle(title); err != nil {
		api.WriteJSONResponse(w, fmt.Sprintf("%v", err))
		return
	}

	response := true
	api.WriteJSONResponse(w, response)

}

func processDashboardPropUpdate(w http.ResponseWriter, r *http.Request, propUpdater DashboardPropUpdater) {
	if updatedDB, err := updateDashboardProps(propUpdater); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, updatedDB)
	}
}

func setNameAPI(w http.ResponseWriter, r *http.Request) {
	var params SetDashboardNameParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processDashboardPropUpdate(w, r, params)
}

func setLayoutAPI(w http.ResponseWriter, r *http.Request) {
	var params SetDashboardLayoutParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processDashboardPropUpdate(w, r, params)
}

func getUserDashboardListAPI(w http.ResponseWriter, r *http.Request) {
	var params GetUserDashboardListParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	if dashboards, err := getUserDashboards(r, params.DatabaseID); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, dashboards)
	}
}
