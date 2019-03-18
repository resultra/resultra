// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package userAuth

import (
	"database/sql"
	"regexp"
)

var regexpUserNameUserName = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]{4,13}[a-zA-Z0-9]$`)
var regexpUserNameRepeatUnderscores = regexp.MustCompile(`^.*__.*$`)

const regexpInvalidMsg string = "Invalid user name: user name must be between 6 and 15 characters, " +
	"start with a letter, then contain letters, numbers and underscores."

func validateWellFormedUserName(userName string) *AuthResponse {
	if !regexpUserNameUserName.MatchString(userName) {
		return newAuthResponse(false, regexpInvalidMsg)
	} else {
		if regexpUserNameRepeatUnderscores.MatchString(userName) {
			return newAuthResponse(false, regexpInvalidMsg)
		}
		return newAuthResponse(true, "Well formed user name")
	}
}

func validateNewUserName(trackerDBHandle *sql.DB, userName string) *AuthResponse {

	isValid, validateErr := validateUniqueUserName(trackerDBHandle, userName)
	if validateErr != nil {
		return newAuthResponse(false, "System error: failed to validate unique user name")
	}
	if !isValid {
		return newAuthResponse(false, "User name is already taken. Choose another user name")
	}

	// TODO(Important) - Validate uniqueness in the database (case insensitive)
	// Also validate user name does not contain key reserved words (e.g. company name)
	return validateWellFormedUserName(userName)
}
