package recordFilter

import (
	"fmt"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/record"
)

type FilterFuncParams struct {
	fieldID        string
	textParamVal   *string
	numberParamVal *float64
}

type FilterRuleFunc func(filterParams FilterFuncParams, record record.Record) (bool, error)

type FilterRuleDef struct {
	RuleID     string         `json:"ruleID"`
	HasParam   bool           `json:"hasParam"`
	DataType   string         `json:"dataType"`
	Label      string         `json:"label"`
	filterFunc FilterRuleFunc `json:"-"`
}

const filterRuleIDNotBlank string = "notBlank"
const filterRuleIDBlank string = "isBlank"

func filterBlankField(filterParams FilterFuncParams, record record.Record) (bool, error) {

	valueIsSet := record.ValueIsSet(filterParams.fieldID)

	if valueIsSet {
		return false, nil
	} else {
		return true, nil
	}
}

func filterNonBlankField(filterParams FilterFuncParams, record record.Record) (bool, error) {

	valueIsSet := record.ValueIsSet(filterParams.fieldID)

	if valueIsSet {
		return true, nil
	} else {
		return false, nil
	}
}

type RuleIDRuleDefMap map[string]FilterRuleDef

var textFieldFilterRuleDefs = RuleIDRuleDefMap{
	filterRuleIDNotBlank: FilterRuleDef{filterRuleIDNotBlank, false, field.FieldTypeText, "Text is set (not blank)", filterNonBlankField},
	filterRuleIDBlank:    FilterRuleDef{filterRuleIDBlank, false, field.FieldTypeText, "Text is not set (blank)", filterBlankField},
}

var numberFieldFilterRuleDefs = RuleIDRuleDefMap{
	filterRuleIDNotBlank: FilterRuleDef{filterRuleIDNotBlank, false, field.FieldTypeNumber, "Value is set (not blank)", filterNonBlankField},
	filterRuleIDBlank:    FilterRuleDef{filterRuleIDBlank, false, field.FieldTypeNumber, "Value is not set (blank)", filterBlankField},
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
				"getRuleDefByFieldType: Failed to retrieve filter rule definition for field type = %v, rule ID = %v",
				fieldType, ruleID)
		} else {
			return &ruleDef, nil
		}
	case field.FieldTypeNumber:
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
