package urlLink

import (
	"github.com/asaskevich/govalidator"
	"resultra/datasheet/server/generic/inputValidation"
	"resultra/datasheet/server/generic/stringValidation"
)

type UrlLinkValidateInputParams struct {
	UrlLinkIDHeader
	InputVal *string `json:"inputVal"`
}

func validateInput(params UrlLinkValidateInputParams) inputValidation.ValidationResult {

	urlLink, err := getUrlLink(params.getParentFormID(), params.getUrlLinkID())
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
