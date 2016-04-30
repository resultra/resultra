package recordFilter

import (
	"appengine"
	"appengine/datastore"
	"fmt"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/generic/datastoreWrapper"
)

const recordFilterRuleEntityKind string = "RecordFilterRule"

type RecordFilterRule struct {
	Field           *datastore.Key
	RuleID          string
	TextRuleParam   string
	NumberRuleParam float64
}

type NewFilterRuleParams struct {
	FieldID         string   `json:"fieldID"`
	RuleID          string   `json:"ruleID"`
	TextRuleParam   *string  `json:"textRuleParam,omitempty"`
	NumberRuleParam *float64 `json:"numberRuleParam,omitempty"`
}

type FilterRuleRef struct {
	FilterRuleID    string         `json:"filterRuleID"`
	FieldRef        field.FieldRef `json:"fieldRef"`
	FilterRuleDef   FilterRuleDef  `json:"filterRuleDef"`
	TextRuleParam   *string        `json:"textRuleParam,omitempty"`
	NumberRuleParam *float64       `json:"numberRuleParam,omitempty"`
}

func NewFilterRule(appEngContext appengine.Context, newRuleParams NewFilterRuleParams) (*FilterRuleRef, error) {

	fieldKey, fieldRef, fieldErr := field.GetExistingFieldRefAndKey(appEngContext, newRuleParams.FieldID)
	if fieldErr != nil {
		return nil, fmt.Errorf("NewFilterRule: Can't get field for filter: datastore error = %v", fieldErr)
	}

	// Verify there is a valid rule definition for the given field's data type, then build up the
	// new filter based upon the rule definition and field type.
	var newFilter RecordFilterRule
	var newRuleDef FilterRuleDef

	switch fieldRef.FieldInfo.Type {

	case field.FieldTypeText:
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

	case field.FieldTypeNumber:
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

	filterRuleID, insertErr := datastoreWrapper.InsertNewRootEntity(appEngContext, recordFilterRuleEntityKind, &newFilter)
	if insertErr != nil {
		return nil, fmt.Errorf("NewFilterRule: Can't create new filter: error inserting into datastore: %v", insertErr)
	}

	// Set the parameter value in returned filter rule reference. What is set in the structure
	// depends on what type of data the filter rule works with and if the filtering rule has a parameter.

	optParamVal, paramValErr := getOptionalParamValueByRuleDef(newRuleDef, newFilter)
	if paramValErr != nil {
		return nil, fmt.Errorf("NewFilterRule: Failed to retrieve filter rule parameter values: err=%v", paramValErr)
	}

	return &FilterRuleRef{filterRuleID, *fieldRef, newRuleDef, optParamVal.textParam, optParamVal.numberParam}, nil

}

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

// Convert one filter rule from the datastore into a reference which is usable by
// API clients
func createOneFilterRef(appEngContext appengine.Context,
	filterRuleID string, filterRule RecordFilterRule) (*FilterRuleRef, error) {

	fieldRef, fieldErr := field.GetFieldFromKey(appEngContext, filterRule.Field)
	if fieldErr != nil {
		return nil, fmt.Errorf("createOneFilterRef: Failed to retrieve field for filter rule: field key=%+v, encode err=%v",
			filterRule.Field, fieldErr)
	}

	filterRuleDef, ruleDefErr := getRuleDefByFieldType(fieldRef.FieldInfo.Type, filterRule.RuleID)
	if ruleDefErr != nil {
		return nil, fmt.Errorf("createOneFilterRef: Failed to retrieve filter rule definition: err=%v", ruleDefErr)
	}

	optParamVal, paramValErr := getOptionalParamValueByRuleDef(*filterRuleDef, filterRule)
	if paramValErr != nil {
		return nil, fmt.Errorf("createOneFilterRef: Failed to retrieve filter rule parameter values: err=%v", paramValErr)
	}

	filterRuleRef := FilterRuleRef{
		FilterRuleID:    filterRuleID,
		FieldRef:        *fieldRef,
		FilterRuleDef:   *filterRuleDef,
		TextRuleParam:   optParamVal.textParam,
		NumberRuleParam: optParamVal.numberParam}

	return &filterRuleRef, nil

}

func GetRecordFilterRefs(appEngContext appengine.Context) ([]FilterRuleRef, error) {

	var allFilterRules []RecordFilterRule
	filterRuleIDs, err := datastoreWrapper.GetAllRootEntities(
		appEngContext, recordFilterRuleEntityKind, &allFilterRules)
	if err != nil {
		return nil, fmt.Errorf("GetRecordFilterRefs: Unable to retrieve record filters from datastore: datastore error =%v", err)
	}

	filterRefs := make([]FilterRuleRef, len(allFilterRules))
	for i, currFilterRule := range allFilterRules {
		filterRuleID := filterRuleIDs[i]

		filterRuleRef, refErr := createOneFilterRef(appEngContext, filterRuleID, currFilterRule)
		if refErr != nil {
			return nil, fmt.Errorf("GetRecordFilterRefs: Unable to create reference to filter rule: error =%v", refErr)
		}

		filterRefs[i] = *filterRuleRef
	}
	return filterRefs, nil
}
