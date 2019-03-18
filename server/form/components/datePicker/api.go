// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package datePicker

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/tracker/server/common/databaseWrapper"
	"resultra/tracker/server/generic/api"
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
	datePickerRouter.HandleFunc("/api/frm/datePicker/setHelpPopupMsg", setHelpPopupMsg)
	datePickerRouter.HandleFunc("/api/frm/datePicker/setConditionalFormats", setConditionalFormats)

	http.Handle("/api/frm/datePicker/", datePickerRouter)
}

func newDatePicker(w http.ResponseWriter, r *http.Request) {

	params := NewDatePickerParams{}
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if checkBoxRef, err := saveNewDatePicker(trackerDBHandle, params); err != nil {
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

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}
	validationResult := validateInput(trackerDBHandle, params)
	api.WriteJSONResponse(w, validationResult)

}

func processDatePickerPropUpdate(w http.ResponseWriter, r *http.Request, propUpdater DatePickerPropUpdater) {
	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}
	if checkBoxRef, err := updateDatePickerProps(trackerDBHandle, propUpdater); err != nil {
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
