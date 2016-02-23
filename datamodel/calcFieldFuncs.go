package datamodel

import (
	"bytes"
	"fmt"
)

const funcNameSum string = "SUM"

func sumEvalFunc(evalContext *EqnEvalContext, funcArgs []EquationNode) (*EquationResult, error) {
	var sumResult float64 = 0.0

	for argIndex, argEqn := range funcArgs {
		argResult, argErr := argEqn.evalEqn(evalContext)
		if argErr != nil {
			return nil, fmt.Errorf("SUM(): Error evaluating argument # %v: arg=%+v, error %v",
				argIndex, argEqn, argErr)
		} else if numberResult, validateErr := argResult.getNumberResult(); validateErr != nil {
			return nil, fmt.Errorf("SUM(): Invalid result found while evaluating argument # %v: arg=%+v, error = %v",
				argIndex, argEqn, validateErr)
		} else {
			sumResult += numberResult
		}
	}

	return numberEqnResult(sumResult), nil

}

const funcNameConcat string = "CONCATENATE"

func concatEvalFunc(evalContext *EqnEvalContext, funcArgs []EquationNode) (*EquationResult, error) {

	var concatBuf bytes.Buffer

	for argIndex, argEqn := range funcArgs {
		argResult, argErr := argEqn.evalEqn(evalContext)
		if argErr != nil {
			return nil, fmt.Errorf("CONCATENATE: Error evaluating argument # %v: arg=%+v, error %v",
				argIndex, argEqn, argErr)
		} else if validateErr := argResult.validateTextResult(); validateErr != nil {
			return nil, fmt.Errorf("CONCATENATE: Invalid result found while evaluating argument # %v: arg=%+v, error = %v",
				argIndex, argEqn, validateErr)
		} else {
			concatBuf.WriteString(*argResult.TextVal)
		}
	}

	return textEqnResult(concatBuf.String()), nil
}

var calcFieldDefinedFuncs = FuncNameFuncInfoMap{
	funcNameSum:    FunctionInfo{funcNameSum, fieldTypeNumber, sumEvalFunc},
	funcNameConcat: FunctionInfo{funcNameConcat, fieldTypeText, concatEvalFunc},
}
