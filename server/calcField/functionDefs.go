package calcField

import (
	"bytes"
	"fmt"
	"log"
	"resultra/datasheet/server/field"
	"time"
)

type FuncSemAnalysisParams struct {
	context  *semanticAnalysisContext
	funcName string
	funcArgs []*EquationNode
}

type EqnEvalFunc func(evalContext *EqnEvalContext, funcArgs []*EquationNode) (*EquationResult, error)
type FuncSemAnalyzeFunc func(semAnalysisParams FuncSemAnalysisParams) (*semanticAnalysisResult, error)

type FunctionInfo struct {
	funcName        string
	resultType      string
	evalFunc        EqnEvalFunc
	semAnalysisFunc FuncSemAnalyzeFunc
}

type FuncNameFuncInfoMap map[string]FunctionInfo

// Semantic analysis functions

func oneOrMoreNumberArgs(params FuncSemAnalysisParams) (*semanticAnalysisResult, error) {

	log.Printf("oneOrMoreNumberArgs: %v", params.funcName)

	if len(params.funcArgs) <= 0 {
		// Even though there's an errors, based upon the function type we know it will return a number. This
		// allows semantic analysis to continue, even though there might be some errors.
		errMsgs := []string{fmt.Sprintf("Not enough arguments to function %v, expecting at at least 1 argument", params.funcName)}
		return &semanticAnalysisResult{analyzeErrors: errMsgs, resultType: field.FieldTypeNumber}, nil
	}

	argErrors := []string{}
	for argIndex, argEqnNode := range params.funcArgs {
		argNum := argIndex + 1
		analyzeResult, analyzeErr := analyzeEqnNode(params.context, argEqnNode)
		if analyzeErr != nil {
			return nil, analyzeErr
		}
		argErrors = append(argErrors, analyzeResult.analyzeErrors...)
		if analyzeResult.resultType != field.FieldTypeNumber {
			argErrors = append(
				argErrors, fmt.Sprintf("Invalid argument type for argument %v of function %v. Expecting a number", argNum, params.funcName))
		}
	}

	return &semanticAnalysisResult{analyzeErrors: argErrors, resultType: field.FieldTypeNumber}, nil

}

func twoNumberArgs(params FuncSemAnalysisParams) (*semanticAnalysisResult, error) {
	if len(params.funcArgs) != 2 {
		errMsgs := []string{fmt.Sprintf("Expecting 2 numerical arguments to function %v", params.funcName)}
		return &semanticAnalysisResult{analyzeErrors: errMsgs, resultType: field.FieldTypeNumber}, nil
	}
	return oneOrMoreNumberArgs(params)
}

func twoNumberArgsBooleanResult(params FuncSemAnalysisParams) (*semanticAnalysisResult, error) {
	if len(params.funcArgs) != 2 {
		errMsgs := []string{fmt.Sprintf("Expecting 2 numerical arguments to function %v", params.funcName)}
		return &semanticAnalysisResult{analyzeErrors: errMsgs, resultType: field.FieldTypeBool}, nil
	}
	errMsgs := []string{}
	return &semanticAnalysisResult{analyzeErrors: errMsgs, resultType: field.FieldTypeBool}, nil
}

func oneOrMoreTextArgs(params FuncSemAnalysisParams) (*semanticAnalysisResult, error) {

	if len(params.funcArgs) <= 0 {
		// Even though there's an errors, based upon the function type we know it will return text. This
		// allows semantic analysis to continue, even though there might be some errors.
		errMsgs := []string{fmt.Sprintf("Not enough arguments to function %v, expecting at at least 1 argument", params.funcName)}
		return &semanticAnalysisResult{analyzeErrors: errMsgs, resultType: field.FieldTypeText}, nil
	}

	argErrors := []string{}
	for argIndex, argEqnNode := range params.funcArgs {
		argNum := argIndex + 1
		analyzeResult, analyzeErr := analyzeEqnNode(params.context, argEqnNode)
		if analyzeErr != nil {
			return nil, analyzeErr
		}
		argErrors = append(argErrors, analyzeResult.analyzeErrors...)
		if analyzeResult.resultType != field.FieldTypeText {
			argErrors = append(
				argErrors, fmt.Sprintf("Invalid argument type for argument %v of function %v. Expecting text", argNum, params.funcName))
		}
	}

	return &semanticAnalysisResult{analyzeErrors: argErrors, resultType: field.FieldTypeText}, nil
}

// Individual function implementations

const FuncNameSum string = "SUM"

func sumEvalFunc(evalContext *EqnEvalContext, funcArgs []*EquationNode) (*EquationResult, error) {

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

const FuncNameIf string = "IF"

func validIfArgs(params FuncSemAnalysisParams) (*semanticAnalysisResult, error) {

	if len(params.funcArgs) != 3 {
		// Even though there's an errors, based upon the function type we know it will return text. This
		// allows semantic analysis to continue, even though there might be some errors.
		errMsgs := []string{fmt.Sprintf("Expecting 3 arguments to function %v, ", params.funcName)}
		return &semanticAnalysisResult{analyzeErrors: errMsgs, resultType: field.FieldTypeTime}, nil
	}

	argErrors := []string{}

	arg1EqnNode := params.funcArgs[0]
	arg1AnalyzeResult, analyzeErr := analyzeEqnNode(params.context, arg1EqnNode)
	if analyzeErr != nil {
		return nil, analyzeErr
	}
	if arg1AnalyzeResult.resultType != field.FieldTypeBool {
		argErrors = append(
			argErrors, fmt.Sprintf("Invalid argument type for argument 1 of function %v. Expecting boolean", params.funcName))
	}

	arg2EqnNode := params.funcArgs[1]
	arg2AnalyzeResult, analyzeErr := analyzeEqnNode(params.context, arg2EqnNode)
	if analyzeErr != nil {
		return nil, analyzeErr
	}

	arg3EqnNode := params.funcArgs[2]
	arg3AnalyzeResult, analyzeErr := analyzeEqnNode(params.context, arg3EqnNode)
	if analyzeErr != nil {
		return nil, analyzeErr
	}

	if arg2AnalyzeResult.resultType != arg3AnalyzeResult.resultType {
		argErrors = append(
			argErrors, fmt.Sprintf("Invalid argument types for function %v. 2nd and 3rd argument types must match", params.funcName))
	}

	return &semanticAnalysisResult{analyzeErrors: argErrors, resultType: field.FieldTypeTime}, nil
}

func ifEvalFunc(evalContext *EqnEvalContext, funcArgs []*EquationNode) (*EquationResult, error) {

	if len(funcArgs) != 3 {
		return nil, fmt.Errorf("IF() - Expecting 3 arguments, got %v", len(funcArgs))
	}

	condEqn := funcArgs[0]
	condResult, condErr := condEqn.EvalEqn(evalContext)
	if condErr != nil {
		return nil, fmt.Errorf("IF(): Error evaluating argument # %v: arg=%+v, error %v", condEqn, condErr)
	} else if condResult.IsUndefined() {
		// If an undefined result is returned, return immediately and propogate the undefined
		// result value up through the equation evaluation.
		return condResult, nil
	} else if condBoolResult, validateErr := condResult.GetBoolResult(); validateErr != nil {
		return nil, fmt.Errorf("IF(): Invalid result found while evaluating argument 1: arg=%+v, error = %v", condEqn, validateErr)
	} else {
		if condBoolResult == true {
			return funcArgs[1].EvalEqn(evalContext)
		} else {
			return funcArgs[2].EvalEqn(evalContext)
		}
	}

}

const FuncNameMinus string = "MINUS"

func minusEvalFunc(evalContext *EqnEvalContext, funcArgs []*EquationNode) (*EquationResult, error) {

	if len(funcArgs) != 2 {
		return nil, fmt.Errorf("MINUS() - Expecting 2 arguments, got %v", len(funcArgs))
	}

	arg1Eqn := funcArgs[0]
	arg1Result, arg1Err := arg1Eqn.EvalEqn(evalContext)
	if arg1Err != nil {
		return nil, fmt.Errorf("MINUS(): Error evaluating argument # %v: arg=%+v, error %v", arg1Eqn, arg1Err)
	} else if arg1Result.IsUndefined() {
		// If an undefined result is returned, return immediately and propogate the undefined
		// result value up through the equation evaluation.
		return arg1Result, nil
	} else if arg1NumberResult, validateErr := arg1Result.GetNumberResult(); validateErr != nil {
		return nil, fmt.Errorf("DATEADD(): Invalid result found while evaluating argument 1: arg=%+v, error = %v", arg1Eqn, validateErr)
	} else {
		arg2Eqn := funcArgs[1]
		arg2Result, arg2Err := arg2Eqn.EvalEqn(evalContext)
		if arg2Err != nil {
			return nil, fmt.Errorf("MINUS(): Error evaluating argument # 2: arg=%+v, error %v", arg2Eqn, arg2Err)
		} else if arg2Result.IsUndefined() {
			// If an undefined result is returned, return immediately and propogate the undefined
			// result value up through the equation evaluation.
			return arg2Result, nil
		} else if arg2NumberResult, validateErr := arg2Result.GetNumberResult(); validateErr != nil {
			return nil, fmt.Errorf("MINUS(): Invalid result found while evaluating argument 2: arg=%+v, error = %v", arg2Eqn, validateErr)
		} else {

			return numberEqnResult(arg1NumberResult - arg2NumberResult), nil

		}
	}

}

const FuncNameDivide string = "DIVIDE"

func divideEvalFunc(evalContext *EqnEvalContext, funcArgs []*EquationNode) (*EquationResult, error) {

	if len(funcArgs) != 2 {
		return nil, fmt.Errorf("DIVIDE() - Expecting 2 arguments, got %v", len(funcArgs))
	}

	arg1Eqn := funcArgs[0]
	arg1Result, arg1Err := arg1Eqn.EvalEqn(evalContext)
	if arg1Err != nil {
		return nil, fmt.Errorf("MINUS(): Error evaluating argument # %v: arg=%+v, error %v", arg1Eqn, arg1Err)
	} else if arg1Result.IsUndefined() {
		// If an undefined result is returned, return immediately and propogate the undefined
		// result value up through the equation evaluation.
		return arg1Result, nil
	} else if arg1NumberResult, validateErr := arg1Result.GetNumberResult(); validateErr != nil {
		return nil, fmt.Errorf("DIVIDE(): Invalid result found while evaluating argument 1: arg=%+v, error = %v", arg1Eqn, validateErr)
	} else {
		arg2Eqn := funcArgs[1]
		arg2Result, arg2Err := arg2Eqn.EvalEqn(evalContext)
		if arg2Err != nil {
			return nil, fmt.Errorf("DIVIDE(): Error evaluating argument # 2: arg=%+v, error %v", arg2Eqn, arg2Err)
		} else if arg2Result.IsUndefined() {
			// If an undefined result is returned, return immediately and propogate the undefined
			// result value up through the equation evaluation.
			return arg2Result, nil
		} else if arg2NumberResult, validateErr := arg2Result.GetNumberResult(); validateErr != nil {
			return nil, fmt.Errorf("MINUS(): Invalid result found while evaluating argument 2: arg=%+v, error = %v", arg2Eqn, validateErr)
		} else {

			if arg2NumberResult == 0.0 {
				return undefinedEqnResult(), nil
			} else {
				return numberEqnResult(arg1NumberResult / arg2NumberResult), nil
			}

		}
	}

}

const FuncNameGreaterThan string = "GREATERTHAN"

func greaterThanEvalFunc(evalContext *EqnEvalContext, funcArgs []*EquationNode) (*EquationResult, error) {

	if len(funcArgs) != 2 {
		return nil, fmt.Errorf("GREATERTHAN() - Expecting 2 arguments, got %v", len(funcArgs))
	}

	arg1Eqn := funcArgs[0]
	arg1Result, arg1Err := arg1Eqn.EvalEqn(evalContext)
	if arg1Err != nil {
		return nil, fmt.Errorf("GREATERTHAN(): Error evaluating argument # %v: arg=%+v, error %v", arg1Eqn, arg1Err)
	} else if arg1Result.IsUndefined() {
		// If an undefined result is returned, return immediately and propogate the undefined
		// result value up through the equation evaluation.
		return arg1Result, nil
	} else if arg1NumberResult, validateErr := arg1Result.GetNumberResult(); validateErr != nil {
		return nil, fmt.Errorf("GREATERTHAN(): Invalid result found while evaluating argument 1: arg=%+v, error = %v", arg1Eqn, validateErr)
	} else {
		arg2Eqn := funcArgs[1]
		arg2Result, arg2Err := arg2Eqn.EvalEqn(evalContext)
		if arg2Err != nil {
			return nil, fmt.Errorf("DIVIDE(): Error evaluating argument # 2: arg=%+v, error %v", arg2Eqn, arg2Err)
		} else if arg2Result.IsUndefined() {
			// If an undefined result is returned, return immediately and propogate the undefined
			// result value up through the equation evaluation.
			return arg2Result, nil
		} else if arg2NumberResult, validateErr := arg2Result.GetNumberResult(); validateErr != nil {
			return nil, fmt.Errorf("GREATERTHAN(): Invalid result found while evaluating argument 2: arg=%+v, error = %v", arg2Eqn, validateErr)
		} else {

			if arg1NumberResult > arg2NumberResult {
				return boolEqnResult(true), nil
			} else {
				return boolEqnResult(false), nil
			}

		}
	}

}

const FuncNameDateAdd string = "DATEADD"

func validDateAddArgs(params FuncSemAnalysisParams) (*semanticAnalysisResult, error) {

	if len(params.funcArgs) != 2 {
		// Even though there's an errors, based upon the function type we know it will return text. This
		// allows semantic analysis to continue, even though there might be some errors.
		errMsgs := []string{fmt.Sprintf("Expecting 2 arguments to function %v, ", params.funcName)}
		return &semanticAnalysisResult{analyzeErrors: errMsgs, resultType: field.FieldTypeTime}, nil
	}

	argErrors := []string{}

	arg1EqnNode := params.funcArgs[0]
	arg1AnalyzeResult, analyzeErr := analyzeEqnNode(params.context, arg1EqnNode)
	if analyzeErr != nil {
		return nil, analyzeErr
	}
	if arg1AnalyzeResult.resultType != field.FieldTypeTime {
		argErrors = append(
			argErrors, fmt.Sprintf("Invalid argument type for argument 1 of function %v. Expecting date/time", params.funcName))
	}

	arg2EqnNode := params.funcArgs[1]
	arg2AnalyzeResult, analyzeErr := analyzeEqnNode(params.context, arg2EqnNode)
	if analyzeErr != nil {
		return nil, analyzeErr
	}
	if arg2AnalyzeResult.resultType != field.FieldTypeNumber {
		argErrors = append(
			argErrors, fmt.Sprintf("Invalid argument type for argument 2 of function %v. Expecting a number", params.funcName))
	}

	return &semanticAnalysisResult{analyzeErrors: argErrors, resultType: field.FieldTypeTime}, nil
}

func dateAddEvalFunc(evalContext *EqnEvalContext, funcArgs []*EquationNode) (*EquationResult, error) {

	if len(funcArgs) != 2 {
		return nil, fmt.Errorf("DATEADD() - Expecting 2 arguments, got %v", len(funcArgs))
	}

	arg1Eqn := funcArgs[0]
	arg1Result, arg1Err := arg1Eqn.EvalEqn(evalContext)
	if arg1Err != nil {
		return nil, fmt.Errorf("DATEADD(): Error evaluating argument # %v: arg=%+v, error %v", arg1Eqn, arg1Err)
	} else if arg1Result.IsUndefined() {
		// If an undefined result is returned, return immediately and propogate the undefined
		// result value up through the equation evaluation.
		return arg1Result, nil
	} else if timeResult, validateErr := arg1Result.GetTimeResult(); validateErr != nil {
		return nil, fmt.Errorf("DATEADD(): Invalid result found while evaluating argument 1: arg=%+v, error = %v", arg1Eqn, validateErr)
	} else {
		arg2Eqn := funcArgs[1]
		arg2Result, arg2Err := arg2Eqn.EvalEqn(evalContext)
		if arg2Err != nil {
			return nil, fmt.Errorf("DATEADD(): Error evaluating argument # 2: arg=%+v, error %v", arg2Eqn, arg2Err)
		} else if arg2Result.IsUndefined() {
			// If an undefined result is returned, return immediately and propogate the undefined
			// result value up through the equation evaluation.
			return arg2Result, nil
		} else if numberResult, validateErr := arg2Result.GetNumberResult(); validateErr != nil {
			return nil, fmt.Errorf("DATEADD(): Invalid result found while evaluating argument 2: arg=%+v, error = %v", arg2Eqn, validateErr)
		} else {

			roundToInt := func(num float64) int {
				if num > 0.0 {
					return int(num + 0.5)
				} else {
					return int(num - 0.5)
				}
			}

			secsToAdd := roundToInt(numberResult * 86400.0)

			timeAfterAdd := timeResult.Add(time.Second * time.Duration(secsToAdd))

			return timeEqnResult(timeAfterAdd), nil

		}
	}
}

const FuncNameProduct string = "PRODUCT"

func productEvalFunc(evalContext *EqnEvalContext, funcArgs []*EquationNode) (*EquationResult, error) {

	var prodResult float64 = 1.0

	if len(funcArgs) < 1 {
		return nil, fmt.Errorf("PRODUCT() - Not enough arguments given to function")
	}

	for argIndex, argEqn := range funcArgs {
		argResult, argErr := argEqn.EvalEqn(evalContext)
		if argErr != nil {
			return nil, fmt.Errorf("PRODUCT(): Error evaluating argument # %v: arg=%+v, error %v",
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

func concatEvalFunc(evalContext *EqnEvalContext, funcArgs []*EquationNode) (*EquationResult, error) {

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
	FuncNameSum:         FunctionInfo{FuncNameSum, field.FieldTypeNumber, sumEvalFunc, oneOrMoreNumberArgs},
	FuncNameMinus:       FunctionInfo{FuncNameMinus, field.FieldTypeNumber, minusEvalFunc, twoNumberArgs},
	FuncNameDivide:      FunctionInfo{FuncNameMinus, field.FieldTypeNumber, divideEvalFunc, twoNumberArgs},
	FuncNameProduct:     FunctionInfo{FuncNameProduct, field.FieldTypeNumber, productEvalFunc, oneOrMoreNumberArgs},
	FuncNameConcat:      FunctionInfo{FuncNameConcat, field.FieldTypeText, concatEvalFunc, oneOrMoreTextArgs},
	FuncNameDateAdd:     FunctionInfo{FuncNameDateAdd, field.FieldTypeTime, dateAddEvalFunc, validDateAddArgs},
	FuncNameGreaterThan: FunctionInfo{FuncNameGreaterThan, field.FieldTypeBool, greaterThanEvalFunc, twoNumberArgsBooleanResult},
	FuncNameIf:          FunctionInfo{FuncNameIf, field.FieldTypeBool, ifEvalFunc, validIfArgs}, // TODO - Support other return types
}
