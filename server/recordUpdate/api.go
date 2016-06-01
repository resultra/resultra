package recordUpdate

import (
	"appengine"
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/generic/api"
)

func init() {
	recordUpdateRouter := mux.NewRouter()

	recordUpdateRouter.HandleFunc("/api/recordUpdate/setBoolFieldValue", setBoolFieldValue)
	recordUpdateRouter.HandleFunc("/api/recordUpdate/setNumberFieldValue", setNumberFieldValue)
	recordUpdateRouter.HandleFunc("/api/recordUpdate/setTextFieldValue", setTextFieldValue)
	recordUpdateRouter.HandleFunc("/api/recordUpdate/setTimeFieldValue", setTimeFieldValue)
	recordUpdateRouter.HandleFunc("/api/recordUpdate/setLongTextFieldValue", setLongTextFieldValue)

	recordUpdateRouter.HandleFunc("/api/recordUpdate/uploadFileToFieldValue", uploadFileAPI)

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

func setLongTextFieldValue(w http.ResponseWriter, r *http.Request) {

	// Reuse same parameter struct as setting text.
	setValParams := SetRecordLongTextValueParams{}
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

func uploadFileAPI(w http.ResponseWriter, req *http.Request) {

	if uploadResponse, uploadErr := uploadFile(req); uploadErr != nil {
		api.WriteErrorResponse(w, uploadErr)
	} else {
		api.WriteJSONResponse(w, *uploadResponse)
	}

}
