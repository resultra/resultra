package inputValidation

type ValidationResult struct {
	ValidationSucceeded bool   `json:"validationSucceeded"`
	ErrorMsg            string `json:"errorMsg,omitempty"`
}

const DefaultErrMsg string = "Value is required"
const SystemErrValidationMsg = "System error validating input"

func SuccessValidationResult() ValidationResult {
	return ValidationResult{true, ""}
}

func FailValidationResult(errorMsg string) ValidationResult {
	return ValidationResult{false, errorMsg}
}
