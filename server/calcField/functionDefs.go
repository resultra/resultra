package calcField

import (
	"bytes"
	"fmt"
	"math"
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
	evalFunc        EqnEvalFunc
	semAnalysisFunc FuncSemAnalyzeFunc
}

type FuncNameFuncInfoMap map[string]FunctionInfo

// Semantic analysis functions

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
	evalFunc := func(args []float64) (*EquationResult, error) {
		sumResult := 0.0
		for _, arg := range args {
			sumResult = sumResult + arg
		}
		return numberEqnResult(sumResult), nil
	}
	return evalOneOrMoreNumberArgFunc(evalContext, funcArgs, evalFunc)

}

const FuncNameMax string = "MAX"

func maxEvalFunc(evalContext *EqnEvalContext, funcArgs []*EquationNode) (*EquationResult, error) {

	evalFunc := func(args []float64) (*EquationResult, error) {
		max := args[0]
		for _, arg := range args {
			if arg > max {
				max = arg
			}
		}
		return numberEqnResult(max), nil
	}
	return evalOneOrMoreNumberArgFunc(evalContext, funcArgs, evalFunc)

}

const FuncNameMin string = "MIN"

func minEvalFunc(evalContext *EqnEvalContext, funcArgs []*EquationNode) (*EquationResult, error) {

	evalFunc := func(args []float64) (*EquationResult, error) {
		min := args[0]
		for _, arg := range args {
			if arg < min {
				min = arg
			}
		}
		return numberEqnResult(min), nil
	}
	return evalOneOrMoreNumberArgFunc(evalContext, funcArgs, evalFunc)

}

const FuncNameProduct string = "PRODUCT"

func productEvalFunc(evalContext *EqnEvalContext, funcArgs []*EquationNode) (*EquationResult, error) {

	evalFunc := func(args []float64) (*EquationResult, error) {
		prodResult := 1.0
		for _, arg := range args {
			prodResult = prodResult * arg
		}
		return numberEqnResult(prodResult), nil
	}
	return evalOneOrMoreNumberArgFunc(evalContext, funcArgs, evalFunc)

}

const FuncNameMinus string = "MINUS"

func minusEvalFunc(evalContext *EqnEvalContext, funcArgs []*EquationNode) (*EquationResult, error) {

	evalFunc := func(num1, num2 float64) (*EquationResult, error) {
		return numberEqnResult(num1 - num2), nil
	}
	return evalTwoNumberArgFunc(evalContext, funcArgs, evalFunc)
}

const FuncNameDivide string = "DIVIDE"

func divideEvalFunc(evalContext *EqnEvalContext, funcArgs []*EquationNode) (*EquationResult, error) {

	evalFunc := func(num1, num2 float64) (*EquationResult, error) {
		if num2 == 0.0 {
			return undefinedEqnResult(), nil
		} else {
			return numberEqnResult(num1 / num2), nil
		}

	}
	return evalTwoNumberArgFunc(evalContext, funcArgs, evalFunc)
}

const FuncNamePower string = "POWER"

func powerEvalFunc(evalContext *EqnEvalContext, funcArgs []*EquationNode) (*EquationResult, error) {

	evalFunc := func(num1, num2 float64) (*EquationResult, error) {
		return numberEqnResult(math.Pow(num1, num2)), nil
	}
	return evalTwoNumberArgFunc(evalContext, funcArgs, evalFunc)
}

const FuncNameMultiply string = "MULTIPLY"

func multiplyEvalFunc(evalContext *EqnEvalContext, funcArgs []*EquationNode) (*EquationResult, error) {
	evalFunc := func(num1, num2 float64) (*EquationResult, error) {
		return numberEqnResult(num1 * num2), nil
	}
	return evalTwoNumberArgFunc(evalContext, funcArgs, evalFunc)
}

const FuncNameAdd string = "ADD"

func addEvalFunc(evalContext *EqnEvalContext, funcArgs []*EquationNode) (*EquationResult, error) {
	evalFunc := func(num1, num2 float64) (*EquationResult, error) {
		return numberEqnResult(num1 + num2), nil
	}
	return evalTwoNumberArgFunc(evalContext, funcArgs, evalFunc)
}

const FuncNameIf string = "IF"

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

const FuncNameGreaterThan string = "GREATERTHAN"

func greaterThanEvalFunc(evalContext *EqnEvalContext, funcArgs []*EquationNode) (*EquationResult, error) {
	evalFunc := func(num1, num2 float64) (*EquationResult, error) {
		if num1 > num2 {
			return boolEqnResult(true), nil
		} else {
			return boolEqnResult(false), nil
		}
	}
	return evalTwoNumberArgFunc(evalContext, funcArgs, evalFunc)
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

const FuncNameDaysBetween string = "DAYSBETWEEN"

func daysBetweenEvalFunc(evalContext *EqnEvalContext, funcArgs []*EquationNode) (*EquationResult, error) {
	evalFunc := func(startTime, endTime time.Time) (*EquationResult, error) {
		elapsedDuration := endTime.Sub(startTime)
		durationDays := elapsedDuration.Hours() / 24.0

		return numberEqnResult(durationDays), nil
	}
	return evalTwoTimeArgFunc(evalContext, funcArgs, evalFunc)
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

const FuncNameCreated = "CREATED"

func createdFunc(evalContext *EqnEvalContext, funcArgs []*EquationNode) (*EquationResult, error) {
	return timeEqnResult(evalContext.Record.CreateTimestampUTC), nil
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

const FuncNameSequenceNum = "SEQUENCENUM"

func sequenceNumFunc(evalContext *EqnEvalContext, funcArgs []*EquationNode) (*EquationResult, error) {
	return numberEqnResult(float64(evalContext.Record.SequenceNum)), nil
}

var CalcFieldDefinedFuncs = FuncNameFuncInfoMap{

	FuncNameSequenceNum: FunctionInfo{FuncNameSequenceNum, sequenceNumFunc, zeroNumberArgs},
	FuncNameCreated:     FunctionInfo{FuncNameCreated, createdFunc, zeroTimeArgsTimeResult},

	FuncNameSum: FunctionInfo{FuncNameSum, sumEvalFunc, oneOrMoreNumberArgs},
	FuncNameAdd: FunctionInfo{FuncNameAdd, addEvalFunc, twoNumberArgs},

	FuncNameProduct: FunctionInfo{FuncNameProduct, productEvalFunc, oneOrMoreNumberArgs},
	FuncNameMax:     FunctionInfo{FuncNameMax, maxEvalFunc, oneOrMoreNumberArgs},
	FuncNameMin:     FunctionInfo{FuncNameMin, minEvalFunc, oneOrMoreNumberArgs},

	FuncNameMinus:       FunctionInfo{FuncNameMinus, minusEvalFunc, twoNumberArgs},
	FuncNameDivide:      FunctionInfo{FuncNameMinus, divideEvalFunc, twoNumberArgs},
	FuncNameMultiply:    FunctionInfo{FuncNameMultiply, multiplyEvalFunc, twoNumberArgs},
	FuncNamePower:       FunctionInfo{FuncNamePower, powerEvalFunc, twoNumberArgs},
	FuncNameConcat:      FunctionInfo{FuncNameConcat, concatEvalFunc, oneOrMoreTextArgs},
	FuncNameDateAdd:     FunctionInfo{FuncNameDateAdd, dateAddEvalFunc, validDateAddArgs},
	FuncNameDaysBetween: FunctionInfo{FuncNameDaysBetween, daysBetweenEvalFunc, twoTimeArgsNumberResult},
	FuncNameGreaterThan: FunctionInfo{FuncNameGreaterThan, greaterThanEvalFunc, twoNumberArgsBooleanResult},
	FuncNameIf:          FunctionInfo{FuncNameIf, ifEvalFunc, validIfArgs},
}
