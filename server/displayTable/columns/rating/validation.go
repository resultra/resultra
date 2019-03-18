// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package rating

import (
	"database/sql"
	"resultra/tracker/server/generic/inputValidation"
)

type RatingValidateInputParams struct {
	RatingIDHeader
	InputVal *float64 `json:"inputVal"`
}

func validateInput(trackerDBHandle *sql.DB, params RatingValidateInputParams) inputValidation.ValidationResult {

	rating, err := getRating(trackerDBHandle, params.getParentTableID(), params.getRatingID())
	if err != nil {
		return inputValidation.FailValidationResult(inputValidation.SystemErrValidationMsg)
	}

	if rating.Properties.Validation.ValueRequired {
		if params.InputVal == nil {
			return inputValidation.FailValidationResult("Rating is required")
		} else {
			return inputValidation.SuccessValidationResult()
		}
	} else {
		return inputValidation.SuccessValidationResult()
	}

}
