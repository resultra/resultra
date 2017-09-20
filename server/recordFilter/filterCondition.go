package recordFilter

import "time"

// A FilterRuleCondition consists of a ConditionID, which is a string depicting how the
// results are filter and optional parameters of different types. The use of
// different optional parameters depends on the RuleID.
// Using pointers and omitempty for the optional parameters results in compact
// JSON storage of the FilterRuleCondition when only a subset of optional parameters is used.
type FilterRuleCondition struct {
	OperatorID  string     `json:"operatorID"`
	TextParam   *string    `json:"textParam,omitempty"`
	NumberParam *float64   `json:"numberParam,omitempty"`
	TagsParam   []string   `json:"tagsParam,omitempty"`
	DateParam   *time.Time `json:"dateParam,omitempty"`
	UsersParam  []string   `json:"usersParam,omitempty"`
}

type FilterConditions []FilterRuleCondition

type FilterConditionMap map[string]FilterRuleCondition

func newFilterConditionMap(conditions []FilterRuleCondition) FilterConditionMap {

	condMap := FilterConditionMap{}
	for _, currCond := range conditions {
		condMap[currCond.OperatorID] = currCond
	}
	return condMap
}

func (condMap FilterConditionMap) getDateConditionParam(conditionID string) *time.Time {

	cond, condFound := condMap[conditionID]
	if !condFound {
		return nil
	}

	return cond.DateParam

}

func (condMap FilterConditionMap) getNumberConditionParam(conditionID string) *float64 {

	cond, condFound := condMap[conditionID]
	if !condFound {
		return nil
	}

	return cond.NumberParam

}

func (condMap FilterConditionMap) getTextConditionParam(conditionID string) *string {

	cond, condFound := condMap[conditionID]
	if !condFound {
		return nil
	}

	return cond.TextParam

}

func (condMap FilterConditionMap) getTagsConditionParam(conditionID string) []string {

	cond, condFound := condMap[conditionID]
	if !condFound {
		return nil
	}

	return cond.TagsParam

}

func (condMap FilterConditionMap) getUsersConditionParam(conditionID string) []string {

	cond, condFound := condMap[conditionID]
	if !condFound {
		return nil
	}

	return cond.UsersParam

}
