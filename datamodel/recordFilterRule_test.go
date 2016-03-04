package datamodel

import (
	"appengine"
	"appengine/aetest"
	"testing"
)

func verifyNewFilterRuleCreation(testSummary string, t *testing.T, appEngContext appengine.Context,
	newRuleParams NewFilterRuleParams) {

	_, err := NewFilterRule(appEngContext, newRuleParams)
	if err != nil {
		t.Errorf("verifyNewFilterRuleCreation (fail): %v: Expected filter rule creation to succeed, but failed with err=%v",
			testSummary, err)
	} else {
		t.Logf("verifyNewFilterRuleCreation (pass): %v: filter rule creation succeeded", testSummary)
	}
}

func verifyNewFilterRuleCreationFailure(testSummary string, t *testing.T, appEngContext appengine.Context,
	newRuleParams NewFilterRuleParams) {

	_, err := NewFilterRule(appEngContext, newRuleParams)
	if err != nil {
		t.Logf("verifyNewFilterRuleCreation (pass): %v: Expected filter rule creation to fail: got err=%v",
			testSummary, err)
	} else {
		t.Logf("verifyNewFilterRuleCreation (fail): %v: filter rule creation unexpectedly succeeded", testSummary)
	}
}

func TestNewFilterRule(t *testing.T) {

	appEngCntxt, err := aetest.NewContext(nil)
	if err != nil {
		t.Fatal(err)
	}

	testField := NewFieldParams{Name: "Test Text Field", Type: fieldTypeText, RefName: "TestRef1"}
	testFieldID, err := NewField(appEngCntxt, testField)
	if err != nil {
		t.Fatal(err)
	}

	verifyNewFilterRuleCreation("Text field with filter rule not requiring a parameter", t, appEngCntxt,
		NewFilterRuleParams{testFieldID, filterRuleIDNotBlank, nil, nil})

	verifyNewFilterRuleCreationFailure("Text field with invalid field ID", t, appEngCntxt,
		NewFilterRuleParams{"badfieldID", filterRuleIDNotBlank, nil, nil})

	verifyNewFilterRuleCreationFailure("Text field with invalid filtering rule ID", t, appEngCntxt,
		NewFilterRuleParams{testFieldID, "invalidRuleID", nil, nil})

	testNumField := NewFieldParams{Name: "Test Number Field", Type: fieldTypeNumber, RefName: "TestRef1"}
	testNumFieldID, numFieldErr := NewField(appEngCntxt, testNumField)
	if numFieldErr != nil {
		t.Fatal(numFieldErr)
	}

	verifyNewFilterRuleCreation("Number field with filter rule not requiring a parameter", t, appEngCntxt,
		NewFilterRuleParams{testNumFieldID, filterRuleIDNotBlank, nil, nil})

	verifyNewFilterRuleCreationFailure("Number field with invalid field ID", t, appEngCntxt,
		NewFilterRuleParams{"badfieldID", filterRuleIDNotBlank, nil, nil})

	verifyNewFilterRuleCreationFailure("Number field with invalid filtering rule ID", t, appEngCntxt,
		NewFilterRuleParams{testNumFieldID, "invalidRuleID", nil, nil})

}
