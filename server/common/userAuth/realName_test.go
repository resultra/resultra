// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package userAuth

import (
	"testing"
)

func verifyOneWellFormedRealName(t *testing.T, inputName string, whatVerify string) {
	nameValidateResp := validateWellFormedRealName(inputName)
	if !nameValidateResp.Success {
		t.Errorf("FAIL: Expecting validation of well-formed real name: %v (%v): error msg = %v",
			inputName, whatVerify, nameValidateResp.Msg)
	} else {
		t.Logf("PASS: Real name validated as expected: %v: %v", inputName, whatVerify)
	}
}

func verifyOneMalFormedRealName(t *testing.T, inputName string, whatVerify string) {
	nameValidateResp := validateWellFormedRealName(inputName)
	if nameValidateResp.Success {
		t.Errorf("FAIL: Name expected to be invalid but it was accepted as valid: %v (%v)",
			inputName, whatVerify)
	} else {
		t.Logf("PASS: Name rejected as expected: %v (%v): error msg = %v", inputName,
			whatVerify, nameValidateResp.Msg)
	}
}

func TestWellFormedRealNames(t *testing.T) {
	verifyOneWellFormedRealName(t, "J", "1 character minimum")
	verifyOneWellFormedRealName(t, "John", "regular name")
	verifyOneWellFormedRealName(t, "Johnson-smith", "hyphenated name")
	verifyOneWellFormedRealName(t, "Dr. Who", "Periods Ok")
	verifyOneWellFormedRealName(t, "de Marco", "spaces OK")
	verifyOneWellFormedRealName(t, "D'Angelo", "apostrophe OK")
	verifyOneWellFormedRealName(t, "Раз два три Jedna dva tři čtyři pět", "Non Latin letters OK")
	verifyOneWellFormedRealName(t, "André Svenson", "Non English name")
	verifyOneWellFormedRealName(t, "John Elkjærd", "Non English name")
	verifyOneWellFormedRealName(t, "张伟", "Chinese name")

	verifyOneMalFormedRealName(t, "123", "Numbers not OK")
	verifyOneMalFormedRealName(t, "", "Empty not OK")
	verifyOneMalFormedRealName(t, " . - '", "At least one letter")
	verifyOneMalFormedRealName(t, "  ", "At least one letter")
	verifyOneMalFormedRealName(t, "<$$$>", "Other special characters not OK")

}
