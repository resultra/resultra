package userAuth

import (
	"testing"
)

func verifyOneValidPassword(t *testing.T, inputPassword string, whatVerify string) {
	pwValidateResp := validatePasswordStrength(inputPassword)
	if !pwValidateResp.ValidPassword {
		t.Errorf("FAIL: Not expecting invalid password: password = %v (%v): error msg = %v",
			inputPassword, whatVerify, pwValidateResp.Msg)
	} else {
		t.Logf("PASS: Password validated as expected: %v: %v", inputPassword, whatVerify)
	}
}

func verifyOneInvalidPassword(t *testing.T, inputPassword string, whatVerify string) {
	pwValidateResp := validatePasswordStrength(inputPassword)
	if pwValidateResp.ValidPassword {
		t.Errorf("FAIL: Password expected to be invalid but it was accepted as valid: %v (%v)",
			inputPassword, whatVerify)
	} else {
		t.Logf("PASS: Password rejected as expected: %v (%v): error msg = %v", inputPassword,
			whatVerify, pwValidateResp.Msg)
	}
}

func TestPasswordStrength(t *testing.T) {
	verifyOneValidPassword(t, "Abc123$%", "8 characters with upper & lower case and special chars")
	verifyOneValidPassword(t, "Abc123$=_df3df98df34j34dfkjdfdf", "Longer passwords ok too")
	verifyOneValidPassword(t, "Abc123de", "8 characters with upper & lower case and a number")

	verifyOneValidPassword(t, "Abcfegd!", "8 characters with upper & lower case and a special character")
	verifyOneValidPassword(t, "Abcfegd@", "8 characters with upper & lower case and a special character")
	verifyOneValidPassword(t, "Abcfegd#", "8 characters with upper & lower case and a special character")
	verifyOneValidPassword(t, "Abcfegd%", "8 characters with upper & lower case and a special character")
	verifyOneValidPassword(t, "Abcfegd^", "8 characters with upper & lower case and a special character")
	verifyOneValidPassword(t, "Abcfegd&", "8 characters with upper & lower case and a special character")
	verifyOneValidPassword(t, "Abcfegd*", "8 characters with upper & lower case and a special character")
	verifyOneValidPassword(t, "Abcfegd(", "8 characters with upper & lower case and a special character")
	verifyOneValidPassword(t, "Abcfegd)", "8 characters with upper & lower case and a special character")
	verifyOneValidPassword(t, "Abcfegd_", "8 characters with upper & lower case and a special character")
	verifyOneValidPassword(t, "Abcfegd-", "8 characters with upper & lower case and a special character")
	verifyOneValidPassword(t, "Abcfegd+", "8 characters with upper & lower case and a special character")
	verifyOneValidPassword(t, "Abcfegd=", "8 characters with upper & lower case and a special character")
	verifyOneValidPassword(t, "Abcfegd{", "8 characters with upper & lower case and a special character")
	verifyOneValidPassword(t, "Abcfegd}", "8 characters with upper & lower case and a special character")
	verifyOneValidPassword(t, "Abcfegd[", "8 characters with upper & lower case and a special character")
	verifyOneValidPassword(t, "Abcfegd]", "8 characters with upper & lower case and a special character")
	verifyOneValidPassword(t, "Abcfegd|", "8 characters with upper & lower case and a special character")
	verifyOneValidPassword(t, `Abcfegd\`, "8 characters with upper & lower case and a special character")
	verifyOneValidPassword(t, "Abcfegd?", "8 characters with upper & lower case and a special character")
	verifyOneValidPassword(t, "Abcfegd/", "8 characters with upper & lower case and a special character")
	verifyOneValidPassword(t, "Abcfegd~", "8 characters with upper & lower case and a special character")
	verifyOneValidPassword(t, "Abcfegd`", "8 characters with upper & lower case and a special character")
	verifyOneValidPassword(t, "Abcfegd<", "8 characters with upper & lower case and a special character")
	verifyOneValidPassword(t, "Abcfegd>", "8 characters with upper & lower case and a special character")
	verifyOneValidPassword(t, "Abcfegd,", "8 characters with upper & lower case and a special character")
	verifyOneValidPassword(t, "Abcfegd.", "8 characters with upper & lower case and a special character")
	verifyOneValidPassword(t, "Abcfegd:", "8 characters with upper & lower case and a special character")
	verifyOneValidPassword(t, "Abcfegd;", "8 characters with upper & lower case and a special character")
	verifyOneValidPassword(t, "Abcfegd0", "8 characters with upper & lower case and a digit")
	verifyOneValidPassword(t, "Abcfegd1", "8 characters with upper & lower case and a digit")
	verifyOneValidPassword(t, "Abcfegd2", "8 characters with upper & lower case and a digit")
	verifyOneValidPassword(t, "Abcfegd3", "8 characters with upper & lower case and a digit")
	verifyOneValidPassword(t, "Abcfegd4", "8 characters with upper & lower case and a digit")
	verifyOneValidPassword(t, "Abcfegd5", "8 characters with upper & lower case and a digit")
	verifyOneValidPassword(t, "Abcfegd6", "8 characters with upper & lower case and a digit")
	verifyOneValidPassword(t, "Abcfegd7", "8 characters with upper & lower case and a digit")
	verifyOneValidPassword(t, "Abcfegd8", "8 characters with upper & lower case and a digit")
	verifyOneValidPassword(t, "Abcfegd9", "8 characters with upper & lower case and a digit")

	verifyOneValidPassword(t, "Abcfegd9", "lower case, upper case, digit (3/4 rules)")
	verifyOneValidPassword(t, "Abcfegd!", "lower case, upper case, special (3/4 rules)")
	verifyOneValidPassword(t, "Abcfegd!", "lower case, upper case, special (3/4 rules)")

	verifyOneValidPassword(t, "Abcf egd9", "space is Ok (not typical but could be part of password generator)")

	verifyOneValidPassword(t, "abc123$%", "no upper & lower case, but both numbers and special char")
	verifyOneValidPassword(t, "ABC123$%", "no upper & lower case, but both numbers and special char")

	verifyOneInvalidPassword(t, "abcdEFGH", "only lower case and upper case")
	verifyOneInvalidPassword(t, "abcd1234", "only lower case and numbers")
	verifyOneInvalidPassword(t, "ABCD1234", "only upper case and numbers")
	verifyOneInvalidPassword(t, "abcd$%^&", "only lower case and special chars")
	verifyOneInvalidPassword(t, "ABCD$%^&", "only upper case and special chars")
	verifyOneInvalidPassword(t, "%^&*1234", "only number and special chars")
	verifyOneInvalidPassword(t, "aA1$", "matches all rules but too short")

	verifyOneInvalidPassword(t, "Abc123$", "only 7 chars")
}
