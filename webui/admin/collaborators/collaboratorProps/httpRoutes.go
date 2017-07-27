package collaboratorProps

import (
	"github.com/gorilla/mux"
)

func RegisterHTTPHandlers(mainRouter *mux.Router) {
	mainRouter.HandleFunc("/admin/collaborator/{databaseID}/{userID}", editCollabPropsPage)
}
