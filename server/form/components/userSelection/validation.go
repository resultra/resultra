package userSelection

import (
	"database/sql"
	"resultra/tracker/server/generic/inputValidation"
)

type ValidateInputParams struct {
	UserSelectionIDHeader
	InputVal string `json:"inputVal"`
}

func validateInput(trackerDBHandle *sql.DB, params ValidateInputParams) inputValidation.ValidationResult {

	userSel, err := getUserSelection(trackerDBHandle, params.getParentFormID(), params.getUserSelectionID())
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
