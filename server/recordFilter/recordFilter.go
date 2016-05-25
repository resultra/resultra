package recordFilter

import (
	"appengine"
	"fmt"
	"resultra/datasheet/server/record"
)

func filterOneRecord(recordRef record.RecordRef, filterRules []FilterRuleRef) (bool, error) {

	for _, currFilterRule := range filterRules {

		// Pass a limited subset of parameters to the functions for doing the actual filtering. This simplifies
		// the code inside actual filtering functions and makes these functions easier to test.
		filterParams :=
			FilterFuncParams{currFilterRule.FieldRef.FieldID,
				currFilterRule.TextRuleParam,
				currFilterRule.NumberRuleParam}

		recordIsFiltered, filterErr := currFilterRule.FilterRuleDef.filterFunc(filterParams, recordRef.FieldValues)

		if filterErr != nil {
			fmt.Errorf("filterOneRecord: Error filtering record: %v", filterErr)
		} else if !recordIsFiltered {

			// Filtering is performed based upon a logical AND of all filtering rules passing; i.e. (F1 & F2 & F3 ...).
			// To be more efficient, we can test for the first to not pass its filtering rule; i.e. !(F1 & F2 & F3 ...)
			// == !F1 || !F2 || !F3
			return false, nil
		}
	}
	return true, nil
}

type GetFilteredRecordsParams struct {
	TableID string `json:"tableID"`
}

func GetFilteredRecords(appEngContext appengine.Context, params GetFilteredRecordsParams) ([]record.RecordRef, error) {

	// TODO - The code below retrieve *all* the records. However, the datastore supports up to 1 filtering criterion
	// for each field, so <=1 of these criterion could be used to filter the records coming from the datastore and
	// before doing any kind of in-memory filtering.
	unfilteredRecordRefs, getRecordErr := record.GetRecords(appEngContext, record.GetRecordsParams{TableID: params.TableID})
	if getRecordErr != nil {
		return nil, fmt.Errorf("GetFilteredRecords: Error retrieving records: %v", getRecordErr)
	}

	/* TODO - Reintegrate with specific forms for forms or dashboard objects */
	/*
		filterRefs, getFilterErr := GetRecordFilterRefs(appEngContext)
		if getFilterErr != nil {
			return nil, fmt.Errorf("GetFilteredRecords: Unable to retrieve filtered records: datastore error =%v", getFilterErr)
		}*/
	filterRefs := []FilterRuleRef{}

	filteredRecordRefs := []record.RecordRef{}
	for _, currRecordRef := range unfilteredRecordRefs {
		recordIsFiltered, filterErr := filterOneRecord(currRecordRef, filterRefs)
		if filterErr != nil {
			return nil, fmt.Errorf("GetFilteredRecords: Error filtering records: %v", filterErr)
		}
		if recordIsFiltered {
			filteredRecordRefs = append(filteredRecordRefs, currRecordRef)
		}
	}

	return filteredRecordRefs, nil
}
