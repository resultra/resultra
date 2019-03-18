// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package stringValidation

import (
	"testing"
)

func TestNameSanitize(t *testing.T) {

	// Leading or trailing whitespace will be stripped
	_, err := SanitizeName("ABC 123")
	if err != nil {
		t.Error(err)
	}

	// Empty names or names with newlines, tabs, or formfeeds are not OK
	_, err = SanitizeName("")
	if err == nil {
		t.Error(err)
	}

	_, err = SanitizeName("N\r\nF")
	if err == nil {
		t.Error(err)
	}

	_, err = SanitizeName("N\t\fF")
	if err == nil {
		t.Error(err)
	}

}
