package emailAddr

import (
	"resultra/datasheet/server/generic/inputValidation"
	"resultra/datasheet/server/generic/stringValidation"
	"resultra/datasheet/server/generic/userAuth"
)

type EmailAddrValidateInputParams struct {
	EmailAddrIDHeader
	InputVal *string `json:"inputVal"`
}

func validateInput(params EmailAddrValidateInputParams) inputValidation.ValidationResult {

	emailAddr, err := getEmailAddr(params.getParentTableID(), params.getEmailAddrID())
	if err != nil {
		return inputValidation.FailValidationResult(inputValidation.SystemErrValidationMsg)
	}

	if emailAddr.Properties.Validation.ValueRequired {
		if params.InputVal == nil || stringValidation.StringAllWhitespace(*params.InputVal) {
			return inputValidation.FailValidationResult("Value is required")
		} else {
			validateResp := userAuth.ValidateWellFormedEmailAddr(*params.InputVal)
			if validateResp.Success {
				return inputValidation.SuccessValidationResult()
			} else {
				return inputValidation.FailValidationResult("Invalid email address")
			}
		}
	} else {

		if params.InputVal == nil || len(*params.InputVal) == 0 {
			return inputValidation.SuccessValidationResult()
		} else {
			validateResp := userAuth.ValidateWellFormedEmailAddr(*params.InputVal)
			if validateResp.Success {
				return inputValidation.SuccessValidationResult()
			} else {
				return inputValidation.FailValidationResult("Invalid email address")
			}
		}
		return inputValidation.SuccessValidationResult()
	}

}
