package userSelection

import (
	"resultra/datasheet/server/generic/inputValidation"
)

type ValidateInputParams struct {
	UserSelectionIDHeader
	InputVal string `json:"inputVal"`
}

func validateInput(params ValidateInputParams) inputValidation.ValidationResult {

	userSel, err := getUserSelection(params.getParentFormID(), params.getUserSelectionID())
	if err != nil {
		return inputValidation.FailValidationResult(inputValidation.SystemErrValidationMsg)
	}

	// TODO - Validate the given string is an actual user ID
	if userSel.Properties.Validation.ValueRequired {
		if len(params.InputVal) == 0 {
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
