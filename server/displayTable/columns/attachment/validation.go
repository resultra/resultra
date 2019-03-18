// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package attachment

import (
	"database/sql"
	"log"
	"resultra/tracker/server/generic/inputValidation"
)

type ValidateInputParams struct {
	AttachmentIDHeader
	Attachments []string `json:"attachments"`
}

func validateInput(trackerDBHandle *sql.DB, params ValidateInputParams) inputValidation.ValidationResult {

	attachComp, err := getAttachment(trackerDBHandle, params.getParentTableID(), params.getAttachmentID())
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
