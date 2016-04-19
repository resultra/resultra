package record

import (
	"appengine"
	"net/http"
	"resultra/datasheet/server/generic/api"
)

func newRecord(w http.ResponseWriter, r *http.Request) {

	appEngCntxt := appengine.NewContext(r)
	newRecordRef, err := NewRecord(appEngCntxt)
	if err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *newRecordRef)
	}

}

func getRecord(w http.ResponseWriter, r *http.Request) {

	params := RecordID{}
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	appEngCntxt := appengine.NewContext(r)

	if recordRef, getErr := GetRecord(appEngCntxt, params); getErr != nil {
		api.WriteErrorResponse(w, getErr)
		return
	} else {
		api.WriteJSONResponse(w, recordRef)
	}

}

func getRecords(w http.ResponseWriter, r *http.Request) {

	// TODO - Once sorting and filtering is implemented, the request
	// will need to include parameters for the sort and filter parameters to use.

	appEngCntxt := appengine.NewContext(r)
	if recordRefs, err := GetRecords(appEngCntxt); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, recordRefs)
	}

}
