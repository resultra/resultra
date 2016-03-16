package controller

import (
	"appengine"
	"net/http"
	"resultra/datasheet/datamodel"
)

func newRecord(w http.ResponseWriter, r *http.Request) {

	appEngCntxt := appengine.NewContext(r)
	newRecordRef, err := datamodel.NewRecord(appEngCntxt)
	if err != nil {
		WriteErrorResponse(w, err)
	} else {
		writeJSONResponse(w, *newRecordRef)
	}

}

func setTextFieldValue(w http.ResponseWriter, r *http.Request) {

	setValParams := datamodel.SetRecordTextValueParams{}
	if err := decodeJSONRequest(r, &setValParams); err != nil {
		WriteErrorResponse(w, err)
		return
	}

	appEngCntxt := appengine.NewContext(r)
	updatedRecordRef, setErr := datamodel.SetRecordTextValue(appEngCntxt, setValParams)
	if setErr != nil {
		WriteErrorResponse(w, setErr)
		return
	} else {
		writeJSONResponse(w, updatedRecordRef)
	}

}

func setNumberFieldValue(w http.ResponseWriter, r *http.Request) {

	setValParams := datamodel.SetRecordNumberValueParams{}
	if err := decodeJSONRequest(r, &setValParams); err != nil {
		WriteErrorResponse(w, err)
		return
	}

	appEngCntxt := appengine.NewContext(r)
	updatedRecordRef, setErr := datamodel.SetRecordNumberValue(appEngCntxt, setValParams)
	if setErr != nil {
		WriteErrorResponse(w, setErr)
		return
	} else {
		writeJSONResponse(w, updatedRecordRef)
	}

}

func getRecord(w http.ResponseWriter, r *http.Request) {

	params := datamodel.GetRecordParams{}
	if err := decodeJSONRequest(r, &params); err != nil {
		WriteErrorResponse(w, err)
		return
	}

	appEngCntxt := appengine.NewContext(r)

	if recordRef, getErr := datamodel.GetRecord(appEngCntxt, params); getErr != nil {
		WriteErrorResponse(w, getErr)
		return
	} else {
		writeJSONResponse(w, recordRef)
	}

}

func getRecords(w http.ResponseWriter, r *http.Request) {

	// TODO - Once sorting and filtering is implemented, the request
	// will need to include parameters for the sort and filter parameters to use.

	appEngCntxt := appengine.NewContext(r)
	if recordRefs, err := datamodel.GetRecords(appEngCntxt); err != nil {
		WriteErrorResponse(w, err)
	} else {
		writeJSONResponse(w, recordRefs)
	}

}

func getFilteredRecords(w http.ResponseWriter, r *http.Request) {

	// TODO - Once filtering is implemented on a per form/dashboard basis,
	// pass in the parent filter.

	appEngCntxt := appengine.NewContext(r)
	if recordRefs, err := datamodel.GetFilteredRecords(appEngCntxt); err != nil {
		WriteErrorResponse(w, err)
	} else {
		writeJSONResponse(w, recordRefs)
	}

}
