// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package numberInput

import (
	"database/sql"
	"fmt"
	"github.com/resultra/resultra/server/generic/inputValidation"
)

type NumberInputValidateInputParams struct {
	NumberInputIDHeader
	InputVal *float64 `json:"inputVal"`
}

func validateInput(trackerDBHandle *sql.DB, params NumberInputValidateInputParams) inputValidation.ValidationResult {

	numberInput, err := getNumberInput(trackerDBHandle, params.getParentTableID(), params.getNumberInputID())
	if err != nil {
		return inputValidation.FailValidationResult(inputValidation.SystemErrValidationMsg)
	}

	valProps := numberInput.Properties.Validation

	switch valProps.Rule {
	case "none":
		return inputValidation.SuccessValidationResult()
	case "required":
		if params.InputVal == nil {
			return inputValidation.FailValidationResult("Value is required")
		} else {
			return inputValidation.SuccessValidationResult()
		}
	case "between":
		if params.InputVal == nil {
			return inputValidation.FailValidationResult("Value is required")
		} else {
			if valProps.MinVal == nil || valProps.MaxVal == nil {
				return inputValidation.FailValidationResult(inputValidation.SystemErrValidationMsg)
			} else {
				inputVal := *params.InputVal
				if (inputVal >= (*valProps.MinVal)) && (inputVal <= (*valProps.MaxVal)) {
					return inputValidation.SuccessValidationResult()
				} else {
					return inputValidation.FailValidationResult(
						fmt.Sprintf("Value must be between %v and %v", *valProps.MinVal, *valProps.MaxVal))
				}
			}

		}
	case "greater":
		if params.InputVal == nil {
			return inputValidation.FailValidationResult("Value is required")
		} else {
			if valProps.CompareVal == nil {
				return inputValidation.FailValidationResult(inputValidation.SystemErrValidationMsg)
			} else {
				inputVal := *params.InputVal
				if inputVal >= (*valProps.CompareVal) {
					return inputValidation.SuccessValidationResult()
				} else {
					return inputValidation.FailValidationResult(
						fmt.Sprintf("Value must be greater than %v", *valProps.CompareVal))
				}
			}

		}

	default:
		return inputValidation.FailValidationResult(inputValidation.SystemErrValidationMsg)
	}

}
