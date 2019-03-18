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
