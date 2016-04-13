package field

import (
	"appengine"
	"net/http"
	"resultra/datasheet/server/common/api"
)

func newField(w http.ResponseWriter, r *http.Request) {

	var newFieldParams NewFieldParams
	if err := api.DecodeJSONRequest(r, &newFieldParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	appEngCntxt := appengine.NewContext(r)
	if fieldID, err := NewField(appEngCntxt, newFieldParams); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, api.JSONParams{"fieldID": fieldID})
	}

}

func getFieldsByType(w http.ResponseWriter, r *http.Request) {

	appEngCntxt := appengine.NewContext(r)
	if fieldsByType, err := GetFieldsByType(appEngCntxt); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, fieldsByType)
	}

}
