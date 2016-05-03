package field

import (
	"appengine"
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/generic/api"
)

func init() {

	fieldRouter := mux.NewRouter()

	fieldRouter.HandleFunc("/api/field/new", newField)
	fieldRouter.HandleFunc("/api/field/getListByType", getFieldsByType)

	http.Handle("/api/field/", fieldRouter)
}

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

	var fieldListParams GetFieldListParams
	if err := api.DecodeJSONRequest(r, &fieldListParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	appEngCntxt := appengine.NewContext(r)
	if fieldsByType, err := GetFieldsByType(appEngCntxt, fieldListParams); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, fieldsByType)
	}

}
