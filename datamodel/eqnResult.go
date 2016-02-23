package datamodel

import (
	"fmt"
	"math"
)

type EquationResult struct {
	ResultType string

	TextVal   *string
	NumberVal *float64
}

func (eqnResult EquationResult) validateTextResult() error {
	if eqnResult.ResultType != fieldTypeText {
		return fmt.Errorf("EquationResult: Invalid result - expecting %v, got %v", fieldTypeText, eqnResult.ResultType)
	} else if eqnResult.TextVal == nil {
		return fmt.Errorf("EquationResult: Malformed result - missing text value (value == nil)")
	} else {
		return nil
	}
}

func (eqnResult EquationResult) getTextResult() (string, error) {
	if validateErr := eqnResult.validateTextResult(); validateErr != nil {
		return "", validateErr
	} else {
		textVal := *eqnResult.TextVal
		return textVal, nil
	}
}

func (eqnResult EquationResult) validateNumberResult() error {
	if eqnResult.ResultType != fieldTypeNumber {
		return fmt.Errorf("EquationResult: Invalid result - expecting %v, got %v", fieldTypeNumber, eqnResult.ResultType)
	} else if eqnResult.NumberVal == nil {
		return fmt.Errorf("EquationResult: Malformed result - missing numeric value (value == nil)")
	} else {
		return nil
	}
}

func (eqnResult EquationResult) getNumberResult() (float64, error) {
	if validateErr := eqnResult.validateNumberResult(); validateErr != nil {
		return math.NaN(), validateErr
	} else {
		numberVal := *eqnResult.NumberVal
		return numberVal, nil
	}
}

func numberEqnResult(val float64) *EquationResult {
	theVal := val
	return &EquationResult{ResultType: fieldTypeNumber, NumberVal: &theVal}
}

func textEqnResult(val string) *EquationResult {
	theVal := val
	return &EquationResult{ResultType: fieldTypeText, TextVal: &theVal}
}
