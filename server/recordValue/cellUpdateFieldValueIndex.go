package recordValue

import (
	"fmt"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/recordUpdate"
	"sort"
	"time"
)

type FieldValueUpdate struct {
	UpdateTimeStamp time.Time
	CellValue       interface{} // Decoded value, type depends on field.
}

type FieldValueUpdateSeries []FieldValueUpdate

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

type CellUpdateFieldValueIndex map[string]FieldValueUpdateSeries

func NewUpdateFieldValueIndex(parentTableID string, recordID string) (*CellUpdateFieldValueIndex, error) {

	recCellUpdates, getErr := recordUpdate.GetRecordCellUpdates(parentTableID, recordID)
	if getErr != nil {
		return nil, fmt.Errorf("NewFieldValueIndex: failure retrieving cell updates for record = %v: error = %v",
			recordID, getErr)
	}

	fieldRefIndex, indexErr := field.GetFieldRefIDIndex(field.GetFieldListParams{ParentTableID: parentTableID})
	if indexErr != nil {
		return nil, fmt.Errorf("NewUpdateFieldValueIndex: Unable to retrieve fields list for table: tableID=%v, error=%v ",
			parentTableID, indexErr)
	}

	// Populate the index with all the updates for the given recordID, broken down by FieldID.
	var fieldValSeriesMap CellUpdateFieldValueIndex
	for _, currUpdate := range recCellUpdates {

		fieldInfo, fieldErr := fieldRefIndex.GetFieldRefByID(currUpdate.FieldID)
		if fieldErr != nil {
			return nil, fmt.Errorf(
				"NewUpdateFieldValueIndex: Unable to retrieve field information for field ID = %v: tableID=%v: error=%v ",
				currUpdate.FieldID, parentTableID, fieldErr)
		}

		decodedCellVal, decodeErr := recordUpdate.DecodeCellValue(fieldInfo.Type, currUpdate.CellValue)
		if decodeErr != nil {
			return nil, fmt.Errorf(
				"NewUpdateFieldValueIndex: Unable to cell value for field ID = %v: tableID=%v: error=%v ",
				currUpdate.FieldID, parentTableID, decodeErr)

		}

		fieldValUpdate := FieldValueUpdate{
			UpdateTimeStamp: currUpdate.UpdateTimeStamp,
			CellValue:       decodedCellVal}

		fieldValSeriesMap[currUpdate.FieldID] = append(fieldValSeriesMap[currUpdate.FieldID], fieldValUpdate)
	}

	// Sort the value updates for each fieldID in reverse chronological order.
	for currFieldID := range fieldValSeriesMap {
		sort.Sort(ByUpdateTime(fieldValSeriesMap[currFieldID]))
	}

	return &fieldValSeriesMap, nil
}
