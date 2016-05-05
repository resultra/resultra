package calcField

import (
	"bytes"
	"fmt"
	"log"
	"resultra/datasheet/server/field"
)

type FuncSemAnalysisParams struct {
	context  *semanticAnalysisContext
	funcName string
	funcArgs []EquationNode
}

type EqnEvalFunc func(evalContext *EqnEvalContext, funcArgs []EquationNode) (*EquationResult, error)
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
		analyzeResult, analyzeErr := analyzeEqnNode(params.context, &argEqnNode)
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
		analyzeResult, analyzeErr := analyzeEqnNode(params.context, &argEqnNode)
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
	FuncNameSum:     FunctionInfo{FuncNameSum, field.FieldTypeNumber, sumEvalFunc, oneOrMoreNumberArgs},
	FuncNameProduct: FunctionInfo{FuncNameProduct, field.FieldTypeNumber, productEvalFunc, oneOrMoreNumberArgs},
	FuncNameConcat:  FunctionInfo{FuncNameConcat, field.FieldTypeText, concatEvalFunc, oneOrMoreTextArgs},
}
