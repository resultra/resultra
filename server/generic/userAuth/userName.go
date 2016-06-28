package userAuth

import (
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

func validateNewUserName(userName string) *AuthResponse {
	// TODO(Important) - Validate uniqueness in the database (case insensitive)
	// Also validate user name does not contain key reserved words (e.g. company name)
	return validateWellFormedUserName(userName)
}
