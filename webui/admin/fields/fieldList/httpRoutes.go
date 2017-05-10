package fieldList

import (
	"github.com/gorilla/mux"
)

func RegisterHTTPHandlers(mainRouter *mux.Router) {
	mainRouter.HandleFunc("/admin/fields/{databaseID}", fieldListAdminPage)
}
