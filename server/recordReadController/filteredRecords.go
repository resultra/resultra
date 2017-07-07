package recordReadController

import (
	"fmt"
	"resultra/datasheet/server/common/recordSortDataModel"
	"resultra/datasheet/server/record"
	"resultra/datasheet/server/recordFilter"
	"resultra/datasheet/server/recordSort"
	"resultra/datasheet/server/recordValue"
	"resultra/datasheet/server/recordValueMappingController"
)

type GetFilteredSortedRecordsParams struct {
	DatabaseID     string                               `json:"databaseID"`
	PreFilterRules recordFilter.RecordFilterRuleSet     `json:"preFilterRules"`
	FilterRules    recordFilter.RecordFilterRuleSet     `json:"filterRules"`
	SortRules      []recordSortDataModel.RecordSortRule `json:"sortRules"`
}

func NewDefaultGetFilteredSortedRecordsParams() GetFilteredSortedRecordsParams {
	return GetFilteredSortedRecordsParams{
		"", recordFilter.NewDefaultRecordFilterRuleSet(), recordFilter.NewDefaultRecordFilterRuleSet(), []recordSortDataModel.RecordSortRule{}}

}

func getCachedOrRemappedRecordValues(databaseID string) ([]recordValue.RecordValueResults, error) {

	cachedValues, valuesFound := recordValue.ResultsCache.Get(databaseID)
	if valuesFound {
		var validType bool
		recordValues, validType := cachedValues.([]recordValue.RecordValueResults)
		if validType {
			return recordValues, nil
		} else {
			return nil, fmt.Errorf("getCachedOrRemappedRecordValues: unexpected type from results cache")
		}
	} else {
		recordValues, mapErr := recordValueMappingController.MapAllRecordUpdatesToFieldValues(databaseID)
		if mapErr != nil {
			return nil, fmt.Errorf("GetFilteredRecords: Error updating records: %v", mapErr)
		}
		recordValue.ResultsCache.Add(databaseID, recordValues)
		return recordValues, nil
	}

}

func GetFilteredSortedRecords(params GetFilteredSortedRecordsParams) ([]recordValue.RecordValueResults, error) {

	// The code below retrieve *all* the mapped record values. A more scalable approach would be to filter the
	// record values in batches, then combine and sort them.
	unfilteredRecordValues, getRecordsErr := getCachedOrRemappedRecordValues(params.DatabaseID)
	if getRecordsErr != nil {
		return nil, fmt.Errorf("GetFilteredRecords: Error updating records: %v", getRecordsErr)
	}

	preFilteredRecords, preFilterErr := recordFilter.FilterRecordValues(params.PreFilterRules, unfilteredRecordValues)
	if preFilterErr != nil {
		return nil, fmt.Errorf("GetFilteredRecords: Error pre-filtering records: %v", preFilterErr)
	}

	filteredRecords, err := recordFilter.FilterRecordValues(params.FilterRules, preFilteredRecords)
	if err != nil {
		return nil, fmt.Errorf("GetFilteredRecords: Error filtering records: %v", err)
	}

	if sortErr := recordSort.SortRecordValues(params.DatabaseID, filteredRecords, params.SortRules); sortErr != nil {
		return nil, fmt.Errorf("GetFilteredSortedRecords: Error sorting records: %v", sortErr)
	}

	return filteredRecords, nil
}

type GetFilteredRecordCountParams struct {
	DatabaseID     string                           `json:"databaseID"`
	PreFilterRules recordFilter.RecordFilterRuleSet `json:"preFilterRules"`
}

func getFilteredRecordCount(params GetFilteredRecordCountParams) (int, error) {
	unfilteredRecordValues, getRecordsErr := getCachedOrRemappedRecordValues(params.DatabaseID)
	if getRecordsErr != nil {
		return -1, fmt.Errorf("GetFilteredRecords: Error updating records: %v", getRecordsErr)
	}

	preFilteredRecords, preFilterErr := recordFilter.FilterRecordValues(params.PreFilterRules, unfilteredRecordValues)
	if preFilterErr != nil {
		return -1, fmt.Errorf("GetFilteredRecords: Error pre-filtering records: %v", preFilterErr)
	}

	return len(preFilteredRecords), nil

}

type GetRecordValResultParams struct {
	ParentDatabaseID string `json:"parentDatabaseID"`
	RecordID         string `json:"recordID"`
}

func getRecordValueResults(params GetRecordValResultParams) (*recordValue.RecordValueResults, error) {

	recCellUpdates, cellUpdatesErr := record.GetRecordCellUpdates(params.RecordID, record.FullyCommittedCellUpdatesChangeSetID)
	if cellUpdatesErr != nil {
		return nil, fmt.Errorf("updateRecordValue: Can't get cell updates: err = %v", cellUpdatesErr)
	}

	updateRecordValResult, mapErr := recordValueMappingController.MapOneRecordUpdatesToFieldValues(
		params.ParentDatabaseID, recCellUpdates, record.FullyCommittedCellUpdatesChangeSetID)
	if mapErr != nil {
		return nil, fmt.Errorf(
			"updateRecordValue: Error mapping field values: err = %v", mapErr)
	}

	return updateRecordValResult, nil

}
