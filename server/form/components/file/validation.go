// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package file

import (
	"database/sql"
	"resultra/tracker/server/generic/inputValidation"
	"resultra/tracker/server/generic/stringValidation"
)

type FileValidateInputParams struct {
	FileIDHeader
	Attachment *string `json:"inputVal"`
}

func validateInput(trackerDBHandle *sql.DB, params FileValidateInputParams) inputValidation.ValidationResult {

	file, err := getFile(trackerDBHandle, params.getParentFormID(), params.getFileID())
	if err != nil {
		return inputValidation.FailValidationResult(inputValidation.SystemErrValidationMsg)
	}

	if file.Properties.Validation.ValueRequired {
		if params.Attachment == nil || stringValidation.StringAllWhitespace(*params.Attachment) {
			return inputValidation.FailValidationResult("Value is required")
		} else {
			return inputValidation.SuccessValidationResult()
		}
	} else {
		return inputValidation.SuccessValidationResult()
	}

}
