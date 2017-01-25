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

	if recordRefs, err := GetFilteredSortedRecords(params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, recordRefs)
	}

}
