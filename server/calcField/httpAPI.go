// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package calcField

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/resultra/resultra/server/common/databaseWrapper"
	"github.com/resultra/resultra/server/field"
	"github.com/resultra/resultra/server/generic/api"
)

func init() {
	calcFieldRouter := mux.NewRouter()

	calcFieldRouter.HandleFunc("/api/calcField/validateFormula", validateFormula)
	calcFieldRouter.HandleFunc("/api/calcField/setFieldFormula", setFieldFormula)
	calcFieldRouter.HandleFunc("/api/calcField/new", newCalcFieldAPI)
	calcFieldRouter.HandleFunc("/api/calcField/getRawFormulaText", getRawFormulaTextAPI)

	http.Handle("/api/calcField/", calcFieldRouter)
}

func newCalcFieldAPI(w http.ResponseWriter, r *http.Request) {

	var newCalcFieldParams NewCalcFieldParams
	if err := api.DecodeJSONRequest(r, &newCalcFieldParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if newField, err := newCalcField(trackerDBHandle, newCalcFieldParams); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *newField)
	}

}

func getRawFormulaTextAPI(w http.ResponseWriter, r *http.Request) {

	var params GetRawFormulaParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if rawFormulaText, err := getRawFormulaText(trackerDBHandle, params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, rawFormulaText)
	}
}

func validateFormula(w http.ResponseWriter, r *http.Request) {

	var validationParams ValidateFormulaParams
	if err := api.DecodeJSONRequest(r, &validationParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	validationResponse := validateFormulaText(trackerDBHandle, validationParams)
	api.WriteJSONResponse(w, *validationResponse)
}

func setFieldFormula(w http.ResponseWriter, r *http.Request) {

	var setFormulaParams SetFormulaParams
	if err := api.DecodeJSONRequest(r, &setFormulaParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if updatedFieldRef, err := field.UpdateFieldProps(trackerDBHandle, setFormulaParams); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, updatedFieldRef)
	}

}
