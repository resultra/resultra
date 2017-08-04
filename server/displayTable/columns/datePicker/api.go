package datePicker

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/generic/api"
)

func init() {
	datePickerRouter := mux.NewRouter()

	datePickerRouter.HandleFunc("/api/tableView/datePicker/new", newDatePicker)

	datePickerRouter.HandleFunc("/api/tableView/datePicker/get", getDatePickerAPI)

	datePickerRouter.HandleFunc("/api/tableView/datePicker/setFormat", setFormat)
	datePickerRouter.HandleFunc("/api/tableView/datePicker/setLabelFormat", setLabelFormat)
	datePickerRouter.HandleFunc("/api/tableView/datePicker/setPermissions", setPermissions)
	datePickerRouter.HandleFunc("/api/tableView/datePicker/setValidation", setValidation)
	datePickerRouter.HandleFunc("/api/tableView/datePicker/setClearValueSupported", setClearValueSupported)
	datePickerRouter.HandleFunc("/api/tableView/datePicker/setHelpPopupMsg", setHelpPopupMsg)
	datePickerRouter.HandleFunc("/api/tableView/datePicker/setConditionalFormats", setConditionalFormats)

	datePickerRouter.HandleFunc("/api/tableView/datePicker/validateInput", validateInputAPI)

	http.Handle("/api/tableView/datePicker/", datePickerRouter)
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

type GetDatePickerParams struct {
	ParentTableID string `json:"parentTableID"`
	DatePickerID  string `json:"datePickerID"`
}

func getDatePickerAPI(w http.ResponseWriter, r *http.Request) {

	var params GetDatePickerParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	datePicker, err := getDatePicker(params.ParentTableID, params.DatePickerID)
	if err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	api.WriteJSONResponse(w, *datePicker)
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

func setHelpPopupMsg(w http.ResponseWriter, r *http.Request) {
	var params HelpPopupMsgParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processDatePickerPropUpdate(w, r, params)
}

func setConditionalFormats(w http.ResponseWriter, r *http.Request) {
	var params ConditionalFormatParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processDatePickerPropUpdate(w, r, params)
}
