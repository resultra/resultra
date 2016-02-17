package controller

import (
	"appengine"
	"net/http"
	"resultra/datasheet/datamodel"
)

func newRecord(w http.ResponseWriter, r *http.Request) {

	var newField datamodel.Field
	if err := decodeJSONRequest(r, &newField); err != nil {
		WriteErrorResponse(w, err)
		return
	}

	emptyRecord := datamodel.Record{}
	appEngCntxt := appengine.NewContext(r)
	if fieldID, err := datamodel.SaveNewRecord(appEngCntxt, emptyRecord); err != nil {
		WriteErrorResponse(w, err)
	} else {
		writeJSONResponse(w, JSONParams{"recordID": fieldID})
	}

}

func setRecordFieldValue(w http.ResponseWriter, r *http.Request) {

	setValParams := datamodel.SetRecordValueParams{}
	if err := decodeJSONRequest(r, &setValParams); err != nil {
		WriteErrorResponse(w, err)
		return
	}

	appEngCntxt := appengine.NewContext(r)
	updatedRecordRef, setErr := datamodel.SetRecordValue(appEngCntxt, setValParams)
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
