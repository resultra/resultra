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
