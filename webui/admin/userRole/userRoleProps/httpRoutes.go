package userRoleProps

import (
	"github.com/gorilla/mux"
)

func RegisterHTTPHandlers(mainRouter *mux.Router) {
	mainRouter.HandleFunc("/admin/userRole/{roleID}", editRolePropsPage)
}
