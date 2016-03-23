package webui

import (
	"appengine"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"resultra/datasheet/controller"
	"resultra/datasheet/datamodel/dashboard"
)

type DesignDashboardPageInfo struct {
	Title         string
	DashboardID   string
	DashboardName string
}

func designDashboard(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	dashboardID := vars["dashboardID"]
	log.Println("Design Dashboard: editing for dashboard with ID = ", dashboardID)

	appEngContext := appengine.NewContext(r)
	dashboardRef, getErr := dashboard.GetDashboardRef(appEngContext, dashboardID)
	if getErr != nil {
		controller.WriteErrorResponse(w, getErr)
		return
	}

	p := DesignDashboardPageInfo{"Design Dashboard", dashboardID, dashboardRef.Name}
	err := htmlTemplates.ExecuteTemplate(w, "designDashboard", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
