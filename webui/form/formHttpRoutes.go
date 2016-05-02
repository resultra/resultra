package form

import (
	"github.com/gorilla/mux"
)

func RegisterHTTPHandlers(mainRouter *mux.Router) {
	mainRouter.HandleFunc("/viewForm/{formID}", viewForm)
	mainRouter.HandleFunc("/designForm/{formID}", designForm)
}
