package controller

import (
	"appengine"
	"net/http"
	"resultra/datasheet/datamodel"
)

func newRecord(w http.ResponseWriter, r *http.Request) {

	var newField datamodel.Field
	if err := decodeJSONRequest(r, &newField); err != nil {
		writeErrorResponse(w, err)
		return
	}

	emptyRecord := datamodel.Record{}
	appEngCntxt := appengine.NewContext(r)
	if fieldID, err := datamodel.SaveNewRecord(appEngCntxt, emptyRecord); err != nil {
		writeErrorResponse(w, err)
	} else {
		writeJSONResponse(w, JSONParams{"recordID": fieldID})
	}

}

func setRecordFieldValue(w http.ResponseWriter, r *http.Request) {

	setValParams := datamodel.SetRecordValueParams{}
	if err := decodeJSONRequest(r, &setValParams); err != nil {
		writeErrorResponse(w, err)
		return
	}

	appEngCntxt := appengine.NewContext(r)
	if setErr := datamodel.SetRecordValue(appEngCntxt, setValParams); setErr != nil {
		writeErrorResponse(w, setErr)
		return
	} else {
		writeJSONResponse(w, JSONParams{})
	}

}

func getRecord(w http.ResponseWriter, r *http.Request) {

	params := datamodel.GetRecordParams{}
	if err := decodeJSONRequest(r, &params); err != nil {
		writeErrorResponse(w, err)
		return
	}

	appEngCntxt := appengine.NewContext(r)

	if recordRef, getErr := datamodel.GetRecord(appEngCntxt, params); getErr != nil {
		writeErrorResponse(w, getErr)
		return
	} else {
		writeJSONResponse(w, recordRef)
	}

}
