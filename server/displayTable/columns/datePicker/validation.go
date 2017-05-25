package datePicker

import (
	"fmt"
	"log"
	"resultra/datasheet/server/generic/inputValidation"
	"time"
)

type DatePickerValidateInputParams struct {
	DatePickerIDHeader
	InputVal *time.Time `json:"inputVal"`
}

func validateInput(params DatePickerValidateInputParams) inputValidation.ValidationResult {

	datePicker, err := getDatePicker(params.getParentTableID(), params.getDatePickerID())
	if err != nil {
		log.Printf("System error validating date picker input: %v", err)
		return inputValidation.FailValidationResult(inputValidation.SystemErrValidationMsg)
	}

	valProps := datePicker.Properties.Validation

	// TODO - Error messages need to be formatted using local time.
	// This can be done by passing in the timezone offset from the client.

	switch valProps.Rule {
	case validationRuleNone:
		return inputValidation.SuccessValidationResult()
	case validationRuleRequired:
		if params.InputVal == nil {
			return inputValidation.FailValidationResult("Date is required")
		} else {
			return inputValidation.SuccessValidationResult()
		}
	case validationRuleFuture:
		if params.InputVal == nil {
			return inputValidation.FailValidationResult("Date is required")
		} else {
			now := time.Now().UTC()
			if (*params.InputVal).After(now) {
				return inputValidation.SuccessValidationResult()
			} else {
				return inputValidation.FailValidationResult("Date must be in the future")
			}
		}
	case validationRulePast:
		if params.InputVal == nil {
			return inputValidation.FailValidationResult("Date is required")
		} else {
			now := time.Now().UTC()
			if (*params.InputVal).Before(now) {
				return inputValidation.SuccessValidationResult()
			} else {
				return inputValidation.FailValidationResult("Date must be in the past")
			}
		}
	case validationRuleBefore:
		if params.InputVal == nil {
			return inputValidation.FailValidationResult("Date is required")
		} else if valProps.CompareDate == nil {
			return inputValidation.FailValidationResult(inputValidation.SystemErrValidationMsg)
		} else {
			if (*params.InputVal).Before(*valProps.CompareDate) {
				return inputValidation.SuccessValidationResult()
			} else {
				errMsg := fmt.Sprintf("Date must be before %v", *valProps.CompareDate)
				return inputValidation.FailValidationResult(errMsg)
			}
		}
	case validationRuleAfter:
		if params.InputVal == nil {
			return inputValidation.FailValidationResult("Date is required")
		} else if valProps.CompareDate == nil {
			return inputValidation.FailValidationResult(inputValidation.SystemErrValidationMsg)
		} else {
			if (*params.InputVal).After(*valProps.CompareDate) {
				return inputValidation.SuccessValidationResult()
			} else {
				errMsg := fmt.Sprintf("Date must be after %v", *valProps.CompareDate)
				return inputValidation.FailValidationResult(errMsg)
			}
		}

	case validationRuleBetween:
		if params.InputVal == nil {
			return inputValidation.FailValidationResult("Date is required")
		} else if (valProps.StartDate == nil) || (valProps.EndDate == nil) {
			return inputValidation.FailValidationResult(inputValidation.SystemErrValidationMsg)
		} else {
			if (*params.InputVal).After(*valProps.StartDate) && (*params.InputVal).After(*valProps.EndDate) {
				return inputValidation.SuccessValidationResult()
			} else {
				errMsg := fmt.Sprintf("Date must be between %v and %v",
					*valProps.StartDate, *valProps.EndDate)
				return inputValidation.FailValidationResult(errMsg)
			}
		}

	default:
		log.Printf("System error validating date picker input: rule= %v", valProps.Rule)
		return inputValidation.SuccessValidationResult()
	}

}
