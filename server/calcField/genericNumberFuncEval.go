package calcField

import (
	"fmt"
	"log"
	"resultra/datasheet/server/field"
)

func zeroNumberArgs(params FuncSemAnalysisParams) (*semanticAnalysisResult, error) {
	if len(params.funcArgs) != 0 {
		errMsgs := []string{fmt.Sprintf("Expecting 0 arguments to function %v", params.funcName)}
		return &semanticAnalysisResult{analyzeErrors: errMsgs, resultType: field.FieldTypeNumber}, nil
	}

	errMsgs := []string{}
	return &semanticAnalysisResult{analyzeErrors: errMsgs, resultType: field.FieldTypeNumber}, nil
}

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

type twoNumberArgEvalFunc func(number1, number2 float64) (*EquationResult, error)

func evalTwoNumberArgFunc(evalContext *EqnEvalContext, funcArgs []*EquationNode, numberEvalFunc twoNumberArgEvalFunc) (*EquationResult, error) {

	if len(funcArgs) != 2 {
		return nil, fmt.Errorf("MINUS() - Expecting 2 arguments, got %v", len(funcArgs))
	}

	arg1Eqn := funcArgs[0]
	arg1Result, arg1Err := arg1Eqn.EvalEqn(evalContext)
	if arg1Err != nil {
		return nil, fmt.Errorf("MINUS(): Error evaluating argument # 1: arg=%+v, error %v", arg1Eqn, arg1Err)
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

			return numberEvalFunc(arg1NumberResult, arg2NumberResult)
		}
	}

}

type oneOrMoreNumberArgFunc func(args []float64) (*EquationResult, error)

func evalOneOrMoreNumberArgFunc(evalContext *EqnEvalContext, funcArgs []*EquationNode,
	numberEvalFunc oneOrMoreNumberArgFunc) (*EquationResult, error) {

	if len(funcArgs) < 1 {
		return nil, fmt.Errorf("Not enough arguments given to function")
	}

	numberArgs := []float64{}
	for argIndex, argEqn := range funcArgs {
		argResult, argErr := argEqn.EvalEqn(evalContext)
		if argErr != nil {
			return nil, fmt.Errorf("Error evaluating argument # %v: arg=%+v, error %v",
				argIndex, argEqn, argErr)
		} else if argResult.IsUndefined() {
			// If any of the arguments are undefined, then the entire result is undefined.
			// If another behavior is needed, then an explicit default value is needed.
			return undefinedEqnResult(), nil
		} else if numberResult, validateErr := argResult.GetNumberResult(); validateErr != nil {
			return nil, fmt.Errorf("Invalid result found while evaluating argument # %v: arg=%+v, error = %v",
				argIndex, argEqn, validateErr)
		} else {
			numberArgs = append(numberArgs, numberResult)
		}
	}

	if len(numberArgs) <= 0 {
		return undefinedEqnResult(), nil
	} else {
		return numberEvalFunc(numberArgs)
	}

}
