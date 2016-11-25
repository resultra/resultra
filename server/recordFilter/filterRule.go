package recordFilter

import (
	"fmt"
	"resultra/datasheet/server/generic/uniqueID"
)

type RecordFilterRule struct {
	FieldID    string                `json:"fieldID"`
	RuleID     string                `json:"ruleID"`
	Conditions []FilterRuleCondition `json:"conditions"`
}

func (srcRule RecordFilterRule) Clone(remappedIDs uniqueID.UniqueIDRemapper) (*RecordFilterRule, error) {

	remappedFieldID, err := remappedIDs.GetExistingRemappedID(srcRule.FieldID)
	if err != nil {
		return nil, fmt.Errorf("RecordFilterRule.Clone: %v", err)
	}

	destFilterRule := srcRule
	destFilterRule.FieldID = remappedFieldID

	return &destFilterRule, nil

}

func CloneFilterRules(remappedIDs uniqueID.UniqueIDRemapper, srcRules []RecordFilterRule) ([]RecordFilterRule, error) {

	destRules := []RecordFilterRule{}

	for _, srcRule := range srcRules {
		destRule, err := srcRule.Clone(remappedIDs)
		if err != nil {
			return nil, fmt.Errorf("CloneFilterRules: %v")
		}
		destRules = append(destRules, *destRule)
	}

	return destRules, nil

}
