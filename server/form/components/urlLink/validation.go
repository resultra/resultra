// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package urlLink

import (
	"database/sql"
	"github.com/asaskevich/govalidator"
	"github.com/resultra/resultra/server/generic/inputValidation"
	"github.com/resultra/resultra/server/generic/stringValidation"
)

type UrlLinkValidateInputParams struct {
	UrlLinkIDHeader
	InputVal *string `json:"inputVal"`
}

func validateInput(trackerDBHandle *sql.DB, params UrlLinkValidateInputParams) inputValidation.ValidationResult {

	urlLink, err := getUrlLink(trackerDBHandle, params.getParentFormID(), params.getUrlLinkID())
	if err != nil {
		return inputValidation.FailValidationResult(inputValidation.SystemErrValidationMsg)
	}

	validateURL := func(url string) inputValidation.ValidationResult {
		validURL := govalidator.IsURL(url)
		if validURL {
			return inputValidation.SuccessValidationResult()
		} else {
			return inputValidation.FailValidationResult("Invalid URL address")
		}
	}

	if urlLink.Properties.Validation.ValueRequired {
		if params.InputVal == nil || stringValidation.StringAllWhitespace(*params.InputVal) {
			return inputValidation.FailValidationResult("Value is required")
		} else {
			return validateURL(*params.InputVal)
		}
	} else {

		if params.InputVal == nil || len(*params.InputVal) == 0 {
			return inputValidation.SuccessValidationResult()
		} else {
			return validateURL(*params.InputVal)
		}
		return inputValidation.SuccessValidationResult()
	}

}
