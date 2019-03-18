// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package userAuth

import (
	"regexp"
)

// The following regular expression is by no means a complete validation for RFC 5322  compliant email addresses.
// However, it is intended to screen out obviously mal-formed email addresses, and be a little bit restrictive
// for type types of email addresses accepted for registration (possibly screening out "RFC 5322 afficianodos")
//
// This regular expression is basically based upon the following, starting at
// https://en.wikipedia.org/wiki/Email_address#Local_part and through the domain part.
//
// Besides doing a basic screening on email address format, for registration an email will generally
// be sent to verify the email address anyway, so sending this email is the ultimate validity test
// of the email.
//
// Note the string is concatenated together since a back-tick is a valid part of the "local part" of the name
//
// The regexp below doesn't handle international characters which is just beginning to be supported
// through RFC 6530
//

var sanityCheckEmailAddrRegexp *regexp.Regexp

var doubleDotsDomainEmailAddrRegexp *regexp.Regexp
var doubleDotsLocalEmailAddrRegexp *regexp.Regexp
var dotInMiddleOfDomainPartRegexp *regexp.Regexp

var localStartDotRegexp *regexp.Regexp
var localEndDotRegexp *regexp.Regexp

var domainStartHyphenRegexp *regexp.Regexp
var domainEndHyphenRegexp *regexp.Regexp

func init() {

	beginAddr := `^`
	localPart := `[a-zA-Z0-9\!\#\$\%\&\'\*\+\-\/\=\?\^\_\` + "`" + `\{\|\}\~\.]+`
	domainPart := `[a-zA-Z0-9\-\.]+`
	domainDotInMiddle := `[^\.].*\..*[^\.]` // no dot in beginning or end of domain part, must be in middle
	endAddr := `$`
	doubleDots := `.*\.\..*`

	localStartDot := `^\.` // use the @ as end marker `.*@) // |(^.*[^\.]@)`
	localEndDot := `^.*\.@`

	domainBeginHyphen := `\-.*`
	domainEndHyphen := `.*\-`

	sanityCheckEmailAddrRegexp = regexp.MustCompile(beginAddr + localPart + `@` + domainPart + endAddr)

	doubleDotsDomainEmailAddrRegexp = regexp.MustCompile(beginAddr + localPart + `@` + doubleDots + endAddr)
	doubleDotsLocalEmailAddrRegexp = regexp.MustCompile(beginAddr + doubleDots + `@` + domainPart + endAddr)

	dotInMiddleOfDomainPartRegexp = regexp.MustCompile(beginAddr + localPart + `@` + domainDotInMiddle + endAddr)

	localStartDotRegexp = regexp.MustCompile(localStartDot)
	localEndDotRegexp = regexp.MustCompile(localEndDot)

	domainStartHyphenRegexp = regexp.MustCompile(beginAddr + localPart + `@` + domainBeginHyphen + endAddr)
	domainEndHyphenRegexp = regexp.MustCompile(beginAddr + localPart + `@` + domainEndHyphen + endAddr)

}

const invalidEmailAddressMsg string = "Invalid email address"

func ValidateWellFormedEmailAddr(emailAddr string) *AuthResponse {

	if !sanityCheckEmailAddrRegexp.MatchString(emailAddr) {
		return newAuthResponse(false, invalidEmailAddressMsg)
	}

	// Double dots ".." not allowed in either the local or domain part
	if doubleDotsDomainEmailAddrRegexp.MatchString(emailAddr) {
		return newAuthResponse(false, invalidEmailAddressMsg)
	}

	if doubleDotsLocalEmailAddrRegexp.MatchString(emailAddr) {
		return newAuthResponse(false, invalidEmailAddressMsg)
	}

	//  In the local part, dots (".") cannot be at the beginning or end
	if localStartDotRegexp.MatchString(emailAddr) {
		return newAuthResponse(false, invalidEmailAddressMsg)
	}
	if localEndDotRegexp.MatchString(emailAddr) {
		return newAuthResponse(false, invalidEmailAddressMsg)
	}

	// In the domain part, hyphens cannot be at the beginning or end
	if domainStartHyphenRegexp.MatchString(emailAddr) {
		return newAuthResponse(false, invalidEmailAddressMsg)
	}
	if domainEndHyphenRegexp.MatchString(emailAddr) {
		return newAuthResponse(false, invalidEmailAddressMsg)
	}

	// In the domain part, there must be at least 1 dot and it must be in the middle -
	// local domains are not permitted for registration.
	if !dotInMiddleOfDomainPartRegexp.MatchString(emailAddr) {
		return newAuthResponse(false, invalidEmailAddressMsg)
	}

	return newAuthResponse(true, "Valid email address")
}
