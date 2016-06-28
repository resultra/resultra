package userAuth

import "unicode"

const minPasswordLength int = 8

type ValidatePasswordResponse struct {
	ValidPassword bool   `json:"validPassword"`
	Msg           string `json:"msg"`
}

// Checks if a password conforms to the following policy.
// Minimum 8 characters in length
// Contains 3/4 of the following items:
//   - Uppercase Letters
//   - Lowercase Letters
//   - Numbers
//   - Symbols
func validatePasswordStrength(password string) ValidatePasswordResponse {

	hasMinChars := (len(password) >= minPasswordLength)

	numUpper := 0
	numLower := 0
	numSpecial := 0
	numLetter := 0
	numDigit := 0

	for _, passChar := range password {

		if unicode.IsUpper(passChar) {
			numUpper++
		}
		if unicode.IsLower(passChar) {
			numLower++
		}
		if unicode.IsLetter(passChar) {
			numLetter++
		}
		if unicode.IsDigit(passChar) {
			numDigit++
		}
		if unicode.IsPunct(passChar) || unicode.IsSymbol(passChar) {
			numSpecial++
		}

	}

	numQualifiers := 0
	if numUpper >= 1 {
		numQualifiers++
	}
	if numLower >= 1 {
		numQualifiers++
	}
	if numDigit >= 1 {
		numQualifiers++
	}
	if numSpecial >= 1 {
		numQualifiers++
	}

	meetsMinQualifier := (numQualifiers >= 3)

	if hasMinChars && meetsMinQualifier {
		return ValidatePasswordResponse{
			ValidPassword: true,
			Msg:           "Valid password"}
	} else {
		return ValidatePasswordResponse{
			ValidPassword: false,
			Msg: "Invalid password: passwords must be at least 8 characters, and " +
				"include a mix of at least 3 of the following: upper case letters, " +
				" lower case letters, numbers and other symbols."}
	}

}
