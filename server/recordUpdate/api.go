package recordUpdate

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/generic/api"
	"resultra/datasheet/server/record"
)

type DummyStructForInclude struct {
	Val int64
}

func init() {
	recordUpdateRouter := mux.NewRouter()

	recordUpdateRouter.HandleFunc("/api/recordUpdate/newRecord", newRecordAPI)

	recordUpdateRouter.HandleFunc("/api/recordUpdate/setBoolFieldValue", setBoolFieldValue)
	recordUpdateRouter.HandleFunc("/api/recordUpdate/setNumberFieldValue", setNumberFieldValue)
	recordUpdateRouter.HandleFunc("/api/recordUpdate/setTextFieldValue", setTextFieldValue)
	recordUpdateRouter.HandleFunc("/api/recordUpdate/setTimeFieldValue", setTimeFieldValue)
	recordUpdateRouter.HandleFunc("/api/recordUpdate/setLongTextFieldValue", setLongTextFieldValue)

	recordUpdateRouter.HandleFunc("/api/recordUpdate/uploadFileToFieldValue", uploadFileAPI)

	http.Handle("/api/recordUpdate/", recordUpdateRouter)
}

func newRecordAPI(w http.ResponseWriter, r *http.Request) {

	params := record.NewRecordParams{}
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	newRecordRef, err := newRecord(params)
	if err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *newRecordRef)
	}

}

func setTextFieldValue(w http.ResponseWriter, r *http.Request) {

	setValParams := record.SetRecordTextValueParams{}
	if err := api.DecodeJSONRequest(r, &setValParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	updatedRecordRef, setErr := updateRecordValue(setValParams)
	if setErr != nil {
		api.WriteErrorResponse(w, setErr)
		return
	} else {
		api.WriteJSONResponse(w, updatedRecordRef)
	}

}

func setLongTextFieldValue(w http.ResponseWriter, r *http.Request) {

	// Reuse same parameter struct as setting text.
	setValParams := record.SetRecordLongTextValueParams{}
	if err := api.DecodeJSONRequest(r, &setValParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	updatedRecordRef, setErr := updateRecordValue(setValParams)
	if setErr != nil {
		api.WriteErrorResponse(w, setErr)
		return
	} else {
		api.WriteJSONResponse(w, updatedRecordRef)
	}

}

func setNumberFieldValue(w http.ResponseWriter, r *http.Request) {

	setValParams := record.SetRecordNumberValueParams{}
	if err := api.DecodeJSONRequest(r, &setValParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	updatedRecordRef, setErr := updateRecordValue(setValParams)
	if setErr != nil {
		api.WriteErrorResponse(w, setErr)
		return
	} else {
		api.WriteJSONResponse(w, updatedRecordRef)
	}

}

func setBoolFieldValue(w http.ResponseWriter, r *http.Request) {

	setValParams := record.SetRecordBoolValueParams{}
	if err := api.DecodeJSONRequest(r, &setValParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	updatedRecordRef, setErr := updateRecordValue(setValParams)
	if setErr != nil {
		api.WriteErrorResponse(w, setErr)
		return
	} else {
		api.WriteJSONResponse(w, updatedRecordRef)
	}

}

func setTimeFieldValue(w http.ResponseWriter, r *http.Request) {

	setValParams := record.SetRecordTimeValueParams{}
	if err := api.DecodeJSONRequest(r, &setValParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	updatedRecordRef, setErr := updateRecordValue(setValParams)
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
