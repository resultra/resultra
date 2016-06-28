package userAuth

import (
	"testing"
)

func verifyOneWellFormedUserName(t *testing.T, inputUserName string, whatVerify string) {
	userNameValidateResp := validateWellFormedUserName(inputUserName)
	if !userNameValidateResp.Success {
		t.Errorf("FAIL: Expecting validation of well-formed user name: %v (%v): error msg = %v",
			inputUserName, whatVerify, userNameValidateResp.Msg)
	} else {
		t.Logf("PASS: Username validated as expected: %v: %v", inputUserName, whatVerify)
	}
}

func verifyOneNonWellFormedUserName(t *testing.T, inputUserName string, whatVerify string) {
	userNameValidateResp := validateWellFormedUserName(inputUserName)
	if userNameValidateResp.Success {
		t.Errorf("FAIL: Username expected to be invalid but it was accepted as valid: %v (%v)",
			inputUserName, whatVerify)
	} else {
		t.Logf("PASS: Username rejected as expected: %v (%v): error msg = %v", inputUserName,
			whatVerify, userNameValidateResp.Msg)
	}
}

func TestWellFormedUserNames(t *testing.T) {
	verifyOneWellFormedUserName(t, "Abc123", "6 character minimum")
	verifyOneWellFormedUserName(t, "Abc123DEF456GHI", "15 character maximum")

	verifyOneWellFormedUserName(t, "Ab_123", "Underscores OK")
	verifyOneNonWellFormedUserName(t, "Ab__c12", "Repeat underscores not OK")
	verifyOneNonWellFormedUserName(t, "Abc123_", "Underscore at end not OK")

	verifyOneNonWellFormedUserName(t, "Abc12", "Below 6 character minimum")
	verifyOneNonWellFormedUserName(t, "Abc123DEF456GHIZ", "Above 15 character maximum")

	verifyOneNonWellFormedUserName(t, "\nABC123", "Invalid characters")
	verifyOneNonWellFormedUserName(t, " ABC123", "Invalid characters")
	verifyOneNonWellFormedUserName(t, "ABC123 ", "Invalid characters")
	verifyOneNonWellFormedUserName(t, "AB C123", "Invalid characters")
	verifyOneNonWellFormedUserName(t, "ABC$123", "Invalid characters")
	verifyOneNonWellFormedUserName(t, "1ABCefg", "Must start with letter")
}
