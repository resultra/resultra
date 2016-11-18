package recordFilter

import (
	"fmt"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/recordValue"
)

func MatchRecord(recValResults recordValue.RecordValueResults, filterRules []RecordFilterRule) (bool, error) {

	for _, currFilterRule := range filterRules {
		ruleField, fieldErr := field.GetField(recValResults.ParentTableID, currFilterRule.FieldID)
		if fieldErr != nil {
			return false, fmt.Errorf("MatchRecord: Can't get field for filter rule = '%v': datastore error=%v",
				currFilterRule.FieldID, fieldErr)
		}

		filterFunc, ruleDefErr := getFilterFuncByFieldType(ruleField.Type, currFilterRule.RuleID)
		if ruleDefErr != nil {
			return false, fmt.Errorf("MatchRecord: Failed to retrieve filter rule definition: err=%v", ruleDefErr)
		}

		conditionMap := newFilterConditionMap(currFilterRule.Conditions)

		filterParams := FilterFuncParams{
			FieldID:      currFilterRule.FieldID,
			Conditions:   currFilterRule.Conditions,
			ConditionMap: conditionMap}
		recordIsFiltered, filterErr := filterFunc(filterParams, recValResults.FieldValues)
		if filterErr != nil {
			return false, fmt.Errorf("MatchRecord: Error filtering: %v", filterErr)
		}

		// Return false if any of the rules fail to match. Filtering is done based upon a logical AND of
		// all filter rules.
		if !recordIsFiltered {
			return false, nil
		}

	}

	return true, nil
}

func FilterRecordValues(filterRules []RecordFilterRule,
	unfilteredRecordValues []recordValue.RecordValueResults) ([]recordValue.RecordValueResults, error) {

	filteredRecords := []recordValue.RecordValueResults{}
	for _, recValue := range unfilteredRecordValues {

		isFiltered, filterErr := MatchRecord(recValue, filterRules)

		if filterErr != nil {
			return nil, fmt.Errorf("GetFilteredSortedRecords: Error filtering record: %v", filterErr)
		}

		if isFiltered {
			filteredRecords = append(filteredRecords, recValue)
		}

	}

	return filteredRecords, nil

}
