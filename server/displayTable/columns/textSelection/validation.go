package textSelection

import (
	"database/sql"
	"log"
	"resultra/tracker/server/generic/inputValidation"
	"resultra/tracker/server/generic/stringValidation"
)

type TextSelectionValidateInputParams struct {
	TextSelectionIDHeader
	InputVal *string `json:"inputVal"`
}

func validateInput(trackerDBHandle *sql.DB, params TextSelectionValidateInputParams) inputValidation.ValidationResult {

	textSelection, err := getTextSelection(trackerDBHandle, params.getParentTableID(), params.getSelectionID())
	if err != nil {
		log.Printf("Error validating text selection input: error = %v", err)
		return inputValidation.FailValidationResult(inputValidation.SystemErrValidationMsg)
	}

	if textSelection.Properties.Validation.ValueRequired {
		if params.InputVal == nil || stringValidation.StringAllWhitespace(*params.InputVal) {
			return inputValidation.FailValidationResult("Value is required")
		} else {
			return inputValidation.SuccessValidationResult()
		}
	} else {
		return inputValidation.SuccessValidationResult()
	}

}
