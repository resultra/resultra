package dashboard

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/common/databaseWrapper"
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

	dashboardRouter.HandleFunc("/api/dashboard/deleteComponent", deleteComponentAPI)

	dashboardRouter.HandleFunc("/api/dashboard/setIncludeInSidebar", setIncludeInSidebar)

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

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if verifyUserErr := userRole.VerifyCurrUserIsDatabaseAdmin(trackerDBHandle,
		r, dashboardParams.DatabaseID); verifyUserErr != nil {
		api.WriteErrorResponse(w, verifyUserErr)
		return
	}

	if dashboardRef, err := NewDashboard(trackerDBHandle, dashboardParams); err != nil {
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

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if dashboardProps, err := GetDashboard(trackerDBHandle, params.DashboardID); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *dashboardProps)
	}

}

func validateNewDashboardNameAPI(w http.ResponseWriter, r *http.Request) {

	dashboardName := r.FormValue("dashboardName")
	databaseID := r.FormValue("databaseID")

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if err := validateNewDashboardName(trackerDBHandle, databaseID, dashboardName); err != nil {
		api.WriteJSONResponse(w, fmt.Sprintf("%v", err))
		return
	}

	response := true
	api.WriteJSONResponse(w, response)

}

func validateDashboardNameAPI(w http.ResponseWriter, r *http.Request) {
	dashboardName := r.FormValue("dashboardName")
	dashboardID := r.FormValue("dashboardID")

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if err := validateDashboardName(trackerDBHandle, dashboardID, dashboardName); err != nil {
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

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if updatedDB, err := updateDashboardProps(trackerDBHandle, propUpdater); err != nil {
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

func setIncludeInSidebar(w http.ResponseWriter, r *http.Request) {
	var params SetIncludeInSidebarParams
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

func deleteComponentAPI(w http.ResponseWriter, r *http.Request) {

	var params DeleteComponentParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if err := deleteComponent(trackerDBHandle, params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, true)
	}

}
