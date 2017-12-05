package workspaceAdmin

import (
	"github.com/gorilla/mux"
)

func RegisterHTTPHandlers(mainRouter *mux.Router) {
	mainRouter.HandleFunc("/workspace-admin/", workspaceAdminPage)
}
