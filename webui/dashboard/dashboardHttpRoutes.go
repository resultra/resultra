package dashboard

import (
	"github.com/gorilla/mux"
	"resultra/datasheet/webui/dashboard/design"
	"resultra/datasheet/webui/dashboard/view"
)

func RegisterHTTPHandlers(mainRouter *mux.Router) {
	mainRouter.HandleFunc("/designDashboard/{dashboardID}", design.DesignDashboard)
	mainRouter.HandleFunc("/viewDashboard/{dashboardID}", view.ViewDashboard)
}
