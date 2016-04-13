package recordFilter

import (
	"github.com/gorilla/mux"
)

func RegisterHTTPHandlers(apiRouter *mux.Router) {
	apiRouter.HandleFunc("/api/getFilteredRecords", getFilteredRecords)
	apiRouter.HandleFunc("/api/newRecordFilterRule", newRecordFilterRule)
	apiRouter.HandleFunc("/api/getRecordFilterRules", getRecordFilterRules)

}
