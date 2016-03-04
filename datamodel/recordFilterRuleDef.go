package datamodel

import (
	"fmt"
)

type FilterRuleDef struct {
	RuleID   string `json:"ruleID"`
	HasParam bool   `json:"hasParam"`
	DataType string `json:"dataType"`
	Label    string `json:"label"`
}

const filterRuleIDNotBlank string = "notBlank"
const filterRuleIDBlank string = "isBlank"

type RuleIDRuleDefMap map[string]FilterRuleDef

var textFieldFilterRuleDefs = RuleIDRuleDefMap{
	filterRuleIDNotBlank: FilterRuleDef{filterRuleIDNotBlank, false, fieldTypeText, "Text is set (not blank)"},
	filterRuleIDBlank:    FilterRuleDef{filterRuleIDBlank, false, fieldTypeText, "Text is not set (blank)"},
}

var numberFieldFilterRuleDefs = RuleIDRuleDefMap{
	filterRuleIDNotBlank: FilterRuleDef{filterRuleIDNotBlank, false, fieldTypeNumber, "Value is set (not blank)"},
	filterRuleIDBlank:    FilterRuleDef{filterRuleIDBlank, false, fieldTypeNumber, "Value is not set (blank)"},
}

var FilterRuleDefs struct {
	TextFieldRules   RuleIDRuleDefMap `json:"textFieldRules"`
	NumberFieldRules RuleIDRuleDefMap `json:"numberFieldRules"`
}

// Get the rule definition based upon the field type
func getRuleDefByFieldType(fieldType string, ruleID string) (*FilterRuleDef, error) {
	switch fieldType {
	case fieldTypeText:
		ruleDef, ruleDefFound := textFieldFilterRuleDefs[ruleID]
		if !ruleDefFound {
			return nil, fmt.Errorf(
				"getRuleDefByFieldType: Failed to retrieve filter rule definition for field type = %v, rule ID = %v",
				fieldType, ruleID)
		} else {
			return &ruleDef, nil
		}
	case fieldTypeNumber:
		ruleDef, ruleDefFound := numberFieldFilterRuleDefs[ruleID]
		if !ruleDefFound {
			return nil, fmt.Errorf(
				"getRuleDefByFieldType: Failed to retrieve filter rule definition for field type = %v, rule ID = %v",
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
