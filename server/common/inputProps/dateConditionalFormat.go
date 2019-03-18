// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package inputProps

import "time"

type DateConditionalFormat struct {
	Condition   string     `json:"condition"`
	ColorScheme string     `json:"colorScheme"`
	NumberParam *float64   `json:"numberParam,omitempty"`
	StartDate   *time.Time `json:"startDate,omitempty"`
	EndDate     *time.Time `json:"endDate,omitempty"`
}
