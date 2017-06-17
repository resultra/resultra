package note

import (
	"resultra/datasheet/server/generic/inputValidation"
	"resultra/datasheet/server/generic/stringValidation"
)

type ValidateInputParams struct {
	NoteIDHeader
	InputVal *string `json:"inputVal"`
}

func validateInput(params ValidateInputParams) inputValidation.ValidationResult {

	editor, err := getNote(params.getParentTableID(), params.getNoteID())
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
