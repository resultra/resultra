package dashboard

import (
	"github.com/gorilla/mux"
	"resultra/datasheet/server/dashboard/components/barChart"
)

func RegisterHTTPHandlers(apiRouter *mux.Router) {

	barChart.RegisterHTTPHandlers(apiRouter)

	apiRouter.HandleFunc("/api/newDashboard", newDashboard)
	apiRouter.HandleFunc("/api/getDashboardData", getDashboardData)

}
