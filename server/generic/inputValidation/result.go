// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package inputValidation

type ValidationResult struct {
	ValidationSucceeded bool   `json:"validationSucceeded"`
	ErrorMsg            string `json:"errorMsg,omitempty"`
}

const DefaultErrMsg string = "Value is required"
const SystemErrValidationMsg = "System error validating input"

func SuccessValidationResult() ValidationResult {
	return ValidationResult{true, ""}
}

func FailValidationResult(errorMsg string) ValidationResult {
	return ValidationResult{false, errorMsg}
}
