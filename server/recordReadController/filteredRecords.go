package recordReadController

import (
	"fmt"
	"resultra/datasheet/server/recordSort"
	"resultra/datasheet/server/recordValue"
)

type GetFilteredSortedRecordsParams struct {
	TableID   string                      `json:"tableID"`
	FilterIDs []string                    `json:filterIDs`
	SortRules []recordSort.RecordSortRule `json:"sortRules"`
}

func GetFilteredSortedRecords(params GetFilteredSortedRecordsParams) ([]recordValue.RecordValueResults, error) {

	// The code below retrieve *all* the mapped record values. A more scalable approach would be to filter the
	// record values in batches, then combine and sort them.
	unfilteredRecordValues, getRecordErr := recordValue.GetAllRecordValueResults(params.TableID)
	if getRecordErr != nil {
		return nil, fmt.Errorf("GetFilteredRecords: Error retrieving records: %v", getRecordErr)
	}

	filteredRecords := []recordValue.RecordValueResults{}
	for _, recValue := range unfilteredRecordValues {
		if recValue.FilterMatches.MatchAllFilterIDs(params.FilterIDs) {
			filteredRecords = append(filteredRecords, recValue)
		}
	}

	if sortErr := recordSort.SortRecordValues(params.TableID, filteredRecords, params.SortRules); sortErr != nil {
		return nil, fmt.Errorf("GetFilteredRecords: Error sorting records: %v", sortErr)
	}

	return filteredRecords, nil
}
