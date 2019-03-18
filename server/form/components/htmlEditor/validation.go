package htmlEditor

import (
	"database/sql"
	"resultra/tracker/server/generic/inputValidation"
	"resultra/tracker/server/generic/stringValidation"
)

type ValidateInputParams struct {
	HtmlEditorIDHeader
	InputVal *string `json:"inputVal"`
}

func validateInput(trackerDBHandle *sql.DB, params ValidateInputParams) inputValidation.ValidationResult {

	editor, err := getHtmlEditor(trackerDBHandle, params.getParentFormID(), params.getHtmlEditorID())
	if err != nil {
		return inputValidation.FailValidationResult(inputValidation.SystemErrValidationMsg)
	}

	if editor.Properties.Validation.ValueRequired {
		if params.InputVal == nil || stringValidation.StringAllWhitespace(*params.InputVal) {
			return inputValidation.FailValidationResult("Note is required")
		} else {
			return inputValidation.SuccessValidationResult()
		}
	} else {
		return inputValidation.SuccessValidationResult()
	}

}
