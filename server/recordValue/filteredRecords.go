package recordValue

import (
	"fmt"
)

type GetFilteredRecordsParams struct {
	TableID   string   `json:"tableID"`
	FilterIDs []string `json:filterIDs`
}

func GetFilteredRecords(params GetFilteredRecordsParams) ([]RecordValueResults, error) {

	// TODO - The code below retrieve *all* the records. However, the datastore supports up to 1 filtering criterion
	// for each field, so <=1 of these criterion could be used to filter the records coming from the datastore and
	// before doing any kind of in-memory filtering.
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
