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

type FilterRuleDef struct {
	RuleID     string         `json:"ruleID"`
	HasParam   bool           `json:"hasParam"`
	DataType   string         `json:"dataType"`
	Label      string         `json:"label"`
	filterFunc FilterRuleFunc `json:"-"`
}

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

type RuleIDRuleDefMap map[string]FilterRuleDef

var textFieldFilterRuleDefs = RuleIDRuleDefMap{
	filterRuleIDNotBlank: FilterRuleDef{filterRuleIDNotBlank, false, field.FieldTypeText, "Text is set (not blank)", filterNonBlankField},
	filterRuleIDBlank:    FilterRuleDef{filterRuleIDBlank, false, field.FieldTypeText, "Text is not set (blank)", filterBlankField},
}

var numberFieldFilterRuleDefs = RuleIDRuleDefMap{
	filterRuleIDNotBlank: FilterRuleDef{filterRuleIDNotBlank, false, field.FieldTypeNumber, "Value is set (not blank)", filterNonBlankField},
	filterRuleIDBlank:    FilterRuleDef{filterRuleIDBlank, false, field.FieldTypeNumber, "Value is not set (blank)", filterBlankField},
	filterRuleIDGreater:  FilterRuleDef{filterRuleIDGreater, false, field.FieldTypeNumber, "Value is greater", filterGreater},
	filterRuleIDLess:     FilterRuleDef{filterRuleIDLess, false, field.FieldTypeNumber, "Value is greater", filterLess},
}

var timeFieldFilterRuleDefs = RuleIDRuleDefMap{
	filterRuleIDCustomDateRange: FilterRuleDef{filterRuleIDCustomDateRange, false, field.FieldTypeTime, "Date Range", filterCustomDateRange},
	filterRuleIDAny:             FilterRuleDef{filterRuleIDAny, false, field.FieldTypeTime, "Any Date", filterAny},
}

var FilterRuleDefs struct {
	TextFieldRules   RuleIDRuleDefMap `json:"textFieldRules"`
	NumberFieldRules RuleIDRuleDefMap `json:"numberFieldRules"`
}

// Get the rule definition based upon the field type
func getRuleDefByFieldType(fieldType string, ruleID string) (*FilterRuleDef, error) {
	switch fieldType {
	case field.FieldTypeText:
		ruleDef, ruleDefFound := textFieldFilterRuleDefs[ruleID]
		if !ruleDefFound {
			return nil, fmt.Errorf(
				"getRuleDefByFieldType: Failed to retrieve filter rule definition for field type = %v, unrecognized rule ID = %v",
				fieldType, ruleID)
		} else {
			return &ruleDef, nil
		}
	case field.FieldTypeNumber:
		ruleDef, ruleDefFound := numberFieldFilterRuleDefs[ruleID]
		if !ruleDefFound {
			return nil, fmt.Errorf(
				"getRuleDefByFieldType: Failed to retrieve filter rule definition for field type = %v, unrecognized rule ID = %v",
				fieldType, ruleID)
		} else {
			return &ruleDef, nil
		}
	case field.FieldTypeTime:
		ruleDef, ruleDefFound := timeFieldFilterRuleDefs[ruleID]
		if !ruleDefFound {
			return nil, fmt.Errorf(
				"getRuleDefByFieldType: Failed to retrieve filter rule definition for field type = %v, unrecognized rule ID = %v",
				fieldType, ruleID)
		} else {
			return &ruleDef, nil
		}
	default:
		return nil, fmt.Errorf(
			"getRuleDefByFieldType: Failed to retrieve filter rule definition: unknown field type = %v",
			fieldType)
	}
}
