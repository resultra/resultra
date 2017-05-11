package formLinkList

import (
	"github.com/gorilla/mux"
)

func RegisterHTTPHandlers(mainRouter *mux.Router) {
	mainRouter.HandleFunc("/admin/formlink/{databaseID}", formLinkAdminPage)
}
