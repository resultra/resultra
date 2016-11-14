package recordFilter

type FilterRuleCondition struct {
	RuleID          string  `json:"ruleID"`
	TextRuleParam   string  `json:"textRuleParam"`
	NumberRuleParam float64 `json:"numberRuleParam"`
}

type RecordFilterRule struct {
	FieldID    string                `json:"fieldID"`
	Conditions []FilterRuleCondition `json:"conditions"`
}
