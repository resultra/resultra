package recordValue

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/generic/api"
)

type DummyStructForInclude struct {
	Val int64
}

func init() {
	recordValueRouter := mux.NewRouter()

	recordValueRouter.HandleFunc("/api/recordValue/getRecordValueResults", getRecordValueResultAPI)
	recordValueRouter.HandleFunc("/api/recordValue/getFilteredSortedRecordValues", getFilteredSortedRecordsAPI)

	http.Handle("/api/recordValue/", recordValueRouter)
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

func getFilteredSortedRecordsAPI(w http.ResponseWriter, r *http.Request) {

	// TODO - Once filtering is implemented on a per form/dashboard basis,
	// pass in the parent filter.
	var params GetFilteredSortedRecordsParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if recordRefs, err := GetFilteredSortedRecords(params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, recordRefs)
	}

}
