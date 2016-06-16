package recordFilter

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/generic/api"
)

func init() {

	filterRouter := mux.NewRouter()

	filterRouter.HandleFunc("/api/filter/new", newFilterAPI)
	filterRouter.HandleFunc("/api/filter/newWithPrefix", newFilterWithPrefixAPI)

	filterRouter.HandleFunc("/api/filter/getList", getFilterListAPI)
	filterRouter.HandleFunc("/api/filter/setName", setFilterNameAPI)
	filterRouter.HandleFunc("/api/filter/newRule", newRecordFilterRule)
	filterRouter.HandleFunc("/api/filter/getRuleList", getRecordFilterRulesAPI)

	filterRouter.HandleFunc("/api/filter/getFilteredRecordValues", getFilteredRecords)

	http.Handle("/api/filter/", filterRouter)
}

func newRecordFilterRule(w http.ResponseWriter, r *http.Request) {

	filterRuleParams := NewFilterRuleParams{}
	if err := api.DecodeJSONRequest(r, &filterRuleParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	newRecordFilterRef, newErr := newFilterRule(filterRuleParams)
	if newErr != nil {
		api.WriteErrorResponse(w, newErr)
		return
	} else {
		api.WriteJSONResponse(w, newRecordFilterRef)
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

	if recordRefs, err := GetFilteredRecords(params); err != nil {
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

	newFilterRef, newErr := newFilter(params)
	if newErr != nil {
		api.WriteErrorResponse(w, newErr)
		return
	} else {
		api.WriteJSONResponse(w, newFilterRef)
	}

}

func newFilterWithPrefixAPI(w http.ResponseWriter, r *http.Request) {

	var params NewFilterWithPrefixParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	newFilterRef, newErr := newFilterWithPrefix(params)
	if newErr != nil {
		api.WriteErrorResponse(w, newErr)
		return
	} else {
		api.WriteJSONResponse(w, newFilterRef)
	}

}

type GetFilterListParams struct {
	ParentTableID string `json:"parentTableID"`
}

func getFilterListAPI(w http.ResponseWriter, r *http.Request) {

	var params GetFilterListParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	filterRefs, getListErr := getFilterList(params.ParentTableID)
	if getListErr != nil {
		api.WriteErrorResponse(w, getListErr)
		return
	} else {
		api.WriteJSONResponse(w, filterRefs)
	}

}

func getRecordFilterRulesAPI(w http.ResponseWriter, r *http.Request) {

	var params GetFilterRulesParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if filterRules, err := getRecordFilterRules(params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, filterRules)
	}

}

func setFilterNameAPI(w http.ResponseWriter, r *http.Request) {
	var params FilterRenameParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if filterRef, err := updateFilterProps(params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, filterRef)
	}

}
