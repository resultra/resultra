package alertList

import (
	"github.com/gorilla/mux"
)

func RegisterHTTPHandlers(mainRouter *mux.Router) {
	mainRouter.HandleFunc("/admin/alerts/{databaseID}", alertListAdminPage)
}
