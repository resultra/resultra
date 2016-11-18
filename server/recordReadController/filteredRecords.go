package recordReadController

import (
	"fmt"
	"resultra/datasheet/server/common/recordSortDataModel"
	"resultra/datasheet/server/recordFilter"
	"resultra/datasheet/server/recordSort"
	"resultra/datasheet/server/recordValue"
)

type GetFilteredSortedRecordsParams struct {
	TableID     string                               `json:"tableID"`
	FilterRules []recordFilter.RecordFilterRule      `json:"filterRules"`
	SortRules   []recordSortDataModel.RecordSortRule `json:"sortRules"`
}

func GetFilteredSortedRecords(params GetFilteredSortedRecordsParams) ([]recordValue.RecordValueResults, error) {

	// The code below retrieve *all* the mapped record values. A more scalable approach would be to filter the
	// record values in batches, then combine and sort them.
	unfilteredRecordValues, getRecordErr := recordValue.GetAllRecordValueResults(params.TableID)
	if getRecordErr != nil {
		return nil, fmt.Errorf("GetFilteredRecords: Error retrieving records: %v", getRecordErr)
	}

	filteredRecords, err := recordFilter.FilterRecordValues(params.FilterRules, unfilteredRecordValues)
	if err != nil {
		return nil, fmt.Errorf("GetFilteredRecords: Error filtering records: %v", err)
	}

	if sortErr := recordSort.SortRecordValues(params.TableID, filteredRecords, params.SortRules); sortErr != nil {
		return nil, fmt.Errorf("GetFilteredSortedRecords: Error sorting records: %v", sortErr)
	}

	return filteredRecords, nil
}
