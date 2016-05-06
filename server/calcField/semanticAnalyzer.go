package calcField

import (
	"appengine"
	"fmt"
	"resultra/datasheet/server/field"
	"strings"
)

type semanticAnalysisContext struct {
	fieldIDsVisited []string // for detecting cycles
	appEngContext   appengine.Context
	definedFuncs    FuncNameFuncInfoMap
}

type semanticAnalysisResult struct {
	analyzeErrors []string
	resultType    string
}

func (semAnalRes semanticAnalysisResult) hasErrors() bool {
	if len(semAnalRes.analyzeErrors) > 0 {
		return true
	} else {
		return false
	}
}

func newErrorAnalysisResult(errMsg string) *semanticAnalysisResult {
	return &semanticAnalysisResult{analyzeErrors: []string{errMsg}, resultType: ""}
}

func newTypedAnalysisResult(resultType string) *semanticAnalysisResult {
	return &semanticAnalysisResult{analyzeErrors: nil, resultType: resultType}
}

func analyzeEqnNode(context *semanticAnalysisContext, eqnNode *EquationNode) (*semanticAnalysisResult, error) {

	if len(eqnNode.FuncName) > 0 {
		// Function names are case-insenstive, so check if the function is
		// defined using upper case.
		upperCaseFuncName := strings.ToUpper(eqnNode.FuncName)
		funcInfo, funcInfoFound := context.definedFuncs[upperCaseFuncName]
		if !funcInfoFound {
			return newErrorAnalysisResult(fmt.Sprintf("Undefined function: %v", eqnNode.FuncName)), nil
		}
		funcSemAnalysisParams := FuncSemAnalysisParams{
			context:  context,
			funcName: eqnNode.FuncName,
			funcArgs: eqnNode.FuncArgs}
		return funcInfo.semAnalysisFunc(funcSemAnalysisParams)

		// TODO - pass argument list to specific analyzer for the function

	} else if len(eqnNode.FieldID) > 0 {

		// Equation references a field. The user sets up these references similar to
		// spreadsheet references using a "reference name", but it is stored in the equation
		// node as a unique field ID. This field reference could be a calculated field or a
		// non-calculated field with literal values.
		// TODO - Once the Field type has a parent, don't use an individual database
		// lookup for each field (database only has strong consistency when
		// entities have a parent.
		fieldRef, err := field.GetFieldRef(context.appEngContext, eqnNode.FieldID)
		if err != nil {
			return nil, fmt.Errorf("Failure retrieving referenced field: %v", err)
		} else {
			return newTypedAnalysisResult(fieldRef.FieldInfo.Type), nil
		}

	} else if eqnNode.TextVal != nil {
		return newTypedAnalysisResult(field.FieldTypeText), nil
	} else if eqnNode.NumberVal != nil {
		return newTypedAnalysisResult(field.FieldTypeNumber), nil
	} else {
		return nil, fmt.Errorf("Unknown error: unexpected result type")
	}

	return nil, fmt.Errorf("analyzeEqnNode: Unknown error: unhandled equation type: %+v", eqnNode)
}

func analyzeSemantics(compileParams formulaCompileParams, rootEqnNode *EquationNode) (*semanticAnalysisResult, error) {

	context := semanticAnalysisContext{
		fieldIDsVisited: []string{},
		appEngContext:   compileParams.appEngContext,
		definedFuncs:    CalcFieldDefinedFuncs}

	// Check the top-level/overall result type to see that it matches the expected type (e.g., bool, number, text)
	analyzeResult, analyzeErr := analyzeEqnNode(&context, rootEqnNode)
	if analyzeResult.hasErrors() {
		return analyzeResult, analyzeErr
	} else {
		if len(analyzeResult.resultType) == 0 {
			return nil, fmt.Errorf("analyzeSemantics: Unexpected results: no errors but not result type either")
		} else {
			if analyzeResult.resultType != compileParams.expectedResultType {
				errMsg := fmt.Sprintf("Unexpected formula result. Expecting %v, but formula returns %v",
					compileParams.expectedResultType, analyzeResult.resultType)
				analyzeResult.analyzeErrors = append(analyzeResult.analyzeErrors, errMsg)
				return analyzeResult, analyzeErr
			} else {
				// No error - result type matches expected result type
				return analyzeResult, analyzeErr
			}

		}
	}

}
