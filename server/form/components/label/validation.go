package label

import (
	"database/sql"
	"resultra/tracker/server/generic/inputValidation"
)

type ValidateInputParams struct {
	LabelIDHeader
	InputVal []string `json:"inputVal"`
}

func validateInput(trackerDBHandle *sql.DB, params ValidateInputParams) inputValidation.ValidationResult {

	userSel, err := getLabel(trackerDBHandle, params.getParentFormID(), params.getLabelID())
	if err != nil {
		return inputValidation.FailValidationResult(inputValidation.SystemErrValidationMsg)
	}

	if userSel.Properties.Validation.ValueRequired {
		if params.InputVal == nil {
			return inputValidation.FailValidationResult("Label is required")
		} else if len(params.InputVal) == 0 {
			return inputValidation.FailValidationResult("Label is required")
		} else {
			return inputValidation.SuccessValidationResult()
		}
	} else {
		return inputValidation.SuccessValidationResult()
	}

}
