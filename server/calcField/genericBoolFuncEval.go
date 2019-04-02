// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package calcField

import (
	"fmt"
	"github.com/resultra/resultra/server/field"
)

func twoNumberArgsBooleanResult(params FuncSemAnalysisParams) (*semanticAnalysisResult, error) {
	if len(params.funcArgs) != 2 {
		errMsgs := []string{fmt.Sprintf("Expecting 2 numerical arguments to function %v", params.funcName)}
		return &semanticAnalysisResult{analyzeErrors: errMsgs, resultType: field.FieldTypeBool}, nil
	}
	errMsgs := []string{}
	return &semanticAnalysisResult{analyzeErrors: errMsgs, resultType: field.FieldTypeBool}, nil
}

func oneBoolArg(params FuncSemAnalysisParams) (*semanticAnalysisResult, error) {
	if len(params.funcArgs) != 1 {
		errMsgs := []string{fmt.Sprintf("Expecting 1 boolean arguments to function %v", params.funcName)}
		return &semanticAnalysisResult{analyzeErrors: errMsgs, resultType: field.FieldTypeBool}, nil
	}

	argErrors := []string{}

	arg1EqnNode := params.funcArgs[0]
	arg1AnalyzeResult, analyzeErr := analyzeEqnNode(params.context, arg1EqnNode)
	if analyzeErr != nil {
		return nil, analyzeErr
	}
	if arg1AnalyzeResult.resultType != field.FieldTypeBool {
		argErrors = append(
			argErrors, fmt.Sprintf("Invalid argument type for argument 1 of function %v. Expecting boolean (logical)", params.funcName))
	}

	return &semanticAnalysisResult{analyzeErrors: argErrors, resultType: field.FieldTypeBool}, nil

}

func oneBoolTimeResultArg(params FuncSemAnalysisParams) (*semanticAnalysisResult, error) {
	if len(params.funcArgs) != 1 {
		errMsgs := []string{fmt.Sprintf("Expecting 1 boolean arguments to function %v", params.funcName)}
		return &semanticAnalysisResult{analyzeErrors: errMsgs, resultType: field.FieldTypeBool}, nil
	}

	argErrors := []string{}

	arg1EqnNode := params.funcArgs[0]
	arg1AnalyzeResult, analyzeErr := analyzeEqnNode(params.context, arg1EqnNode)
	if analyzeErr != nil {
		return nil, analyzeErr
	}
	if arg1AnalyzeResult.resultType != field.FieldTypeBool {
		argErrors = append(
			argErrors, fmt.Sprintf("Invalid argument type for argument 1 of function %v. Expecting boolean (logical)", params.funcName))
	}

	return &semanticAnalysisResult{analyzeErrors: argErrors, resultType: field.FieldTypeTime}, nil

}

func anySingleArgBoolResult(params FuncSemAnalysisParams) (*semanticAnalysisResult, error) {
	if len(params.funcArgs) != 1 {
		errMsgs := []string{fmt.Sprintf("Expecting 1 argument to function %v", params.funcName)}
		return &semanticAnalysisResult{analyzeErrors: errMsgs, resultType: field.FieldTypeBool}, nil
	}

	argErrors := []string{}
	return &semanticAnalysisResult{analyzeErrors: argErrors, resultType: field.FieldTypeBool}, nil
}

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
			argErrors, fmt.Sprintf("Invalid argument type for argument 1 of function %v. Expecting boolean (logical)", params.funcName))
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

	return &semanticAnalysisResult{analyzeErrors: argErrors, resultType: arg2AnalyzeResult.resultType}, nil
}

func oneOrMoreBoolArgs(params FuncSemAnalysisParams) (*semanticAnalysisResult, error) {

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
		if analyzeResult.resultType != field.FieldTypeBool {
			argErrors = append(
				argErrors, fmt.Sprintf("Invalid argument type for argument %v of function %v. Expecting a number", argNum, params.funcName))
		}
	}

	return &semanticAnalysisResult{analyzeErrors: argErrors, resultType: field.FieldTypeBool}, nil

}

type oneOrMoreBoolArgFunc func(args []bool) (*EquationResult, error)

func evalOneOrMoreBoolArgFunc(evalContext *EqnEvalContext, funcArgs []*EquationNode,
	boolEvalFunc oneOrMoreBoolArgFunc, requireAllArgsDefined bool) (*EquationResult, error) {

	if len(funcArgs) < 1 {
		return nil, fmt.Errorf("Not enough arguments given to function")
	}

	boolArgs := []bool{}
	for argIndex, argEqn := range funcArgs {
		argResult, argErr := argEqn.EvalEqn(evalContext)
		if argErr != nil {
			return nil, fmt.Errorf("Error evaluating argument # %v: arg=%+v, error %v",
				argIndex, argEqn, argErr)
		} else if argResult.IsUndefined() {
			// No-op - undefined results aren't passed along to the function for evaluation
			if requireAllArgsDefined {
				// TBD - this behavior is expected by the ALLTRUE() function; i.e. if all the arguments
				// are not defined and return true, then a false value is returned. There might be a cleaner
				// way to handle this generically for ALLTRUE() and ANYTRUE()
				return boolEqnResult(false), nil
			}
		} else if boolResult, validateErr := argResult.GetBoolResult(); validateErr != nil {
			return nil, fmt.Errorf("Invalid result found while evaluating argument # %v: arg=%+v, error = %v",
				argIndex, argEqn, validateErr)
		} else {
			boolArgs = append(boolArgs, boolResult)
		}
	}

	return boolEvalFunc(boolArgs)

}
