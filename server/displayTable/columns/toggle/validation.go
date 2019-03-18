// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package toggle

import (
	"database/sql"
)

type ToggleValidateInputParams struct {
	ToggleIDHeader
	InputVal *bool `json:"inputVal"`
}

type validationResult struct {
	ValidationSucceeded bool   `json:"validationSucceeded"`
	ErrorMsg            string `json:"errorMsg,omitempty"`
}

const defaultErrMsg string = "Value is required"
const systemErrValidationMsg = "System error validating input"

func successValidationResult() validationResult {
	return validationResult{true, ""}
}

func failValidationResult(errorMsg string) validationResult {
	return validationResult{false, errorMsg}
}

func validateInput(trackerDBHandle *sql.DB, params ToggleValidateInputParams) validationResult {

	toggle, err := getToggle(trackerDBHandle, params.getParentTableID(), params.getToggleID())
	if err != nil {
		return failValidationResult(systemErrValidationMsg)
	}

	if toggle.Properties.Validation.ValueRequired {
		if params.InputVal != nil {
			return successValidationResult()
		} else {
			return failValidationResult(defaultErrMsg)
		}
	}

	return successValidationResult()
}
