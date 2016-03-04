package controller

import (
	"appengine"
	"net/http"
	"resultra/datasheet/datamodel"
)

func newRecordFilterRule(w http.ResponseWriter, r *http.Request) {

	filterRuleParams := datamodel.NewFilterRuleParams{}
	if err := decodeJSONRequest(r, &filterRuleParams); err != nil {
		WriteErrorResponse(w, err)
		return
	}

	appEngCntxt := appengine.NewContext(r)
	newRecordFilterRef, newErr := datamodel.NewFilterRule(appEngCntxt, filterRuleParams)
	if newErr != nil {
		WriteErrorResponse(w, newErr)
		return
	} else {
		writeJSONResponse(w, newRecordFilterRef)
	}

}

func getRecordFilterRules(w http.ResponseWriter, r *http.Request) {

	appEngCntxt := appengine.NewContext(r)
	if filterRefs, err := datamodel.GetRecordFilterRefs(appEngCntxt); err != nil {
		WriteErrorResponse(w, err)
	} else {
		writeJSONResponse(w, filterRefs)
	}

}
