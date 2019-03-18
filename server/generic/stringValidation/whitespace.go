// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
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
