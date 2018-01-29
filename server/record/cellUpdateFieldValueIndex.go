package record

import (
	"fmt"
	"resultra/datasheet/server/field"
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
