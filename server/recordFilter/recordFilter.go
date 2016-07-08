package recordFilter

import (
	"fmt"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/record"
)

type OptonalFilterRuleParamVal struct {
	textParam   *string
	numberParam *float64
}

// Convert the optional paramater values as stored in the datastore (via filterRule),
// and convert them to pointer values which can be ommitted from output when
// converting to JSON.
func getOptionalParamValueByRuleDef(filterRuleDef FilterRuleDef,
	filterRule RecordFilterRule) (*OptonalFilterRuleParamVal, error) {

	optParamVal := OptonalFilterRuleParamVal{}

	if filterRuleDef.HasParam {
		switch filterRuleDef.DataType {
		case field.FieldTypeText:
			optParamVal.textParam = &filterRule.TextRuleParam
		case field.FieldTypeNumber:
			optParamVal.numberParam = &filterRule.NumberRuleParam
		default:
			return nil, fmt.Errorf("getOptionalParamValueByRuleDef: unknown rule definition data type = %v",
				filterRuleDef.DataType)
		} // switch
	} // if has param

	return &optParamVal, nil
}

func filterOneRecord(parentTableID string, recFieldVals record.RecFieldValues, filterRules []RecordFilterRule) (bool, error) {

	for _, currFilterRule := range filterRules {

		ruleField, fieldErr := field.GetField(parentTableID, currFilterRule.FieldID)
		if fieldErr != nil {
			return false, fmt.Errorf("filterOneRecord: Can't get field for filter rule = '%v': datastore error=%v",
				currFilterRule.FieldID, fieldErr)
		}

		filterRuleDef, ruleDefErr := getRuleDefByFieldType(ruleField.Type, currFilterRule.RuleID)
		if ruleDefErr != nil {
			return false, fmt.Errorf("createOneFilterRef: Failed to retrieve filter rule definition: err=%v", ruleDefErr)
		}

		// Pass a limited subset of parameters to the functions for doing the actual filtering. This simplifies
		// the code inside actual filtering functions and makes these functions easier to test.
		optParamVal, paramErr := getOptionalParamValueByRuleDef(*filterRuleDef, currFilterRule)
		if paramErr != nil {
			return false, fmt.Errorf("filterOneRecord: Error filtering record: %v", paramErr)
		}

		filterParams :=
			FilterFuncParams{currFilterRule.FieldID,
				optParamVal.textParam,
				optParamVal.numberParam}

		recordIsFiltered, filterErr := filterRuleDef.filterFunc(filterParams, recFieldVals)

		if filterErr != nil {
			return false, fmt.Errorf("filterOneRecord: Error filtering record: %v", filterErr)
		} else if !recordIsFiltered {

			// Filtering is performed based upon a logical AND of all filtering rules passing; i.e. (F1 & F2 & F3 ...).
			// To be more efficient, we can test for the first to not pass its filtering rule; i.e. !(F1 & F2 & F3 ...)
			// == !F1 || !F2 || !F3
			return false, nil
		}
	}
	return true, nil
}

func testOneFilterOneRecFieldVals(parentTableID string, recFieldVals record.RecFieldValues, filterID string) (bool, error) {

	getFilterRulesParam := GetFilterRulesParams{ParentFilterID: filterID}
	filterRules, rulesErr := getRecordFilterRules(getFilterRulesParam)
	if rulesErr != nil {
		return false, rulesErr
	}

	recordIsFiltered, filterErr := filterOneRecord(parentTableID, recFieldVals, filterRules)
	if filterErr != nil {
		return false, fmt.Errorf("GetFilteredRecords: Error filtering records: %v", filterErr)
	}

	return recordIsFiltered, nil

}

func GenerateFilterMatchResults(recFieldVals record.RecFieldValues, parentTableID string) (RecFilterMatchResults, error) {

	filters, getFilterErr := getFilterList(parentTableID)
	if getFilterErr != nil {
		return nil, fmt.Errorf("GenerateFilterMatchResults: Unable to retrieve filters for table ID = %v: error = %v", parentTableID, getFilterErr)
	}

	matchResults := RecFilterMatchResults{}
	for _, currFilter := range filters {
		matchResult, matchErr := testOneFilterOneRecFieldVals(parentTableID, recFieldVals, currFilter.FilterID)
		if matchErr != nil {
			return nil, fmt.Errorf(
				"GenerateFilterMatchResults: Error matching filter for table ID = %v, filter ID = %v: error = %v",
				parentTableID, currFilter.FilterID, matchErr)
		}
		matchResults[currFilter.FilterID] = matchResult
	}
	return matchResults, nil

}
