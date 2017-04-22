package datePicker

import (
	//	"fmt"
	"resultra/datasheet/server/generic/inputValidation"
	"time"
)

type DatePickerValidateInputParams struct {
	DatePickerIDHeader
	InputVal *time.Time `json:"inputVal"`
}

func validateInput(params DatePickerValidateInputParams) inputValidation.ValidationResult {

	if params.InputVal == nil {
		return inputValidation.FailValidationResult("Value is required")
	}

	return inputValidation.SuccessValidationResult()

}
