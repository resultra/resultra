package numberInput

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/generic/api"
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

	http.Handle("/api/frm/numberInput/", numberInputRouter)
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

func processNumberInputPropUpdate(w http.ResponseWriter, r *http.Request, propUpdater NumberInputPropUpdater) {
	if numberInputRef, err := updateNumberInputProps(propUpdater); err != nil {
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
