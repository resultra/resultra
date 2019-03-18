package userTag

import (
	"database/sql"
	"resultra/tracker/server/generic/inputValidation"
)

type ValidateInputParams struct {
	UserTagIDHeader
	InputVal []string `json:"inputVal"`
}

func validateInput(trackerDBHandle *sql.DB, params ValidateInputParams) inputValidation.ValidationResult {

	userSel, err := getUserTag(trackerDBHandle, params.getParentFormID(), params.getUserTagID())
	if err != nil {
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
