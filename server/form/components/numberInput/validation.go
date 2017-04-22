package numberInput

import (
	"resultra/datasheet/server/generic/inputValidation"
)

type NumberInputValidateInputParams struct {
	NumberInputIDHeader
	InputVal *float64 `json:"inputVal"`
}

func validateInput(params NumberInputValidateInputParams) inputValidation.ValidationResult {

	if params.InputVal == nil {
		return inputValidation.FailValidationResult("Value is required")
	}

	return inputValidation.SuccessValidationResult()
}
