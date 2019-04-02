// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package common

import (
	"github.com/resultra/resultra/server/recordFilter"
)

type ComponentVisibilityProperties struct {
	VisibilityConditions recordFilter.RecordFilterRuleSet `json:"visibilityConditions"`
}

func NewDefaultComponentVisibilityProperties() ComponentVisibilityProperties {

	props := ComponentVisibilityProperties{
		VisibilityConditions: recordFilter.NewDefaultRecordFilterRuleSet()}

	return props
}
