package calcField

import (
	"fmt"
	"resultra/datasheet/server/field"
	"time"
)

func oneOrMoreTimeArgs(params FuncSemAnalysisParams) (*semanticAnalysisResult, error) {

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
		if analyzeResult.resultType != field.FieldTypeTime {
			argErrors = append(
				argErrors, fmt.Sprintf("Invalid argument type for argument %v of function %v. Expecting a date/time argument", argNum, params.funcName))
		}
	}

	return &semanticAnalysisResult{analyzeErrors: argErrors, resultType: field.FieldTypeNumber}, nil

}

func twoTimeArgsNumberResult(params FuncSemAnalysisParams) (*semanticAnalysisResult, error) {
	if len(params.funcArgs) != 2 {
		// Even though there's an errors, based upon the function type we know it will return a number. This
		// allows semantic analysis to continue, even though there might be some errors.
		errMsgs := []string{fmt.Sprintf("Not enough arguments to function %v, expecting at at least 1 argument", params.funcName)}
		return &semanticAnalysisResult{analyzeErrors: errMsgs, resultType: field.FieldTypeNumber}, nil
	}
	return oneOrMoreTimeArgs(params)
}

type twoTimeArgEvalFunc func(time1 time.Time, time2 time.Time) (*EquationResult, error)

func evalTwoTimeArgFunc(evalContext *EqnEvalContext, funcArgs []*EquationNode, timeEvalFunc twoTimeArgEvalFunc) (*EquationResult, error) {
	if len(funcArgs) != 2 {
		return nil, fmt.Errorf("Expecting 2 arguments, got %v", len(funcArgs))
	}
	arg1Eqn := funcArgs[0]
	arg1Result, arg1Err := arg1Eqn.EvalEqn(evalContext)
	if arg1Err != nil {
		return nil, fmt.Errorf("Error evaluating argument # 1: arg=%+v, error %v", arg1Eqn, arg1Err)
	} else if arg1Result.IsUndefined() {
		// If an undefined result is returned, return immediately and propogate the undefined
		// result value up through the equation evaluation.
		return arg1Result, nil
	} else if arg1TimeResult, validateErr := arg1Result.GetTimeResult(); validateErr != nil {
		return nil, fmt.Errorf("Invalid result found while evaluating argument 1: arg=%+v, error = %v", arg1Eqn, validateErr)
	} else {
		arg2Eqn := funcArgs[1]
		arg2Result, arg2Err := arg2Eqn.EvalEqn(evalContext)
		if arg2Err != nil {
			return nil, fmt.Errorf("Error evaluating argument # 2: arg=%+v, error %v", arg2Eqn, arg2Err)
		} else if arg2Result.IsUndefined() {
			// If an undefined result is returned, return immediately and propogate the undefined
			// result value up through the equation evaluation.
			return arg2Result, nil
		} else if arg2TimeResult, validateErr := arg2Result.GetTimeResult(); validateErr != nil {
			return nil, fmt.Errorf("Invalid result found while evaluating argument 2: arg=%+v, error = %v", arg2Eqn, validateErr)
		} else {
			return timeEvalFunc(arg1TimeResult, arg2TimeResult)
		}
	}
}
