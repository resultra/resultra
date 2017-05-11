package valueListList

import (
	"github.com/gorilla/mux"
)

func RegisterHTTPHandlers(mainRouter *mux.Router) {
	mainRouter.HandleFunc("/admin/valuelists/{databaseID}", valueListAdminPage)
}
