package controller

import (
	"appengine"
	"net/http"
	"resultra/datasheet/datamodel"
)

func newField(w http.ResponseWriter, r *http.Request) {

	var newFieldParams datamodel.NewFieldParams
	if err := decodeJSONRequest(r, &newFieldParams); err != nil {
		WriteErrorResponse(w, err)
		return
	}

	appEngCntxt := appengine.NewContext(r)
	if fieldID, err := datamodel.NewField(appEngCntxt, newFieldParams); err != nil {
		WriteErrorResponse(w, err)
	} else {
		writeJSONResponse(w, JSONParams{"fieldID": fieldID})
	}

}

func newCalcField(w http.ResponseWriter, r *http.Request) {

	var newCalcField datamodel.NewCalcFieldParams
	if err := decodeJSONRequest(r, &newCalcField); err != nil {
		WriteErrorResponse(w, err)
		return
	}

	appEngCntxt := appengine.NewContext(r)
	if fieldID, err := datamodel.NewCalcField(appEngCntxt, newCalcField); err != nil {
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
