package recordFilter

import (
	"fmt"
	"log"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/databaseWrapper"
)

type RecordFilter struct {
	ParentTableID string `json:"parentTableID"`
	FilterID      string `json:"filterID"`
	Name          string `json:"name"`
	/* Rules are managed as child entities */
}

type RecordFilterRule struct {
	ParentFilterID  string  `json:"parentFilterID"`
	FilterRuleID    string  `json:"filterRuleID"`
	FieldID         string  `json:"fieldID"`
	RuleID          string  `json:"ruleID"`
	TextRuleParam   string  `json:"textRuleParam"`
	NumberRuleParam float64 `json:"numberRuleParam"`
}

type NewFilterParams struct {
	ParentTableID string `json:"parentTableID"`
	Name          string `json:"name"`
}

func getExistingFilterNames(parentTableID string) ([]string, error) {

	var allFilters, getFiltersErr = getFilterList(parentTableID)
	if getFiltersErr != nil {
		return nil, fmt.Errorf("Unable to retrieve filters: %v", getFiltersErr)
	}

	existingNames := []string{}
	for _, currFilter := range allFilters {
		existingNames = append(existingNames, currFilter.Name)
	}
	return existingNames, nil

}

func validateUnusedFilterName(parentTableID string, filterName string) error {

	existingNames, getNamesErr := getExistingFilterNames(parentTableID)
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

func genUnusedFilterName(parentTableID string, namePrefix string) (string, error) {

	existingNames, getNamesErr := getExistingFilterNames(parentTableID)
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

func newFilter(params NewFilterParams) (*RecordFilter, error) {

	sanitizedName, sanitizeErr := generic.SanitizeName(params.Name)
	if sanitizeErr != nil {
		return nil, fmt.Errorf("newFilter (sanitize name): %v", sanitizeErr)
	}

	if validateNameErr := validateUnusedFilterName(params.ParentTableID, sanitizedName); validateNameErr != nil {
		return nil, fmt.Errorf("newFilter: %v", validateNameErr)
	}

	newFilter := RecordFilter{
		ParentTableID: params.ParentTableID,
		FilterID:      databaseWrapper.GlobalUniqueID(),
		Name:          sanitizedName}

	if _, insertErr := databaseWrapper.DBHandle().Exec(`INSERT INTO filters (table_id, filter_id, name) VALUES ($1,$2,$3)`,
		newFilter.ParentTableID, newFilter.FilterID, newFilter.Name); insertErr != nil {
		return nil, fmt.Errorf("newFilter: Can't create filter: error = %v", insertErr)
	}

	log.Printf("newFilter: Created new filter: %+v", newFilter)

	return &newFilter, nil
}

func getFilter(parentTableID string, filterID string) (*RecordFilter, error) {

	var filterGetDest RecordFilter
	getErr := databaseWrapper.DBHandle().QueryRow(`SELECT table_id,filter_id,name FROM filters WHERE table_id=$1 AND filter_id=$2 LIMIT 1`,
		parentTableID, filterID).Scan(&filterGetDest.ParentTableID,
		&filterGetDest.FilterID,
		&filterGetDest.Name)
	if getErr != nil {
		return nil, fmt.Errorf("getFilter: Unabled to get filter: id = %v: datastore err=%v", filterID, getErr)
	}

	return &filterGetDest, nil
}

func updateExistingFilter(updatedFilter *RecordFilter) (*RecordFilter, error) {

	_, updateErr := databaseWrapper.DBHandle().Exec(`UPDATE filters SET name=$1 WHERE table_id=$2 AND filter_id=$3`,
		updatedFilter.Name, updatedFilter.ParentTableID, updatedFilter.FilterID)
	if updateErr != nil {
		return nil, fmt.Errorf("updateExistingFilter: Unabled to update filter: id = %v: datastore err=%v", updatedFilter.FilterID, updateErr)
	}

	return updatedFilter, nil

}

type NewFilterWithPrefixParams struct {
	ParentTableID string `json:"parentTableID"`
	NamePrefix    string `json:"namePrefix"`
}

func newFilterWithPrefix(params NewFilterWithPrefixParams) (*RecordFilter, error) {

	sanitizedPrefix, sanitizeErr := generic.SanitizeName(params.NamePrefix)
	if sanitizeErr != nil {
		return nil, fmt.Errorf("newFilterWithPrefix: %v", sanitizeErr)
	}
	newFilterName, nameGenErr := genUnusedFilterName(params.ParentTableID, sanitizedPrefix)
	if nameGenErr != nil {
		return nil, fmt.Errorf("newFilterWithPrefix: %v", nameGenErr)
	}
	log.Printf("newFilterWithPrefix: New filter name: %v", newFilterName)
	return newFilter(NewFilterParams{ParentTableID: params.ParentTableID, Name: newFilterName})

}

type GetFilterRulesParams struct {
	ParentFilterID string `json:"parentFilterID"`
}

func getRecordFilterRules(params GetFilterRulesParams) ([]RecordFilterRule, error) {

	rows, queryErr := databaseWrapper.DBHandle().Query(`SELECT filter_id, rule_id, field_id,rule_def_id,text_param,number_param 
					FROM filter_rules WHERE filter_id=$1`, params.ParentFilterID)
	if queryErr != nil {
		return nil, fmt.Errorf("getRecordFilterRules: Failure querying database: %v", queryErr)
	}
	allFilterRules := []RecordFilterRule{}
	for rows.Next() {
		var currRule RecordFilterRule
		if scanErr := rows.Scan(&currRule.ParentFilterID,
			&currRule.FilterRuleID,
			&currRule.FieldID,
			&currRule.RuleID,
			&currRule.TextRuleParam,
			&currRule.NumberRuleParam); scanErr != nil {
			return nil, fmt.Errorf("getRecordFilterRules: Failure querying database: %v", scanErr)
		}
		allFilterRules = append(allFilterRules, currRule)
	}

	return allFilterRules, nil
}

func getFilterList(parentTableID string) ([]RecordFilter, error) {

	rows, queryErr := databaseWrapper.DBHandle().Query(`SELECT table_id,filter_id,name FROM filters WHERE table_id=$1`, parentTableID)
	if queryErr != nil {
		return nil, fmt.Errorf("getFilterList: Failure querying database: %v", queryErr)
	}

	allFilters := []RecordFilter{}
	for rows.Next() {
		var currFilter RecordFilter
		if scanErr := rows.Scan(&currFilter.ParentTableID,
			&currFilter.FilterID,
			&currFilter.Name); scanErr != nil {
			return nil, fmt.Errorf("getFilterList: Failure querying database: %v", scanErr)
		}
		allFilters = append(allFilters, currFilter)
	}

	return allFilters, nil

}

type NewFilterRuleParams struct {
	ParentFilterID     string   `json:"parentFilterID"`
	FieldParentTableID string   `json:'fieldParentTableID'`
	FieldID            string   `json:"fieldID"`
	RuleID             string   `json:"ruleID"`
	TextRuleParam      *string  `json:"textRuleParam,omitempty"`
	NumberRuleParam    *float64 `json:"numberRuleParam,omitempty"`
}

func newFilterRule(newRuleParams NewFilterRuleParams) (*RecordFilterRule, error) {

	filterOnField, fieldErr := field.GetField(newRuleParams.FieldParentTableID, newRuleParams.FieldID)
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
					FilterRuleID:   databaseWrapper.GlobalUniqueID(),
					FieldID:        newRuleParams.FieldID,
					RuleID:         newRuleParams.RuleID,
					TextRuleParam:  *newRuleParams.TextRuleParam}
			} else {
				newFilterRule = RecordFilterRule{
					ParentFilterID: newRuleParams.ParentFilterID,
					FilterRuleID:   databaseWrapper.GlobalUniqueID(),
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
					FilterRuleID:    databaseWrapper.GlobalUniqueID(),
					FieldID:         newRuleParams.FieldID,
					RuleID:          newRuleParams.RuleID,
					NumberRuleParam: *newRuleParams.NumberRuleParam}
			} else {
				newFilterRule = RecordFilterRule{
					ParentFilterID:  newRuleParams.ParentFilterID,
					FilterRuleID:    databaseWrapper.GlobalUniqueID(),
					FieldID:         newRuleParams.FieldID,
					RuleID:          newRuleParams.RuleID,
					NumberRuleParam: 0}
			}
		}

	default:
		return nil, fmt.Errorf("NewFilterRule: Filtering not supported on field type: %v", filterOnField.Type)
	}

	if _, insertErr := databaseWrapper.DBHandle().Exec(`INSERT INTO filter_rules 
				(filter_id, rule_id, field_id,rule_def_id,text_param,number_param) 
				VALUES ($1,$2,$3,$4,$5,$6)`,
		newFilterRule.ParentFilterID,
		newFilterRule.FilterRuleID,
		newFilterRule.FieldID,
		newFilterRule.RuleID,
		newFilterRule.TextRuleParam,
		newFilterRule.NumberRuleParam); insertErr != nil {
		return nil, fmt.Errorf("saveNewTable: insert failed: error = %v", insertErr)
	}

	return &newFilterRule, nil

}
