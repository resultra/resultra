package dashboard

import (
	"github.com/gorilla/mux"
)

func RegisterHTTPHandlers(mainRouter *mux.Router) {
	mainRouter.HandleFunc("/designDashboard/{dashboardID}", designDashboard)
}
