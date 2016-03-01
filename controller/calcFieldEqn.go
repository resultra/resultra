package controller

import (
	"appengine"
	"net/http"
	"resultra/datasheet/datamodel"
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
	if err := decodeJSONRequest(r, &calcFieldValidationParams); err != nil {
		WriteErrorResponse(w, err)
		return
	}

	appEngCntxt := appengine.NewContext(r)
	if validateErr := datamodel.ValidateCalcFieldEqn(appEngCntxt, calcFieldValidationParams.EqnText); validateErr != nil {
		writeJSONResponse(w, CalcFieldValidationResponse{false, validateErr.Error()})
	} else {
		writeJSONResponse(w, CalcFieldValidationResponse{true, ""})
	}

}
