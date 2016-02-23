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

	if evalEqnResult, evalErr := funcEqn.evalEqn(&EqnEvalContext{appEngCntxt, calcFieldDefinedFuncs}); evalErr != nil {
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
