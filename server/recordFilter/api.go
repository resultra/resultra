package recordFilter

import (
	"appengine"
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/generic/api"
)

func init() {

	filterRouter := mux.NewRouter()

	filterRouter.HandleFunc("/api/filter/new", newFilterAPI)

	http.Handle("/api/filter/", filterRouter)
}

func RegisterHTTPHandlers(apiRouter *mux.Router) {
	apiRouter.HandleFunc("/api/getFilteredRecords", getFilteredRecords)
	apiRouter.HandleFunc("/api/newRecordFilterRule", newRecordFilterRule)
	apiRouter.HandleFunc("/api/getRecordFilterRules", getRecordFilterRules)

}

func newRecordFilterRule(w http.ResponseWriter, r *http.Request) {

	filterRuleParams := NewFilterRuleParams{}
	if err := api.DecodeJSONRequest(r, &filterRuleParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	appEngCntxt := appengine.NewContext(r)
	newRecordFilterRef, newErr := NewFilterRule(appEngCntxt, filterRuleParams)
	if newErr != nil {
		api.WriteErrorResponse(w, newErr)
		return
	} else {
		api.WriteJSONResponse(w, newRecordFilterRef)
	}

}

func getRecordFilterRules(w http.ResponseWriter, r *http.Request) {

	appEngCntxt := appengine.NewContext(r)
	if filterRefs, err := GetRecordFilterRefs(appEngCntxt); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, filterRefs)
	}

}

func getFilteredRecords(w http.ResponseWriter, r *http.Request) {

	// TODO - Once filtering is implemented on a per form/dashboard basis,
	// pass in the parent filter.
	var params GetFilteredRecordsParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	appEngCntxt := appengine.NewContext(r)
	if recordRefs, err := GetFilteredRecords(appEngCntxt, params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, recordRefs)
	}

}

func newFilterAPI(w http.ResponseWriter, r *http.Request) {

	var params NewFilterParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	appEngCntxt := appengine.NewContext(r)
	newFilterRef, newErr := newFilter(appEngCntxt, params)
	if newErr != nil {
		api.WriteErrorResponse(w, newErr)
		return
	} else {
		api.WriteJSONResponse(w, newFilterRef)
	}

}
