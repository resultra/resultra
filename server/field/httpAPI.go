package field

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/generic/api"
	"resultra/datasheet/server/userRole"
)

func init() {

	fieldRouter := mux.NewRouter()

	fieldRouter.HandleFunc("/api/field/new", newField)
	fieldRouter.HandleFunc("/api/field/getListByType", getFieldsByType)

	http.Handle("/api/field/", fieldRouter)
}

func newField(w http.ResponseWriter, r *http.Request) {

	var newFieldParams NewNonCalcFieldParams
	if err := api.DecodeJSONRequest(r, &newFieldParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if verifyErr := userRole.VerifyCurrUserIsDatabaseAdmin(
		r, newFieldParams.ParentDatabaseID); verifyErr != nil {
		api.WriteErrorResponse(w, verifyErr)
		return
	}

	if newField, err := NewNonCalcField(newFieldParams); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *newField)
	}

}

func getFieldsByType(w http.ResponseWriter, r *http.Request) {

	var fieldListParams GetFieldListParams
	if err := api.DecodeJSONRequest(r, &fieldListParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if fieldsByType, err := GetFieldsByType(fieldListParams); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, fieldsByType)
	}

}
