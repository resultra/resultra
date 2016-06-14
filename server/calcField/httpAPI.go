package calcField

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/generic/api"
)

func init() {
	calcFieldRouter := mux.NewRouter()

	calcFieldRouter.HandleFunc("/api/calcField/validateFormula", validateFormula)
	calcFieldRouter.HandleFunc("/api/calcField/setFieldFormula", setFieldFormula)
	calcFieldRouter.HandleFunc("/api/calcField/new", newCalcFieldAPI)
	calcFieldRouter.HandleFunc("/api/calcField/getRawFormulaText", getRawFormulaTextAPI)

	http.Handle("/api/calcField/", calcFieldRouter)
}

func newCalcFieldAPI(w http.ResponseWriter, r *http.Request) {

	var newCalcFieldParams NewCalcFieldParams
	if err := api.DecodeJSONRequest(r, &newCalcFieldParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if fieldID, err := newCalcField(newCalcFieldParams); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, api.JSONParams{"fieldID": fieldID})
	}

}

func getRawFormulaTextAPI(w http.ResponseWriter, r *http.Request) {

	var params GetRawFormulaParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if rawFormulaText, err := getRawFormulaText(params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, rawFormulaText)
	}
}

func validateFormula(w http.ResponseWriter, r *http.Request) {

	var validationParams ValidateFormulaParams
	if err := api.DecodeJSONRequest(r, &validationParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	validationResponse := validateFormulaText(validationParams)
	api.WriteJSONResponse(w, *validationResponse)
}

func setFieldFormula(w http.ResponseWriter, r *http.Request) {

	var setFormulaParams SetFormulaParams
	if err := api.DecodeJSONRequest(r, &setFormulaParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if updatedFieldRef, err := field.UpdateFieldProps(setFormulaParams); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, updatedFieldRef)
	}

}
