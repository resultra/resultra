package recordUpdate

import (
	"github.com/gorilla/mux"
)

func RegisterHTTPHandlers(apiRouter *mux.Router) {
	apiRouter.HandleFunc("/api/setTextFieldValue", setTextFieldValue)
	apiRouter.HandleFunc("/api/setNumberFieldValue", setNumberFieldValue)

}
