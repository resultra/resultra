package file

import (
	"resultra/datasheet/server/generic/inputValidation"
	"resultra/datasheet/server/generic/stringValidation"
)

type FileValidateInputParams struct {
	FileIDHeader
	Attachment *string `json:"inputVal"`
}

func validateInput(params FileValidateInputParams) inputValidation.ValidationResult {

	file, err := getFile(params.getParentFormID(), params.getFileID())
	if err != nil {
		return inputValidation.FailValidationResult(inputValidation.SystemErrValidationMsg)
	}

	if file.Properties.Validation.ValueRequired {
		if params.Attachment == nil || stringValidation.StringAllWhitespace(*params.Attachment) {
			return inputValidation.FailValidationResult("Value is required")
		} else {
			return inputValidation.SuccessValidationResult()
		}
	} else {
		return inputValidation.SuccessValidationResult()
	}

}
