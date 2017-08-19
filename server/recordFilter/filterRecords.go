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
type FilterRuleContext struct {
	filterRule   RecordFilterRule
	filterFunc   FilterRuleFunc
	filterParams FilterFuncParams
}

func CreateFilterRuleContexts(filterRules []RecordFilterRule) ([]FilterRuleContext, error) {

	contexts := []FilterRuleContext{}

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

		context := FilterRuleContext{
			filterRule:   currFilterRule,
			filterFunc:   filterFunc,
			filterParams: filterParams}

		contexts = append(contexts, context)
	}

	return contexts, nil
}

func MatchOneRecordFromFieldValues(matchLogic string, filterContexts []FilterRuleContext, recordVals record.RecFieldValues) (bool, error) {

	if matchLogic == RecordFilterMatchLogicAll {
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

	} else {
		// Match any filtering condition.
		for _, currContext := range filterContexts {

			recordIsFiltered, err := currContext.filterFunc(currContext.filterParams, recordVals)
			if err != nil {
				return false, fmt.Errorf("matchOneRecord: Error filtering: %v", err)
			}

			// Return false if any of the rules fail to match. Filtering is done based upon a logical AND of
			// all filter rules.
			if recordIsFiltered {
				return true, nil
			}

		}
		// Matching is based upon a logical OR of the filtering rules. So, if matching over all the conditions
		// hasn't matched at least one condition by now, the boolean OR is false.
		return false, nil
	}

}

func MatchOneRecord(matchLogic string, filterContexts []FilterRuleContext, recValResults recordValue.RecordValueResults) (bool, error) {
	return MatchOneRecordFromFieldValues(matchLogic, filterContexts, recValResults.FieldValues)
}

func FilterRecordValues(filterRules RecordFilterRuleSet,
	unfilteredRecordValues []recordValue.RecordValueResults) ([]recordValue.RecordValueResults, error) {

	filterContexts, err := CreateFilterRuleContexts(filterRules.FilterRules)
	if err != nil {
		return nil, fmt.Errorf("FilterRecordValues: Error setting up for filtering: %v", err)
	}

	filteredRecords := []recordValue.RecordValueResults{}
	for _, recValue := range unfilteredRecordValues {

		isFiltered, filterErr := MatchOneRecord(filterRules.MatchLogic, filterContexts, recValue)

		if filterErr != nil {
			return nil, fmt.Errorf("FilterRecordValues: Error filtering record: %v", filterErr)
		}

		if isFiltered {
			filteredRecords = append(filteredRecords, recValue)
		}

	}

	return filteredRecords, nil

}
