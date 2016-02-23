package datamodel

import (
	"appengine"
	"bytes"
	"fmt"
	"strconv"
)

type EqnEvalFunc func(evalContext *EqnEvalContext, funcArgs []EquationNode) (*EquationResult, error)

type FunctionInfo struct {
	funcName   string
	resultType string
	evalFunc   EqnEvalFunc
}

type FuncNameFuncInfoMap map[string]FunctionInfo

type EqnEvalContext struct {
	appEngContext appengine.Context
	definedFuncs  FuncNameFuncInfoMap

	// Record into which the results will be calculated. This is also the record
	// which is referenced for field values, in the case a calculated field references
	// other fields.
	resultRecord RecordRef
}

// This function needs to use a FieldRefIDIndex instead of get for every field. However,
// since the Fields are currently stored without an ancestore, they are not consistent.
// TODO - migrate this function back to using FieldRefIDIndex once Fields are setup
// with an ancestore.
func (equation EquationNode) UserText(appEngContext appengine.Context) (string, error) {

	var resultBuf bytes.Buffer

	if len(equation.FuncName) > 0 {

		resultBuf.WriteString(equation.FuncName)
		resultBuf.WriteString("(")

		handledFirstArg := false
		for _, arg := range equation.FuncArgs {
			if handledFirstArg {
				// After the first argument has been processed, prefix
				// the next argument with a comma.
				resultBuf.WriteString(",")
			}
			handledFirstArg = true

			argEquationText, err := arg.UserText(appEngContext)
			if err != nil {
				return "", err
			} else {
				resultBuf.WriteString(argEquationText)
			}

		} // for each argument

		resultBuf.WriteString(")")
	} else if len(equation.FieldID) > 0 {

		fieldRef, err := GetField(appEngContext, GetFieldParams{equation.FieldID})
		if err != nil {
			return "", err
		} else {
			resultBuf.WriteString(fieldRef.FieldInfo.RefName)
		}

	} else if equation.TextVal != nil {
		resultBuf.WriteString("\"")
		resultBuf.WriteString(*equation.TextVal)
		resultBuf.WriteString("\"")
	} else if equation.NumberVal != nil {
		resultBuf.WriteString(strconv.FormatFloat(*equation.NumberVal, 'f', 6, 64))
	}
	return resultBuf.String(), nil
}

func (equation EquationNode) evalEqn(evalContext *EqnEvalContext) (*EquationResult, error) {

	if len(equation.FuncName) > 0 {
		funcInfo, funcInfoFound := evalContext.definedFuncs[equation.FuncName]
		if !funcInfoFound {
			return nil, fmt.Errorf("EvalEqn: Undefined function: %v", equation.FuncName)
		}
		if funcEvalResult, funcErr := funcInfo.evalFunc(evalContext, equation.FuncArgs); funcErr != nil {
			// Function failed to compute
			return nil, funcErr
		} else {
			// TBD - Is it necessary to check the result type from the function
			// to ensure it matches the expected result type of this equation.
			return funcEvalResult, nil
		}

	} else if len(equation.FieldID) > 0 {

		// Equation references a field
		// TODO - Once the Field type has a parent, don't use an individual database
		// lookup for each field (database only has strong consistency when
		// entities have a parent.
		fieldRef, err := GetField(evalContext.appEngContext, GetFieldParams{equation.FieldID})
		if err != nil {
			return nil, fmt.Errorf("evalEqn: failure retrieving referenced field: %+v", err)
		} else {
			return fieldRef.evalEqn(evalContext)
		}

	} else if equation.TextVal != nil {
		// Text literal
		return textEqnResult(*equation.TextVal), nil
	} else if equation.NumberVal != nil {
		// Number literal
		return numberEqnResult(*equation.NumberVal), nil
	} else {
		return nil, fmt.Errorf("evalEqn: malformed calculated field equation : system error: %+v", equation)
	}

}
