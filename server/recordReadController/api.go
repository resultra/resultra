package recordReadController

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/generic/api"
)

type DummyStructForInclude struct {
	Val int64
}

func init() {
	recordReadRouter := mux.NewRouter()

	recordReadRouter.HandleFunc("/api/recordRead/getFilteredSortedRecordValues", getFilteredSortedRecordsAPI)

	recordReadRouter.HandleFunc("/api/recordRead/getRecordValueResults", getRecordValueResultAPI)

	http.Handle("/api/recordRead/", recordReadRouter)
}

func getFilteredSortedRecordsAPI(w http.ResponseWriter, r *http.Request) {

	// TODO - Once filtering is implemented on a per form/dashboard basis,
	// pass in the parent filter.
	params := NewDefaultGetFilteredSortedRecordsParams()
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	// By default, recompute/refresh the record values before returning.
	if recordRefs, err := GetFilteredSortedRecords(params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, recordRefs)
	}

}

func getRecordValueResultAPI(w http.ResponseWriter, r *http.Request) {

	params := GetRecordValResultParams{}
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	recordValResults, err := getRecordValueResults(params)
	if err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *recordValResults)
	}

}
