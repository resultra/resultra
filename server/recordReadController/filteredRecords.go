package recordReadController

import (
	"database/sql"
	"fmt"
	"resultra/tracker/server/common/recordSortDataModel"
	"resultra/tracker/server/generic/timestamp"
	"resultra/tracker/server/record"
	"resultra/tracker/server/recordFilter"
	"resultra/tracker/server/recordSort"
	"resultra/tracker/server/recordValue"
	"resultra/tracker/server/recordValueMappingController"
)

type GetFilteredSortedRecordsParams struct {
	DatabaseID     string                               `json:"databaseID"`
	PreFilterRules recordFilter.RecordFilterRuleSet     `json:"preFilterRules"`
	FilterRules    recordFilter.RecordFilterRuleSet     `json:"filterRules"`
	SortRules      []recordSortDataModel.RecordSortRule `json:"sortRules"`
}

func NewDefaultGetFilteredSortedRecordsParams() GetFilteredSortedRecordsParams {
	return GetFilteredSortedRecordsParams{
		"", recordFilter.NewDefaultRecordFilterRuleSet(),
		recordFilter.NewDefaultRecordFilterRuleSet(),
		[]recordSortDataModel.RecordSortRule{}}

}

func getCachedOrRemappedRecordValues(trackerDBHandle *sql.DB,
	currUserID string, databaseID string) ([]recordValue.RecordValueResults, error) {

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

		calcFieldAsOfTime := timestamp.CurrentTimestampUTC()

		recordValues, mapErr := recordValueMappingController.MapAllRecordUpdatesToFieldValues(
			trackerDBHandle, currUserID, databaseID, calcFieldAsOfTime)
		if mapErr != nil {
			return nil, fmt.Errorf("GetFilteredRecords: Error updating records: %v", mapErr)
		}
		recordValue.ResultsCache.Add(databaseID, recordValues)
		return recordValues, nil
	}

}

func GetFilteredSortedRecords(trackerDBHandle *sql.DB,
	currUserID string, params GetFilteredSortedRecordsParams) ([]recordValue.RecordValueResults, error) {

	// The code below retrieve *all* the mapped record values. A more scalable approach would be to filter the
	// record values in batches, then combine and sort them.
	unfilteredRecordValues, getRecordsErr := getCachedOrRemappedRecordValues(trackerDBHandle, currUserID, params.DatabaseID)
	if getRecordsErr != nil {
		return nil, fmt.Errorf("GetFilteredRecords: Error updating records: %v", getRecordsErr)
	}

	preFilteredRecords, preFilterErr := recordFilter.FilterRecordValues(trackerDBHandle, currUserID, params.PreFilterRules, unfilteredRecordValues)
	if preFilterErr != nil {
		return nil, fmt.Errorf("GetFilteredRecords: Error pre-filtering records: %v", preFilterErr)
	}

	filteredRecords, err := recordFilter.FilterRecordValues(trackerDBHandle, currUserID, params.FilterRules, preFilteredRecords)
	if err != nil {
		return nil, fmt.Errorf("GetFilteredRecords: Error filtering records: %v", err)
	}

	if sortErr := recordSort.SortRecordValues(trackerDBHandle, params.DatabaseID, filteredRecords, params.SortRules); sortErr != nil {
		return nil, fmt.Errorf("GetFilteredSortedRecords: Error sorting records: %v", sortErr)
	}

	return filteredRecords, nil
}

type GetFilteredRecordCountParams struct {
	DatabaseID     string                           `json:"databaseID"`
	PreFilterRules recordFilter.RecordFilterRuleSet `json:"preFilterRules"`
}

func getFilteredRecordCount(trackerDBHandle *sql.DB, currUserID string, params GetFilteredRecordCountParams) (int, error) {

	unfilteredRecordValues, getRecordsErr := getCachedOrRemappedRecordValues(trackerDBHandle,
		currUserID, params.DatabaseID)
	if getRecordsErr != nil {
		return -1, fmt.Errorf("GetFilteredRecords: Error updating records: %v", getRecordsErr)
	}

	preFilteredRecords, preFilterErr := recordFilter.FilterRecordValues(trackerDBHandle, currUserID,
		params.PreFilterRules, unfilteredRecordValues)
	if preFilterErr != nil {
		return -1, fmt.Errorf("GetFilteredRecords: Error pre-filtering records: %v", preFilterErr)
	}

	return len(preFilteredRecords), nil

}

type GetRecordValResultParams struct {
	ParentDatabaseID string `json:"parentDatabaseID"`
	RecordID         string `json:"recordID"`
}

func getRecordValueResults(trackerDBHandle *sql.DB, currUserID string, params GetRecordValResultParams) (*recordValue.RecordValueResults, error) {

	recCellUpdates, cellUpdatesErr := record.GetRecordCellUpdates(trackerDBHandle, params.RecordID, record.FullyCommittedCellUpdatesChangeSetID)
	if cellUpdatesErr != nil {
		return nil, fmt.Errorf("updateRecordValue: Can't get cell updates: err = %v", cellUpdatesErr)
	}

	updateRecordValResult, mapErr := recordValueMappingController.MapOneRecordUpdatesToLatestFieldValues(
		trackerDBHandle, currUserID, params.ParentDatabaseID, recCellUpdates, record.FullyCommittedCellUpdatesChangeSetID)
	if mapErr != nil {
		return nil, fmt.Errorf(
			"updateRecordValue: Error mapping field values: err = %v", mapErr)
	}

	return updateRecordValResult, nil

}

type TestRecordIsFilteredParams struct {
	DatabaseID     string                           `json:"databaseID"`
	PreFilterRules recordFilter.RecordFilterRuleSet `json:"preFilterRules"`
	FilterRules    recordFilter.RecordFilterRuleSet `json:"filterRules"`
	RecordID       string                           `json:"recordID"`
}

func NewDefaultTestRecordIsFilteredParams() TestRecordIsFilteredParams {
	return TestRecordIsFilteredParams{
		"", recordFilter.NewDefaultRecordFilterRuleSet(),
		recordFilter.NewDefaultRecordFilterRuleSet(),
		""}

}

func testRecordIsFiltered(trackerDBHandle *sql.DB,
	currUserID string, params TestRecordIsFilteredParams) (bool, error) {

	recValResult, resultErr := recordValueMappingController.MapSingleRecordLatestValueResult(
		trackerDBHandle, currUserID, params.DatabaseID, params.RecordID)
	unfilteredRecordValues := []recordValue.RecordValueResults{}
	if resultErr != nil {
		return false, fmt.Errorf("testRecordIsFiltered: Error getting record value: %v", resultErr)
	}
	unfilteredRecordValues = append(unfilteredRecordValues, *recValResult)

	preFilteredRecords, preFilterErr := recordFilter.FilterRecordValues(trackerDBHandle,
		currUserID, params.PreFilterRules, unfilteredRecordValues)
	if preFilterErr != nil {
		return false, fmt.Errorf("testRecordIsFiltered: Error pre-filtering records: %v", preFilterErr)
	}

	filteredRecords, err := recordFilter.FilterRecordValues(trackerDBHandle, currUserID, params.FilterRules, preFilteredRecords)
	if err != nil {
		return false, fmt.Errorf("testRecordIsFiltered: Error filtering records: %v", err)
	}

	if len(filteredRecords) > 0 {
		return true, nil
	}

	return false, nil

}
