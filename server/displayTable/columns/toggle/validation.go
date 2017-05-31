package toggle

type ToggleValidateInputParams struct {
	ToggleIDHeader
	InputVal *bool `json:"inputVal"`
}

type validationResult struct {
	ValidationSucceeded bool   `json:"validationSucceeded"`
	ErrorMsg            string `json:"errorMsg,omitempty"`
}

const defaultErrMsg string = "Value is required"
const systemErrValidationMsg = "System error validating input"

func successValidationResult() validationResult {
	return validationResult{true, ""}
}

func failValidationResult(errorMsg string) validationResult {
	return validationResult{false, errorMsg}
}

func validateInput(params ToggleValidateInputParams) validationResult {

	toggle, err := getToggle(params.getParentTableID(), params.getToggleID())
	if err != nil {
		return failValidationResult(systemErrValidationMsg)
	}

	if toggle.Properties.Validation.ValueRequired {
		if params.InputVal != nil {
			return successValidationResult()
		} else {
			return failValidationResult(defaultErrMsg)
		}
	}

	return successValidationResult()
}
