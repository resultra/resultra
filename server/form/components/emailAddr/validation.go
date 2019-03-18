// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package emailAddr

import (
	"database/sql"
	"resultra/tracker/server/generic/inputValidation"
	"resultra/tracker/server/generic/stringValidation"
	"resultra/tracker/server/common/userAuth"
)

type EmailAddrValidateInputParams struct {
	EmailAddrIDHeader
	InputVal *string `json:"inputVal"`
}

func validateInput(trackerDBHandle *sql.DB, params EmailAddrValidateInputParams) inputValidation.ValidationResult {

	emailAddr, err := getEmailAddr(trackerDBHandle, params.getParentFormID(), params.getEmailAddrID())
	if err != nil {
		return inputValidation.FailValidationResult(inputValidation.SystemErrValidationMsg)
	}

	if emailAddr.Properties.Validation.ValueRequired {
		if params.InputVal == nil || stringValidation.StringAllWhitespace(*params.InputVal) {
			return inputValidation.FailValidationResult("Value is required")
		} else {
			validateResp := userAuth.ValidateWellFormedEmailAddr(*params.InputVal)
			if validateResp.Success {
				return inputValidation.SuccessValidationResult()
			} else {
				return inputValidation.FailValidationResult("Invalid email address")
			}
		}
	} else {

		if params.InputVal == nil || len(*params.InputVal) == 0 {
			return inputValidation.SuccessValidationResult()
		} else {
			validateResp := userAuth.ValidateWellFormedEmailAddr(*params.InputVal)
			if validateResp.Success {
				return inputValidation.SuccessValidationResult()
			} else {
				return inputValidation.FailValidationResult("Invalid email address")
			}
		}
		return inputValidation.SuccessValidationResult()
	}

}
