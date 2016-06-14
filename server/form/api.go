package form

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/generic/api"
)

func init() {

	formRouter := mux.NewRouter()

	formRouter.HandleFunc("/api/frm/getFormInfo", getFormInfoAPI)
	formRouter.HandleFunc("/api/frm/new", newFormAPI)
	formRouter.HandleFunc("/api/frm/get", getFormAPI)

	http.Handle("/api/frm/", formRouter)
}

func newFormAPI(w http.ResponseWriter, r *http.Request) {

	var params NewFormParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if formRef, err := newForm(params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *formRef)
	}

}

func getFormAPI(w http.ResponseWriter, r *http.Request) {

	var params GetFormParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if theForm, err := GetForm(params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *theForm)
	}

}

func getFormInfoAPI(w http.ResponseWriter, r *http.Request) {

	var params GetFormInfoParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if formInfo, err := getFormInfo(params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *formInfo)
	}

}
