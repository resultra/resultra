package form

import (
	"github.com/gorilla/mux"
)

func RegisterHTTPHandlers(mainRouter *mux.Router) {
	mainRouter.HandleFunc("/viewForm/{layoutID}", viewForm)
	mainRouter.HandleFunc("/designForm/{layoutID}", designForm)
}
