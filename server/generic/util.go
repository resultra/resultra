package generic

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

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

func EncodeJSONString(val interface{}) (string, error) {
	b, err := json.Marshal(val)
	if err != nil {
		return "", fmt.Errorf("Error encoding JSON: %v", err)
	}
	return string(b), nil
}
