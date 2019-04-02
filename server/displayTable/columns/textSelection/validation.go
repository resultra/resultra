// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package textSelection

import (
	"database/sql"
	"log"
	"github.com/resultra/resultra/server/generic/inputValidation"
	"github.com/resultra/resultra/server/generic/stringValidation"
)

type TextSelectionValidateInputParams struct {
	TextSelectionIDHeader
	InputVal *string `json:"inputVal"`
}

func validateInput(trackerDBHandle *sql.DB, params TextSelectionValidateInputParams) inputValidation.ValidationResult {

	textSelection, err := getTextSelection(trackerDBHandle, params.getParentTableID(), params.getSelectionID())
	if err != nil {
		log.Printf("Error validating text selection input: error = %v", err)
		return inputValidation.FailValidationResult(inputValidation.SystemErrValidationMsg)
	}

	if textSelection.Properties.Validation.ValueRequired {
		if params.InputVal == nil || stringValidation.StringAllWhitespace(*params.InputVal) {
			return inputValidation.FailValidationResult("Value is required")
		} else {
			return inputValidation.SuccessValidationResult()
		}
	} else {
		return inputValidation.SuccessValidationResult()
	}

}
