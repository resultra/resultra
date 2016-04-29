package calcField

import (
	"appengine"
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/generic/api"
)

func init() {
	calcFieldRouter := mux.NewRouter()

	calcFieldRouter.HandleFunc("/api/calcField/validateFormula", validateFormula)
	calcFieldRouter.HandleFunc("/api/calcField/setFieldFormula", setFieldFormula)
	calcFieldRouter.HandleFunc("/api/calcField/newCalcField", newCalcFieldAPI)

	http.Handle("/api/calcField/", calcFieldRouter)
}

func newCalcFieldAPI(w http.ResponseWriter, r *http.Request) {

	var newCalcFieldParams NewCalcFieldParams
	if err := api.DecodeJSONRequest(r, &newCalcFieldParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	appEngCntxt := appengine.NewContext(r)
	if fieldID, err := newCalcField(appEngCntxt, newCalcFieldParams); err != nil {
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
