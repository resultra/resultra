package image

import (
	"resultra/datasheet/server/generic/inputValidation"
)

type ValidateInputParams struct {
	ImageIDHeader
	InputVal *string `json:"inputVal"`
}

func validateInput(params ValidateInputParams) inputValidation.ValidationResult {

	attachComp, err := getImage(params.getParentFormID(), params.getImageID())
	if err != nil {
		return inputValidation.FailValidationResult(inputValidation.SystemErrValidationMsg)
	}

	if attachComp.Properties.Validation.ValueRequired {
		if params.InputVal == nil {
			return inputValidation.FailValidationResult("Value is required")
		} else {
			return inputValidation.SuccessValidationResult()
		}
	} else {
		return inputValidation.SuccessValidationResult()
	}

}
