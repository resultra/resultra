package stringValidation

import (
	"fmt"
	"regexp"
	"strings"
)

var itemNameRegexp = regexp.MustCompile(`^[\p{L}0-9][\p{L}0-9 \'\.\-]{0,31}$`)

func WellFormedItemName(itemName string) bool {

	if !itemNameRegexp.MatchString(itemName) {
		return false
	}
	return true
}

// Names can include anything except newlines, form feeds, and tabs - spaces are OK,
// but leading and trailing whitespace (including spaces) is stripped (see below)
var validNameRegexp = regexp.MustCompile("^[^\t\n\f\r ][^\t\n\f\r]*$")

func SanitizeName(unsanitizedName string) (string, error) {

	stripWhite := strings.TrimSpace(unsanitizedName) // strip leading & trailing whitespace

	if !validNameRegexp.MatchString(stripWhite) {
		return "", fmt.Errorf("Invalid name: '%v': Cannot be empty and must not contain newlines or tabs", unsanitizedName)
	}
	return stripWhite, nil
}
