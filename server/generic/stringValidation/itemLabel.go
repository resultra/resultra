package stringValidation

import (
	"fmt"
	"regexp"
)

var itemLabelRegexp = regexp.MustCompile(`^[\p{L}0-9][\p{L}0-9 \'\.\-\&\(\)\%\$\#\/\?\*\!\+\-\{\}\=\_\^\&\@\%\~\|\/\\]{0,63}$`)

func WellFormedItemLabel(itemLabel string) bool {

	if !itemLabelRegexp.MatchString(itemLabel) {
		return false
	}
	return true
}

func validateItemLabel(itemLabel string) error {

	if !WellFormedItemLabel(itemLabel) {
		return fmt.Errorf("Invalid label")
	}
	return nil
}

func ValidateOptionalItemLabel(itemLabel string) error {
	if len(itemLabel) == 0 {
		return nil
	}

	return validateItemLabel(itemLabel)
}
