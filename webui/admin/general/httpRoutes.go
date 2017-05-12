package general

import (
	"github.com/gorilla/mux"
)

func RegisterHTTPHandlers(mainRouter *mux.Router) {
	mainRouter.HandleFunc("/admin/{databaseID}", generalAdminPage)

	mainRouter.HandleFunc("/admin/general/{databaseID}", generalAdminPage)
}
