package recordFilter

import (
	"time"
)

type RecordFilterRule struct {
	FieldID    string                `json:"fieldID"`
	RuleID     string                `json:"ruleID"`
	Conditions []FilterRuleCondition `json:"conditions"`
}

// A FilterRuleCondition consists of a RuleID, which is a string depicting how the
// results are filter and optional parameters of different types. The use of
// different optional parameters depends on the RuleID.
// Using pointers and omitempty for the optional parameters results in compact
// JSON storage of the FilterRuleCondition when only a subset of optional parameters is used.
type FilterRuleCondition struct {
	OperatorID  string     `json:"operatorID"`
	TextParam   *string    `json:"textParam,omitempty"`
	NumberParam *float64   `json:"numberParam,omitempty"`
	DateParam   *time.Time `json:"dateParam,omitempty"`
}
