// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package numberInput

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/resultra/resultra/server/common/databaseWrapper"
	"github.com/resultra/resultra/server/generic/api"
)

func init() {
	numberInputRouter := mux.NewRouter()

	numberInputRouter.HandleFunc("/api/tableView/numberInput/new", newNumberInput)

	numberInputRouter.HandleFunc("/api/tableView/numberInput/get", getNumberInputAPI)

	numberInputRouter.HandleFunc("/api/tableView/numberInput/setValueFormat", setValueFormat)
	numberInputRouter.HandleFunc("/api/tableView/numberInput/setLabelFormat", setLabelFormat)
	numberInputRouter.HandleFunc("/api/tableView/numberInput/setPermissions", setPermissions)
	numberInputRouter.HandleFunc("/api/tableView/numberInput/setShowSpinner", setShowSpinner)
	numberInputRouter.HandleFunc("/api/tableView/numberInput/setClearValueSupported", setClearValueSupported)
	numberInputRouter.HandleFunc("/api/tableView/numberInput/setHelpPopupMsg", setHelpPopupMsg)

	numberInputRouter.HandleFunc("/api/tableView/numberInput/setConditionalFormats", setConditionalFormats)

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

type GetNumberInputParams struct {
	ParentTableID string `json:"parentTableID"`
	NumberInputID string `json:"numberInputID"`
}

func getNumberInputAPI(w http.ResponseWriter, r *http.Request) {

	var params GetNumberInputParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	numberInput, err := getNumberInput(trackerDBHandle, params.ParentTableID, params.NumberInputID)
	if err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	api.WriteJSONResponse(w, *numberInput)
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
