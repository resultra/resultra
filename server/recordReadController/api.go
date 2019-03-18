package recordReadController

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"resultra/tracker/server/common/databaseWrapper"
	"resultra/tracker/server/common/userAuth"
	"resultra/tracker/server/generic/api"
)

type DummyStructForInclude struct {
	Val int64
}

func init() {
	recordReadRouter := mux.NewRouter()

	recordReadRouter.HandleFunc("/api/recordRead/getFilteredSortedRecordValues", getFilteredSortedRecordsAPI)
	recordReadRouter.HandleFunc("/api/recordRead/testRecordIsFiltered", testRecordIsFilteredAPI)

	recordReadRouter.HandleFunc("/api/recordRead/getFilteredRecordCount", getFilteredRecordCountAPI)

	recordReadRouter.HandleFunc("/api/recordRead/getRecordValueResults", getRecordValueResultAPI)

	http.Handle("/api/recordRead/", recordReadRouter)
}

func getFilteredSortedRecordsAPI(w http.ResponseWriter, r *http.Request) {

	params := NewDefaultGetFilteredSortedRecordsParams()
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	currUserID, userErr := userAuth.GetCurrentUserID(r)
	if userErr != nil {
		api.WriteJSONResponse(w, fmt.Errorf("Can't verify user authentication"))
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	// By default, recompute/refresh the record values before returning.
	if recordRefs, err := GetFilteredSortedRecords(trackerDBHandle, currUserID, params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, recordRefs)
	}

}

func testRecordIsFilteredAPI(w http.ResponseWriter, r *http.Request) {

	params := NewDefaultTestRecordIsFilteredParams()
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	currUserID, userErr := userAuth.GetCurrentUserID(r)
	if userErr != nil {
		api.WriteJSONResponse(w, fmt.Errorf("Can't verify user authentication"))
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	// By default, recompute/refresh the record values before returning.
	if recordIsFiltered, err := testRecordIsFiltered(trackerDBHandle, currUserID, params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, recordIsFiltered)
	}

}

func getRecordValueResultAPI(w http.ResponseWriter, r *http.Request) {

	params := GetRecordValResultParams{}
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	currUserID, userErr := userAuth.GetCurrentUserID(r)
	if userErr != nil {
		api.WriteJSONResponse(w, fmt.Errorf("Can't verify user authentication"))
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	recordValResults, err := getRecordValueResults(trackerDBHandle, currUserID, params)
	if err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *recordValResults)
	}

}

func getFilteredRecordCountAPI(w http.ResponseWriter, r *http.Request) {

	params := GetFilteredRecordCountParams{}
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	currUserID, userErr := userAuth.GetCurrentUserID(r)
	if userErr != nil {
		api.WriteJSONResponse(w, fmt.Errorf("Can't verify user authentication"))
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	countResults, err := getFilteredRecordCount(trackerDBHandle, currUserID, params)
	if err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, countResults)
	}

}
