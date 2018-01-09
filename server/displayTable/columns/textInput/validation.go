package textInput

import (
	"database/sql"
	"log"
	"resultra/datasheet/server/generic/inputValidation"
	"resultra/datasheet/server/generic/stringValidation"
)

type TextInputValidateInputParams struct {
	TextInputIDHeader
	InputVal *string `json:"inputVal"`
}

func validateInput(trackerDBHandle *sql.DB, params TextInputValidateInputParams) inputValidation.ValidationResult {

	textInput, err := getTextInput(trackerDBHandle, params.getParentTableID(), params.getTextInputID())
	if err != nil {
		log.Printf("Error validating text box input: error = %v", err)
		return inputValidation.FailValidationResult(inputValidation.SystemErrValidationMsg)
	}

	if textInput.Properties.Validation.ValueRequired {
		if params.InputVal == nil || stringValidation.StringAllWhitespace(*params.InputVal) {
			return inputValidation.FailValidationResult("Value is required")
		} else {
			return inputValidation.SuccessValidationResult()
		}
	} else {
		return inputValidation.SuccessValidationResult()
	}

}
