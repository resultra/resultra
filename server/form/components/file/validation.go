package file

import (
	"database/sql"
	"resultra/tracker/server/generic/inputValidation"
	"resultra/tracker/server/generic/stringValidation"
)

type FileValidateInputParams struct {
	FileIDHeader
	Attachment *string `json:"inputVal"`
}

func validateInput(trackerDBHandle *sql.DB, params FileValidateInputParams) inputValidation.ValidationResult {

	file, err := getFile(trackerDBHandle, params.getParentFormID(), params.getFileID())
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
