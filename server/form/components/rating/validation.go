package rating

import (
	"resultra/datasheet/server/generic/inputValidation"
)

type RatingValidateInputParams struct {
	RatingIDHeader
	InputVal *float64 `json:"inputVal"`
}

func validateInput(params RatingValidateInputParams) inputValidation.ValidationResult {

	rating, err := getRating(params.getParentFormID(), params.getRatingID())
	if err != nil {
		return inputValidation.FailValidationResult(inputValidation.SystemErrValidationMsg)
	}

	if rating.Properties.Validation.ValueRequired {
		if params.InputVal == nil {
			return inputValidation.FailValidationResult("Value is required")
		} else {
			return inputValidation.SuccessValidationResult()
		}
	} else {
		return inputValidation.SuccessValidationResult()
	}

}
