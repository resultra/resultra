package datamodel

import (
	"appengine"
	"fmt"
	"log"
)

type EqnEvalFunc func(evalContext *EqnEvalContext, funcArgs []EquationNode) (*EquationResult, error)

type FunctionInfo struct {
	funcName   string
	resultType string
	evalFunc   EqnEvalFunc
}

type FuncNameFuncInfoMap map[string]FunctionInfo

type EqnEvalContext struct {
	appEngContext appengine.Context
	definedFuncs  FuncNameFuncInfoMap

	// Record into which the results will be calculated. This is also the record
	// which is referenced for field values, in the case a calculated field references
	// other fields.
	resultRecord RecordRef
}

// Get the literal value from a number field, or undefined if it doesn't exist.
func (recordRef RecordRef) GetNumberRecordEqnResult(appEngContext appengine.Context, fieldID string) (*EquationResult, error) {

	if fieldValidateErr := validateFieldForRecordValue(appEngContext, fieldID, fieldTypeNumber); fieldValidateErr != nil {
		return nil, fmt.Errorf("Can't get value from record = %+v and fieldID = %v: "+
			"Can't validate field with value type: validation error = %v", recordRef, fieldID, fieldValidateErr)
	}

	val, foundVal := recordRef.FieldValues[fieldID]
	if !foundVal {
		// Note this is the only place which will return an undefined equation result (along with the
		// the similar function for other value types). This is because record values for non-calculated
		// fields is the only place where an undefined (or blank) value could originate, because a
		// user hasn't entered a value yet for the field.
		log.Printf("GetNumberRecordEqnResult: Undefined equation result for field: %v", fieldID)
		return undefinedEqnResult(), nil
	} else {
		if numberVal, foundNumber := val.(float64); !foundNumber {
			return nil, fmt.Errorf("Type mismatch retrieving value from record with ID = %v, field = %v:"+
				" expecting number, got %v", recordRef.RecordID, fieldID, val)
		} else {
			return numberEqnResult(numberVal), nil
		}
	} // else (if found a value for the given field ID)
}

// Get the literal value from a text field, or undefined if it doesn't exist.
func (recordRef RecordRef) GetTextRecordEqnResult(appEngContext appengine.Context, fieldID string) (*EquationResult, error) {

	if fieldValidateErr := validateFieldForRecordValue(appEngContext, fieldID, fieldTypeText); fieldValidateErr != nil {
		return nil, fmt.Errorf("Can't get value from record = %+v and fieldID = %v: "+
			"Can't validate field with value type: validation error = %v", recordRef, fieldID, fieldValidateErr)
	}

	val, foundVal := recordRef.FieldValues[fieldID]
	if !foundVal {
		// Note this is the only place which will return an undefined equation result (along with the
		// the similar function for other value types). This is because record values for non-calculated
		// fields is the only place where an undefined (or blank) value could originate, because a
		// user hasn't entered a value yet for the field.
		log.Printf("GetTextRecordEqnResult: Undefined equation result for field: %v", fieldID)
		return undefinedEqnResult(), nil
	} else {
		if textVal, foundText := val.(string); !foundText {
			return nil, fmt.Errorf("Type mismatch retrieving value from record with ID = %v, field = %v:"+
				" expecting string, got %v", recordRef.RecordID, fieldID, val)

		} else {
			return textEqnResult(textVal), nil
		}
	} // else (if found a value for the given field ID)
}

func (fieldRef FieldRef) evalEqn(evalContext *EqnEvalContext) (*EquationResult, error) {

	field := fieldRef.FieldInfo

	if field.IsCalcField {
		// Calculated field - return the result of the calculation
		decodedEqn, decodeErr := decodeEquation(field.CalcFieldEqn)
		if decodeErr != nil {
			return nil, fmt.Errorf("Failure decoding equation for evaluation: %v", decodeErr)
		} else {
			calcFieldResult, calcFieldErr := decodedEqn.evalEqn(evalContext)
			if calcFieldErr != nil {
				return calcFieldResult, calcFieldErr
			} else if calcFieldResult.isUndefined() {
				// If an undefined result is returned, return immediately and propogate the undefined
				// result value up through the equation evaluation.
				return calcFieldResult, nil
			} else if calcFieldResult.ResultType != field.Type {
				return nil, fmt.Errorf("evalEqn: type mismatch in result calculated for field: "+
					" expecting %v, got %v: field = %+v", field.Type, calcFieldResult.ResultType, field)
			} else {
				return calcFieldResult, nil
			}
		}
	} else { // literal field values
		switch field.Type {
		case fieldTypeText:
			return evalContext.resultRecord.GetTextRecordEqnResult(
				evalContext.appEngContext, fieldRef.FieldID)
		case fieldTypeNumber:
			return evalContext.resultRecord.GetNumberRecordEqnResult(
				evalContext.appEngContext, fieldRef.FieldID)
			//		case fieldTypeDate:
		default:
			return nil, fmt.Errorf("Unknown field result type: %v", field.Type)

		} // switch

	} // field value is a literal, just return it
}

func (equation EquationNode) evalEqn(evalContext *EqnEvalContext) (*EquationResult, error) {

	if len(equation.FuncName) > 0 {
		funcInfo, funcInfoFound := evalContext.definedFuncs[equation.FuncName]
		if !funcInfoFound {
			return nil, fmt.Errorf("EvalEqn: Undefined function: %v", equation.FuncName)
		}
		if funcEvalResult, funcErr := funcInfo.evalFunc(evalContext, equation.FuncArgs); funcErr != nil {
			// Function failed to compute
			return nil, funcErr
		} else {
			// TBD - Is it necessary to check the result type from the function
			// to ensure it matches the expected result type of this equation.
			return funcEvalResult, nil
		}

	} else if len(equation.FieldID) > 0 {

		// Equation references a field. The user sets up these references similar to
		// spreadsheet references using a "reference name", but it is stored in the equation
		// node as a unique field ID. This field reference could be a calculated field or a
		// non-calculated field with literal values.
		// TODO - Once the Field type has a parent, don't use an individual database
		// lookup for each field (database only has strong consistency when
		// entities have a parent.
		fieldRef, err := GetField(evalContext.appEngContext, GetFieldParams{equation.FieldID})
		if err != nil {
			return nil, fmt.Errorf("evalEqn: failure retrieving referenced field: %+v", err)
		} else {
			return fieldRef.evalEqn(evalContext)
		}

	} else if equation.TextVal != nil {
		// Text literal given directly in the equation itself  (rather than a field value)
		return textEqnResult(*equation.TextVal), nil
	} else if equation.NumberVal != nil {
		// Number literal given directly in the equation itself (rather than a field value)
		return numberEqnResult(*equation.NumberVal), nil
	} else {
		return nil, fmt.Errorf("evalEqn: malformed calculated field equation : system error: %+v", equation)
	}

}

// updateCalcFieldValues is (currently) the top-most entry point into the calculated field
// equation evaluation functionality.
func updateCalcFieldValues(appEngContext appengine.Context, recordRef *RecordRef) error {
	fieldRefs, getFieldErr := GetAllFieldRefs(appEngContext)
	if getFieldErr != nil {
		return fmt.Errorf("Error updating field values - can't get fields: %v", getFieldErr)
	}

	for _, fieldRef := range fieldRefs {
		if fieldRef.FieldInfo.IsCalcField {

			log.Printf("updateCalcFieldValues: Updating calculated field %v", fieldRef.FieldInfo.RefName)

			rootFieldEqnNode, decodeErr := decodeEquation(fieldRef.FieldInfo.CalcFieldEqn)
			if decodeErr != nil {
				return fmt.Errorf("Can't decode equation for field = %+v: decode error = %v", fieldRef, decodeErr)
			}

			fieldEqnResult, evalErr := rootFieldEqnNode.evalEqn(
				&EqnEvalContext{appEngContext, calcFieldDefinedFuncs, *recordRef})
			if evalErr != nil {
				return fmt.Errorf("Unexpected error evaluating equation for field=%v: error=%+v",
					fieldRef.FieldID, evalErr)
			} else if fieldEqnResult.isUndefined() {
				// If the evaluation result is undefined, don't proceed to actually set the value for the
				// calculated field (i.e., it will remain blank/undefined). An undefined value is a valid
				// condition when not all the depended upon field's values are defined.
				log.Printf("Undefined field eqn result for field = %v", fieldRef.FieldInfo.RefName)
				// TODO - Don't return nil - need to continue processing other calculated fields even
				// if one field happens to be undefined.
				return nil
			}

			if fieldEqnResult.ResultType != fieldRef.FieldInfo.Type {
				return fmt.Errorf("Error evaluating equation for field=%+v: eqn=%+v: type mismatch on equation result:"+
					"expected=%v, got=%v", fieldRef, rootFieldEqnNode,
					fieldRef.FieldInfo.Type, fieldEqnResult.ResultType)
			} // if type mismatch between calculated equation result and calculated field's type

			switch fieldRef.FieldInfo.Type {

			case fieldTypeText:
				textResult, textResultErr := fieldEqnResult.getTextResult()
				if textResultErr != nil {
					return fmt.Errorf("Unexpected error evaluating equation for field=%v: eqn=%+v: error=%v "+
						"unexpected error getting number result: raw result=%+v",
						fieldRef.FieldID, rootFieldEqnNode, textResultErr, fieldEqnResult)

				}
				log.Printf("updateCalcFieldValues: Setting calculated field value: field=%v, value=%v", fieldRef.FieldInfo.RefName, textResult)
				recordRef.FieldValues[fieldRef.FieldID] = textResult

			case fieldTypeNumber:
				numberResult, numberResultErr := fieldEqnResult.getNumberResult()
				if numberResultErr != nil {
					return fmt.Errorf("Unexpected error evaluating equation for field=%v: eqn=%+v: error=%v "+
						"unexpected error getting number result: raw result=%+v",
						fieldRef.FieldID, rootFieldEqnNode, numberResultErr, fieldEqnResult)

				}
				log.Printf("updateCalcFieldValues: Setting calculated field value: field=%v, value=%v", fieldRef.FieldInfo.RefName, numberResult)
				recordRef.FieldValues[fieldRef.FieldID] = numberResult
				// TODO case fieldTypeDate

			default:
				return fmt.Errorf("Unexpected error evaluating equation for field=%+v: eqn=%+v: unsupported field type %v",
					fieldRef, rootFieldEqnNode, evalErr, fieldRef.FieldInfo.Type)
			} // switch field type

		} // If calculated field
	} // for each fieldRef
	return nil
}
