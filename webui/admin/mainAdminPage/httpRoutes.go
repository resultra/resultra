package mainAdminPage

import (
	"github.com/gorilla/mux"
)

func RegisterHTTPHandlers(mainRouter *mux.Router) {
	mainRouter.HandleFunc("/admin/{databaseID}", mainAdminPage)
	mainRouter.HandleFunc("/admin/offPageContent", mainAdminPageOffPageContent)
}
