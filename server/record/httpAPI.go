package record

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/generic/api"
	"resultra/datasheet/server/generic/uniqueID"
)

func init() {
	recordRouter := mux.NewRouter()

	recordRouter.HandleFunc("/api/record/getAll", getRecords)

	recordRouter.HandleFunc("/api/record/setDraftStatus", setDraftStatusAPI)

	recordRouter.HandleFunc("/api/record/allocateChangeSetID", allocateChangeSetIDAPI)

	recordRouter.HandleFunc("/api/record/getFieldValChangeInfo", getFieldValChangeInfoAPI)

	http.Handle("/api/record/", recordRouter)
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

func getFieldValChangeInfoAPI(w http.ResponseWriter, r *http.Request) {

	var params GetFieldValChangeInfoParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if valChangeInfo, err := getFieldValChangeInfo(r, params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, valChangeInfo)
	}

}

func setDraftStatusAPI(w http.ResponseWriter, r *http.Request) {

	var params SetDraftStatusParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if err := setDraftStatus(params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, true)
	}
}

type changeSetIDAllocationResponse struct {
	ChangeSetID string `json:"changeSetID"`
}

func allocateChangeSetIDAPI(w http.ResponseWriter, r *http.Request) {

	changeSetID := uniqueID.GenerateSnowflakeID()

	response := changeSetIDAllocationResponse{changeSetID}

	api.WriteJSONResponse(w, response)
}
