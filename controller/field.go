package controller

import (
	"appengine"
	"net/http"
	"resultra/datasheet/datamodel"
)

func newField(w http.ResponseWriter, r *http.Request) {

	var newField datamodel.Field
	if err := decodeJSONRequest(r, &newField); err != nil {
		WriteErrorResponse(w, err)
		return
	}

	appEngCntxt := appengine.NewContext(r)
	if fieldID, err := datamodel.NewField(appEngCntxt, newField); err != nil {
		WriteErrorResponse(w, err)
	} else {
		writeJSONResponse(w, JSONParams{"fieldID": fieldID})
	}

}

func getFieldsByType(w http.ResponseWriter, r *http.Request) {

	appEngCntxt := appengine.NewContext(r)
	if fieldsByType, err := datamodel.GetFieldsByType(appEngCntxt); err != nil {
		WriteErrorResponse(w, err)
	} else {
		writeJSONResponse(w, fieldsByType)
	}

}
