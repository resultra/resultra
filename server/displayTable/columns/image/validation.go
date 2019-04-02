// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package image

import (
	"database/sql"
	"github.com/resultra/resultra/server/generic/inputValidation"
	"github.com/resultra/resultra/server/generic/stringValidation"
)

type ImageValidateInputParams struct {
	ImageIDHeader
	InputVal *string `json:"inputVal"`
}

func validateInput(trackerDBHandle *sql.DB, params ImageValidateInputParams) inputValidation.ValidationResult {

	image, err := getImage(trackerDBHandle, params.getParentTableID(), params.getImageID())
	if err != nil {
		return inputValidation.FailValidationResult(inputValidation.SystemErrValidationMsg)
	}

	if image.Properties.Validation.ValueRequired {
		if params.InputVal == nil || stringValidation.StringAllWhitespace(*params.InputVal) {
			return inputValidation.FailValidationResult("Value is required")
		} else {
			return inputValidation.SuccessValidationResult()
		}
	} else {
		return inputValidation.SuccessValidationResult()
	}

}
