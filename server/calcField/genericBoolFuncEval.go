package calcField

import (
	"fmt"
	"resultra/datasheet/server/field"
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
