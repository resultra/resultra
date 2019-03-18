// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package stringValidation

import (
	"fmt"
	"regexp"
	"strings"
)

var itemNameRegexp = regexp.MustCompile(`^[\p{L}0-9][\p{L}0-9 \'\.\-\&\(\)\%\$\#\/\?\*\^\!\%]{0,63}$`)

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
