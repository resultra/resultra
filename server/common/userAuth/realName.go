// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package userAuth

import (
	"regexp"
)

// Names can include Unicode letters, periods, hyphens and apostrophes
// Leading and trailing whitespace should be stripped off before validating.
var realNameRegexp = regexp.MustCompile("^[\\p{L} \\'\\.\\-]{1,256}$")
var noLetterInNameRegexp = regexp.MustCompile("^[ \\'\\.\\-]+$")

const regexpInvalidNameMsg string = "Name cannot be empty and can only include include letters, periods, hyphens and apostrophes."

func validateWellFormedRealName(realName string) *AuthResponse {

	// Check match against basic name match
	if !realNameRegexp.MatchString(realName) {
		return newAuthResponse(false, regexpInvalidNameMsg)
	}

	// Name must have some letter(s) - can't just be periods and apostrophes
	if noLetterInNameRegexp.MatchString(realName) {
		return newAuthResponse(false, regexpInvalidNameMsg)
	}

	return newAuthResponse(true, "Well formed real name")
}
