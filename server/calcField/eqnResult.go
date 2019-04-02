// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package calcField

import (
	"fmt"
	"math"
	"github.com/resultra/resultra/server/field"
	"github.com/resultra/resultra/server/generic/timestamp"
	"time"
)

const eqnResultTypeUndefined string = "undefined"

type EquationResult struct {
	ResultType string

	TextVal   *string
	NumberVal *float64
	TimeVal   *time.Time
	BoolVal   *bool
}

func (eqnResult EquationResult) IsUndefined() bool {
	if eqnResult.ResultType == eqnResultTypeUndefined {
		return true
	} else {
		return false
	}
}

func (eqnResult EquationResult) validateTextResult() error {
	if eqnResult.ResultType != field.FieldTypeText {
		return fmt.Errorf("EquationResult: Invalid result - expecting %v, got %v", field.FieldTypeText, eqnResult.ResultType)
	} else if eqnResult.TextVal == nil {
		return fmt.Errorf("EquationResult: Malformed result - missing text value (value == nil)")
	} else {
		return nil
	}
}

func (eqnResult EquationResult) validateTimeResult() error {
	if eqnResult.ResultType != field.FieldTypeTime {
		return fmt.Errorf("EquationResult: Invalid result - expecting %v, got %v", field.FieldTypeTime, eqnResult.ResultType)
	} else if eqnResult.TimeVal == nil {
		return fmt.Errorf("EquationResult: Malformed result - missing text value (value == nil)")
	} else {
		return nil
	}
}

func (eqnResult EquationResult) validateBoolResult() error {
	if eqnResult.ResultType != field.FieldTypeBool {
		return fmt.Errorf("EquationResult: Invalid result - expecting %v, got %v", field.FieldTypeBool, eqnResult.ResultType)
	} else if eqnResult.BoolVal == nil {
		return fmt.Errorf("EquationResult: Malformed result - missing boolean value (value == nil)")
	} else {
		return nil
	}
}

func (eqnResult EquationResult) GetTextResult() (string, error) {
	if validateErr := eqnResult.validateTextResult(); validateErr != nil {
		return "", validateErr
	} else {
		textVal := *eqnResult.TextVal
		return textVal, nil
	}
}

func (eqnResult EquationResult) validateNumberResult() error {
	if eqnResult.ResultType != field.FieldTypeNumber {
		return fmt.Errorf("EquationResult: Invalid result - expecting %v, got %v", field.FieldTypeNumber, eqnResult.ResultType)
	} else if eqnResult.NumberVal == nil {
		return fmt.Errorf("EquationResult: Malformed result - missing numeric value (value == nil)")
	} else {
		return nil
	}
}

func (eqnResult EquationResult) GetNumberResult() (float64, error) {
	if validateErr := eqnResult.validateNumberResult(); validateErr != nil {
		return math.NaN(), validateErr
	} else {
		numberVal := *eqnResult.NumberVal
		return numberVal, nil
	}
}

func (eqnResult EquationResult) GetTimeResult() (time.Time, error) {
	if validateErr := eqnResult.validateTimeResult(); validateErr != nil {
		return timestamp.CurrentTimestampUTC(), validateErr
	} else {
		timeVal := *eqnResult.TimeVal
		return timeVal, nil
	}

}

func (eqnResult EquationResult) GetBoolResult() (bool, error) {
	if validateErr := eqnResult.validateBoolResult(); validateErr != nil {
		return false, validateErr
	} else {
		boolVal := *eqnResult.BoolVal
		return boolVal, nil
	}
}

func (eqnResult EquationResult) IsTrueResult() bool {
	if eqnResult.IsUndefined() {
		return false
	}
	condBoolResult, validateErr := eqnResult.GetBoolResult()
	if validateErr != nil {
		return false
	}
	if condBoolResult == true {
		return true
	} else {
		return false
	}

}

func undefinedEqnResult() *EquationResult {
	return &EquationResult{ResultType: eqnResultTypeUndefined}
}

func numberEqnResult(val float64) *EquationResult {
	theVal := val
	return &EquationResult{ResultType: field.FieldTypeNumber, NumberVal: &theVal}
}

func textEqnResult(val string) *EquationResult {
	theVal := val
	return &EquationResult{ResultType: field.FieldTypeText, TextVal: &theVal}
}

func timeEqnResult(val time.Time) *EquationResult {
	theVal := val
	return &EquationResult{ResultType: field.FieldTypeTime, TimeVal: &theVal}
}

func boolEqnResult(val bool) *EquationResult {
	theVal := val
	return &EquationResult{ResultType: field.FieldTypeBool, BoolVal: &theVal}
}
