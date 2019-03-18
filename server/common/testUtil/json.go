// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package testUtil

import (
	"encoding/json"
	"testing"
)

func EncodeJSONString(t *testing.T, val interface{}) string {
	b, err := json.Marshal(val)
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}
