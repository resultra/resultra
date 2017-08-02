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
const filterRuleIDCustomDateRange string = "dateRange"
const filterRuleIDGreater string = "greater"
const filterRuleIDLess string = "less"
const filterRuleIDAny string = "any"
const filterRuleTrue string = "true"
const filterRuleNotTrue string = "notTrue"
const filterRuleFalse string = "false"
const filterRuleNotFalse string = "notFalse"
const filterRuleIDContains string = "contains"

const conditionDateRangeMinDate string = "minDate"
const conditionDateRangeMaxDate string = "maxDate"
const conditionGreater string = "greater"
const conditionLess string = "less"
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

type RuleIDFilterFuncMap map[string]FilterRuleFunc

var textFieldFilterRuleDefs = RuleIDFilterFuncMap{
	filterRuleIDNotBlank: filterNonBlankField,
	filterRuleIDBlank:    filterBlankField,
	filterRuleIDContains: filterTextContains}

var numberFieldFilterRuleDefs = RuleIDFilterFuncMap{
	filterRuleIDAny:      filterAny,
	filterRuleIDNotBlank: filterNonBlankField,
	filterRuleIDBlank:    filterBlankField,
	filterRuleIDGreater:  filterGreater,
	filterRuleIDLess:     filterLess}

var timeFieldFilterRuleDefs = RuleIDFilterFuncMap{
	filterRuleIDCustomDateRange: filterCustomDateRange,
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
	default:
		return nil, fmt.Errorf(
			"getRuleDefByFieldType: Failed to retrieve filter rule definition: unknown field type = %v",
			fieldType)
	}
}
