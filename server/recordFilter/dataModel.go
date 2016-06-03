package recordFilter

import (
	"appengine"
	"fmt"
	"log"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/datastoreWrapper"
	"resultra/datasheet/server/generic/uniqueID"
)

const recordFilterRuleEntityKind string = "RecordFilterRule"
const recordFilterEntityKind string = "RecordFilter"

type RecordFilter struct {
	ParentTableID string `json:"parentTableID"`
	FilterID      string `json:"filterID"`
	Name          string `json:"name"`
	/* Rules are managed as child entities */
}

const filterParentTableIDFieldName string = "ParentTableID"
const recordFilterIDFieldName string = "FilterID"

type RecordFilterRule struct {
	ParentFilterID  string  `json:"parentFilterID"`
	FilterRuleID    string  `json:"filterRuleID"`
	FieldID         string  `json:"fieldID"`
	RuleID          string  `json:"ruleID"`
	TextRuleParam   string  `json:"textRuleParam"`
	NumberRuleParam float64 `json:"numberRuleParam"`
}

const filterRuleParentFilterIDFieldName string = "ParentFilterID"

type NewFilterParams struct {
	ParentTableID string `json:"parentTableID"`
	Name          string `json:"name"`
}

func getExistingFilterNames(appEngContext appengine.Context, parentTableID string) ([]string, error) {

	var allFilters, getFiltersErr = getFilterList(appEngContext, parentTableID)
	if getFiltersErr != nil {
		return nil, fmt.Errorf("Unable to retrieve filters: %v", getFiltersErr)
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

func genUnusedFilterName(appEngContext appengine.Context, parentTableID string, namePrefix string) (string, error) {

	existingNames, getNamesErr := getExistingFilterNames(appEngContext, parentTableID)
	if getNamesErr != nil {
		return "", getNamesErr
	}

	currNameLookup := map[string]bool{}
	for _, currName := range existingNames {
		currNameLookup[currName] = true
	}

	// Use prefix all by itself if it's not already in use.
	if _, nameFound := currNameLookup[namePrefix]; nameFound != true {
		return namePrefix, nil
	}

	currNameSuffix := 1
	for currNameSuffix < 100 {
		candidateName := fmt.Sprintf("%v %v", namePrefix, currNameSuffix)
		log.Printf("genUnusedFilterName: candidate name = %v", candidateName)
		if _, nameFound := currNameLookup[candidateName]; nameFound != true {
			return candidateName, nil
		}
		currNameSuffix = currNameSuffix + 1
	}

	log.Printf("genUnusedFilterName: failure generating name")
	return "", fmt.Errorf("Failure generating unused name for filter")

}

func newFilter(appEngContext appengine.Context, params NewFilterParams) (*RecordFilter, error) {

	sanitizedName, sanitizeErr := generic.SanitizeName(params.Name)
	if sanitizeErr != nil {
		return nil, fmt.Errorf("newFilter (sanitize name): %v", sanitizeErr)
	}

	if validateNameErr := validateUnusedFilterName(appEngContext, params.ParentTableID, sanitizedName); validateNameErr != nil {
		return nil, fmt.Errorf("newFilter: %v", validateNameErr)
	}

	newFilter := RecordFilter{
		ParentTableID: params.ParentTableID,
		FilterID:      uniqueID.GenerateUniqueID(),
		Name:          sanitizedName}

	insertErr := datastoreWrapper.InsertNewRootEntity(appEngContext, recordFilterEntityKind, &newFilter)
	if insertErr != nil {
		return nil, fmt.Errorf("Can't create new text box component: error inserting into datastore: %v", insertErr)
	}

	log.Printf("newFilter: Created new filter: %+v", newFilter)

	return &newFilter, nil
}

func getFilter(appEngContext appengine.Context, filterID string) (*RecordFilter, error) {
	var filterGetDest RecordFilter

	if getErr := datastoreWrapper.GetEntityByUUID(appEngContext, recordFilterEntityKind,
		recordFilterIDFieldName, filterID, &filterGetDest); getErr != nil {
		return nil, fmt.Errorf("getBarChart: Unable to image container from datastore: error = %v", getErr)
	}

	return &filterGetDest, nil
}

func updateExistingFilter(appEngContext appengine.Context, filterID string, updatedFilter *RecordFilter) (*RecordFilter, error) {

	if updateErr := datastoreWrapper.UpdateExistingEntityByUUID(appEngContext,
		filterID, recordFilterEntityKind, recordFilterIDFieldName, updatedFilter); updateErr != nil {
		return nil, fmt.Errorf("updateExistingHtmlEditor: Error updating filter: error = %v", updateErr)
	}

	return updatedFilter, nil

}

type NewFilterWithPrefixParams struct {
	ParentTableID string `json:"parentTableID"`
	NamePrefix    string `json:"namePrefix"`
}

func newFilterWithPrefix(appEngContext appengine.Context, params NewFilterWithPrefixParams) (*RecordFilter, error) {
	sanitizedPrefix, sanitizeErr := generic.SanitizeName(params.NamePrefix)
	if sanitizeErr != nil {
		return nil, fmt.Errorf("newFilterWithPrefix: %v", sanitizeErr)
	}
	newFilterName, nameGenErr := genUnusedFilterName(appEngContext, params.ParentTableID, sanitizedPrefix)
	if nameGenErr != nil {
		return nil, fmt.Errorf("newFilterWithPrefix: %v", nameGenErr)
	}
	log.Printf("newFilterWithPrefix: New filter name: %v", newFilterName)
	return newFilter(appEngContext, NewFilterParams{ParentTableID: params.ParentTableID, Name: newFilterName})

}

type GetFilterRulesParams struct {
	ParentFilterID string `json:"parentFilterID"`
}

func getRecordFilterRules(appEngContext appengine.Context, params GetFilterRulesParams) ([]RecordFilterRule, error) {

	var allFilterRules []RecordFilterRule
	getErr := datastoreWrapper.GetAllChildEntitiesWithParentUUID(appEngContext, params.ParentFilterID,
		recordFilterRuleEntityKind, filterRuleParentFilterIDFieldName, &allFilterRules)
	if getErr != nil {
		return nil, fmt.Errorf("Unable to retrieve filter rule definitions: filter id=%v", params.ParentFilterID)
	}

	return allFilterRules, nil
}

func getFilterList(appEngContext appengine.Context, parentTableID string) ([]RecordFilter, error) {

	var allFilters []RecordFilter

	getErr := datastoreWrapper.GetAllChildEntitiesWithParentUUID(appEngContext, parentTableID,
		recordFilterEntityKind, filterParentTableIDFieldName, &allFilters)
	if getErr != nil {
		return nil, fmt.Errorf("Unable to retrieve filters: parent table id=%v", parentTableID)
	}

	return allFilters, nil

}

type NewFilterRuleParams struct {
	ParentFilterID  string   `jsaon:"parentFilterID"`
	FieldID         string   `json:"fieldID"`
	RuleID          string   `json:"ruleID"`
	TextRuleParam   *string  `json:"textRuleParam,omitempty"`
	NumberRuleParam *float64 `json:"numberRuleParam,omitempty"`
}

func newFilterRule(appEngContext appengine.Context, newRuleParams NewFilterRuleParams) (*RecordFilterRule, error) {

	filterOnField, fieldErr := field.GetField(appEngContext, newRuleParams.FieldID)
	if fieldErr != nil {
		return nil, fmt.Errorf("NewImage: Can't create filter rule with field ID = '%v': datastore error=%v",
			newRuleParams.FieldID, fieldErr)
	}

	// Verify there is a valid rule definition for the given field's data type, then build up the
	// new filter based upon the rule definition and field type.
	var newFilterRule RecordFilterRule

	switch filterOnField.Type {

	case field.FieldTypeText:
		ruleDef, ruleDefFound := textFieldFilterRuleDefs[newRuleParams.RuleID]
		if !ruleDefFound {
			return nil, fmt.Errorf("NewFilterRule: No filtering rule found for rule ID = %v", newRuleParams.RuleID)
		} else {
			if ruleDef.HasParam {
				if (newRuleParams.TextRuleParam == nil) || (len(*newRuleParams.TextRuleParam) == 0) {
					return nil, fmt.Errorf("NewFilterRule: Missing filter paramater for rule ID = %v", newRuleParams.RuleID)
				}
				newFilterRule = RecordFilterRule{
					ParentFilterID: newRuleParams.ParentFilterID,
					FilterRuleID:   uniqueID.GenerateUniqueID(),
					FieldID:        newRuleParams.FieldID,
					RuleID:         newRuleParams.RuleID,
					TextRuleParam:  *newRuleParams.TextRuleParam}
			} else {
				newFilterRule = RecordFilterRule{
					ParentFilterID: newRuleParams.ParentFilterID,
					FilterRuleID:   uniqueID.GenerateUniqueID(),
					FieldID:        newRuleParams.FieldID,
					RuleID:         newRuleParams.RuleID,
					TextRuleParam:  ""}
			}
		}

	case field.FieldTypeNumber:
		ruleDef, ruleDefFound := numberFieldFilterRuleDefs[newRuleParams.RuleID]
		if !ruleDefFound {
			return nil, fmt.Errorf("NewFilterRule: No filtering rule found for rule ID = %v", newRuleParams.RuleID)
		} else {
			if ruleDef.HasParam {
				if newRuleParams.NumberRuleParam == nil {
					return nil, fmt.Errorf("NewFilterRule: Missing filter paramater for rule ID = %v", newRuleParams.RuleID)
				}
				newFilterRule = RecordFilterRule{
					ParentFilterID:  newRuleParams.ParentFilterID,
					FilterRuleID:    uniqueID.GenerateUniqueID(),
					FieldID:         newRuleParams.FieldID,
					RuleID:          newRuleParams.RuleID,
					NumberRuleParam: *newRuleParams.NumberRuleParam}
			} else {
				newFilterRule = RecordFilterRule{
					ParentFilterID:  newRuleParams.ParentFilterID,
					FilterRuleID:    uniqueID.GenerateUniqueID(),
					FieldID:         newRuleParams.FieldID,
					RuleID:          newRuleParams.RuleID,
					NumberRuleParam: 0}
			}
		}

	default:
		return nil, fmt.Errorf("NewFilterRule: Filtering not supported on field type: %v", filterOnField.Type)
	}

	insertErr := datastoreWrapper.InsertNewRootEntity(appEngContext, recordFilterRuleEntityKind, &newFilterRule)
	if insertErr != nil {
		return nil, fmt.Errorf("Can't create new text box component: error inserting into datastore: %v", insertErr)
	}

	return &newFilterRule, nil

}
