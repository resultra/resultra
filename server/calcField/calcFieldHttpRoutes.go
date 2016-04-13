package calcField

import (
	"github.com/gorilla/mux"
)

func RegisterHTTPHandlers(apiRouter *mux.Router) {

	apiRouter.HandleFunc("/api/validateCalcFieldEqn", validateCalcFieldEqn)
	apiRouter.HandleFunc("/api/newCalcField", newCalcField)

}
