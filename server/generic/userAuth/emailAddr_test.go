package userAuth

import (
	"testing"
)

func verifyOneWellFormedEmailAddr(t *testing.T, emailAddr string, whatVerify string) {
	validateResp := validateWellFormedEmailAddr(emailAddr)
	if !validateResp.Success {
		t.Errorf("FAIL: Expecting validation of well-formed email address: %v (%v): error msg = %v",
			emailAddr, whatVerify, validateResp.Msg)
	} else {
		t.Logf("PASS: email address validated as expected: %v: %v", emailAddr, whatVerify)
	}
}

func verifyOneMalFormedEmailAddr(t *testing.T, emailAddr string, whatVerify string) {
	validateResp := validateWellFormedEmailAddr(emailAddr)
	if validateResp.Success {
		t.Errorf("FAIL: Address expected to be invalid but it was accepted as valid: %v (%v)",
			emailAddr, whatVerify)
	} else {
		t.Logf("PASS: Email address rejected as expected: %v (%v): error msg = %v", emailAddr,
			whatVerify, validateResp.Msg)
	}
}

func TestWellFormedEmailAddress(t *testing.T) {
	verifyOneWellFormedEmailAddr(t, `prettyandsimple@example.com`, "simple")
	verifyOneWellFormedEmailAddr(t, `very.common@example.com`, "simple")
	verifyOneWellFormedEmailAddr(t, `John_Doe@example.com`, "simple - underscores OK")
	verifyOneWellFormedEmailAddr(t, `JohnDoe123@example.com`, "simple - numbers OK in local part")
	verifyOneWellFormedEmailAddr(t, `JohnDoe@example42.com`, "simple - numbers OK in domain part")
	verifyOneWellFormedEmailAddr(t, `JohnDoe@example-42.com`, "simple - hyphens OK in domain part")
	verifyOneWellFormedEmailAddr(t, `JohnDoe@example-42-and-beyond.com`, "hyphens OK in domain part")
	verifyOneWellFormedEmailAddr(t, `disposable.style.email.with+symbol@example.com`, "simple")

	verifyOneWellFormedEmailAddr(t, `#!$%&'*+-/=?^_{}|~@example.org`, "special chars except backtick")
	verifyOneWellFormedEmailAddr(t, "me`@example.com", "includes backtick")
	verifyOneWellFormedEmailAddr(t, "s@example.com", "one letter local part")
	verifyOneWellFormedEmailAddr(t, "test@a.com", "one letter local part")
	verifyOneWellFormedEmailAddr(t, "s@example.com", "longer top-level domain")

	verifyOneMalFormedEmailAddr(t, `Abc.example.com`, "No @ sign")
	verifyOneMalFormedEmailAddr(t, `prettyandsimple@`, "No Domain part")
	verifyOneMalFormedEmailAddr(t, `prettyandsimple`, "No Domain part")
	verifyOneMalFormedEmailAddr(t, `@example.com`, "No local part")
	verifyOneMalFormedEmailAddr(t, ``, "empty address")
	verifyOneMalFormedEmailAddr(t, ` `, "empty address")

	verifyOneMalFormedEmailAddr(t, `A@b@c@example.com`, "Multiple @ sign")
	verifyOneMalFormedEmailAddr(t, `a"b(c)d,e:f;g<h>i[j\k]l@example.com`, "Invalid special chars")
	verifyOneMalFormedEmailAddr(t, `<foo@example.com>`, "Valid address, but invalid format for registration")
	verifyOneMalFormedEmailAddr(t, `Foo <foo@example.com>`, "Valid address, but invalid format for registration")
	verifyOneMalFormedEmailAddr(t, `prettyandsimple@example.com `, "trailing space")
	verifyOneMalFormedEmailAddr(t, ` prettyandsimple@example.com `, "leading space")
	verifyOneMalFormedEmailAddr(t, `prettyand simple@example.com`, "space in middle")
	verifyOneMalFormedEmailAddr(t, `prettyand simple@example.com`, "space in middle")

	verifyOneMalFormedEmailAddr(t, `prettyandsimple@example..com`, "double dots in domain part")
	verifyOneMalFormedEmailAddr(t, `prettya..ndsimple@example.com`, "double dots in local part")
	verifyOneMalFormedEmailAddr(t, `prettya...ndsimple@example.com`, "triple dots in local part")
	verifyOneMalFormedEmailAddr(t, `prettyandsimple@example.com.`, "dot at end of domain part")
	verifyOneMalFormedEmailAddr(t, `prettyandsimple@.example.com`, "dot at beginning of domain part")
	verifyOneMalFormedEmailAddr(t, `prettyandsimple.@example.com`, "dot at end of local part")
	verifyOneMalFormedEmailAddr(t, `.prettyandsimple@example.com`, "dot at beginning of local part")
	verifyOneMalFormedEmailAddr(t, `test@-example.com`, "hyphen at beginning of domain part")
	verifyOneMalFormedEmailAddr(t, `test@example.com-`, "hyphen at end of domain part")

	// The following are technically valid, but not permitted for registration.
	// There's no business reason not to accept quoted email addresses, but they are very uncommon
	// and it probably won't hurt not to accept.
	verifyOneMalFormedEmailAddr(t, `"much.more unusual"@example.com`, "technically valid, but not supported for registration")
	verifyOneMalFormedEmailAddr(t, `"very.unusual.@.unusual.com"@example.com`, "technically valid, but not supported for registration")

	// The following are technically valid, but not permitted for registration
	verifyOneMalFormedEmailAddr(t, `admin@mailserver1`, "local domain not allowed")
	verifyOneMalFormedEmailAddr(t, `admin@localhost`, "local domain not allowed")
	verifyOneMalFormedEmailAddr(t, `user@[IPv6:2001:db8::1]`, "technically valid, but not supported for registration")
	verifyOneMalFormedEmailAddr(t, `jsmith@[192.168.2.1]`, "technically valid, but not supported for registration")
	verifyOneMalFormedEmailAddr(t, `" "@example.org`, "technically valid, but not supported for registration")
	verifyOneMalFormedEmailAddr(t, `john.smith@(comment)example.com`, "technically valid, but not supported for registration")
	verifyOneMalFormedEmailAddr(t, `john.smith(comment)@example.com`, "technically valid, but not supported for registration")

}
