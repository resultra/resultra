package recordFilter

import (
	"fmt"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/recordValue"
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

func filterOneRecord(recValResults recordValue.RecordValueResults, filterRules []RecordFilterRule) (bool, error) {

	for _, currFilterRule := range filterRules {

		ruleField, fieldErr := field.GetField(recValResults.ParentTableID, currFilterRule.FieldID)
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

		recordIsFiltered, filterErr := filterRuleDef.filterFunc(filterParams, recValResults)

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

type GetFilteredRecordsParams struct {
	TableID   string   `json:"tableID"`
	FilterIDs []string `json:filterIDs`
}

// Since filtering is done using a logical AND of all the filters and their rules, it simplifies actual
// filtering to filter using a combined set of rules. Then, if any rule across all the filters doesn't pass,
// then the record is filtered.
func getCombinedFilterRules(filterIDs []string) ([]RecordFilterRule, error) {

	allFilterRules := []RecordFilterRule{}

	for _, currFilterID := range filterIDs {

		getFilterRulesParam := GetFilterRulesParams{ParentFilterID: currFilterID}
		currFilterRules, rulesErr := getRecordFilterRules(getFilterRulesParam)
		if rulesErr != nil {
			return nil, rulesErr
		}
		allFilterRules = append(allFilterRules, currFilterRules...)
	}

	return allFilterRules, nil
}

func GetFilteredRecords(params GetFilteredRecordsParams) ([]recordValue.RecordValueResults, error) {

	// TODO - The code below retrieve *all* the records. However, the datastore supports up to 1 filtering criterion
	// for each field, so <=1 of these criterion could be used to filter the records coming from the datastore and
	// before doing any kind of in-memory filtering.
	unfilteredRecordValues, getRecordErr := recordValue.GetAllRecordValueResults(params.TableID)
	if getRecordErr != nil {
		return nil, fmt.Errorf("GetFilteredRecords: Error retrieving records: %v", getRecordErr)
	}

	filterRules, rulesErr := getCombinedFilterRules(params.FilterIDs)
	if rulesErr != nil {
		return nil, fmt.Errorf("GetFilteredRecords: Error retrieving filter rules: %v", rulesErr)
	}

	filteredRecords := []recordValue.RecordValueResults{}
	for _, currRecordValResult := range unfilteredRecordValues {
		recordIsFiltered, filterErr := filterOneRecord(currRecordValResult, filterRules)
		if filterErr != nil {
			return nil, fmt.Errorf("GetFilteredRecords: Error filtering records: %v", filterErr)
		}
		if recordIsFiltered {
			filteredRecords = append(filteredRecords, currRecordValResult)
		}
	}

	return filteredRecords, nil
}
