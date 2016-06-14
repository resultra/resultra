package record

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/generic/api"
)

func init() {
	recordRouter := mux.NewRouter()

	recordRouter.HandleFunc("/api/record/new", newRecord)
	recordRouter.HandleFunc("/api/record/get", getRecord)
	recordRouter.HandleFunc("/api/record/getAll", getRecords)
	recordRouter.HandleFunc("/api/record/getFieldValUrl", getFieldValUrlAPI)

	http.Handle("/api/record/", recordRouter)
}

func newRecord(w http.ResponseWriter, r *http.Request) {

	params := NewRecordParams{}
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	newRecordRef, err := NewRecord(params)
	if err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *newRecordRef)
	}

}

type GetRecordParams struct {
	ParentTableID string `json:"parentTableID"`
	RecordID      string `json:"recordID"`
}

func getRecord(w http.ResponseWriter, r *http.Request) {

	var params GetRecordParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if record, getErr := GetRecord(params.ParentTableID, params.RecordID); getErr != nil {
		api.WriteErrorResponse(w, getErr)
		return
	} else {
		api.WriteJSONResponse(w, record)
	}

}

func getRecords(w http.ResponseWriter, r *http.Request) {

	// TODO - Once sorting and filtering is implemented, the request
	// will need to include parameters for the sort and filter parameters to use.

	var params GetRecordsParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if recordRefs, err := GetRecords(params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, recordRefs)
	}

}

func getFieldValUrlAPI(w http.ResponseWriter, r *http.Request) {

	// TODO - Once sorting and filtering is implemented, the request
	// will need to include parameters for the sort and filter parameters to use.

	var params GetFieldValUrlParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if urlResponse, err := getFieldValUrl(params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, urlResponse)
	}

}
