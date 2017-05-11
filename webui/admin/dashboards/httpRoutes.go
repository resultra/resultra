package dashboards

import (
	"github.com/gorilla/mux"
)

func RegisterHTTPHandlers(mainRouter *mux.Router) {
	mainRouter.HandleFunc("/admin/dashboards/{databaseID}", dashboardAdminPage)
}
