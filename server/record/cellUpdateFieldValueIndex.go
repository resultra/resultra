package record

import (
	"fmt"
	"resultra/tracker/server/field"
	"sort"
	"time"
)

type FieldValueUpdate struct {
	UpdateTimeStamp time.Time
	CellValue       interface{} // Decoded value, type depends on field.
}

type FieldValueUpdateSeries []FieldValueUpdate

func (series FieldValueUpdateSeries) valueAsOf(asOfTime time.Time) interface{} {

	var theVal interface{}
	theVal = nil

	if len(series) <= 0 {
		return theVal
	}

	// Most of the time this function is called, asOfTime will retrieve the latest value, rather than one
	// of the previous values. So, handle this as a special case to prevent unnecessary computation.
	if series[0].UpdateTimeStamp.Before(asOfTime) {
		theVal = series[0].CellValue
		return theVal
	}

	// The values are sorted in reverse chronological order. However,
	// to retrieve the value as of a given time, the values need to be
	// traversed in chronological order, finding the last value whose
	// timestamp is <= asOfTime.
	for i := len(series) - 1; i >= 0; i-- {
		valUpdateTime := series[i].UpdateTimeStamp
		if valUpdateTime.Before(asOfTime) {
			theVal = series[i].CellValue
		} else if valUpdateTime.Equal(asOfTime) {
			theVal = series[i].CellValue
		} else {
			// Update time is after 'asOfTime' => we've moved past asOfTime in the
			// series, so return whatever was the last value before this value.
			return theVal
		}
	}
	return theVal
}

// Custom sort function for the FieldValueUpate
type ByUpdateTime []FieldValueUpdate

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

type FieldUpdateFieldValueIndex map[string]FieldValueUpdateSeries

type CellUpdateFieldValueIndex struct {
	fieldUpdateFieldValueIndex FieldUpdateFieldValueIndex
	CellUpdates                []CellUpdate
}

type CellUpdateTimeIterFunc func(asOfTime time.Time) (bool, error)

// Iterate over the cell updates which occur before startMaxTime and invoke a call-back with the given
// cell update's time. The iteration occurs in reverse chronological order. This iteration is needed
// for calculated field formulas/equations with time-sensitive results.
func (cellUpdateFieldValIndex CellUpdateFieldValueIndex) IterateCellUpdateTimesInReverseChronologicalOrder(startMaxTime time.Time,
	iterCallback CellUpdateTimeIterFunc) error {

	// Visit each update in reverse chronological order. The cell updates are sorted in chronological order below.
	for i := len(cellUpdateFieldValIndex.CellUpdates) - 1; i >= 0; i-- {
		currCellUpdate := cellUpdateFieldValIndex.CellUpdates[i]
		if currCellUpdate.UpdateTimeStamp.Before(startMaxTime) || currCellUpdate.UpdateTimeStamp.Equal(startMaxTime) {
			continueIter, err := iterCallback(currCellUpdate.UpdateTimeStamp)
			if err != nil {
				return err
			} else if !continueIter {
				return nil
			}
		}
	}

	return nil
}

// There will only be cell updates in the datastore for non-calculated fields. For these fields,
// this function returns the latest (most recent) values.
func (cellUpdateFieldValIndex CellUpdateFieldValueIndex) LatestNonCalcFieldValues() *RecFieldValues {

	recFieldValues := RecFieldValues{}

	for fieldID, updateSeries := range cellUpdateFieldValIndex.fieldUpdateFieldValueIndex {
		if len(updateSeries) > 0 {
			recFieldValues[fieldID] = updateSeries[0].CellValue
		}
	}

	return &recFieldValues
}

func (cellUpdateFieldValIndex CellUpdateFieldValueIndex) NonCalcFieldValuesAsOf(asOfTime time.Time) RecFieldValues {

	recFieldValues := RecFieldValues{}

	for fieldID, updateSeries := range cellUpdateFieldValIndex.fieldUpdateFieldValueIndex {
		if len(updateSeries) > 0 {
			asOfVal := updateSeries.valueAsOf(asOfTime)
			if asOfVal != nil {
				recFieldValues[fieldID] = asOfVal
			}
		}
	}

	return recFieldValues
}

func NewUpdateFieldValueIndexForCellUpdates(recCellUpdates *RecordCellUpdates, fieldsByID map[string]field.Field) (*CellUpdateFieldValueIndex, error) {
	// Populate the index with all the updates for the given recordID, broken down by FieldID.
	fieldValSeriesMap := FieldUpdateFieldValueIndex{}
	cellUpdates := []CellUpdate{}
	for _, currUpdate := range recCellUpdates.CellUpdates {

		fieldInfo, foundField := fieldsByID[currUpdate.FieldID]
		if !foundField {
			return nil, fmt.Errorf("NewUpdateFieldValueIndex: Unable to get field: %v", currUpdate.FieldID)
		}

		decodedCellVal, decodeErr := DecodeCellValue(fieldInfo.Type, currUpdate.CellValue)
		if decodeErr != nil {
			return nil, decodeErr

		}

		fieldValUpdate := FieldValueUpdate{
			UpdateTimeStamp: currUpdate.UpdateTimeStamp,
			CellValue:       decodedCellVal}

		fieldValSeriesMap[currUpdate.FieldID] = append(fieldValSeriesMap[currUpdate.FieldID], fieldValUpdate)
		cellUpdates = append(cellUpdates, currUpdate)
	}

	// Sort the value updates for each fieldID in reverse chronological order.
	for currFieldID := range fieldValSeriesMap {
		sort.Sort(ByUpdateTime(fieldValSeriesMap[currFieldID]))
	}

	// Cell updates need to be in chronological order to be processed
	sort.Sort(CellUpdateByUpdateTime(cellUpdates))

	cellUpdateIndex := CellUpdateFieldValueIndex{
		fieldUpdateFieldValueIndex: fieldValSeriesMap,
		CellUpdates:                cellUpdates}

	return &cellUpdateIndex, nil

}
