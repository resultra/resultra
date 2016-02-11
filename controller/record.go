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
