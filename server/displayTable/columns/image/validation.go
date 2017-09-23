package image

import (
	"resultra/datasheet/server/generic/inputValidation"
	"resultra/datasheet/server/generic/stringValidation"
)

type ImageValidateInputParams struct {
	ImageIDHeader
	InputVal *string `json:"inputVal"`
}

func validateInput(params ImageValidateInputParams) inputValidation.ValidationResult {

	image, err := getImage(params.getParentTableID(), params.getImageID())
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
