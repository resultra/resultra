package textBox

import (
	"resultra/datasheet/server/generic/inputValidation"
	"resultra/datasheet/server/generic/stringValidation"
)

type TextBoxValidateInputParams struct {
	TextBoxIDHeader
	InputVal *string `json:"inputVal"`
}

func validateInput(params TextBoxValidateInputParams) inputValidation.ValidationResult {

	textBox, err := getTextBox(params.getParentFormID(), params.getTextBoxID())
	if err != nil {
		return inputValidation.FailValidationResult(inputValidation.SystemErrValidationMsg)
	}

	if textBox.Properties.Validation.ValueRequired {
		if params.InputVal == nil || stringValidation.StringAllWhitespace(*params.InputVal) {
			return inputValidation.FailValidationResult("Value is required")
		} else {
			return inputValidation.SuccessValidationResult()
		}
	} else {
		return inputValidation.SuccessValidationResult()
	}

}
