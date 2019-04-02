// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package label

import (
	"database/sql"
	"github.com/resultra/resultra/server/generic/inputValidation"
)

type ValidateInputParams struct {
	LabelIDHeader
	InputVal []string `json:"inputVal"`
}

func validateInput(trackerDBHandle *sql.DB, params ValidateInputParams) inputValidation.ValidationResult {

	userSel, err := getLabel(trackerDBHandle, params.getParentFormID(), params.getLabelID())
	if err != nil {
		return inputValidation.FailValidationResult(inputValidation.SystemErrValidationMsg)
	}

	if userSel.Properties.Validation.ValueRequired {
		if params.InputVal == nil {
			return inputValidation.FailValidationResult("Label is required")
		} else if len(params.InputVal) == 0 {
			return inputValidation.FailValidationResult("Label is required")
		} else {
			return inputValidation.SuccessValidationResult()
		}
	} else {
		return inputValidation.SuccessValidationResult()
	}

}
