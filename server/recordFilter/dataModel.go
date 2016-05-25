package recordFilter

import (
	"appengine"
	"appengine/datastore"
	"fmt"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/datastoreWrapper"
	"resultra/datasheet/server/table"
)

const recordFilterRuleEntityKind string = "RecordFilterRule"
const recordFilterEntityKind string = "RecordFilter"

type RecordFilter struct {
	Name string
	/* Rules are managed as child entities */
}

type RecordFilterRef struct {
	FilterID string `json:"filterID"`
	Name     string `json:"name"`
}

type RecordFilterRule struct {
	FilterRuleID    string
	Field           *datastore.Key
	RuleID          string
	TextRuleParam   string
	NumberRuleParam float64
}

type FilterRuleRef struct {
	FilterRuleID    string         `json:"filterRuleID"`
	FieldRef        field.FieldRef `json:"fieldRef"`
	FilterRuleDef   FilterRuleDef  `json:"filterRuleDef"`
	TextRuleParam   *string        `json:"textRuleParam,omitempty"`
	NumberRuleParam *float64       `json:"numberRuleParam,omitempty"`
}

type NewFilterParams struct {
	ParentTableID string `json:"parentTableID"`
	Name          string `json:"name"`
}

func getExistingFilterNames(appEngContext appengine.Context, parentTableID string) ([]string, error) {

	var allFilters []RecordFilter
	_, err := datastoreWrapper.GetAllChildEntities(appEngContext, parentTableID,
		recordFilterEntityKind, &allFilters)
	if err != nil {
		return nil, fmt.Errorf("validateUnusedFilterName: Unable to retrieve fields from datastore: datastore error =%v", err)
	}

	existingNames := []string{}
	for _, currFilter := range allFilters {
		existingNames = append(existingNames, currFilter.Name)
	}
	return existingNames, nil

}

func validateUnusedFilterName(appEngContext appengine.Context, parentTableID string, filterName string) error {

	existingNames, getNamesErr := getExistingFilterNames(appEngContext, parentTableID)
	if getNamesErr != nil {
		return getNamesErr
	}
	for _, currName := range existingNames {
		if currName == filterName {
			return fmt.Errorf("validateUnusedFilterName: Filter name '%v' already used", filterName)
		}
	}
	return nil
}

func newFilter(appEngContext appengine.Context, params NewFilterParams) (*RecordFilterRef, error) {

	if tableIDErr := datastoreWrapper.ValidateEntityKind(params.ParentTableID, table.TableEntityKind); tableIDErr != nil {
		return nil, fmt.Errorf("newFilter: %v", tableIDErr)
	}

	sanitizedName, sanitizeErr := generic.SanitizeName(params.Name)
	if sanitizeErr != nil {
		return nil, fmt.Errorf("newFilter: %v", sanitizeErr)
	}

	if validateNameErr := validateUnusedFilterName(appEngContext, params.ParentTableID, sanitizedName); validateNameErr != nil {
		return nil, fmt.Errorf("newFilter: %v", validateNameErr)
	}

	newFilter := RecordFilter{Name: sanitizedName}

	filterID, insertErr := datastoreWrapper.InsertNewChildEntity(appEngContext, params.ParentTableID, recordFilterEntityKind, &newFilter)
	if insertErr != nil {
		return nil, insertErr
	}

	filterRef := RecordFilterRef{FilterID: filterID, Name: newFilter.Name}

	return &filterRef, nil
}

type GetFilterRulesParams struct {
	ParentFilterID string `json:"parentFilterID"`
}

func getRecordFilterRuleRefs(appEngContext appengine.Context, params GetFilterRulesParams) ([]FilterRuleRef, error) {

	if err := datastoreWrapper.ValidateEntityKind(params.ParentFilterID, recordFilterEntityKind); err != nil {
		return nil, fmt.Errorf("getRecordFilterRefs: Invalid filter ID: %v", err)
	}

	var allFilterRules []RecordFilterRule
	filterRuleIDs, err := datastoreWrapper.GetAllChildEntities(
		appEngContext, params.ParentFilterID, recordFilterRuleEntityKind, &allFilterRules)
	if err != nil {
		return nil, fmt.Errorf("getRecordFilterRefs: Unable to retrieve record filters from datastore: datastore error =%v", err)
	}

	filterRefs := make([]FilterRuleRef, len(allFilterRules))
	for i, currFilterRule := range allFilterRules {
		filterRuleID := filterRuleIDs[i]

		filterRuleRef, refErr := createOneFilterRef(appEngContext, filterRuleID, currFilterRule)
		if refErr != nil {
			return nil, fmt.Errorf("getRecordFilterRefs: Unable to create reference to filter rule: error =%v", refErr)
		}

		filterRefs[i] = *filterRuleRef
	}
	return filterRefs, nil
}

type GetFilterListParams struct {
	ParentTableID string `json:"parentTableID"`
}

func getFilterList(appEngContext appengine.Context, params GetFilterListParams) ([]RecordFilterRef, error) {

	if err := datastoreWrapper.ValidateEntityKind(params.ParentTableID, table.TableEntityKind); err != nil {
		return nil, fmt.Errorf("getFilterList: Invalid table ID: %v", err)
	}

	var allFilters []RecordFilter
	filterIDs, err := datastoreWrapper.GetAllChildEntities(appEngContext, params.ParentTableID,
		recordFilterEntityKind, &allFilters)
	if err != nil {
		return nil, fmt.Errorf("GetRecordFilterRefs: Unable to retrieve record filters from datastore: datastore error =%v", err)
	}

	filterRefs := make([]RecordFilterRef, len(allFilters))
	for currFilterIndex, currFilter := range allFilters {
		filterID := filterIDs[currFilterIndex]

		filterRef := RecordFilterRef{
			FilterID: filterID,
			Name:     currFilter.Name}

		filterRefs[currFilterIndex] = filterRef

	}

	return filterRefs, nil

}

type NewFilterRuleParams struct {
	ParentFilterID  string   `jsaon:"parentFilterID"`
	FieldID         string   `json:"fieldID"`
	RuleID          string   `json:"ruleID"`
	TextRuleParam   *string  `json:"textRuleParam,omitempty"`
	NumberRuleParam *float64 `json:"numberRuleParam,omitempty"`
}

func newFilterRule(appEngContext appengine.Context, newRuleParams NewFilterRuleParams) (*FilterRuleRef, error) {

	if validateParentErr := datastoreWrapper.ValidateSameParentEntity(newRuleParams.FieldID,
		newRuleParams.ParentFilterID, table.TableEntityKind); validateParentErr != nil {
		return nil, fmt.Errorf("newFilterRule: Can't create new rule: %v", validateParentErr)
	}

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

	filterRuleID, insertErr := datastoreWrapper.InsertNewChildEntity(appEngContext,
		newRuleParams.ParentFilterID, recordFilterRuleEntityKind, &newFilter)
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
	ruleID string, filterRule RecordFilterRule) (*FilterRuleRef, error) {

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
		FilterRuleID:    ruleID,
		FieldRef:        *fieldRef,
		FilterRuleDef:   *filterRuleDef,
		TextRuleParam:   optParamVal.textParam,
		NumberRuleParam: optParamVal.numberParam}

	return &filterRuleRef, nil

}
