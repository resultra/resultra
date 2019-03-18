package recordFilter

import (
	"fmt"
	"resultra/tracker/server/trackerDatabase"
)

type RecordFilterRule struct {
	FieldID    string                `json:"fieldID"`
	RuleID     string                `json:"ruleID"`
	Conditions []FilterRuleCondition `json:"conditions"`
}

func (srcRule RecordFilterRule) Clone(cloneParams *trackerDatabase.CloneDatabaseParams) (*RecordFilterRule, error) {

	remappedFieldID, err := cloneParams.IDRemapper.GetExistingRemappedID(srcRule.FieldID)
	if err != nil {
		return nil, fmt.Errorf("RecordFilterRule.Clone: %v", err)
	}

	destFilterRule := srcRule
	destFilterRule.FieldID = remappedFieldID

	return &destFilterRule, nil

}

func CloneFilterRules(cloneParams *trackerDatabase.CloneDatabaseParams, srcRules []RecordFilterRule) ([]RecordFilterRule, error) {

	destRules := []RecordFilterRule{}

	for _, srcRule := range srcRules {
		destRule, err := srcRule.Clone(cloneParams)
		if err != nil {
			return nil, fmt.Errorf("CloneFilterRules: %v")
		}
		destRules = append(destRules, *destRule)
	}

	return destRules, nil

}
