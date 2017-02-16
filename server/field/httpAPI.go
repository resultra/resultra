package field

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/generic/api"
	"resultra/datasheet/server/userRole"
)

func init() {

	fieldRouter := mux.NewRouter()

	fieldRouter.HandleFunc("/api/field/new", newField)
	fieldRouter.HandleFunc("/api/field/getListByType", getFieldsByType)
	fieldRouter.HandleFunc("/api/field/get", getField)

	fieldRouter.HandleFunc("/api/field/validateExistingFieldName", validateExistingFieldNameAPI)
	fieldRouter.HandleFunc("/api/field/validateNewFieldName", validateNewFieldNameAPI)

	fieldRouter.HandleFunc("/api/field/setName", setNameAPI)

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

type GetFieldParams struct {
	FieldID string `json:"fieldID"`
}

func getField(w http.ResponseWriter, r *http.Request) {

	var params GetFieldParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if fieldInfo, err := GetField(params.FieldID); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, fieldInfo)
	}

}

func validateExistingFieldNameAPI(w http.ResponseWriter, r *http.Request) {

	fieldName := r.FormValue("fieldName")
	fieldID := r.FormValue("fieldID")

	if err := validateExistingFieldName(fieldID, fieldName); err != nil {
		api.WriteJSONResponse(w, fmt.Sprintf("%v", err))
		return
	}

	response := true
	api.WriteJSONResponse(w, response)

}

func validateNewFieldNameAPI(w http.ResponseWriter, r *http.Request) {

	fieldName := r.FormValue("fieldName")
	databaseID := r.FormValue("databaseID")

	if err := validateNewFieldName(databaseID, fieldName); err != nil {
		api.WriteJSONResponse(w, fmt.Sprintf("%v", err))
		return
	}

	response := true
	api.WriteJSONResponse(w, response)

}

func setNameAPI(w http.ResponseWriter, r *http.Request) {
	var params SetFieldNameParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processFieldPropUpdate(w, r, params)
}
