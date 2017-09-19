package recordFilter

import (
	"fmt"
	"log"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/record"
	"strings"
)

type FilterFuncParams struct {
	Conditions   []FilterRuleCondition
	FieldID      string
	ConditionMap FilterConditionMap
}

type FilterRuleFunc func(filterParams FilterFuncParams, recFieldVals record.RecFieldValues) (bool, error)

const filterRuleIDNotBlank string = "notBlank"
const filterRuleIDBlank string = "isBlank"

const filterRuleIDTags string = "tags"
const conditionMatchTags string = "tags"

const filterRuleIDCustomDateRange string = "dateRange"

const filterRuleIDGreater string = "greater"
const filterRuleIDGreaterEqual string = "greaterEqual"
const filterRuleIDLess string = "less"
const filterRuleIDLessEqual string = "lessEqual"
const filterRuleIDEqual string = "equal"
const conditionGreater string = "greater"
const conditionLess string = "less"
const conditionGreaterEqual string = "greaterEqual"
const conditionLessEqual string = "lessEqual"
const conditionEqual string = "equal"

const filterRuleIDAny string = "any"
const filterRuleTrue string = "true"
const filterRuleNotTrue string = "notTrue"
const filterRuleFalse string = "false"
const filterRuleNotFalse string = "notFalse"
const filterRuleIDContains string = "contains"
const filterRuleAfter string = "after"
const filterRuleBefore string = "before"

const conditionDateRangeMinDate string = "minDate"
const conditionDateRangeMaxDate string = "maxDate"
const conditionTextContains string = "contains"

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

	greaterComparisonVal := filterParams.ConditionMap.getNumberConditionParam(conditionGreater)
	if greaterComparisonVal == nil {
		return false, nil
	}

	log.Printf("filterGreater: comparison value = %v", *greaterComparisonVal)

	numberVal, valFound := recFieldVals.GetNumberFieldValue(filterParams.FieldID)
	if !valFound {
		return false, nil
	}
	log.Printf("filterGreater: value = %v", numberVal)

	if numberVal > *greaterComparisonVal {
		return true, nil
	} else {
		return false, nil
	}

}

func filterGreaterEqual(filterParams FilterFuncParams, recFieldVals record.RecFieldValues) (bool, error) {

	greaterComparisonVal := filterParams.ConditionMap.getNumberConditionParam(conditionGreaterEqual)
	if greaterComparisonVal == nil {
		return false, nil
	}

	log.Printf("filterGreater: comparison value = %v", *greaterComparisonVal)

	numberVal, valFound := recFieldVals.GetNumberFieldValue(filterParams.FieldID)
	if !valFound {
		return false, nil
	}
	log.Printf("filterGreater: value = %v", numberVal)

	if numberVal >= *greaterComparisonVal {
		return true, nil
	} else {
		return false, nil
	}

}

func filterLess(filterParams FilterFuncParams, recFieldVals record.RecFieldValues) (bool, error) {
	lessComparisonVal := filterParams.ConditionMap.getNumberConditionParam(conditionLess)
	if lessComparisonVal == nil {
		return false, nil
	}

	log.Printf("filterLess: comparison value = %v", *lessComparisonVal)

	numberVal, valFound := recFieldVals.GetNumberFieldValue(filterParams.FieldID)
	if !valFound {
		return false, nil
	}
	log.Printf("filterLess: value = %v", numberVal)

	if numberVal < *lessComparisonVal {
		return true, nil
	} else {
		return false, nil
	}
}

func filterLessEqual(filterParams FilterFuncParams, recFieldVals record.RecFieldValues) (bool, error) {
	lessComparisonVal := filterParams.ConditionMap.getNumberConditionParam(conditionLessEqual)
	if lessComparisonVal == nil {
		return false, nil
	}

	log.Printf("filterLess: comparison value = %v", *lessComparisonVal)

	numberVal, valFound := recFieldVals.GetNumberFieldValue(filterParams.FieldID)
	if !valFound {
		return false, nil
	}
	log.Printf("filterLess: value = %v", numberVal)

	if numberVal <= *lessComparisonVal {
		return true, nil
	} else {
		return false, nil
	}
}

func filterEqual(filterParams FilterFuncParams, recFieldVals record.RecFieldValues) (bool, error) {
	equalComparisonVal := filterParams.ConditionMap.getNumberConditionParam(conditionEqual)
	if equalComparisonVal == nil {
		return false, nil
	}

	numberVal, valFound := recFieldVals.GetNumberFieldValue(filterParams.FieldID)
	if !valFound {
		return false, nil
	}

	if numberVal == *equalComparisonVal {
		return true, nil
	} else {
		return false, nil
	}
}

func filterAny(filterParams FilterFuncParams, recFieldVals record.RecFieldValues) (bool, error) {
	return true, nil // Always return true
}

func filterCustomDateRange(filterParams FilterFuncParams, recFieldVals record.RecFieldValues) (bool, error) {

	startDate := filterParams.ConditionMap.getDateConditionParam(conditionDateRangeMinDate)
	endDate := filterParams.ConditionMap.getDateConditionParam(conditionDateRangeMaxDate)
	if startDate == nil || endDate == nil {
		return false, nil
	}
	log.Printf("date range filter: start date = %v, end date = %v", *startDate, *endDate)

	timeVal, valFound := recFieldVals.GetTimeFieldValue(filterParams.FieldID)
	if !valFound {
		return false, nil
	}

	log.Printf("date range filter: date val = %v", timeVal)

	if timeVal.After(*startDate) && timeVal.Before(*endDate) {
		return true, nil
	} else {
		return false, nil
	}

	return true, nil // stubbed out
}

func filterBefore(filterParams FilterFuncParams, recFieldVals record.RecFieldValues) (bool, error) {

	endDate := filterParams.ConditionMap.getDateConditionParam(conditionDateRangeMaxDate)
	if endDate == nil {
		return false, nil
	}

	timeVal, valFound := recFieldVals.GetTimeFieldValue(filterParams.FieldID)
	if !valFound {
		return false, nil
	}

	log.Printf("date range filter: date val = %v", timeVal)

	if timeVal.Before(*endDate) {
		return true, nil
	} else {
		return false, nil
	}

	return true, nil // stubbed out
}

func filterAfter(filterParams FilterFuncParams, recFieldVals record.RecFieldValues) (bool, error) {

	startDate := filterParams.ConditionMap.getDateConditionParam(conditionDateRangeMinDate)
	if startDate == nil {
		return false, nil
	}

	timeVal, valFound := recFieldVals.GetTimeFieldValue(filterParams.FieldID)
	if !valFound {
		return false, nil
	}

	log.Printf("date range filter: date val = %v", timeVal)

	if timeVal.After(*startDate) {
		return true, nil
	} else {
		return false, nil
	}

	return true, nil // stubbed out
}

func filterTrue(filterParams FilterFuncParams, recFieldVals record.RecFieldValues) (bool, error) {

	boolVal, valFound := recFieldVals.GetBoolFieldValue(filterParams.FieldID)
	if !valFound {
		return false, nil
	}

	if boolVal == true {
		return true, nil
	} else {
		return false, nil
	}

}

func filterNotTrue(filterParams FilterFuncParams, recFieldVals record.RecFieldValues) (bool, error) {

	isTrue, err := filterTrue(filterParams, recFieldVals)
	if err != nil {
		return false, err
	} else {
		if !isTrue {
			return true, nil
		} else {
			return false, nil
		}
	}
}

func filterFalse(filterParams FilterFuncParams, recFieldVals record.RecFieldValues) (bool, error) {

	boolVal, valFound := recFieldVals.GetBoolFieldValue(filterParams.FieldID)
	if !valFound {
		return false, nil
	}

	if boolVal == false {
		return true, nil
	} else {
		return false, nil
	}

}

func filterNotFalse(filterParams FilterFuncParams, recFieldVals record.RecFieldValues) (bool, error) {

	isFalse, err := filterFalse(filterParams, recFieldVals)
	if err != nil {
		return false, err
	} else {
		if !isFalse {
			return true, nil
		} else {
			return false, nil
		}
	}
}

func filterTextContains(filterParams FilterFuncParams, recFieldVals record.RecFieldValues) (bool, error) {

	textVal, valFound := recFieldVals.GetTextFieldValue(filterParams.FieldID)
	if !valFound {
		return false, nil
	}

	matchParam := filterParams.ConditionMap.getTextConditionParam(conditionTextContains)

	if matchParam == nil {
		return false, nil
	} else {
		if strings.Contains(textVal, *matchParam) {
			return true, nil
		} else {
			return false, nil
		}
	}

}

func filterTags(filterParams FilterFuncParams, recFieldVals record.RecFieldValues) (bool, error) {

	matchTags := filterParams.ConditionMap.getTagsConditionParam(conditionMatchTags)
	if matchTags == nil {
		return false, nil
	} else if len(matchTags) == 0 {
		// If there are no tags to match against, match all items
		return true, nil
	}

	tagVals, valFound := recFieldVals.GetTagsFieldValue(filterParams.FieldID)
	if !valFound || tagVals == nil {
		return false, nil
	}

	matchAnyTag := func() bool {
		for _, currTagVal := range tagVals {
			for _, currMatchTagVal := range matchTags {
				if currTagVal == currMatchTagVal {
					return true
				}
			}
		}
		return false
	}

	return matchAnyTag(), nil

}

type RuleIDFilterFuncMap map[string]FilterRuleFunc

var textFieldFilterRuleDefs = RuleIDFilterFuncMap{
	filterRuleIDNotBlank: filterNonBlankField,
	filterRuleIDBlank:    filterBlankField,
	filterRuleIDContains: filterTextContains}

var numberFieldFilterRuleDefs = RuleIDFilterFuncMap{
	filterRuleIDAny:          filterAny,
	filterRuleIDNotBlank:     filterNonBlankField,
	filterRuleIDBlank:        filterBlankField,
	filterRuleIDGreater:      filterGreater,
	filterRuleIDGreaterEqual: filterGreaterEqual,
	filterRuleIDLess:         filterLess,
	filterRuleIDLessEqual:    filterLessEqual,
	filterRuleIDEqual:        filterEqual}

var timeFieldFilterRuleDefs = RuleIDFilterFuncMap{
	filterRuleIDCustomDateRange: filterCustomDateRange,
	filterRuleAfter:             filterAfter,
	filterRuleBefore:            filterBefore,
	filterRuleIDAny:             filterAny,
	filterRuleIDNotBlank:        filterNonBlankField,
	filterRuleIDBlank:           filterBlankField}

var boolFieldFilterRuleDefs = RuleIDFilterFuncMap{
	filterRuleIDAny:      filterAny,
	filterRuleIDNotBlank: filterNonBlankField,
	filterRuleIDBlank:    filterBlankField,
	filterRuleTrue:       filterTrue,
	filterRuleNotTrue:    filterNotTrue,
	filterRuleFalse:      filterFalse,
	filterRuleNotFalse:   filterNotFalse}

var tagFieldFilterRuleDefs = RuleIDFilterFuncMap{
	filterRuleIDTags: filterTags}

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
	case field.FieldTypeBool:
		filterFunc, funcFound := boolFieldFilterRuleDefs[ruleID]
		if !funcFound {
			return nil, fmt.Errorf(
				"getRuleDefByFieldType: Failed to retrieve filter rule definition for field type = %v, unrecognized rule ID = %v",
				fieldType, ruleID)
		} else {
			return filterFunc, nil
		}
	case field.FieldTypeLabel:
		filterFunc, funcFound := tagFieldFilterRuleDefs[ruleID]
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
