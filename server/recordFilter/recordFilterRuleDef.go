package recordFilter

import (
	"fmt"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/record"
)

type FilterFuncParams struct {
	Conditions []FilterRuleCondition
	FieldID    string
}

type FilterRuleFunc func(filterParams FilterFuncParams, recFieldVals record.RecFieldValues) (bool, error)

const filterRuleIDNotBlank string = "notBlank"
const filterRuleIDBlank string = "isBlank"
const filterRuleIDCustomDateRange string = "dateRange"
const filterRuleIDGreater string = "greater"
const filterRuleIDLess string = "less"
const filterRuleIDAny string = "any"

func filterBlankField(filterParams FilterFuncParams, recFieldVals record.RecFieldValues) (bool, error) {

	valueIsSet := recFieldVals.ValueIsSet(filterParams.FieldID)

	if valueIsSet {
		return false, nil
	} else {
		return true, nil
	}
}

func filterNonBlankField(filterParams FilterFuncParams, recFieldVals record.RecFieldValues) (bool, error) {

	valueIsSet := recFieldVals.ValueIsSet(filterParams.FieldID)

	if valueIsSet {
		return true, nil
	} else {
		return false, nil
	}
}

func filterGreater(filterParams FilterFuncParams, recFieldVals record.RecFieldValues) (bool, error) {
	return true, nil // stubbed out
}

func filterLess(filterParams FilterFuncParams, recFieldVals record.RecFieldValues) (bool, error) {
	return true, nil // stubbed out
}

func filterAny(filterParams FilterFuncParams, recFieldVals record.RecFieldValues) (bool, error) {
	return true, nil // stubbed out
}

func filterCustomDateRange(filterParams FilterFuncParams, recFieldVals record.RecFieldValues) (bool, error) {
	return true, nil // stubbed out
}

type RuleIDFilterFuncMap map[string]FilterRuleFunc

var textFieldFilterRuleDefs = RuleIDFilterFuncMap{
	filterRuleIDNotBlank: filterNonBlankField,
	filterRuleIDBlank:    filterBlankField}

var numberFieldFilterRuleDefs = RuleIDFilterFuncMap{
	filterRuleIDNotBlank: filterNonBlankField,
	filterRuleIDBlank:    filterBlankField,
	filterRuleIDGreater:  filterGreater,
	filterRuleIDLess:     filterLess}

var timeFieldFilterRuleDefs = RuleIDFilterFuncMap{
	filterRuleIDCustomDateRange: filterCustomDateRange,
	filterRuleIDAny:             filterAny}

// Get the rule definition based upon the field type
func getFilterFuncByFieldType(fieldType string, ruleID string) (FilterRuleFunc, error) {
	switch fieldType {
	case field.FieldTypeText:
		filterFunc, funcFound := textFieldFilterRuleDefs[ruleID]
		if !funcFound {
			return nil, fmt.Errorf(
				"getRuleDefByFieldType: Failed to retrieve filter rule definition for field type = %v, unrecognized rule ID = %v",
				fieldType, ruleID)
		} else {
			return filterFunc, nil
		}
	case field.FieldTypeNumber:
		filterFunc, funcFound := numberFieldFilterRuleDefs[ruleID]
		if !funcFound {
			return nil, fmt.Errorf(
				"getRuleDefByFieldType: Failed to retrieve filter rule definition for field type = %v, unrecognized rule ID = %v",
				fieldType, ruleID)
		} else {
			return filterFunc, nil
		}
	case field.FieldTypeTime:
		filterFunc, funcFound := timeFieldFilterRuleDefs[ruleID]
		if !funcFound {
			return nil, fmt.Errorf(
				"getRuleDefByFieldType: Failed to retrieve filter rule definition for field type = %v, unrecognized rule ID = %v",
				fieldType, ruleID)
		} else {
			return filterFunc, nil
		}
	default:
		return nil, fmt.Errorf(
			"getRuleDefByFieldType: Failed to retrieve filter rule definition: unknown field type = %v",
			fieldType)
	}
}
