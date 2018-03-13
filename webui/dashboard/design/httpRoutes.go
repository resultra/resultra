package design

import (
	"github.com/gorilla/mux"
)

func RegisterHTTPHandlers(mainRouter *mux.Router) {
	mainRouter.HandleFunc("/admin/dashboard/designDashboardMainContent/{dashboardID}", designDashboardMainContent)
	mainRouter.HandleFunc("/admin/dashboard/designDashboardOffpageContent/{dashboardID}", designDashboardOffpageContent)
	mainRouter.HandleFunc("/admin/dashboard/designDashboardSidebarContent/{dashboardID}", designDashboardSidebarContent)

}
