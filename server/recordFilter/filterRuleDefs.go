// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package recordFilter

import (
	"fmt"
	"resultra/tracker/server/field"
	"resultra/tracker/server/record"
	"strings"
)

type FilterFuncParams struct {
	Conditions   []FilterRuleCondition
	FieldID      string
	ConditionMap FilterConditionMap
	CurrUserID   string
}

type FilterRuleFunc func(filterParams FilterFuncParams, recFieldVals record.RecFieldValues) (bool, error)

const filterRuleIDNotBlank string = "notBlank"
const filterRuleIDBlank string = "isBlank"

const filterRuleIDTags string = "tags"
const conditionMatchTags string = "tags"
const conditionMatchLogic string = "logic"
const matchLogicAll string = "all"
const matchLogicAny string = "any"

const filterRuleIDCurrentUser string = "currentUser"
const filterRuleIDSpecificUsers string = "specificUsers"
const conditionUsers string = "users"

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

	numberVal, valFound := recFieldVals.GetNumberFieldValue(filterParams.FieldID)
	if !valFound {
		return false, nil
	}

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

	numberVal, valFound := recFieldVals.GetNumberFieldValue(filterParams.FieldID)
	if !valFound {
		return false, nil
	}

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

	numberVal, valFound := recFieldVals.GetNumberFieldValue(filterParams.FieldID)
	if !valFound {
		return false, nil
	}

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

	numberVal, valFound := recFieldVals.GetNumberFieldValue(filterParams.FieldID)
	if !valFound {
		return false, nil
	}

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

	timeVal, valFound := recFieldVals.GetTimeFieldValue(filterParams.FieldID)
	if !valFound {
		return false, nil
	}

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

	matchOneTag := func(matchTag string) bool {
		for _, currTagVal := range tagVals {
			if currTagVal == matchTag {
				return true
			}
		}
		return false
	}

	matchAnyTag := func() bool {
		// Test if any tag in the filter matches a value in the current record
		for _, currMatchTagVal := range matchTags {
			if matchOneTag(currMatchTagVal) {
				return true
			}
		}
		return false
	}

	matchAllTags := func() bool {
		// Test if every tag in the filter matches a value in the current record
		for _, currMatchTagVal := range matchTags {
			if !matchOneTag(currMatchTagVal) {
				return false
			}
		}
		return true
	}

	matchLogic := filterParams.ConditionMap.getTextConditionParam(conditionMatchLogic)

	if (matchLogic == nil) || (*matchLogic == matchLogicAny) {
		return matchAnyTag(), nil
	} else {
		return matchAllTags(), nil
	}

}

func filterBlankTagField(filterParams FilterFuncParams, recFieldVals record.RecFieldValues) (bool, error) {

	tagVals, valFound := recFieldVals.GetTagsFieldValue(filterParams.FieldID)
	if !valFound || tagVals == nil {
		return true, nil
	} else if len(tagVals) == 0 {
		return true, nil
	} else {
		return false, nil
	}

}

func filterNonBlankTagField(filterParams FilterFuncParams, recFieldVals record.RecFieldValues) (bool, error) {
	isBlank, err := filterBlankTagField(filterParams, recFieldVals)
	return !isBlank, err
}

func filterSpecificUsers(filterParams FilterFuncParams, recFieldVals record.RecFieldValues) (bool, error) {

	matchUsers := filterParams.ConditionMap.getUsersConditionParam(conditionUsers)
	if matchUsers == nil {
		return false, nil
	} else if len(matchUsers) == 0 {
		// If there are no tags to match against, match all items
		return true, nil
	}

	userVals, valFound := recFieldVals.GetUsersFieldValue(filterParams.FieldID)
	if !valFound || userVals == nil {
		return false, nil
	}

	matchOneUser := func(matchUserID string) bool {
		for _, currUserIDVal := range userVals {
			if currUserIDVal == matchUserID {
				return true
			}
		}
		return false
	}

	matchAnyUser := func() bool {
		// Test if any tag in the filter matches a value in the current record
		for _, currMatchUserID := range matchUsers {
			if matchOneUser(currMatchUserID) {
				return true
			}
		}
		return false
	}

	return matchAnyUser(), nil

}

func filterBlankUserField(filterParams FilterFuncParams, recFieldVals record.RecFieldValues) (bool, error) {

	userVals, valFound := recFieldVals.GetUsersFieldValue(filterParams.FieldID)
	if !valFound || userVals == nil {
		return true, nil
	} else if len(userVals) == 0 {
		return true, nil
	} else {
		return false, nil
	}

}

func filterNonBlankUserField(filterParams FilterFuncParams, recFieldVals record.RecFieldValues) (bool, error) {
	isBlank, err := filterBlankUserField(filterParams, recFieldVals)
	return !isBlank, err
}

func filterCurrentUser(filterParams FilterFuncParams, recFieldVals record.RecFieldValues) (bool, error) {

	userVals, valFound := recFieldVals.GetUsersFieldValue(filterParams.FieldID)
	if !valFound || userVals == nil {
		return false, nil
	}

	for _, currUserIDVal := range userVals {
		if currUserIDVal == filterParams.CurrUserID {
			return true, nil
		}
	}
	return false, nil

}

type RuleIDFilterFuncMap map[string]FilterRuleFunc

var textFieldFilterRuleDefs = RuleIDFilterFuncMap{
	filterRuleIDAny:      filterAny,
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
	filterRuleIDTags:     filterTags,
	filterRuleIDNotBlank: filterNonBlankTagField,
	filterRuleIDBlank:    filterBlankTagField}

var userFieldFilterRuleDefs = RuleIDFilterFuncMap{
	filterRuleIDSpecificUsers: filterSpecificUsers,
	filterRuleIDCurrentUser:   filterCurrentUser,
	filterRuleIDNotBlank:      filterNonBlankUserField,
	filterRuleIDBlank:         filterBlankUserField}

var noteFieldFilterRuleDefs = RuleIDFilterFuncMap{
	filterRuleIDAny:      filterAny,
	filterRuleIDNotBlank: filterNonBlankField,
	filterRuleIDBlank:    filterBlankField,
	filterRuleIDContains: filterTextContains}

var emailFieldFilterRuleDefs = RuleIDFilterFuncMap{
	filterRuleIDAny:      filterAny,
	filterRuleIDNotBlank: filterNonBlankField,
	filterRuleIDBlank:    filterBlankField,
	filterRuleIDContains: filterTextContains}

var urlFieldFilterRuleDefs = RuleIDFilterFuncMap{
	filterRuleIDAny:      filterAny,
	filterRuleIDNotBlank: filterNonBlankField,
	filterRuleIDBlank:    filterBlankField,
	filterRuleIDContains: filterTextContains}

var imageFieldFilterRuleDefs = RuleIDFilterFuncMap{
	filterRuleIDAny:      filterAny,
	filterRuleIDNotBlank: filterNonBlankField,
	filterRuleIDBlank:    filterBlankField}

var fileFieldFilterRuleDefs = RuleIDFilterFuncMap{
	filterRuleIDAny:      filterAny,
	filterRuleIDNotBlank: filterNonBlankField,
	filterRuleIDBlank:    filterBlankField}

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
	case field.FieldTypeUser:
		filterFunc, funcFound := userFieldFilterRuleDefs[ruleID]
		if !funcFound {
			return nil, fmt.Errorf(
				"getRuleDefByFieldType: Failed to retrieve filter rule definition for field type = %v, unrecognized rule ID = %v",
				fieldType, ruleID)
		} else {
			return filterFunc, nil
		}
	case field.FieldTypeLongText:
		filterFunc, funcFound := noteFieldFilterRuleDefs[ruleID]
		if !funcFound {
			return nil, fmt.Errorf(
				"getRuleDefByFieldType: Failed to retrieve filter rule definition for field type = %v, unrecognized rule ID = %v",
				fieldType, ruleID)
		} else {
			return filterFunc, nil
		}
	case field.FieldTypeImage:
		filterFunc, funcFound := imageFieldFilterRuleDefs[ruleID]
		if !funcFound {
			return nil, fmt.Errorf(
				"getRuleDefByFieldType: Failed to retrieve filter rule definition for field type = %v, unrecognized rule ID = %v",
				fieldType, ruleID)
		} else {
			return filterFunc, nil
		}

	case field.FieldTypeFile:
		filterFunc, funcFound := imageFieldFilterRuleDefs[ruleID]
		if !funcFound {
			return nil, fmt.Errorf(
				"getRuleDefByFieldType: Failed to retrieve filter rule definition for field type = %v, unrecognized rule ID = %v",
				fieldType, ruleID)
		} else {
			return filterFunc, nil
		}
	case field.FieldTypeEmail:
		filterFunc, funcFound := emailFieldFilterRuleDefs[ruleID]
		if !funcFound {
			return nil, fmt.Errorf(
				"getRuleDefByFieldType: Failed to retrieve filter rule definition for field type = %v, unrecognized rule ID = %v",
				fieldType, ruleID)
		} else {
			return filterFunc, nil
		}
	case field.FieldTypeURL:
		filterFunc, funcFound := urlFieldFilterRuleDefs[ruleID]
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
