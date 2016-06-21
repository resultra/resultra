package recordSort

import (
	"fmt"
	"log"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/cassandraWrapper"
)

const sortDirectionAsc string = "asc"
const sortDirectionDesc string = "desc"

type RecordSortRule struct {
	SortFieldID   string `json:"fieldID"`
	SortDirection string `json:"direction"`
}

type FormSortRules struct {
	ParentFormID string           `json:"parentFormID"`
	SortRules    []RecordSortRule `json:sortRules`
}

type NewSortRuleParams FormSortRules

func validSortDirection(sortDir string) bool {
	if (sortDir == sortDirectionAsc) || (sortDir == sortDirectionDesc) {
		return true
	} else {
		return false
	}
}

func saveFormSortRules(params NewSortRuleParams) (*FormSortRules, error) {

	if len(params.ParentFormID) == 0 {
		return nil, fmt.Errorf("saveSortRule: Missing parent form ID")
	}

	for _, sortRule := range params.SortRules {
		if !validSortDirection(sortRule.SortDirection) {
			return nil, fmt.Errorf("saveSortRule: Invalid sort direction '%v'", sortRule.SortDirection)
		}
		if len(sortRule.SortFieldID) == 0 {
			return nil, fmt.Errorf("saveSortRule: Missing field ID")
		}
	}

	dbSession, sessionErr := cassandraWrapper.CreateSession()
	if sessionErr != nil {
		return nil, fmt.Errorf("saveSortRule: Can't create database session: error = %v", sessionErr)
	}
	defer dbSession.Close()

	encodedSortRules, encodeErr := generic.EncodeJSONString(params.SortRules)
	if encodeErr != nil {
		return nil, fmt.Errorf("saveSortRule: Unable to encode sort rules %+v: error = %v", params.SortRules, encodeErr)
	}

	if insertErr := dbSession.Query(`INSERT INTO form_sort_rules 
			(form_id, sort_rules) VALUES (?,?)`,
		params.ParentFormID, encodedSortRules).Exec(); insertErr != nil {
		return nil, fmt.Errorf("saveSortRule: Can't save sort rule for form %+v: error = %v", params, insertErr)
	}

	sortRules := FormSortRules{
		ParentFormID: params.ParentFormID,
		SortRules:    params.SortRules}

	log.Printf("saveSortRule: Saved form sort rules: %+v", sortRules)

	return &sortRules, nil
}

func GetFormSortRules(parentFormID string) (*FormSortRules, error) {

	if len(parentFormID) == 0 {
		return nil, fmt.Errorf("GetFormSortRules: Missing parent form ID")
	}

	dbSession, sessionErr := cassandraWrapper.CreateSession()
	if sessionErr != nil {
		return nil, fmt.Errorf("saveSortRule: Can't create database session: error = %v", sessionErr)
	}
	defer dbSession.Close()

	decodedSortRules := []RecordSortRule{}

	encodedSortRules := ""
	getErr := dbSession.Query(`SELECT sort_rules FROM form_sort_rules
			WHERE form_id=? LIMIT 1`, parentFormID).Scan(&encodedSortRules)
	if getErr != nil {
		return nil, fmt.Errorf("Unabled to get form sort rules for form: datastore err=%v", getErr)
	}

	decodeErr := generic.DecodeJSONString(encodedSortRules, &decodedSortRules)
	if decodeErr != nil {
		return nil, fmt.Errorf("GetFormSortRules: Unable to decode sort rules %v: error = %v",
			encodedSortRules, decodeErr)
	}

	sortRules := FormSortRules{
		ParentFormID: parentFormID,
		SortRules:    decodedSortRules}

	return &sortRules, nil

}
