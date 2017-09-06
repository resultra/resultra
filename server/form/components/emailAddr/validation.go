package emailAddr

import (
	"resultra/datasheet/server/generic/inputValidation"
	"resultra/datasheet/server/generic/stringValidation"
)

type EmailAddrValidateInputParams struct {
	EmailAddrIDHeader
	InputVal *string `json:"inputVal"`
}

func validateInput(params EmailAddrValidateInputParams) inputValidation.ValidationResult {

	emailAddr, err := getEmailAddr(params.getParentFormID(), params.getEmailAddrID())
	if err != nil {
		return inputValidation.FailValidationResult(inputValidation.SystemErrValidationMsg)
	}

	if emailAddr.Properties.Validation.ValueRequired {
		if params.InputVal == nil || stringValidation.StringAllWhitespace(*params.InputVal) {
			return inputValidation.FailValidationResult("Value is required")
		} else {
			return inputValidation.SuccessValidationResult()
		}
	} else {
		return inputValidation.SuccessValidationResult()
	}

}
