package recordUpdate

import (
	"appengine"
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/generic/api"
)

import ()

func RegisterHTTPHandlers(apiRouter *mux.Router) {
	apiRouter.HandleFunc("/api/setTextFieldValue", setTextFieldValue)
	apiRouter.HandleFunc("/api/setNumberFieldValue", setNumberFieldValue)
	apiRouter.HandleFunc("/api/setBoolFieldValue", setBoolFieldValue)

}

func init() {
	recordUpdateRouter := mux.NewRouter()

	recordUpdateRouter.HandleFunc("/api/recordUpdate/setTimeFieldValue", setTimeFieldValue)

	http.Handle("/api/recordUpdate/", recordUpdateRouter)
}

func setTextFieldValue(w http.ResponseWriter, r *http.Request) {

	setValParams := SetRecordTextValueParams{}
	if err := api.DecodeJSONRequest(r, &setValParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	appEngCntxt := appengine.NewContext(r)
	updatedRecordRef, setErr := UpdateRecordValue(appEngCntxt, setValParams)
	if setErr != nil {
		api.WriteErrorResponse(w, setErr)
		return
	} else {
		api.WriteJSONResponse(w, updatedRecordRef)
	}

}

func setNumberFieldValue(w http.ResponseWriter, r *http.Request) {

	setValParams := SetRecordNumberValueParams{}
	if err := api.DecodeJSONRequest(r, &setValParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	appEngCntxt := appengine.NewContext(r)
	updatedRecordRef, setErr := UpdateRecordValue(appEngCntxt, setValParams)
	if setErr != nil {
		api.WriteErrorResponse(w, setErr)
		return
	} else {
		api.WriteJSONResponse(w, updatedRecordRef)
	}

}

func setBoolFieldValue(w http.ResponseWriter, r *http.Request) {

	setValParams := SetRecordBoolValueParams{}
	if err := api.DecodeJSONRequest(r, &setValParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	appEngCntxt := appengine.NewContext(r)
	updatedRecordRef, setErr := UpdateRecordValue(appEngCntxt, setValParams)
	if setErr != nil {
		api.WriteErrorResponse(w, setErr)
		return
	} else {
		api.WriteJSONResponse(w, updatedRecordRef)
	}

}

func setTimeFieldValue(w http.ResponseWriter, r *http.Request) {

	setValParams := SetRecordTimeValueParams{}
	if err := api.DecodeJSONRequest(r, &setValParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	appEngCntxt := appengine.NewContext(r)
	updatedRecordRef, setErr := UpdateRecordValue(appEngCntxt, setValParams)
	if setErr != nil {
		api.WriteErrorResponse(w, setErr)
		return
	} else {
		api.WriteJSONResponse(w, updatedRecordRef)
	}

}
