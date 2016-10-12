package recordSort

import (
	"database/sql"
	"fmt"
	"log"
	"resultra/datasheet/server/common/recordSortDataModel"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/databaseWrapper"
)

type FormSortRules struct {
	ParentFormID string                               `json:"parentFormID"`
	SortRules    []recordSortDataModel.RecordSortRule `json:"sortRules"`
}

type NewSortRuleParams FormSortRules

func saveFormSortRules(params NewSortRuleParams) (*FormSortRules, error) {

	if len(params.ParentFormID) == 0 {
		return nil, fmt.Errorf("saveSortRule: Missing parent form ID")
	}

	for _, sortRule := range params.SortRules {
		if !recordSortDataModel.ValidSortDirection(sortRule.SortDirection) {
			return nil, fmt.Errorf("saveSortRule: Invalid sort direction '%v'", sortRule.SortDirection)
		}
		if len(sortRule.SortFieldID) == 0 {
			return nil, fmt.Errorf("saveSortRule: Missing field ID")
		}
	}

	encodedSortRules, encodeErr := generic.EncodeJSONString(params.SortRules)
	if encodeErr != nil {
		return nil, fmt.Errorf("saveSortRule: Unable to encode sort rules %+v: error = %v", params.SortRules, encodeErr)
	}

	// Delete any previous sort rules
	if _, deleteErr := databaseWrapper.DBHandle().Exec(`DELETE FROM form_sort_rules WHERE form_id=$1`,
		params.ParentFormID); deleteErr != nil {
		return nil, fmt.Errorf("saveSortRule: Can't save sort rule for form %+v: error = %v", params, deleteErr)
	}

	if _, insertErr := databaseWrapper.DBHandle().Exec(`INSERT INTO form_sort_rules 
			(form_id, sort_rules) VALUES ($1,$2)`,
		params.ParentFormID, encodedSortRules); insertErr != nil {
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

	decodedSortRules := []recordSortDataModel.RecordSortRule{}
	encodedSortRules := ""
	getErr := databaseWrapper.DBHandle().QueryRow(`SELECT sort_rules FROM form_sort_rules
			WHERE form_id=$1 LIMIT 1`, parentFormID).Scan(&encodedSortRules)

	if getErr != nil {
		if getErr == sql.ErrNoRows {
			return &FormSortRules{
				ParentFormID: parentFormID,
				SortRules:    decodedSortRules}, nil
		} else {
			return nil, fmt.Errorf("Unabled to get form sort rules for form: datastore err=%v", getErr)
		}
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
