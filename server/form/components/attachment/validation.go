package attachment

import (
	"database/sql"
	"log"
	"resultra/tracker/server/generic/inputValidation"
)

type ValidateInputParams struct {
	ImageIDHeader
	Attachments []string `json:"attachments"`
}

func validateInput(trackerDBHandle *sql.DB, params ValidateInputParams) inputValidation.ValidationResult {

	attachComp, err := getImage(trackerDBHandle, params.getParentFormID(), params.getImageID())
	if err != nil {
		log.Printf("Error getting attachment component for form validation: %v", err)
		return inputValidation.FailValidationResult(inputValidation.SystemErrValidationMsg)
	}

	if attachComp.Properties.Validation.ValueRequired {
		if len(params.Attachments) == 0 {
			return inputValidation.FailValidationResult("Attachment is required")
		} else {
			return inputValidation.SuccessValidationResult()
		}
	} else {
		return inputValidation.SuccessValidationResult()
	}

}
