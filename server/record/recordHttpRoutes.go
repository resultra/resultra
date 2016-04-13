package record

import (
	"github.com/gorilla/mux"
)

func RegisterHTTPHandlers(apiRouter *mux.Router) {

	apiRouter.HandleFunc("/api/newRecord", newRecord)
	apiRouter.HandleFunc("/api/getRecord", getRecord)
	apiRouter.HandleFunc("/api/getRecords", getRecords)

}
