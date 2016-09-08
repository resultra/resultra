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

var itemNameRegexp = regexp.MustCompile(`^[\p{L}0-9][\p{L}0-9 \'\.\-]{0,31}$`)

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

func DecodeJSONString(encodedStr string, decodedVal interface{}) error {

	encodedBytes := []byte(encodedStr)
	if err := json.Unmarshal(encodedBytes, decodedVal); err != nil {
		return fmt.Errorf("DecodeJSONString:Error decoding server JSON: encoded =  %v: decode error = %v", encodedStr, err)
	}
	return nil
}

func WellFormedItemName(itemName string) bool {

	if !itemNameRegexp.MatchString(itemName) {
		return false
	}
	return true
}

// A "reference name" for a field can only contain
// TODO - Can't start with "true or false" - add this when supporting boolean values
var wellFormedFormulaReferenceNameRegexp = regexp.MustCompile("^[a-zA-Z_][a-zA-Z0-9_]*$")

func WellFormedFormulaReferenceName(referenceName string) bool {
	if !wellFormedFormulaReferenceNameRegexp.MatchString(referenceName) {
		return false
	} else {
		return true
	}
}
