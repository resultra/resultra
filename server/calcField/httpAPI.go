package calcField

import (
	"appengine"
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/generic/api"
)

func RegisterHTTPHandlers(apiRouter *mux.Router) {

	apiRouter.HandleFunc("/api/validateCalcFieldEqn", validateCalcFieldEqn)
	apiRouter.HandleFunc("/api/newCalcField", newCalcField)

}

func init() {
	calcFieldRouter := mux.NewRouter()

	calcFieldRouter.HandleFunc("/api/calcField/validateFormula", validateFormula)
	calcFieldRouter.HandleFunc("/api/calcField/setFieldFormula", setFieldFormula)

	http.Handle("/api/calcField/", calcFieldRouter)
}

type CalcFieldValidationParams struct {
	EqnText    string `json:"eqnText"`
	IsNewField bool   `json:"isNewField"`
}

type CalcFieldValidationResponse struct {
	IsValidEqn bool   `json:"isValidEqn"`
	ErrorMsg   string `json:"errorMsg"`
}

func validateCalcFieldEqn(w http.ResponseWriter, r *http.Request) {

	var calcFieldValidationParams CalcFieldValidationParams
	if err := api.DecodeJSONRequest(r, &calcFieldValidationParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	appEngCntxt := appengine.NewContext(r)
	if validateErr := validateCalcFieldEqnText(appEngCntxt, calcFieldValidationParams.EqnText); validateErr != nil {
		api.WriteJSONResponse(w, CalcFieldValidationResponse{false, validateErr.Error()})
	} else {
		api.WriteJSONResponse(w, CalcFieldValidationResponse{true, ""})
	}

}

func newCalcField(w http.ResponseWriter, r *http.Request) {

	var newCalcField NewCalcFieldParams
	if err := api.DecodeJSONRequest(r, &newCalcField); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	appEngCntxt := appengine.NewContext(r)
	if fieldID, err := NewCalcField(appEngCntxt, newCalcField); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, api.JSONParams{"fieldID": fieldID})
	}

}

func validateFormula(w http.ResponseWriter, r *http.Request) {

	var validationParams ValidateFormulaParams
	if err := api.DecodeJSONRequest(r, &validationParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	appEngCntxt := appengine.NewContext(r)
	validationResponse := validateFormulaText(appEngCntxt, validationParams)
	api.WriteJSONResponse(w, *validationResponse)
}

func setFieldFormula(w http.ResponseWriter, r *http.Request) {

	var setFormulaParams SetFormulaParams
	if err := api.DecodeJSONRequest(r, &setFormulaParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	appEngCntxt := appengine.NewContext(r)
	if updatedFieldRef, err := field.UpdateFieldProps(appEngCntxt, setFormulaParams); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, updatedFieldRef)
	}

}
