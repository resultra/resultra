package recordValue

import (
	"fmt"
	"resultra/datasheet/server/recordSort"
)

type GetFilteredSortedRecordsParams struct {
	TableID   string                      `json:"tableID"`
	FilterIDs []string                    `json:filterIDs`
	SortRules []recordSort.RecordSortRule `json:"sortRules"`
}

func GetFilteredSortedRecords(params GetFilteredSortedRecordsParams) ([]RecordValueResults, error) {

	// The code below retrieve *all* the mapped record values. A more scalable approach would be to filter the
	// record values in batches, then combine and sort them.
	unfilteredRecordValues, getRecordErr := GetAllRecordValueResults(params.TableID)
	if getRecordErr != nil {
		return nil, fmt.Errorf("GetFilteredRecords: Error retrieving records: %v", getRecordErr)
	}

	filteredRecords := []RecordValueResults{}
	for _, recValue := range unfilteredRecordValues {
		if recValue.FilterMatches.MatchAllFilterIDs(params.FilterIDs) {
			filteredRecords = append(filteredRecords, recValue)
		}
	}

	return filteredRecords, nil
}
