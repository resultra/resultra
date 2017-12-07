package users

import (
	"github.com/gorilla/mux"
)

func RegisterHTTPHandlers(mainRouter *mux.Router) {
	mainRouter.HandleFunc("/workspace-admin/users", userAdminPage)
	mainRouter.HandleFunc("/workspace-admin/user/{userID}", userPropsPage)
}
