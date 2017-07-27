package collaboratorList

import (
	"github.com/gorilla/mux"
)

func RegisterHTTPHandlers(mainRouter *mux.Router) {
	mainRouter.HandleFunc("/admin/collaborators/{databaseID}", userAdminPage)
}
