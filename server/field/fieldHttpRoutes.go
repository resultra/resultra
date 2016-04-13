package field

import (
	"github.com/gorilla/mux"
)

func RegisterHTTPHandlers(apiRouter *mux.Router) {

	apiRouter.HandleFunc("/api/newField", newField)
	apiRouter.HandleFunc("/api/getFieldsByType", getFieldsByType)

}
