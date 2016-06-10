package form

import (
	"github.com/gorilla/mux"
)

func RegisterHTTPHandlers(mainRouter *mux.Router) {
	mainRouter.HandleFunc("/viewForm/{tableID}/{formID}", viewForm)
	mainRouter.HandleFunc("/designForm/{tableID}/{formID}", designForm)
}
