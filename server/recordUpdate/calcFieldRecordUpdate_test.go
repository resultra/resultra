package recordUpdate

import (
	"appengine/aetest"
	"encoding/json"
	"resultra/datasheet/server/calcField"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/record"
	//	"resultra/datasheet/server/recordUpdate"
	"testing"
)

func encodeJSONString(t *testing.T, val interface{}) string {
	b, err := json.Marshal(val)
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}

func TestEquation(t *testing.T) {

	appEngCntxt, err := aetest.NewContext(nil)
	if err != nil {
		t.Fatal(err)
	}

	testField1 := field.NewFieldParams{Name: "Test Field 1", Type: "text", RefName: "FieldRef1"}
	fieldID, err := field.NewField(appEngCntxt, testField1)
	if err != nil {
		t.Fatal(err)
	}

	// Empty string with a text value.
	textVal := ""
	textEquation := calcField.EquationNode{TextVal: &textVal}
	t.Logf(encodeJSONString(t, textEquation))

	userText, userTextErr := textEquation.UserText(appEngCntxt)
	if userTextErr != nil {
		t.Error(userTextErr)
	} else {
		t.Logf(userText)
	}

	// Non-empty string with a text value
	textVal = "Foo"
	textEquation = calcField.EquationNode{TextVal: &textVal}
	t.Logf(encodeJSONString(t, textEquation))

	userText, userTextErr = textEquation.UserText(appEngCntxt)
	if userTextErr != nil {
		t.Error(userTextErr)
	} else {
		t.Logf(userText)
	}

	numVal := 24.2
	numberEquation := calcField.EquationNode{NumberVal: &numVal}
	t.Logf(encodeJSONString(t, numberEquation))

	userText, userTextErr = numberEquation.UserText(appEngCntxt)
	if userTextErr != nil {
		t.Error(userTextErr)
	} else {
		t.Logf(userText)
	}

	funcName := calcField.FuncNameConcat
	arg1Val := "arg1"
	arg1 := calcField.EquationNode{TextVal: &arg1Val}
	arg2Val := "arg2"
	arg2 := calcField.EquationNode{TextVal: &arg2Val}
	args := []calcField.EquationNode{arg1, arg2}
	funcEqn := calcField.EquationNode{FuncName: funcName, FuncArgs: args}
	jsonEncodeEqn := encodeJSONString(t, funcEqn)
	t.Logf(jsonEncodeEqn)

	expectedEqnJSON := "{\"funcName\":\"CONCATENATE\",\"funcArgs\":[{\"textVal\":\"arg1\"},{\"textVal\":\"arg2\"}]}"
	if jsonEncodeEqn != expectedEqnJSON {
		t.Errorf("Conversion to JSON equation failed: expected %v, got %v", jsonEncodeEqn, expectedEqnJSON)
	}

	userText, userTextErr = funcEqn.UserText(appEngCntxt)
	if userTextErr != nil {
		t.Error(userTextErr)
	} else {
		expected := "CONCATENATE(\"arg1\",\"arg2\")"
		if userText != expected {
			t.Errorf("Conversion to user equation string failed: expected = %v, got %v",
				expected, userText)
		}
		t.Logf(userText)
	}

	// This test doesn't retrieve record values, so a dummy record will suffice
	dummyRecordRef := record.RecordRef{"dummyFieldID", record.Record{}}

	if evalEqnResult, evalErr := funcEqn.EvalEqn(&calcField.EqnEvalContext{appEngCntxt, calcField.CalcFieldDefinedFuncs, dummyRecordRef}); evalErr != nil {
		t.Errorf("Unexpected error evaluating equation: %+v, eqn=%v", evalErr, userText)
	} else {
		textRes, validateErr := evalEqnResult.GetTextResult()
		if validateErr != nil {
			t.Errorf("Unexpected error evaluating equation: got unexpected result=%+v: err = %v",
				evalEqnResult, validateErr)
		} else {
			t.Logf("Concatenate results: %v", textRes)
			expected := "arg1arg2"
			if textRes != expected {
				t.Errorf("Conversion to user equation string failed: expected = %v, got %v",
					expected, textRes)
			}

		}
	}

	// The permanent fieldID is stored in the EquationNode, but the reference
	// name is used by the user.
	fieldRefEqn := calcField.EquationNode{FieldID: fieldID}
	t.Logf(encodeJSONString(t, fieldRefEqn))

	userText, userTextErr = fieldRefEqn.UserText(appEngCntxt)
	if userTextErr != nil {
		t.Error(userTextErr)
	} else {
		expected := "FieldRef1"
		if userText != expected {
			t.Errorf("Conversion to user equation string failed: expected = %v, got %v",
				expected, userText)
		}
		t.Logf(userText)
	}

}

func TestTextFieldReference(t *testing.T) {

	appEngCntxt, err := aetest.NewContext(nil)
	if err != nil {
		t.Fatal(err)
	}

	testField1 := field.NewFieldParams{Name: "Test Field 1", Type: "text", RefName: "FieldRef1"}
	fieldID1, field1Err := field.NewField(appEngCntxt, testField1)
	if field1Err != nil {
		t.Fatal(field1Err)
	}

	testField2 := field.NewFieldParams{Name: "Test Field 2", Type: "text", RefName: "FieldRef2"}
	fieldID2, field2Err := field.NewField(appEngCntxt, testField2)
	if field2Err != nil {
		t.Fatal(field2Err)
	}

	testRecordRef, recordErr := record.NewRecord(appEngCntxt)
	if recordErr != nil {
		t.Fatal(recordErr)
	}

	funcName := calcField.FuncNameConcat
	arg1 := calcField.FieldRefEqnNode(fieldID1)
	arg2 := calcField.FieldRefEqnNode(fieldID2)
	funcEqn := calcField.FuncEqnNode(funcName, []calcField.EquationNode{*arg1, *arg2})

	var updatedRecordRef *record.RecordRef
	var updateErr error
	if updatedRecordRef, updateErr = UpdateRecordValue(appEngCntxt,
		SetRecordTextValueParams{RecordUpdateHeader{testRecordRef.RecordID, fieldID1}, "fieldOneVal"}); updateErr != nil {
		t.Fatal(updateErr)
	}
	if updatedRecordRef, updateErr = UpdateRecordValue(appEngCntxt,
		SetRecordTextValueParams{RecordUpdateHeader{testRecordRef.RecordID, fieldID2}, "fieldTwoVal"}); updateErr != nil {
		t.Fatal(updateErr)
	}

	if evalEqnResult, evalErr := funcEqn.EvalEqn(&calcField.EqnEvalContext{appEngCntxt,
		calcField.CalcFieldDefinedFuncs, *updatedRecordRef}); evalErr != nil {
		t.Errorf("Unexpected error evaluating equation: %+v, eqn=%+v", evalErr, funcEqn)
	} else {
		catResult, catErr := evalEqnResult.GetTextResult()
		if catErr != nil {
			t.Errorf("Unexpected error from CONCATENATE with field references: %v", catErr)
		} else {
			t.Logf("TestFieldReference: concatenate results: %v", catResult)
			expected := "fieldOneValfieldTwoVal"
			if catResult != expected {
				t.Errorf("Unexpected result from CONCATENATE with field references: expecting %v, got %v",
					expected, catResult)
			}
		}
	}

}

func TestNumberFieldReference(t *testing.T) {

	appEngCntxt, err := aetest.NewContext(nil)
	if err != nil {
		t.Fatal(err)
	}

	testField1 := field.NewFieldParams{Name: "Test Field 1", Type: "number", RefName: "FieldRef1"}
	fieldID1, field1Err := field.NewField(appEngCntxt, testField1)
	if field1Err != nil {
		t.Fatal(field1Err)
	}

	testField2 := field.NewFieldParams{Name: "Test Field 2", Type: "number", RefName: "FieldRef2"}
	fieldID2, field2Err := field.NewField(appEngCntxt, testField2)
	if field2Err != nil {
		t.Fatal(field2Err)
	}

	testRecordRef, recordErr := record.NewRecord(appEngCntxt)
	if recordErr != nil {
		t.Fatal(recordErr)
	}

	funcName := calcField.FuncNameSum
	arg1 := calcField.FieldRefEqnNode(fieldID1)
	arg2 := calcField.FieldRefEqnNode(fieldID2)
	funcEqn := calcField.FuncEqnNode(funcName, []calcField.EquationNode{*arg1, *arg2})

	var updatedRecordRef *record.RecordRef
	var updateErr error
	if updatedRecordRef, updateErr = UpdateRecordValue(appEngCntxt,
		SetRecordNumberValueParams{RecordUpdateHeader{testRecordRef.RecordID, fieldID1}, 32.2}); updateErr != nil {
		t.Fatal(updateErr)
	}
	if updatedRecordRef, updateErr = UpdateRecordValue(appEngCntxt,
		SetRecordNumberValueParams{RecordUpdateHeader{testRecordRef.RecordID, fieldID2}, 42.4}); updateErr != nil {
		t.Fatal(updateErr)
	}

	if evalEqnResult, evalErr := funcEqn.EvalEqn(&calcField.EqnEvalContext{appEngCntxt,
		calcField.CalcFieldDefinedFuncs, *updatedRecordRef}); evalErr != nil {
		t.Errorf("Unexpected error evaluating equation: %+v, eqn=%+v", evalErr, funcEqn)
	} else {
		sumResult, sumErr := evalEqnResult.GetNumberResult()
		if sumErr != nil {
			t.Errorf("Unexpected error from SUM with field references: %v", sumErr)
		} else {
			t.Logf("TestFieldReference: sum results: %v", sumResult)
			expected := 74.6
			if sumResult != expected {
				t.Errorf("Unexpected result from SUM with field references: expecting %v, got %v",
					expected, sumResult)
			}
		}
	}

}

func TestCalculatedFieldSum(t *testing.T) {

	//appEngCntxt, err := aetest.NewContext(nil)
	// Currently the Field doesn't have an ancestor in the datastore, so there is not strong consistency,
	// meaning it won't be available for query until a period of time. To force consistency, the appEngCntx
	// is constructed with the right options ... TODO: Go back to using default appEngCntxt once
	// an ancestor has been established for Field.
	appEngCntxt, err := aetest.NewContext(&aetest.Options{StronglyConsistentDatastore: true})

	if err != nil {
		t.Fatal(err)
	}

	testField1 := field.NewFieldParams{Name: "Test Field 1", Type: "number", RefName: "FieldRef1"}
	fieldID1, field1Err := field.NewField(appEngCntxt, testField1)
	if field1Err != nil {
		t.Fatal(field1Err)
	}

	testField2 := field.NewFieldParams{Name: "Test Field 2", Type: "number", RefName: "FieldRef2"}
	fieldID2, field2Err := field.NewField(appEngCntxt, testField2)
	if field2Err != nil {
		t.Fatal(field2Err)
	}

	testRecordRef, recordErr := record.NewRecord(appEngCntxt)
	if recordErr != nil {
		t.Fatal(recordErr)
	}

	funcName := calcField.FuncNameSum
	arg1 := calcField.FieldRefEqnNode(fieldID1)
	arg2 := calcField.FieldRefEqnNode(fieldID2)
	funcEqn := calcField.FuncEqnNode(funcName, []calcField.EquationNode{*arg1, *arg2})

	calcFieldParms := calcField.NewCalcFieldParams{Name: "Test Field 2", Type: "number", RefName: "CalcField", FieldEqn: *funcEqn}
	calcFieldID, calcFieldErr := calcField.NewCalcField(appEngCntxt, calcFieldParms)
	if calcFieldErr != nil {
		t.Fatal(calcFieldErr)
	}

	var updatedRecordRef *record.RecordRef
	var updateErr error
	if updatedRecordRef, updateErr = UpdateRecordValue(appEngCntxt,
		SetRecordNumberValueParams{RecordUpdateHeader{testRecordRef.RecordID, fieldID1}, 32.2}); updateErr != nil {
		t.Fatal(updateErr)
	}
	if updatedRecordRef, updateErr = UpdateRecordValue(appEngCntxt,
		SetRecordNumberValueParams{RecordUpdateHeader{testRecordRef.RecordID, fieldID2}, 42.4}); updateErr != nil {
		t.Fatal(updateErr)
	}

	// After setting 2 values summed for the equation, get the value for the calculated field

	calcResult, calcErr := calcField.GetNumberRecordEqnResult(appEngCntxt, *updatedRecordRef, calcFieldID)
	if calcErr != nil {
		t.Fatal(calcErr)
	} else if calcResult.IsUndefined() {
		t.Fatalf("Error getting calculated field result - result shouldn't be undefined")
	} else {
		calcVal, resultErr := calcResult.GetNumberResult()
		if resultErr != nil {
			t.Fatalf("Error getting calculated field result - can't get numberical result from equation evaluation result: %v", resultErr)
		} else {
			expectedVal := 74.6
			t.Logf("Result for calculated field: %v", calcVal)
			if calcVal != expectedVal {
				t.Errorf("Calculated field doesn't match expected value: expected = %v, got %v", expectedVal, calcVal)
			}
		}
	}

}
