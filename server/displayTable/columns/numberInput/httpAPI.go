package numberInput

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/generic/api"
)

func init() {
	numberInputRouter := mux.NewRouter()

	numberInputRouter.HandleFunc("/api/tableView/numberInput/new", newNumberInput)
	numberInputRouter.HandleFunc("/api/tableView/numberInput/setValueFormat", setValueFormat)
	numberInputRouter.HandleFunc("/api/tableView/numberInput/setLabelFormat", setLabelFormat)
	numberInputRouter.HandleFunc("/api/tableView/numberInput/setPermissions", setPermissions)
	numberInputRouter.HandleFunc("/api/tableView/numberInput/setShowSpinner", setShowSpinner)
	numberInputRouter.HandleFunc("/api/tableView/numberInput/setSpinnerStepSize", setSpinnerStepSize)

	numberInputRouter.HandleFunc("/api/tableView/numberInput/setValidation", setValidation)
	numberInputRouter.HandleFunc("/api/tableView/numberInput/validateInput", validateInputAPI)

	http.Handle("/api/tableView/numberInput/", numberInputRouter)
}

func newNumberInput(w http.ResponseWriter, r *http.Request) {

	numberInputParams := NewNumberInputParams{}
	if err := api.DecodeJSONRequest(r, &numberInputParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if numberInputRef, err := saveNewNumberInput(numberInputParams); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *numberInputRef)
	}

}

func validateInputAPI(w http.ResponseWriter, r *http.Request) {

	var params NumberInputValidateInputParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	validationResp := validateInput(params)
	api.WriteJSONResponse(w, validationResp)
}

func processNumberInputPropUpdate(w http.ResponseWriter, r *http.Request, propUpdater NumberInputPropUpdater) {
	if numberInputRef, err := updateNumberInputProps(propUpdater); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, numberInputRef)
	}
}

func setValueFormat(w http.ResponseWriter, r *http.Request) {
	var params NumberInputValueFormatParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processNumberInputPropUpdate(w, r, params)
}

func setLabelFormat(w http.ResponseWriter, r *http.Request) {
	var params NumberInputLabelFormatParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processNumberInputPropUpdate(w, r, params)
}

func setPermissions(w http.ResponseWriter, r *http.Request) {
	var params NumberInputPermissionParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processNumberInputPropUpdate(w, r, params)
}

func setShowSpinner(w http.ResponseWriter, r *http.Request) {
	var params ShowValueSpinnerParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processNumberInputPropUpdate(w, r, params)
}

func setSpinnerStepSize(w http.ResponseWriter, r *http.Request) {
	var params ValueSpinnerStepSizeParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processNumberInputPropUpdate(w, r, params)
}

func setValidation(w http.ResponseWriter, r *http.Request) {
	var params NumberInputValidationParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processNumberInputPropUpdate(w, r, params)
}
