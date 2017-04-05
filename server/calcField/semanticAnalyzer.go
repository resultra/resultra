package calcField

import (
	"fmt"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/global"
	"strings"
)

type semanticAnalysisContext struct {
	resultFieldID string // for detecting cycles
	definedFuncs  FuncNameFuncInfoMap
	globalIndex   global.GlobalIDGlobalIndex
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

func newUndefinedAnalysisResult() *semanticAnalysisResult {
	return &semanticAnalysisResult{analyzeErrors: nil, resultType: ""}
}

func checkEqnCycles(context *semanticAnalysisContext, eqnNode *EquationNode) (bool, error) {
	// The only types of equation nodes which need to be checked are field references
	// and other "non-terminal" nodes in the equation tree. If the equation refers
	// to a value literal, there is no need to check for cycles.
	// All the other elements in the compiled formulas equation tree refere to
	if len(eqnNode.FieldID) > 0 {
		eqnField, fieldErr := field.GetField(eqnNode.FieldID)
		if fieldErr != nil {
			return false, fmt.Errorf("Failure retrieving referenced field: %v", fieldErr)
		} else {
			return checkFieldCycles(context, eqnField)
		}
	} else if len(eqnNode.FuncName) > 0 {
		for _, funcArgEqn := range eqnNode.FuncArgs {
			argHasCycle, argCycleErr := checkEqnCycles(context, funcArgEqn)
			if argCycleErr != nil {
				return false, argCycleErr
			} else if argHasCycle {
				return true, nil
			}
		} // for each argument in the function call
		// No cycle dependencies in the function call
		return false, nil
	} else {
		// The remaining types of equation nodes are terminal/literal values, so there's no
		// need to check for cycles.
		return false, nil
	}
}

// TODO - While doing the recursion to check for cycles, keep a "chain" of field references.
// If a cycle is found, this can be passed back to give the user more information to resolve
// the cycle: e.g. "FieldA -> FieldB -> FieldC -> FieldA"
func checkFieldCycles(context *semanticAnalysisContext, checkField *field.Field) (bool, error) {

	if !checkField.IsCalcField {
		// If the fiels is not a calculated field, it is "terminal" and contains a literal
		// value.
		return false, nil
	}

	if checkField.FieldID == context.resultFieldID {
		// Cycle found
		return true, nil
	}

	// Retrieve the compiled equation for the field, then recursively check if there are cycles
	// in that equation. There is no need for full-blown semantic analysis, since semantic analysis
	// was already performed before the equation was saved.
	decodedEqn, decodeErr := decodeEquation(checkField.CalcFieldEqn)
	if decodeErr != nil {
		return false, fmt.Errorf("Failure decoding equation for formula cycle detection: %v", decodeErr)
	} else {
		return checkEqnCycles(context, decodedEqn)
	}

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

		// Perform semantic analysis which is specific to the function (e.g. check number of arguments, type of arguments)
		return funcInfo.semAnalysisFunc(funcSemAnalysisParams)

	} else if len(eqnNode.FieldID) > 0 {

		// Equation references a field. The user sets up these references similar to
		// spreadsheet references using a "reference name", but it is stored in the equation
		// node as a unique field ID. This field reference could be a calculated field or a
		// non-calculated field with literal values.
		// TODO - Once the Field type has a parent, don't use an individual database
		// lookup for each field (database only has strong consistency when
		// entities have a parent.
		eqnField, err := field.GetField(eqnNode.FieldID)
		if err != nil {
			return nil, fmt.Errorf("Failure retrieving referenced field: %v", err)
		} else {
			if len(context.resultFieldID) == 0 {
				// If the resultFieldID is empty, then the compilation is being done for a new
				// field, so no circular reference checking is needed.
				return newTypedAnalysisResult(eqnField.Type), nil
			} else if eqnNode.FieldID == context.resultFieldID {
				// Check if the formula refers to itself
				return newErrorAnalysisResult(fmt.Sprintf("Circular reference: a calculated field cannot refer to itself using [%v]",
					eqnField.RefName)), nil
			} else {
				if eqnField.IsCalcField {
					// The field reference is to a calculated field, so we need to retrieve its equation and
					// check there are no direct or indirect references to the field with id = context.resultFieldID
					cycleFound, cycleCheckErr := checkFieldCycles(context, eqnField)
					if cycleCheckErr != nil {
						return nil, cycleCheckErr
					} else {
						if cycleFound {
							return newErrorAnalysisResult(
								fmt.Sprintf("Circular reference: a calculated field's cannot refer back to itself indirectly through [%v]",
									eqnField.RefName)), nil
						} else {
							return newTypedAnalysisResult(eqnField.Type), nil
						}
					}
				} else {
					// The field reference is to a non-calculated field. Since there is an literal value in the field
					// (i.e., no formula), no further checking for cycles is needed ... just return the fields type.
					return newTypedAnalysisResult(eqnField.Type), nil
				}
			}
		}

	} else if len(eqnNode.GlobalID) > 0 {
		globalInfo, globalInfoFound := (context.globalIndex)[eqnNode.GlobalID]
		if !globalInfoFound {
			return nil, fmt.Errorf("Failure retrieving referenced global %v", eqnNode.GlobalID)
		}
		switch globalInfo.Type {
		case global.GlobalTypeText:
			return newTypedAnalysisResult(field.FieldTypeText), nil
		case global.GlobalTypeNumber:
			return newTypedAnalysisResult(field.FieldTypeNumber), nil
		default:
			return nil, fmt.Errorf("Unknown global result type: %v", globalInfo.Type)

		}

		return nil, fmt.Errorf("Unknown error: global result type not supported yet")
	} else if eqnNode.TextVal != nil {
		return newTypedAnalysisResult(field.FieldTypeText), nil
	} else if eqnNode.NumberVal != nil {
		return newTypedAnalysisResult(field.FieldTypeNumber), nil
	} else if eqnNode.TimeVal != nil {
		return newTypedAnalysisResult(field.FieldTypeTime), nil
	} else {
		// If the EquationNode doesn't have any of it's properties defined, it is considered undefined.
		// This corresponds to an empty formula, which will always return an undefined result.
		return newUndefinedAnalysisResult(), nil
	}

	return nil, fmt.Errorf("analyzeEqnNode: Unknown error: unhandled equation type: %+v", eqnNode)
}

func analyzeSemantics(compileParams formulaCompileParams, rootEqnNode *EquationNode) (*semanticAnalysisResult, error) {

	globalIndex, globalIndexErr := global.GetIndexedGlobals(compileParams.databaseID)
	if globalIndexErr != nil {
		return nil, fmt.Errorf("analyzeSemantics: Unable to retrieve indexed globals: error =%v", globalIndexErr)
	}

	context := semanticAnalysisContext{
		resultFieldID: compileParams.resultFieldID,
		globalIndex:   globalIndex,
		definedFuncs:  CalcFieldDefinedFuncs}

	// Check the top-level/overall result type to see that it matches the expected type (e.g., bool, number, text)
	analyzeResult, analyzeErr := analyzeEqnNode(&context, rootEqnNode)
	if analyzeErr != nil {
		return nil, fmt.Errorf("analyzeSemantics: Unexpected results: analyze failed with error = %v", analyzeErr)
	} else if analyzeResult.hasErrors() {
		return analyzeResult, analyzeErr
	} else {
		if len(analyzeResult.resultType) == 0 {
			// The formula is an empty formula which will always return an undefined result.
			return analyzeResult, analyzeErr
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
