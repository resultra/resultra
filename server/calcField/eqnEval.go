package calcField

import (
	"database/sql"
	"fmt"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/global"
	"resultra/datasheet/server/record"
	"time"
)

type EqnEvalContext struct {

	// The record/item for which the calculated field equations are being evaluated.
	Record record.Record

	DefinedFuncs FuncNameFuncInfoMap

	ParentDatabaseID string

	// Set of field values into which the results will be calculated.
	// This is expected to be pre-populated with values from non-calculated fields.
	ResultFieldVals *record.RecFieldValues

	GlobalVals  global.GlobalValues
	GlobalIndex global.GlobalIDGlobalIndex

	FieldsByID map[string]field.Field

	CellUpdateFieldValIndex *record.CellUpdateFieldValueIndex
	EvalEqnAsOfTimestamp    time.Time
}

func validateEqnFieldForRecordValue(evalContext *EqnEvalContext, fieldID string,
	expectedFieldType string, allowCalcField bool) error {

	eqnField, foundField := evalContext.FieldsByID[fieldID]
	if !foundField {
		return fmt.Errorf("validateEqnFieldForRecordValue: Can't find field: %v", fieldID)
	}
	return record.ValidateFieldForRecordValue(eqnField, expectedFieldType, allowCalcField)

}

// Get the literal value from a number field, or undefined if it doesn't exist.
func getNumberRecordEqnResult(evalContext *EqnEvalContext, fieldID string) (*EquationResult, error) {

	// Since the calculated field values are stored in the Record just the same as values directly entered by end-users,
	// it is OK to retrieve the literal values from these fields just like non-calculated fields.
	allowCalcField := true
	if fieldValidateErr := validateEqnFieldForRecordValue(evalContext, fieldID,
		field.FieldTypeNumber, allowCalcField); fieldValidateErr != nil {
		return nil, fmt.Errorf("Can't get value from record with fieldID = %v: "+
			"Can't validate field with value type: validation error = %v", fieldID, fieldValidateErr)
	}

	val, foundVal := (*evalContext.ResultFieldVals)[fieldID]
	if !foundVal {
		// Note this is the only place which will return an undefined equation result (along with the
		// the similar function for other value types). This is because record values for non-calculated
		// fields is the only place where an undefined (or blank) value could originate, because a
		// user hasn't entered a value yet for the field.
		return undefinedEqnResult(), nil

	} else if val == nil {
		// If the value is defined, but set to a nil value, this means the value has been cleared. For purposes
		// of equation evaluation, this is the same as an undefined value.
		return undefinedEqnResult(), nil

	} else {
		if numberVal, foundNumber := val.(float64); !foundNumber {
			return nil, fmt.Errorf("Type mismatch retrieving value from record field id = %v:"+
				" expecting number, got %v", fieldID, val)
		} else {
			return numberEqnResult(numberVal), nil
		}
	} // else (if found a value for the given field ID)
}

func getBoolRecordEqnResult(evalContext *EqnEvalContext, fieldID string) (*EquationResult, error) {

	// Since the calculated field values are stored in the Record just the same as values directly entered by end-users,
	// it is OK to retrieve the literal values from these fields just like non-calculated fields.
	allowCalcField := true
	if fieldValidateErr := validateEqnFieldForRecordValue(evalContext, fieldID,
		field.FieldTypeBool, allowCalcField); fieldValidateErr != nil {
		return nil, fmt.Errorf("Can't get value from record with fieldID = %v: "+
			"Can't validate field with value type: validation error = %v", fieldID, fieldValidateErr)
	}

	val, foundVal := (*evalContext.ResultFieldVals)[fieldID]
	if !foundVal {
		// Note this is the only place which will return an undefined equation result (along with the
		// the similar function for other value types). This is because record values for non-calculated
		// fields is the only place where an undefined (or blank) value could originate, because a
		// user hasn't entered a value yet for the field.
		return undefinedEqnResult(), nil

	} else if val == nil {
		// If the value is defined, but set to a nil value, this means the value has been cleared. For purposes
		// of equation evaluation, this is the same as an undefined value.
		return undefinedEqnResult(), nil

	} else {
		if boolVal, foundBool := val.(bool); !foundBool {
			return nil, fmt.Errorf("getBoolRecordEqnResult: Type mismatch retrieving value from record field id = %v:"+
				" expecting bool, got %v", fieldID, val)
		} else {
			return boolEqnResult(boolVal), nil
		}
	} // else (if found a value for the given field ID)
}

// Get the literal value from a text field, or undefined if it doesn't exist.
func getTextRecordEqnResult(evalContext *EqnEvalContext, fieldID string) (*EquationResult, error) {

	// Since the calculated field values are stored in the Record just the same as values directly entered by end-users,
	// it is OK to retrieve the literal values from these fields just like non-calculated fields.
	allowCalcField := true
	if fieldValidateErr := validateEqnFieldForRecordValue(evalContext, fieldID, field.FieldTypeText, allowCalcField); fieldValidateErr != nil {
		return nil, fmt.Errorf("Can't get value from record with fieldID = %v: "+
			"Can't validate field with value type: validation error = %v", fieldID, fieldValidateErr)
	}

	val, foundVal := (*evalContext.ResultFieldVals)[fieldID]
	if !foundVal {
		// Note this is the only place which will return an undefined equation result (along with the
		// the similar function for other value types). This is because record values for non-calculated
		// fields is the only place where an undefined (or blank) value could originate, because a
		// user hasn't entered a value yet for the field.
		return undefinedEqnResult(), nil
	} else if val == nil {
		// If the value is defined, but set to a nil value, this means the value has been cleared. For purposes
		// of equation evaluation, this is the same as an undefined value.
		return undefinedEqnResult(), nil
	} else {
		if textVal, foundText := val.(string); !foundText {
			return nil, fmt.Errorf("Type mismatch retrieving value from record field id = %v:"+
				" expecting string, got %v", fieldID, val)

		} else {
			return textEqnResult(textVal), nil
		}
	} // else (if found a value for the given field ID)
}

func getTimeRecordEqnResult(evalContext *EqnEvalContext, fieldID string) (*EquationResult, error) {

	// Since the calculated field values are stored in the Record just the same as values directly entered by end-users,
	// it is OK to retrieve the literal values from these fields just like non-calculated fields.
	allowCalcField := true
	if fieldValidateErr := validateEqnFieldForRecordValue(evalContext, fieldID, field.FieldTypeTime, allowCalcField); fieldValidateErr != nil {
		return nil, fmt.Errorf("Can't get value from record with fieldID = %v: "+
			"Can't validate field with value type: validation error = %v", fieldID, fieldValidateErr)
	}

	val, foundVal := (*evalContext.ResultFieldVals)[fieldID]
	if !foundVal {
		// Note this is the only place which will return an undefined equation result (along with the
		// the similar function for other value types). This is because record values for non-calculated
		// fields is the only place where an undefined (or blank) value could originate, because a
		// user hasn't entered a value yet for the field.
		return undefinedEqnResult(), nil
	} else if val == nil {
		// If the value is defined, but set to a nil value, this means the value has been cleared. For purposes
		// of equation evaluation, this is the same as an undefined value.
		return undefinedEqnResult(), nil
	} else {
		if dateVal, foundDate := val.(time.Time); !foundDate {
			return nil, fmt.Errorf("Type mismatch retrieving value from record field id = %v:"+
				" expecting date/time, got %v", fieldID, val)

		} else {
			return timeEqnResult(dateVal), nil
		}
	} // else (if found a value for the given field ID)
}

func getGlobalValResult(evalContext *EqnEvalContext, globalID string) (*EquationResult, error) {

	globalInfo, globalInfoFound := (evalContext.GlobalIndex)[globalID]
	if !globalInfoFound {
		return nil, fmt.Errorf("getGlobalValResult: Can't find global information for global ID = %v", globalID)
	}

	val, foundVal := (evalContext.GlobalVals)[globalID]
	if !foundVal {
		// Note this is the only place which will return an undefined equation result (along with the
		// the similar function for other value types). This is because record values for non-calculated
		// fields is the only place where an undefined (or blank) value could originate, because a
		// user hasn't entered a value yet for the field.
		return undefinedEqnResult(), nil
	} else {
		switch globalInfo.Type {
		case global.GlobalTypeText:
			if textVal, foundText := val.(string); !foundText {
				return nil, fmt.Errorf("Type mismatch retrieving value from global id = %v:"+
					" expecting string, got %v", globalID, val)

			} else {
				return textEqnResult(textVal), nil
			}
		case global.GlobalTypeNumber:
			if numVal, foundNum := val.(float64); !foundNum {
				return nil, fmt.Errorf("Type mismatch retrieving value from global id = %v:"+
					" expecting number, got %v", globalID, val)

			} else {
				return numberEqnResult(numVal), nil
			}
		default:
			return nil, fmt.Errorf("Unknown global result type: %v", globalInfo.Type)

		}
	} // else (if found a value for the given field ID)

}

func EvalEqn(evalContext *EqnEvalContext, evalField field.Field) (*EquationResult, error) {

	if evalField.IsCalcField {
		// Calculated field - return the result of the calculation
		decodedEqn, decodeErr := decodeEquation(evalField.CalcFieldEqn)
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
			} else if calcFieldResult.ResultType != evalField.Type {
				return nil, fmt.Errorf("EvalEqn: type mismatch in result calculated for field: "+
					" expecting %v, got %v: field = %+v", evalField.Type, calcFieldResult.ResultType, evalField)
			} else {
				return calcFieldResult, nil
			}
		}
	} else { // literal field values
		switch evalField.Type {
		case field.FieldTypeText:
			return getTextRecordEqnResult(evalContext, evalField.FieldID)
		case field.FieldTypeNumber:
			return getNumberRecordEqnResult(evalContext, evalField.FieldID)
		case field.FieldTypeTime:
			return getTimeRecordEqnResult(evalContext, evalField.FieldID)
		case field.FieldTypeBool:
			return getBoolRecordEqnResult(evalContext, evalField.FieldID)
		default:
			return nil, fmt.Errorf("Unknown field result type: %v", evalField.Type)

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
		field, foundField := evalContext.FieldsByID[equation.FieldID]
		if !foundField {
			return nil, fmt.Errorf("EvalEqn: failure retrieving referenced field: %v", equation.FieldID)
		} else {
			return EvalEqn(evalContext, field)
		}

	} else if len(equation.GlobalID) > 0 {
		return getGlobalValResult(evalContext, equation.GlobalID)
	} else if equation.TextVal != nil {
		// Text literal given directly in the equation itself  (rather than a field value)
		return textEqnResult(*equation.TextVal), nil
	} else if equation.NumberVal != nil {
		// Number literal given directly in the equation itself (rather than a field value)
		return numberEqnResult(*equation.NumberVal), nil
	} else if equation.TimeVal != nil {
		return timeEqnResult(*equation.TimeVal), nil
	} else if equation.BoolVal != nil {
		return boolEqnResult(*equation.BoolVal), nil
	} else {
		// The EquationNode corresponds to an empty formula. In this case, always return an
		// undefined result.
		return undefinedEqnResult(), nil
	}

}

// Update the calculated value for one calculated field
func updateOneCalcFieldValue(evalContext *EqnEvalContext, evalField field.Field) error {

	if !evalField.IsCalcField {
		return fmt.Errorf("updateOneCalcFieldValue: Calculated field expected: got non-calculated field = %v", evalField.RefName)
	}

	// The calculated field's equation is stored as regular JSON (since the datastore doesn't support
	// recursive structures). So, before evaluating the equation, the equation must first be decoded
	// into a (root) equation node.
	rootFieldEqnNode, decodeErr := decodeEquation(evalField.CalcFieldEqn)
	if decodeErr != nil {
		return fmt.Errorf("Can't decode equation for field = %+v: decode error = %v", evalField, decodeErr)
	}

	// Perform the actual evaluation/calculation.
	fieldEqnResult, evalErr := rootFieldEqnNode.EvalEqn(evalContext)
	if evalErr != nil {
		return fmt.Errorf("Unexpected error evaluating equation for field=%v: error=%+v",
			evalField.RefName, evalErr)
	} else if fieldEqnResult.IsUndefined() {
		// If the evaluation result is undefined, don't proceed to actually set the value for the
		// calculated field (i.e., it will remain blank/undefined). An undefined value is a valid
		// condition when not all the depended upon field's values are defined.
		return nil
	}

	// Validate the result type matches the calculated field's expected result type.
	if fieldEqnResult.ResultType != evalField.Type {
		return fmt.Errorf("Error evaluating equation for field=%+v: eqn=%+v: type mismatch on equation result:"+
			"expected=%v, got=%v", evalField, rootFieldEqnNode,
			evalField.Type, fieldEqnResult.ResultType)
	} // if type mismatch between calculated equation result and calculated field's type

	// Set the calculated value in the record, depending on the type of value.
	switch evalField.Type {

	case field.FieldTypeText:
		textResult, textResultErr := fieldEqnResult.GetTextResult()
		if textResultErr != nil {
			return fmt.Errorf("Unexpected error evaluating equation for field=%v: eqn=%+v: error=%v "+
				"unexpected error getting number result: raw result=%+v",
				evalField.FieldID, rootFieldEqnNode, textResultErr, fieldEqnResult)

		}
		// TODO - encapsulate the actual setting of raw values into record.go
		(*evalContext.ResultFieldVals)[evalField.FieldID] = textResult
		return nil

	case field.FieldTypeNumber:
		numberResult, numberResultErr := fieldEqnResult.GetNumberResult()
		if numberResultErr != nil {
			return fmt.Errorf("Unexpected error evaluating equation for field=%v: eqn=%+v: error=%v "+
				"unexpected error getting number result: raw result=%+v",
				evalField.FieldID, rootFieldEqnNode, numberResultErr, fieldEqnResult)

		}
		// TODO - encapsulate the actual setting of raw values into record.go
		(*evalContext.ResultFieldVals)[evalField.FieldID] = numberResult
		return nil

	case field.FieldTypeBool:
		boolResult, boolResultErr := fieldEqnResult.GetBoolResult()
		if boolResultErr != nil {
			return fmt.Errorf("Unexpected error evaluating equation for field=%v: eqn=%+v: error=%v "+
				"unexpected error getting number result: raw result=%+v",
				evalField.FieldID, rootFieldEqnNode, boolResultErr, fieldEqnResult)

		}
		(*evalContext.ResultFieldVals)[evalField.FieldID] = boolResult
		return nil
		// TODO case FieldTypeDate

	case field.FieldTypeTime:
		timeResult, timeResultErr := fieldEqnResult.GetTimeResult()
		if timeResultErr != nil {
			return fmt.Errorf("Unexpected error evaluating equation for field=%v: eqn=%+v: error=%v "+
				"unexpected error getting number result: raw result=%+v",
				evalField.FieldID, rootFieldEqnNode, timeResultErr, fieldEqnResult)

		}
		// TODO - encapsulate the actual setting of raw values into record.go
		(*evalContext.ResultFieldVals)[evalField.FieldID] = timeResult
		return nil

	default:
		return fmt.Errorf("Unexpected error evaluating equation for field=%+v: eqn=%+v: unsupported field type %v: error = %v",
			evalField, rootFieldEqnNode, evalField.Type, evalErr)
	} // switch field type

	return nil
}

type CalcFieldUpdateConfig struct {
	TrackerDBHandle  *sql.DB
	ParentDatabaseID string
	GlobalIndex      global.GlobalIDGlobalIndex
	GlobalVals       global.GlobalValues
	Fields           []field.Field
	CurrUserID       string
	FieldsByID       map[string]field.Field
}

func CreateCalcFieldUpdateConfig(trackerDBHandle *sql.DB,
	currUserID string, parentDatabaseID string) (*CalcFieldUpdateConfig, error) {

	globalVals, globalValErr := global.GetGlobalValues(trackerDBHandle,
		global.GetGlobalValuesParams{ParentDatabaseID: parentDatabaseID})
	if globalValErr != nil {
		return nil, fmt.Errorf("UpdateCalcFieldValues: Unable to retrieve global values: error =%v", globalValErr)
	}

	globalIndex, globalIndexErr := global.GetIndexedGlobals(trackerDBHandle, parentDatabaseID)
	if globalIndexErr != nil {
		return nil, fmt.Errorf("UpdateCalcFieldValues: Unable to retrieve indexed globals: error =%v", globalIndexErr)
	}

	fields, getErr := field.GetAllFields(trackerDBHandle,
		field.GetFieldListParams{ParentDatabaseID: parentDatabaseID})
	if getErr != nil {
		return nil, fmt.Errorf("UpdateCalcFieldValues: Unable to retrieve fields from datastore: datastore error =%v", getErr)
	}

	fieldsByID := map[string]field.Field{}
	for _, currField := range fields {
		fieldsByID[currField.FieldID] = currField
	}

	config := CalcFieldUpdateConfig{
		TrackerDBHandle:  trackerDBHandle,
		ParentDatabaseID: parentDatabaseID,
		GlobalVals:       *globalVals,
		GlobalIndex:      globalIndex,
		CurrUserID:       currUserID,
		Fields:           fields,
		FieldsByID:       fieldsByID}

	return &config, nil

}

// UpdateCalcFieldValues is (currently) the top-most entry point into the calculated field
// equation evaluation functionality. This is called after record updates (see recordUpdate package)
// to refresh calculated values.
func UpdateCalcFieldValues(config *CalcFieldUpdateConfig, calcRecord record.Record, cellUpdateFieldValIndex *record.CellUpdateFieldValueIndex,
	calcAsOfTimestamp time.Time, resultFieldVals *record.RecFieldValues) error {

	// By definition, calculated field results are undefined before a record even comes into existance.
	if calcAsOfTimestamp.Before(calcRecord.CreateTimestampUTC) {
		return nil
	}

	eqnEvalContext := EqnEvalContext{
		ParentDatabaseID:        config.ParentDatabaseID,
		Record:                  calcRecord,
		ResultFieldVals:         resultFieldVals,
		CellUpdateFieldValIndex: cellUpdateFieldValIndex,
		EvalEqnAsOfTimestamp:    calcAsOfTimestamp,
		DefinedFuncs:            CalcFieldDefinedFuncs,
		GlobalVals:              config.GlobalVals,
		GlobalIndex:             config.GlobalIndex,
		FieldsByID:              config.FieldsByID}

	for _, currField := range config.Fields {

		if currField.IsCalcField {
			if calcErr := updateOneCalcFieldValue(&eqnEvalContext, currField); calcErr != nil {
				// Some of the calculated fields may evaluate to an undefined values, in which case
				// an error won't be returned, but we can continue evaluating the other calculated
				// fields. However, if there is an actual error calculating the field values, processing
				// needs to stop and the error propagated.
				return fmt.Errorf("Error updating calculated field values - error for field = %v: %v",
					currField.RefName, calcErr)
			}
		} // If calculated field, proceed to update its calculation
	} // for each fieldRef
	return nil
}
