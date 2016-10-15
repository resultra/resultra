package generic

import (
	"encoding/json"
	"fmt"
	"regexp"
)

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
