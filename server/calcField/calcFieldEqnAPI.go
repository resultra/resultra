package calcField

import (
	"appengine"
	"net/http"
	"resultra/datasheet/server/generic/api"
)

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
