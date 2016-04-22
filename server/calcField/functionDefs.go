package calcField

import (
	"bytes"
	"fmt"
	"resultra/datasheet/server/field"
)

const FuncNameSum string = "SUM"

func sumEvalFunc(evalContext *EqnEvalContext, funcArgs []EquationNode) (*EquationResult, error) {

	var sumResult float64 = 0.0

	for argIndex, argEqn := range funcArgs {
		argResult, argErr := argEqn.EvalEqn(evalContext)
		if argErr != nil {
			return nil, fmt.Errorf("SUM(): Error evaluating argument # %v: arg=%+v, error %v",
				argIndex, argEqn, argErr)
		} else if argResult.IsUndefined() {
			// If an undefined result is returned, return immediately and propogate the undefined
			// result value up through the equation evaluation.
			return argResult, nil
		} else if numberResult, validateErr := argResult.GetNumberResult(); validateErr != nil {
			return nil, fmt.Errorf("SUM(): Invalid result found while evaluating argument # %v: arg=%+v, error = %v",
				argIndex, argEqn, validateErr)
		} else {
			sumResult += numberResult
		}
	}

	return numberEqnResult(sumResult), nil

}

const FuncNameProduct string = "PRODUCT"

func productEvalFunc(evalContext *EqnEvalContext, funcArgs []EquationNode) (*EquationResult, error) {

	var prodResult float64 = 1.0

	if len(funcArgs) < 1 {
		return nil, fmt.Errorf("PRODUCT() - Not enough arguments given to function")
	}

	for argIndex, argEqn := range funcArgs {
		argResult, argErr := argEqn.EvalEqn(evalContext)
		if argErr != nil {
			return nil, fmt.Errorf("SUM(): Error evaluating argument # %v: arg=%+v, error %v",
				argIndex, argEqn, argErr)
		} else if argResult.IsUndefined() {
			// If an undefined result is returned, return immediately and propogate the undefined
			// result value up through the equation evaluation.
			return argResult, nil
		} else if numberResult, validateErr := argResult.GetNumberResult(); validateErr != nil {
			return nil, fmt.Errorf("PRODUCT(): Invalid result found while evaluating argument # %v: arg=%+v, error = %v",
				argIndex, argEqn, validateErr)
		} else {
			prodResult *= numberResult
		}
	}

	return numberEqnResult(prodResult), nil

}

const FuncNameConcat string = "CONCATENATE"

func concatEvalFunc(evalContext *EqnEvalContext, funcArgs []EquationNode) (*EquationResult, error) {

	var concatBuf bytes.Buffer

	for argIndex, argEqn := range funcArgs {
		argResult, argErr := argEqn.EvalEqn(evalContext)
		if argErr != nil {
			return nil, fmt.Errorf("CONCATENATE: Error evaluating argument # %v: arg=%+v, error %v",
				argIndex, argEqn, argErr)
		} else if argResult.IsUndefined() {
			// If an undefined result is returned, return immediately and propogate the undefined
			// result value up through the equation evaluation.
			return argResult, nil
		} else if validateErr := argResult.validateTextResult(); validateErr != nil {
			return nil, fmt.Errorf("CONCATENATE: Invalid result found while evaluating argument # %v: arg=%+v, error = %v",
				argIndex, argEqn, validateErr)
		} else {
			concatBuf.WriteString(*argResult.TextVal)
		}
	}

	return textEqnResult(concatBuf.String()), nil
}

var CalcFieldDefinedFuncs = FuncNameFuncInfoMap{
	FuncNameSum:     FunctionInfo{FuncNameSum, field.FieldTypeNumber, sumEvalFunc},
	FuncNameProduct: FunctionInfo{FuncNameProduct, field.FieldTypeNumber, productEvalFunc},
	FuncNameConcat:  FunctionInfo{FuncNameConcat, field.FieldTypeText, concatEvalFunc},
}
