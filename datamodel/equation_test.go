package datamodel

import (
	"appengine/aetest"
	"encoding/json"
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

	testField1 := Field{Name: "Test Field 1", Type: "text", RefName: "FieldRef1"}
	fieldID, err := NewField(appEngCntxt, testField1)
	if err != nil {
		t.Fatal(err)
	}

	// Empty string with a text value.
	textVal := ""
	textEquation := EquationNode{TextVal: &textVal}
	t.Logf(encodeJSONString(t, textEquation))

	userText, userTextErr := textEquation.UserText(appEngCntxt)
	if userTextErr != nil {
		t.Error(userTextErr)
	} else {
		t.Logf(userText)
	}

	// Non-empty string with a text value
	textVal = "Foo"
	textEquation = EquationNode{TextVal: &textVal}
	t.Logf(encodeJSONString(t, textEquation))

	userText, userTextErr = textEquation.UserText(appEngCntxt)
	if userTextErr != nil {
		t.Error(userTextErr)
	} else {
		t.Logf(userText)
	}

	numVal := 24.2
	numberEquation := EquationNode{NumberVal: &numVal}
	t.Logf(encodeJSONString(t, numberEquation))

	userText, userTextErr = numberEquation.UserText(appEngCntxt)
	if userTextErr != nil {
		t.Error(userTextErr)
	} else {
		t.Logf(userText)
	}

	funcName := funcNameConcat
	arg1Val := "arg1"
	arg1 := EquationNode{TextVal: &arg1Val}
	arg2Val := "arg2"
	arg2 := EquationNode{TextVal: &arg2Val}
	args := []EquationNode{arg1, arg2}
	funcEqn := EquationNode{FuncName: funcName, FuncArgs: args}
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
	dummyRecordRef := RecordRef{"dummyFieldID", Record{}}

	if evalEqnResult, evalErr := funcEqn.evalEqn(&EqnEvalContext{appEngCntxt, calcFieldDefinedFuncs, dummyRecordRef}); evalErr != nil {
		t.Errorf("Unexpected error evaluating equation: %+v, eqn=%v", evalErr, userText)
	} else {
		textRes, validateErr := evalEqnResult.getTextResult()
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
	fieldRefEqn := EquationNode{FieldID: fieldID}
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

	testField1 := Field{Name: "Test Field 1", Type: "text", RefName: "FieldRef1"}
	fieldID1, field1Err := NewField(appEngCntxt, testField1)
	if field1Err != nil {
		t.Fatal(field1Err)
	}

	testField2 := Field{Name: "Test Field 2", Type: "text", RefName: "FieldRef2"}
	fieldID2, field2Err := NewField(appEngCntxt, testField2)
	if field2Err != nil {
		t.Fatal(field2Err)
	}

	testRecordRef, recordErr := NewRecord(appEngCntxt)
	if recordErr != nil {
		t.Fatal(recordErr)
	}

	funcName := funcNameConcat
	arg1 := fieldRefEqnNode(fieldID1)
	arg2 := fieldRefEqnNode(fieldID2)
	funcEqn := funcEqnNode(funcName, []EquationNode{*arg1, *arg2})

	var updatedRecordRef *RecordRef
	var updateErr error
	if updatedRecordRef, updateErr = SetRecordTextValue(appEngCntxt,
		SetRecordTextValueParams{testRecordRef.RecordID, fieldID1, "fieldOneVal"}); updateErr != nil {
		t.Fatal(updateErr)
	}
	if updatedRecordRef, updateErr = SetRecordTextValue(appEngCntxt,
		SetRecordTextValueParams{testRecordRef.RecordID, fieldID2, "fieldTwoVal"}); updateErr != nil {
		t.Fatal(updateErr)
	}

	if evalEqnResult, evalErr := funcEqn.evalEqn(&EqnEvalContext{appEngCntxt,
		calcFieldDefinedFuncs, *updatedRecordRef}); evalErr != nil {
		t.Errorf("Unexpected error evaluating equation: %+v, eqn=%+v", evalErr, funcEqn)
	} else {
		catResult, catErr := evalEqnResult.getTextResult()
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

	testField1 := Field{Name: "Test Field 1", Type: "number", RefName: "FieldRef1"}
	fieldID1, field1Err := NewField(appEngCntxt, testField1)
	if field1Err != nil {
		t.Fatal(field1Err)
	}

	testField2 := Field{Name: "Test Field 2", Type: "number", RefName: "FieldRef2"}
	fieldID2, field2Err := NewField(appEngCntxt, testField2)
	if field2Err != nil {
		t.Fatal(field2Err)
	}

	testRecordRef, recordErr := NewRecord(appEngCntxt)
	if recordErr != nil {
		t.Fatal(recordErr)
	}

	funcName := funcNameSum
	arg1 := fieldRefEqnNode(fieldID1)
	arg2 := fieldRefEqnNode(fieldID2)
	funcEqn := funcEqnNode(funcName, []EquationNode{*arg1, *arg2})

	var updatedRecordRef *RecordRef
	var updateErr error
	if updatedRecordRef, updateErr = SetRecordNumberValue(appEngCntxt,
		SetRecordNumberValueParams{testRecordRef.RecordID, fieldID1, 32.2}); updateErr != nil {
		t.Fatal(updateErr)
	}
	if updatedRecordRef, updateErr = SetRecordNumberValue(appEngCntxt,
		SetRecordNumberValueParams{testRecordRef.RecordID, fieldID2, 42.4}); updateErr != nil {
		t.Fatal(updateErr)
	}

	if evalEqnResult, evalErr := funcEqn.evalEqn(&EqnEvalContext{appEngCntxt,
		calcFieldDefinedFuncs, *updatedRecordRef}); evalErr != nil {
		t.Errorf("Unexpected error evaluating equation: %+v, eqn=%+v", evalErr, funcEqn)
	} else {
		sumResult, sumErr := evalEqnResult.getNumberResult()
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

	testField1 := Field{Name: "Test Field 1", Type: "number", RefName: "FieldRef1"}
	fieldID1, field1Err := NewField(appEngCntxt, testField1)
	if field1Err != nil {
		t.Fatal(field1Err)
	}

	testField2 := Field{Name: "Test Field 2", Type: "number", RefName: "FieldRef2"}
	fieldID2, field2Err := NewField(appEngCntxt, testField2)
	if field2Err != nil {
		t.Fatal(field2Err)
	}

	testRecordRef, recordErr := NewRecord(appEngCntxt)
	if recordErr != nil {
		t.Fatal(recordErr)
	}

	funcName := funcNameSum
	arg1 := fieldRefEqnNode(fieldID1)
	arg2 := fieldRefEqnNode(fieldID2)
	funcEqn := funcEqnNode(funcName, []EquationNode{*arg1, *arg2})

	jsonEncodeEqn := encodeJSONString(t, funcEqn)
	calcField := Field{Name: "Test Field 2", Type: "number", RefName: "CalcField", IsCalcField: true, CalcFieldEqn: jsonEncodeEqn}
	calcFieldID, calcFieldErr := NewField(appEngCntxt, calcField)
	if calcFieldErr != nil {
		t.Fatal(calcFieldErr)
	}

	var updatedRecordRef *RecordRef
	var updateErr error
	if updatedRecordRef, updateErr = SetRecordNumberValue(appEngCntxt,
		SetRecordNumberValueParams{testRecordRef.RecordID, fieldID1, 32.2}); updateErr != nil {
		t.Fatal(updateErr)
	}
	if updatedRecordRef, updateErr = SetRecordNumberValue(appEngCntxt,
		SetRecordNumberValueParams{testRecordRef.RecordID, fieldID2, 42.4}); updateErr != nil {
		t.Fatal(updateErr)
	}

	if calcErr := updateCalcFieldValues(appEngCntxt, updatedRecordRef); calcErr != nil {
		t.Fatalf("Error updating calculated fields to reflect update: err = %v", calcErr)
	}

	// After setting 2 values summed for the equation, get the value for the calculated field

	/* The evaluation of calculated field equations is currently returning an error for undefined fields.
	   		 Instead, there needs to be support for an undefined results, which will cascade up through the
	   	     recursion. The code below only works becuase there is a call to updateCalcFieldValues *after*
		    we know all the depended upon fields are defined. */
	calcResult, calcErr := updatedRecordRef.GetNumberRecordValue(appEngCntxt, calcFieldID)
	if calcErr != nil {
		t.Fatal(calcErr)
	}
	t.Logf("Result for calculated field: %v", calcResult)

}
