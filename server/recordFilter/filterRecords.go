package recordFilter

import (
	"fmt"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/record"
	"resultra/datasheet/server/recordValue"
)

// A filterRuleContext includes all the runtime information which is common for each filter and
// can be reused for filtering different records. For efficiency, these filterRuleContext's are set up once, then
// reused for all the records being filtered.
type filterRuleContext struct {
	filterRule   RecordFilterRule
	filterFunc   FilterRuleFunc
	filterParams FilterFuncParams
}

func CreateFilterRuleContexts(filterRules []RecordFilterRule) ([]filterRuleContext, error) {

	contexts := []filterRuleContext{}

	for _, currFilterRule := range filterRules {
		ruleField, fieldErr := field.GetField(currFilterRule.FieldID)
		if fieldErr != nil {
			return nil, fmt.Errorf("createFilterRuleContexts: Can't get field for filter rule = '%v': datastore error=%v",
				currFilterRule.FieldID, fieldErr)
		}

		filterFunc, ruleDefErr := getFilterFuncByFieldType(ruleField.Type, currFilterRule.RuleID)
		if ruleDefErr != nil {
			return nil, fmt.Errorf("createFilterRuleContexts: Failed to retrieve filter rule definition: err=%v", ruleDefErr)
		}

		conditionMap := newFilterConditionMap(currFilterRule.Conditions)

		filterParams := FilterFuncParams{
			FieldID:      currFilterRule.FieldID,
			Conditions:   currFilterRule.Conditions,
			ConditionMap: conditionMap}

		context := filterRuleContext{
			filterRule:   currFilterRule,
			filterFunc:   filterFunc,
			filterParams: filterParams}

		contexts = append(contexts, context)
	}

	return contexts, nil
}

func MatchOneRecordFromFieldValues(filterContexts []filterRuleContext, recordVals record.RecFieldValues) (bool, error) {
	for _, currContext := range filterContexts {

		recordIsFiltered, err := currContext.filterFunc(currContext.filterParams, recordVals)
		if err != nil {
			return false, fmt.Errorf("matchOneRecord: Error filtering: %v", err)
		}

		// Return false if any of the rules fail to match. Filtering is done based upon a logical AND of
		// all filter rules.
		if !recordIsFiltered {
			return false, nil
		}

	}

	// Matching a record is based upon a logical AND of all the results from the filters. If filtering gets to here,
	// then none of the filters have failed to match. The filtering logic will also get here if there are no filter rules,
	// and there is by default a match.
	return true, nil

}

func MatchOneRecord(filterContexts []filterRuleContext, recValResults recordValue.RecordValueResults) (bool, error) {
	return MatchOneRecordFromFieldValues(filterContexts, recValResults.FieldValues)
}

func FilterRecordValues(filterRules RecordFilterRuleSet,
	unfilteredRecordValues []recordValue.RecordValueResults) ([]recordValue.RecordValueResults, error) {

	filterContexts, err := CreateFilterRuleContexts(filterRules.FilterRules)
	if err != nil {
		return nil, fmt.Errorf("FilterRecordValues: Error setting up for filtering: %v", err)
	}

	filteredRecords := []recordValue.RecordValueResults{}
	for _, recValue := range unfilteredRecordValues {

		isFiltered, filterErr := MatchOneRecord(filterContexts, recValue)

		if filterErr != nil {
			return nil, fmt.Errorf("FilterRecordValues: Error filtering record: %v", filterErr)
		}

		if isFiltered {
			filteredRecords = append(filteredRecords, recValue)
		}

	}

	return filteredRecords, nil

}
