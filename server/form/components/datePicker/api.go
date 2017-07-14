package datePicker

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/generic/api"
)

func init() {
	datePickerRouter := mux.NewRouter()

	datePickerRouter.HandleFunc("/api/frm/datePicker/new", newDatePicker)
	datePickerRouter.HandleFunc("/api/frm/datePicker/resize", resizeDatePicker)
	datePickerRouter.HandleFunc("/api/frm/datePicker/setFormat", setFormat)
	datePickerRouter.HandleFunc("/api/frm/datePicker/setLabelFormat", setLabelFormat)
	datePickerRouter.HandleFunc("/api/frm/datePicker/setVisibility", setVisibility)
	datePickerRouter.HandleFunc("/api/frm/datePicker/setPermissions", setPermissions)
	datePickerRouter.HandleFunc("/api/frm/datePicker/setClearValueSupported", setClearValueSupported)
	datePickerRouter.HandleFunc("/api/frm/datePicker/setValidation", setValidation)
	datePickerRouter.HandleFunc("/api/frm/datePicker/validateInput", validateInputAPI)

	http.Handle("/api/frm/datePicker/", datePickerRouter)
}

func newDatePicker(w http.ResponseWriter, r *http.Request) {

	params := NewDatePickerParams{}
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if checkBoxRef, err := saveNewDatePicker(params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *checkBoxRef)
	}

}

func validateInputAPI(w http.ResponseWriter, r *http.Request) {

	params := DatePickerValidateInputParams{}
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	validationResult := validateInput(params)
	api.WriteJSONResponse(w, validationResult)

}

func processDatePickerPropUpdate(w http.ResponseWriter, r *http.Request, propUpdater DatePickerPropUpdater) {
	if checkBoxRef, err := updateDatePickerProps(propUpdater); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, checkBoxRef)
	}
}

func resizeDatePicker(w http.ResponseWriter, r *http.Request) {
	var resizeParams DatePickerResizeParams
	if err := api.DecodeJSONRequest(r, &resizeParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processDatePickerPropUpdate(w, r, resizeParams)
}

func setFormat(w http.ResponseWriter, r *http.Request) {
	var params DatePickerFormatParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processDatePickerPropUpdate(w, r, params)
}

func setLabelFormat(w http.ResponseWriter, r *http.Request) {
	var params DatePickerLabelFormatParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processDatePickerPropUpdate(w, r, params)
}

func setVisibility(w http.ResponseWriter, r *http.Request) {
	var params DatePickerVisibilityParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processDatePickerPropUpdate(w, r, params)
}

func setPermissions(w http.ResponseWriter, r *http.Request) {
	var params DatePickerPermissionParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processDatePickerPropUpdate(w, r, params)
}

func setValidation(w http.ResponseWriter, r *http.Request) {
	var params DatePickerValidationParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processDatePickerPropUpdate(w, r, params)
}

func setClearValueSupported(w http.ResponseWriter, r *http.Request) {
	var params DatePickerClearValueSupportedParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processDatePickerPropUpdate(w, r, params)
}
