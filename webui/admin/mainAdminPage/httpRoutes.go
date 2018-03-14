package mainAdminPage

import (
	"github.com/gorilla/mux"
)

func RegisterHTTPHandlers(mainRouter *mux.Router) {
	mainRouter.HandleFunc("/admin/mainAdminPage/mainPageContent/{databaseID}", mainAdminPage)
	mainRouter.HandleFunc("/admin/mainAdminPage/offPageContent", mainAdminPageOffPageContent)
}
