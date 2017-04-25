package stringValidation

import (
	"strings"
)

func StringAllWhitespace(s string) bool {
	// If there's nothing left after trimming all the (unicode) whitespace, then the string is all whitespace
	if len(strings.TrimSpace(s)) == 0 {
		return true
	} else {
		return false
	}
}
