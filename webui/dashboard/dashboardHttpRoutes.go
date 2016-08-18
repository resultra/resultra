package dashboard

import (
	"github.com/gorilla/mux"
	"resultra/datasheet/webui/dashboard/design"
)

func RegisterHTTPHandlers(mainRouter *mux.Router) {
	mainRouter.HandleFunc("/designDashboard/{dashboardID}", design.DesignDashboard)
}
