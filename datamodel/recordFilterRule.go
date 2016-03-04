package datamodel

import (
	"appengine"
	"appengine/datastore"
	"fmt"
)

const recordFilterRuleEntityKind string = "RecordFilterRule"

type RecordFilterRule struct {
	Field           *datastore.Key
	RuleID          string
	TextRuleParam   string
	NumberRuleParam float64
}

type FilterRuleDef struct {
	RuleID   string `json:"ruleID"`
	HasParam bool   `json:"hasParam"`
	DataType string `json:"dataType"`
	Label    string `json:"label"`
}

type RuleIDRuleDefMap map[string]FilterRuleDef

var textFieldFilterRuleDefs = RuleIDRuleDefMap{
	"notBlank": FilterRuleDef{"notBlank", false, fieldTypeText, "Text is set (not blank)"},
	"isBlank":  FilterRuleDef{"isBlank", false, fieldTypeText, "Text is not set (blank)"},
}

var numberFieldFilterRuleDefs = RuleIDRuleDefMap{
	"notBlank": FilterRuleDef{"notBlank", false, fieldTypeNumber, "Value is set (not blank)"},
	"isBlank":  FilterRuleDef{"isBlank", false, fieldTypeNumber, "Value is not set (blank)"},
}

var FilterRuleDefs struct {
	TextFieldRules   RuleIDRuleDefMap `json:"textFieldRules"`
	NumberFieldRules RuleIDRuleDefMap `json:"numberFieldRules"`
}

type NewFilterRuleParams struct {
	FieldID         string   `json:"fieldID"`
	RuleID          string   `json:"ruleID"`
	TextRuleParam   *string  `json:"textRuleParam,omitempty"`
	NumberRuleParam *float64 `json:"numberRuleParam,omitempty"`
}

type FilterRuleRef struct {
	FilterRuleID    string        `json:"filterRuleID"`
	FieldRef        FieldRef      `json:"fieldRef"`
	FilterRuleDef   FilterRuleDef `json:"filterRuleDef"`
	TextRuleParam   *string       `json:"textRuleParam,omitempty"`
	NumberRuleParam *float64      `json:"numberRuleParam,omitempty"`
}

func NewFilterRule(appEngContext appengine.Context, newRuleParams NewFilterRuleParams) (*FilterRuleRef, error) {

	fieldKey, fieldRef, fieldErr := GetExistingFieldRefAndKey(appEngContext, GetFieldParams{newRuleParams.FieldID})
	if fieldErr != nil {
		return nil, fmt.Errorf("NewFilterRule: Can't get field for filter: datastore error = %v", fieldErr)
	}

	// Verify there is a valid rule definition for the given field's data type, then build up the
	// new filter based upon the rule definition and field type.
	var newFilter RecordFilterRule
	var newRuleDef FilterRuleDef

	switch fieldRef.FieldInfo.Type {

	case fieldTypeText:
		ruleDef, ruleDefFound := textFieldFilterRuleDefs[newRuleParams.RuleID]
		newRuleDef = ruleDef
		if !ruleDefFound {
			return nil, fmt.Errorf("NewFilterRule: No filtering rule found for rule ID = %v", newRuleParams.RuleID)
		} else {
			if ruleDef.HasParam {
				if (newRuleParams.TextRuleParam == nil) || (len(*newRuleParams.TextRuleParam) == 0) {
					return nil, fmt.Errorf("NewFilterRule: Missing filter paramater for rule ID = %v", newRuleParams.RuleID)
				}
				newFilter = RecordFilterRule{Field: fieldKey, RuleID: newRuleParams.RuleID,
					TextRuleParam: *newRuleParams.TextRuleParam}
			} else {
				newFilter = RecordFilterRule{Field: fieldKey, RuleID: newRuleParams.RuleID, TextRuleParam: ""}
			}
		}

	case fieldTypeNumber:
		ruleDef, ruleDefFound := numberFieldFilterRuleDefs[newRuleParams.RuleID]
		newRuleDef = ruleDef
		if !ruleDefFound {
			return nil, fmt.Errorf("NewFilterRule: No filtering rule found for rule ID = %v", newRuleParams.RuleID)
		} else {
			if ruleDef.HasParam {
				if newRuleParams.NumberRuleParam == nil {
					return nil, fmt.Errorf("NewFilterRule: Missing filter paramater for rule ID = %v", newRuleParams.RuleID)
				}
				newFilter = RecordFilterRule{Field: fieldKey, RuleID: newRuleParams.RuleID,
					NumberRuleParam: *newRuleParams.NumberRuleParam}
			} else {
				newFilter = RecordFilterRule{Field: fieldKey, RuleID: newRuleParams.RuleID, NumberRuleParam: 0}
			}
		}

	default:
		return nil, fmt.Errorf("NewFilterRule: Filtering not supported on field type: %v", fieldRef.FieldInfo.Type)
	}

	// TODO - Replace nil with database parent

	filterRuleID, insertErr := insertNewEntity(appEngContext, recordFilterRuleEntityKind, nil, &newFilter)
	if insertErr != nil {
		return nil, fmt.Errorf("NewFilterRule: Can't create new filter: error inserting into datastore: %v", insertErr)
	}

	// Set the parameter value in returned filter rule reference. What is set in the structure
	// depends on what type of data the filter rule works with and if the filtering rule has a parameter.
	var textParam *string = nil
	var numberParam *float64 = nil

	if newRuleDef.HasParam {
		if newRuleDef.DataType == fieldTypeText {
			textParam = &newFilter.TextRuleParam
		} else {
			numberParam = &newFilter.NumberRuleParam
		}
	}

	return &FilterRuleRef{filterRuleID, *fieldRef, newRuleDef, textParam, numberParam}, nil

}
