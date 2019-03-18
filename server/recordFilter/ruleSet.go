// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package recordFilter

import (
	"fmt"
	"resultra/tracker/server/trackerDatabase"
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

func (srcRuleSet RecordFilterRuleSet) Clone(cloneParams *trackerDatabase.CloneDatabaseParams) (*RecordFilterRuleSet, error) {

	destRules, err := CloneFilterRules(cloneParams, srcRuleSet.FilterRules)
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
