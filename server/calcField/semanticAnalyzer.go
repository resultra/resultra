package calcField

import (
	"appengine"
	"fmt"
	"resultra/datasheet/server/field"
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
		funcInfo, funcInfoFound := context.definedFuncs[eqnNode.FuncName]
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

func analyzeSemantics(appEngContext appengine.Context, rootEqnNode *EquationNode) (*semanticAnalysisResult, error) {

	// TODO - Check the overall result type against what is epected for the formula. This depends on the type
	// of field the formula is assigned to.

	context := semanticAnalysisContext{
		fieldIDsVisited: []string{},
		appEngContext:   appEngContext,
		definedFuncs:    CalcFieldDefinedFuncs}

	return analyzeEqnNode(&context, rootEqnNode)
}
