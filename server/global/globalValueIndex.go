package global

import (
	"fmt"
	"sort"
	"time"
)

type GlobalValueUpdate struct {
	UpdateTimeStamp time.Time
	Value           interface{} // Decoded value, type depends on field.
}

type GlobalValueUpdateSeries []GlobalValueUpdate

// Custom sort function for the FieldValueUpate
type ByUpdateTime []GlobalValueUpdate

func (s ByUpdateTime) Len() int {
	return len(s)
}
func (s ByUpdateTime) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Sort in reverse chronological order; i.e. the most recent dates come first.
func (s ByUpdateTime) Less(i, j int) bool {
	return s[i].UpdateTimeStamp.After(s[j].UpdateTimeStamp)
}

type GlobalValueUpdateSeriesIndex map[string]GlobalValueUpdateSeries

type GlobalValues map[string]interface{}

// There will only be cell updates in the datastore for non-calculated fields. For these fields,
// this function returns the latest (most recent) values.
func (globalValSeriesIndex GlobalValueUpdateSeriesIndex) LatestValues() *GlobalValues {

	globalValues := GlobalValues{}

	for globalID, updateSeries := range globalValSeriesIndex {
		if len(updateSeries) > 0 {
			globalValues[globalID] = updateSeries[0].Value
		}
	}

	return &globalValues
}

func NewGlobalValueIndex(parentDatabaseID string) (*GlobalValueUpdateSeriesIndex, error) {

	valUpdates, getErr := getValUpdates(parentDatabaseID)
	if getErr != nil {
		return nil, fmt.Errorf("NewGlobalValueIndex: failure retrieving value updates  for database = %v error = %v",
			parentDatabaseID, getErr)
	}

	globals, getGlobalsErr := GetGlobals(parentDatabaseID)
	if getGlobalsErr != nil {
		return nil, fmt.Errorf("NewGlobalValueIndex: failure retrieving globals for database = %v error = %v",
			parentDatabaseID, getGlobalsErr)
	}
	globalIndex := map[string]Global{}
	for _, currGlobal := range globals {
		globalIndex[currGlobal.GlobalID] = currGlobal
	}

	// Populate the index with all the updates for the given recordID, broken down by FieldID.
	globalValSeriesMap := GlobalValueUpdateSeriesIndex{}
	for _, currUpdate := range valUpdates {

		globalInfo, foundInfo := globalIndex[currUpdate.GlobalID]
		if !foundInfo {
			return nil, fmt.Errorf("NewGlobalValueIndex: failure retrieving global information for value update = %+v", currUpdate)
		}

		decodedVal, decodeErr := decodeGlobalValue(globalInfo.Type, currUpdate.Value)
		if decodeErr != nil {
			return nil, fmt.Errorf(
				"NewGlobalValueIndex: Unable to decode value for global = %+v: error=%v ",
				globalInfo, decodeErr)
		}

		valUpdate := GlobalValueUpdate{
			UpdateTimeStamp: currUpdate.UpdateTimestamp,
			Value:           decodedVal}

		globalValSeriesMap[currUpdate.GlobalID] = append(globalValSeriesMap[currUpdate.GlobalID], valUpdate)
	}

	// Sort the value updates for each global ID in reverse chronological order.
	for currGlobalID := range globalValSeriesMap {
		sort.Sort(ByUpdateTime(globalValSeriesMap[currGlobalID]))
	}

	return &globalValSeriesMap, nil
}

type GetGlobalValuesParams struct {
	ParentDatabaseID string `json:"parentDatabaseID"`
}

func GetGlobalValues(params GetGlobalValuesParams) (*GlobalValues, error) {
	globalValIndex, err := NewGlobalValueIndex(params.ParentDatabaseID)
	if err != nil {
		return nil, fmt.Errorf("getGlobalValues: failure retrieving global value index: %v", err)
	}

	latestVals := globalValIndex.LatestValues()

	return latestVals, nil

}
