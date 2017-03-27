package recordFilter

import (
	"fmt"
	"resultra/datasheet/server/generic/uniqueID"
)

const RecordFilterMatchLogicAny string = "any"
const RecordFilterMatchLogicAll string = "all"

type RecordFilterRuleSet struct {
	MatchLogic  string             `json:"matchLogic"`
	FilterRules []RecordFilterRule `json:"filterRules"`
}

func NewDefaultRecordFilterRuleSet() RecordFilterRuleSet {
	ruleSet := RecordFilterRuleSet{
		MatchLogic:  RecordFilterMatchLogicAll,
		FilterRules: []RecordFilterRule{}}
	return ruleSet
}

func (srcRuleSet RecordFilterRuleSet) Clone(remappedIDs uniqueID.UniqueIDRemapper) (*RecordFilterRuleSet, error) {

	destRules, err := CloneFilterRules(remappedIDs, srcRuleSet.FilterRules)
	if err != nil {
		return nil, fmt.Errorf("RecordFilterRuleSet.Clone: %v", err)
	}

	destRuleSet := RecordFilterRuleSet{
		MatchLogic:  srcRuleSet.MatchLogic,
		FilterRules: destRules}

	return &destRuleSet, nil

}

func (srcRuleSet RecordFilterRuleSet) IsEmptyRuleSet() bool {
	if len(srcRuleSet.FilterRules) > 0 {
		return false
	} else {
		return true
	}
}
