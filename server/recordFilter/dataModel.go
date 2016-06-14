package recordFilter

import (
	"appengine"
	"fmt"
	"github.com/gocql/gocql"
	"log"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/cassandraWrapper"
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
		FilterID:      gocql.TimeUUID().String(),
		Name:          sanitizedName}

	dbSession, sessionErr := cassandraWrapper.CreateSession()
	if sessionErr != nil {
		return nil, fmt.Errorf("NewRecord: Can't create record: unable to create record: error = %v", sessionErr)
	}
	defer dbSession.Close()

	if insertErr := dbSession.Query(`INSERT INTO filters (tableID, filter_id, name) VALUES (?,?,?)`,
		newFilter.ParentTableID, newFilter.FilterID, newFilter.Name).Exec(); insertErr != nil {
		return nil, fmt.Errorf("newFilter: Can't create filter: error = %v", insertErr)
	}

	log.Printf("newFilter: Created new filter: %+v", newFilter)

	return &newFilter, nil
}

func getFilter(appEngContext appengine.Context, parentTableID string, filterID string) (*RecordFilter, error) {

	dbSession, sessionErr := cassandraWrapper.CreateSession()
	if sessionErr != nil {
		return nil, fmt.Errorf("GetRecord: Can't create database: unable to create database session: error = %v", sessionErr)
	}
	defer dbSession.Close()

	var filterGetDest RecordFilter
	getErr := dbSession.Query(`SELECT tableID,filter_id,name FROM filters WHERE tableID=? AND filter_id=? LIMIT 1`,
		parentTableID, filterID).Scan(&filterGetDest.ParentTableID,
		&filterGetDest.FilterID,
		&filterGetDest.Name)
	if getErr != nil {
		return nil, fmt.Errorf("getFilter: Unabled to get filter: id = %v: datastore err=%v", filterID, getErr)
	}

	return &filterGetDest, nil
}

func updateExistingFilter(appEngContext appengine.Context, updatedFilter *RecordFilter) (*RecordFilter, error) {

	dbSession, sessionErr := cassandraWrapper.CreateSession()
	if sessionErr != nil {
		return nil, fmt.Errorf("updateExistingFilter: Can't create database: unable to create database session: error = %v", sessionErr)
	}
	defer dbSession.Close()

	updateErr := dbSession.Query(`UPDATE filters SET name=? WHERE tableID=? AND filter_id=?`,
		updatedFilter.Name, updatedFilter.ParentTableID, updatedFilter.FilterID).Exec()
	if updateErr != nil {
		return nil, fmt.Errorf("updateExistingFilter: Unabled to update filter: id = %v: datastore err=%v", updatedFilter.FilterID, updateErr)
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

	dbSession, sessionErr := cassandraWrapper.CreateSession()
	if sessionErr != nil {
		return nil, fmt.Errorf("getTableList: Unable to create database session: error = %v", sessionErr)
	}
	defer dbSession.Close()

	ruleIter := dbSession.Query(`SELECT filter_id, rule_id, field_id,rule_def_id,text_param,number_param 
					FROM filter_rules WHERE filter_id=?`,
		params.ParentFilterID).Iter()

	var currRule RecordFilterRule
	allFilterRules := []RecordFilterRule{}
	for ruleIter.Scan(&currRule.ParentFilterID,
		&currRule.FilterRuleID,
		&currRule.FieldID,
		&currRule.RuleID,
		&currRule.TextRuleParam,
		&currRule.NumberRuleParam) {
		allFilterRules = append(allFilterRules, currRule)
	}
	if closeErr := ruleIter.Close(); closeErr != nil {
		return nil, fmt.Errorf("getRecordFilterRules: Failure querying database: %v", closeErr)
	}

	return allFilterRules, nil
}

func getFilterList(appEngContext appengine.Context, parentTableID string) ([]RecordFilter, error) {

	dbSession, sessionErr := cassandraWrapper.CreateSession()
	if sessionErr != nil {
		return nil, fmt.Errorf("getTableList: Unable to create database session: error = %v", sessionErr)
	}
	defer dbSession.Close()

	filterIter := dbSession.Query(`SELECT tableID,filter_id,name FROM filters WHERE tableID=?`,
		parentTableID).Iter()

	var currFilter RecordFilter
	allFilters := []RecordFilter{}
	for filterIter.Scan(&currFilter.ParentTableID,
		&currFilter.FilterID,
		&currFilter.Name) {
		allFilters = append(allFilters, currFilter)
	}
	if closeErr := filterIter.Close(); closeErr != nil {
		return nil, fmt.Errorf("getFilterList: Failure querying database: %v", closeErr)
	}

	return allFilters, nil

}

type NewFilterRuleParams struct {
	ParentFilterID     string   `jsaon:"parentFilterID"`
	FieldParentTableID string   `json:'fieldParentTableID'`
	FieldID            string   `json:"fieldID"`
	RuleID             string   `json:"ruleID"`
	TextRuleParam      *string  `json:"textRuleParam,omitempty"`
	NumberRuleParam    *float64 `json:"numberRuleParam,omitempty"`
}

func newFilterRule(appEngContext appengine.Context, newRuleParams NewFilterRuleParams) (*RecordFilterRule, error) {

	filterOnField, fieldErr := field.GetField(appEngContext, newRuleParams.FieldParentTableID, newRuleParams.FieldID)
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
					FilterRuleID:   gocql.TimeUUID().String(),
					FieldID:        newRuleParams.FieldID,
					RuleID:         newRuleParams.RuleID,
					TextRuleParam:  *newRuleParams.TextRuleParam}
			} else {
				newFilterRule = RecordFilterRule{
					ParentFilterID: newRuleParams.ParentFilterID,
					FilterRuleID:   gocql.TimeUUID().String(),
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
					FilterRuleID:    gocql.TimeUUID().String(),
					FieldID:         newRuleParams.FieldID,
					RuleID:          newRuleParams.RuleID,
					NumberRuleParam: *newRuleParams.NumberRuleParam}
			} else {
				newFilterRule = RecordFilterRule{
					ParentFilterID:  newRuleParams.ParentFilterID,
					FilterRuleID:    gocql.TimeUUID().String(),
					FieldID:         newRuleParams.FieldID,
					RuleID:          newRuleParams.RuleID,
					NumberRuleParam: 0}
			}
		}

	default:
		return nil, fmt.Errorf("NewFilterRule: Filtering not supported on field type: %v", filterOnField.Type)
	}

	dbSession, sessionErr := cassandraWrapper.CreateSession()
	if sessionErr != nil {
		return nil, fmt.Errorf("NewRecord: Can't create record: unable to create record: error = %v", sessionErr)
	}
	defer dbSession.Close()

	if insertErr := dbSession.Query(`INSERT INTO filter_rules (filter_id, rule_id, field_id,rule_def_id,text_param,number_param) VALUES (?,?,?,?,?,?)`,
		newFilterRule.ParentFilterID,
		newFilterRule.FilterRuleID,
		newFilterRule.FieldID,
		newFilterRule.RuleID,
		newFilterRule.TextRuleParam,
		newFilterRule.NumberRuleParam).Exec(); insertErr != nil {
		return nil, fmt.Errorf("newFilter: Can't create filter rule: error = %v", insertErr)
	}

	return &newFilterRule, nil

}
