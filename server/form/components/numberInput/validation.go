package numberInput

import (
	"database/sql"
	"fmt"
	"resultra/tracker/server/generic/inputValidation"
)

type NumberInputValidateInputParams struct {
	NumberInputIDHeader
	InputVal *float64 `json:"inputVal"`
}

func validateInput(trackerDBHandle *sql.DB, params NumberInputValidateInputParams) inputValidation.ValidationResult {

	numberInput, err := getNumberInput(trackerDBHandle, params.getParentFormID(), params.getNumberInputID())
	if err != nil {
		return inputValidation.FailValidationResult(inputValidation.SystemErrValidationMsg)
	}

	valProps := numberInput.Properties.Validation

	switch valProps.Rule {
	case "none":
		return inputValidation.SuccessValidationResult()
	case "required":
		if params.InputVal == nil {
			return inputValidation.FailValidationResult("Value is required")
		} else {
			return inputValidation.SuccessValidationResult()
		}
	case "between":
		if params.InputVal == nil {
			return inputValidation.FailValidationResult("Value is required")
		} else {
			if valProps.MinVal == nil || valProps.MaxVal == nil {
				return inputValidation.FailValidationResult(inputValidation.SystemErrValidationMsg)
			} else {
				inputVal := *params.InputVal
				if (inputVal >= (*valProps.MinVal)) && (inputVal <= (*valProps.MaxVal)) {
					return inputValidation.SuccessValidationResult()
				} else {
					return inputValidation.FailValidationResult(
						fmt.Sprintf("Value must be between %v and %v", *valProps.MinVal, *valProps.MaxVal))
				}
			}

		}
	case "greater":
		if params.InputVal == nil {
			return inputValidation.FailValidationResult("Value is required")
		} else {
			if valProps.CompareVal == nil {
				return inputValidation.FailValidationResult(inputValidation.SystemErrValidationMsg)
			} else {
				inputVal := *params.InputVal
				if inputVal >= (*valProps.CompareVal) {
					return inputValidation.SuccessValidationResult()
				} else {
					return inputValidation.FailValidationResult(
						fmt.Sprintf("Value must be greater than %v", *valProps.CompareVal))
				}
			}

		}

	default:
		return inputValidation.FailValidationResult(inputValidation.SystemErrValidationMsg)
	}

}
