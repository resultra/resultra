package userTag

import (
	"log"
	"resultra/datasheet/server/generic/inputValidation"
)

type ValidateInputParams struct {
	UserTagIDHeader
	InputVal []string `json:"inputVal"`
}

func validateInput(params ValidateInputParams) inputValidation.ValidationResult {

	userSel, err := getUserTag(params.getParentTableID(), params.getUserTagID())
	if err != nil {
		log.Printf("user selection: validate input: %v", err)
		return inputValidation.FailValidationResult(inputValidation.SystemErrValidationMsg)
	}

	if userSel.Properties.Validation.ValueRequired {
		if params.InputVal == nil {
			return inputValidation.FailValidationResult("Selection is required")
		} else if len(params.InputVal) == 0 {
			return inputValidation.FailValidationResult("Selection is required")
		} else {
			return inputValidation.SuccessValidationResult()
		}
	} else {
		return inputValidation.SuccessValidationResult()
	}

}
