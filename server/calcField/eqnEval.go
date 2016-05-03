package calcField

import (
	"appengine"
	"fmt"
	"log"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/generic/datastoreWrapper"
	"resultra/datasheet/server/record"
	"resultra/datasheet/server/table"
)

type EqnEvalFunc func(evalContext *EqnEvalContext, funcArgs []EquationNode) (*EquationResult, error)

type FunctionInfo struct {
	funcName   string
	resultType string
	evalFunc   EqnEvalFunc
}

type FuncNameFuncInfoMap map[string]FunctionInfo

type EqnEvalContext struct {
	AppEngContext appengine.Context
	DefinedFuncs  FuncNameFuncInfoMap

	// Record into which the results will be calculated. This is also the record
	// which is referenced for field values, in the case a calculated field references
	// other fields.
	ResultRecord record.RecordRef
}

// Get the literal value from a number field, or undefined if it doesn't exist.
func GetNumberRecordEqnResult(appEngContext appengine.Context, recordRef record.RecordRef, fieldID string) (*EquationResult, error) {

	// Since the calculated field values are stored in the Record just the same as values directly entered by end-users,
	// it is OK to retrieve the literal values from these fields just like non-calculated fields.
	allowCalcField := true
	if fieldValidateErr := record.ValidateFieldForRecordValue(appEngContext, fieldID, field.FieldTypeNumber, allowCalcField); fieldValidateErr != nil {
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
func GetTextRecordEqnResult(appEngContext appengine.Context, recordRef record.RecordRef, fieldID string) (*EquationResult, error) {

	// Since the calculated field values are stored in the Record just the same as values directly entered by end-users,
	// it is OK to retrieve the literal values from these fields just like non-calculated fields.
	allowCalcField := true
	if fieldValidateErr := record.ValidateFieldForRecordValue(appEngContext, fieldID, field.FieldTypeText, allowCalcField); fieldValidateErr != nil {
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

func EvalEqn(evalContext *EqnEvalContext, fieldRef field.FieldRef) (*EquationResult, error) {

	fieldInfo := fieldRef.FieldInfo

	if fieldInfo.IsCalcField {
		// Calculated field - return the result of the calculation
		decodedEqn, decodeErr := decodeEquation(fieldInfo.CalcFieldEqn)
		if decodeErr != nil {
			return nil, fmt.Errorf("Failure decoding equation for evaluation: %v", decodeErr)
		} else {
			calcFieldResult, calcFieldErr := decodedEqn.EvalEqn(evalContext)
			if calcFieldErr != nil {
				return calcFieldResult, calcFieldErr
			} else if calcFieldResult.IsUndefined() {
				// If an undefined result is returned, return immediately and propogate the undefined
				// result value up through the equation evaluation.
				return calcFieldResult, nil
			} else if calcFieldResult.ResultType != fieldInfo.Type {
				return nil, fmt.Errorf("EvalEqn: type mismatch in result calculated for field: "+
					" expecting %v, got %v: field = %+v", fieldInfo.Type, calcFieldResult.ResultType, fieldInfo)
			} else {
				return calcFieldResult, nil
			}
		}
	} else { // literal field values
		switch fieldInfo.Type {
		case field.FieldTypeText:
			return GetTextRecordEqnResult(
				evalContext.AppEngContext, evalContext.ResultRecord, fieldRef.FieldID)
		case field.FieldTypeNumber:
			return GetNumberRecordEqnResult(
				evalContext.AppEngContext, evalContext.ResultRecord, fieldRef.FieldID)
			//		case FieldTypeDate:
		default:
			return nil, fmt.Errorf("Unknown field result type: %v", fieldInfo.Type)

		} // switch

	} // field value is a literal, just return it
}

func (equation EquationNode) EvalEqn(evalContext *EqnEvalContext) (*EquationResult, error) {

	if len(equation.FuncName) > 0 {
		funcInfo, funcInfoFound := evalContext.DefinedFuncs[equation.FuncName]
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
		fieldRef, err := field.GetFieldRef(evalContext.AppEngContext, equation.FieldID)
		if err != nil {
			return nil, fmt.Errorf("EvalEqn: failure retrieving referenced field: %+v", err)
		} else {
			return EvalEqn(evalContext, *fieldRef)
		}

	} else if equation.TextVal != nil {
		// Text literal given directly in the equation itself  (rather than a field value)
		return textEqnResult(*equation.TextVal), nil
	} else if equation.NumberVal != nil {
		// Number literal given directly in the equation itself (rather than a field value)
		return numberEqnResult(*equation.NumberVal), nil
	} else {
		return nil, fmt.Errorf("EvalEqn: malformed calculated field equation : system error: %+v", equation)
	}

}

// Update the calculated value for one calculated field
func updateOneCalcFieldValue(appEngContext appengine.Context, recordRef *record.RecordRef, fieldRef field.FieldRef) error {

	if !fieldRef.FieldInfo.IsCalcField {
		return fmt.Errorf("updateOneCalcFieldValue: Calculated field expected: got non-calculated field = %v", fieldRef.FieldInfo.RefName)
	}

	log.Printf("updateCalcFieldValues: Updating calculated field %v", fieldRef.FieldInfo.RefName)

	// The calculated field's equation is stored as regular JSON (since the datastore doesn't support
	// recursive structures). So, before evaluating the equation, the equation must first be decoded
	// into a (root) equation node.
	rootFieldEqnNode, decodeErr := decodeEquation(fieldRef.FieldInfo.CalcFieldEqn)
	if decodeErr != nil {
		return fmt.Errorf("Can't decode equation for field = %+v: decode error = %v", fieldRef, decodeErr)
	}

	// Perform the actual evaluation/calculation.
	fieldEqnResult, evalErr := rootFieldEqnNode.EvalEqn(
		&EqnEvalContext{appEngContext, CalcFieldDefinedFuncs, *recordRef})
	if evalErr != nil {
		return fmt.Errorf("Unexpected error evaluating equation for field=%v: error=%+v",
			fieldRef.FieldInfo.RefName, evalErr)
	} else if fieldEqnResult.IsUndefined() {
		log.Printf("Undefined field eqn result for calculated field = %v", fieldRef.FieldInfo.RefName)
		// If the evaluation result is undefined, don't proceed to actually set the value for the
		// calculated field (i.e., it will remain blank/undefined). An undefined value is a valid
		// condition when not all the depended upon field's values are defined.
		return nil
	}

	// Validate the result type matches the calculated field's expected result type.
	if fieldEqnResult.ResultType != fieldRef.FieldInfo.Type {
		return fmt.Errorf("Error evaluating equation for field=%+v: eqn=%+v: type mismatch on equation result:"+
			"expected=%v, got=%v", fieldRef, rootFieldEqnNode,
			fieldRef.FieldInfo.Type, fieldEqnResult.ResultType)
	} // if type mismatch between calculated equation result and calculated field's type

	// Set the calculated value in the record, depending on the type of value.
	switch fieldRef.FieldInfo.Type {

	case field.FieldTypeText:
		textResult, textResultErr := fieldEqnResult.GetTextResult()
		if textResultErr != nil {
			return fmt.Errorf("Unexpected error evaluating equation for field=%v: eqn=%+v: error=%v "+
				"unexpected error getting number result: raw result=%+v",
				fieldRef.FieldID, rootFieldEqnNode, textResultErr, fieldEqnResult)

		}
		log.Printf("updateCalcFieldValues: Setting calculated field value: field=%v, value=%v", fieldRef.FieldInfo.RefName, textResult)
		// TODO - encapsulate the actual setting of raw values into record.go
		recordRef.FieldValues[fieldRef.FieldID] = textResult
		return nil

	case field.FieldTypeNumber:
		numberResult, numberResultErr := fieldEqnResult.GetNumberResult()
		if numberResultErr != nil {
			return fmt.Errorf("Unexpected error evaluating equation for field=%v: eqn=%+v: error=%v "+
				"unexpected error getting number result: raw result=%+v",
				fieldRef.FieldID, rootFieldEqnNode, numberResultErr, fieldEqnResult)

		}
		log.Printf("updateCalcFieldValues: Setting calculated field value: field=%v, value=%v", fieldRef.FieldInfo.RefName, numberResult)
		// TODO - encapsulate the actual setting of raw values into record.go
		recordRef.FieldValues[fieldRef.FieldID] = numberResult
		return nil
		// TODO case FieldTypeDate

	default:
		return fmt.Errorf("Unexpected error evaluating equation for field=%+v: eqn=%+v: unsupported field type %v",
			fieldRef, rootFieldEqnNode, evalErr, fieldRef.FieldInfo.Type)
	} // switch field type

	return nil
}

// UpdateCalcFieldValues is (currently) the top-most entry point into the calculated field
// equation evaluation functionality. This is called after record updates (see recordUpdate package)
// to refresh calculated values.
func UpdateCalcFieldValues(appEngContext appengine.Context, recordRef *record.RecordRef) error {

	parentTableID, getParentErr := datastoreWrapper.GetParentID(recordRef.RecordID, table.TableEntityKind)
	if getParentErr != nil {
		return fmt.Errorf("Error updating field values - can't get record's parent table: %v", getParentErr)
	}

	fieldRefs, getFieldErr := field.GetAllFieldRefs(appEngContext, field.GetFieldListParams{ParentTableID: parentTableID})
	if getFieldErr != nil {
		return fmt.Errorf("Error updating field values - can't get fields: %v", getFieldErr)
	}

	for _, fieldRef := range fieldRefs {

		if fieldRef.FieldInfo.IsCalcField {
			if calcErr := updateOneCalcFieldValue(appEngContext, recordRef, fieldRef); calcErr != nil {
				// Some of the calculated fields may evaluate to an undefined values, in which case
				// an error won't be returned, but we can continue evaluating the other calculated
				// fields. However, if there is an actual error calculating the field values, processing
				// needs to stop and the error propagated.
				return fmt.Errorf("Error updating calculated field values - error for field = %v: %v",
					fieldRef.FieldInfo.RefName, calcErr)
			}
		} // If calculated field, proceed to update its calculation
	} // for each fieldRef
	return nil
}
