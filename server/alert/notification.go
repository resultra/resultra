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
