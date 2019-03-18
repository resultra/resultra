// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package alert

import (
	"time"
)

type AlertNotification struct {
	AlertID          string         `json:"alertID"`
	RecordID         string         `json:"recordID"`
	Timestamp        time.Time      `json:"timestamp"`
	Caption          string         `json:"caption"`
	DateBefore       *time.Time     `json:"dateBefore,omitempty"`
	DateAfter        *time.Time     `json:"dateAfter,omitempty"`
	TriggerCondition AlertCondition `json:"triggerCondition"`
}

// Custom sort function
type NotificationByTime []AlertNotification

func (s NotificationByTime) Len() int {
	return len(s)
}
func (s NotificationByTime) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Sort in  reverse chronological order; i.e. the most recent dates come first.
func (s NotificationByTime) Less(i, j int) bool {
	return s[i].Timestamp.After(s[j].Timestamp)
}
