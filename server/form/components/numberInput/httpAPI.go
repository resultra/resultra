package numberInput

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/tracker/server/common/databaseWrapper"
	"resultra/tracker/server/generic/api"
)

func init() {
	numberInputRouter := mux.NewRouter()

	numberInputRouter.HandleFunc("/api/frm/numberInput/new", newNumberInput)
	numberInputRouter.HandleFunc("/api/frm/numberInput/resize", resizeNumberInput)
	numberInputRouter.HandleFunc("/api/frm/numberInput/setValueFormat", setValueFormat)
	numberInputRouter.HandleFunc("/api/frm/numberInput/setLabelFormat", setLabelFormat)
	numberInputRouter.HandleFunc("/api/frm/numberInput/setVisibility", setVisibility)
	numberInputRouter.HandleFunc("/api/frm/numberInput/setPermissions", setPermissions)
	numberInputRouter.HandleFunc("/api/frm/numberInput/setShowSpinner", setShowSpinner)
	numberInputRouter.HandleFunc("/api/frm/numberInput/setSpinnerStepSize", setSpinnerStepSize)
	numberInputRouter.HandleFunc("/api/frm/numberInput/setClearValueSupported", setClearValueSupported)
	numberInputRouter.HandleFunc("/api/frm/numberInput/setHelpPopupMsg", setHelpPopupMsg)
	numberInputRouter.HandleFunc("/api/frm/numberInput/setConditionalFormats", setConditionalFormats)

	numberInputRouter.HandleFunc("/api/frm/numberInput/setValidation", setValidation)
	numberInputRouter.HandleFunc("/api/frm/numberInput/validateInput", validateInputAPI)

	http.Handle("/api/frm/numberInput/", numberInputRouter)
}

func newNumberInput(w http.ResponseWriter, r *http.Request) {

	numberInputParams := NewNumberInputParams{}
	if err := api.DecodeJSONRequest(r, &numberInputParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if numberInputRef, err := saveNewNumberInput(trackerDBHandle, numberInputParams); err != nil {
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
	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	validationResp := validateInput(trackerDBHandle, params)
	api.WriteJSONResponse(w, validationResp)
}

func processNumberInputPropUpdate(w http.ResponseWriter, r *http.Request, propUpdater NumberInputPropUpdater) {
	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}
	if numberInputRef, err := updateNumberInputProps(trackerDBHandle, propUpdater); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, numberInputRef)
	}
}

func resizeNumberInput(w http.ResponseWriter, r *http.Request) {
	var resizeParams NumberInputResizeParams
	if err := api.DecodeJSONRequest(r, &resizeParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processNumberInputPropUpdate(w, r, resizeParams)
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

func setVisibility(w http.ResponseWriter, r *http.Request) {
	var params NumberInputVisibilityParams
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

func setClearValueSupported(w http.ResponseWriter, r *http.Request) {
	var params NumberInputClearValueSupportedParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processNumberInputPropUpdate(w, r, params)
}

func setHelpPopupMsg(w http.ResponseWriter, r *http.Request) {
	var params HelpPopupMsgParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processNumberInputPropUpdate(w, r, params)
}

func setConditionalFormats(w http.ResponseWriter, r *http.Request) {
	var params ConditionalFormatParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processNumberInputPropUpdate(w, r, params)
}
