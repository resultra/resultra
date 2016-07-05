package record

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/generic/api"
	"resultra/datasheet/server/generic/cloudStorageWrapper"
)

func init() {
	recordRouter := mux.NewRouter()

	recordRouter.HandleFunc("/api/record/getAll", getRecords)

	recordRouter.HandleFunc("/api/record/getFieldValUrl", getFieldValUrlAPI)
	recordRouter.HandleFunc("/api/record/getFile/{fileName}", getRecordFileAPI)

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

func getFieldValUrlAPI(w http.ResponseWriter, r *http.Request) {

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

func getRecordFileAPI(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	fileName := vars["fileName"]

	http.ServeFile(w, r, cloudStorageWrapper.LocalAttachmentFileUploadDir+fileName)

}
