package htmlEditor

import (
	"resultra/datasheet/server/generic/inputValidation"
)

type ValidateInputParams struct {
	HtmlEditorIDHeader
	InputVal *string `json:"inputVal"`
}

func validateInput(params ValidateInputParams) inputValidation.ValidationResult {

	editor, err := getHtmlEditor(params.getParentFormID(), params.getHtmlEditorID())
	if err != nil {
		return inputValidation.FailValidationResult(inputValidation.SystemErrValidationMsg)
	}

	if editor.Properties.Validation.ValueRequired {
		if params.InputVal == nil {
			return inputValidation.FailValidationResult("Value is required")
		} else {
			return inputValidation.SuccessValidationResult()
		}
	} else {
		return inputValidation.SuccessValidationResult()
	}

}
