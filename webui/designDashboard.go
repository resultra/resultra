package webui

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
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

	dashboardName := "Test Dashboard"
	p := DesignDashboardPageInfo{"Design Dashboard", dashboardID, dashboardName}
	err := htmlTemplates.ExecuteTemplate(w, "designDashboard", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
