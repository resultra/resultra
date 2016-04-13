package dashboard

import (
	"appengine"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"resultra/datasheet/server/common/api"
	"resultra/datasheet/server/dashboard"
)

var designDashboardTemplates = template.Must(template.ParseFiles(
	"common/common.html",
	"dashboard/barChart/barChartProps.html",
	"dashboard/barChart/newBarChartDialog.html",
	"dashboard/dashboardCommon.html",
	"dashboard/dashboardProps.html",
	"dashboard/designDashboard.html"))

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
		api.WriteErrorResponse(w, getErr)
		return
	}

	p := DesignDashboardPageInfo{
		Title:         "Design Dashboard",
		DashboardID:   dashboardID,
		DashboardName: dashboardRef.Name}
	err := designDashboardTemplates.ExecuteTemplate(w, "designDashboard", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
