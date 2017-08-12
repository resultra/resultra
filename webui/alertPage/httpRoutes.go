package alertPage

import (
	"github.com/gorilla/mux"
)

func RegisterHTTPHandlers(mainRouter *mux.Router) {

	mainRouter.HandleFunc("/alerts/{databaseID}", alertPage)

}
