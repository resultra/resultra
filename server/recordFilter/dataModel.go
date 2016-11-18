package recordFilter

type RecordFilterRule struct {
	FieldID    string                `json:"fieldID"`
	RuleID     string                `json:"ruleID"`
	Conditions []FilterRuleCondition `json:"conditions"`
}
