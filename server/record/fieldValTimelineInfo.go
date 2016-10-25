package record

import (
	"fmt"
	"resultra/datasheet/server/field"
	"sort"
	"time"
)

type FieldValTimelineChangeInfo struct {
	UpdateTimeStamp time.Time   `json:"updateTime"`
	UpdatedValue    interface{} `json:"updatedValue"` // Decoded value, type depends on field.
}

// Custom sort function for the FieldValTimelineChangeInfo
type ByTimelineUpdateTime []FieldValTimelineChangeInfo

func (s ByTimelineUpdateTime) Len() int {
	return len(s)
}
func (s ByTimelineUpdateTime) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Sort in reverse chronological order; i.e. the most recent dates come first.
func (s ByTimelineUpdateTime) Less(i, j int) bool {
	return s[i].UpdateTimeStamp.After(s[j].UpdateTimeStamp)
}

func GetFieldValUpdateTimelineInfo(parentTableID string, recordID string, fieldID string) ([]FieldValTimelineChangeInfo, error) {

	fieldCellUpdates, getErr := GetRecordFieldCellUpdates(parentTableID, recordID, fieldID)
	if getErr != nil {
		return nil, fmt.Errorf("GetCellUpdateTimelineInfo: failure retrieving cell updates for record = %v: error = %v",
			recordID, getErr)
	}

	fieldRefIndex, indexErr := field.GetFieldRefIDIndex(field.GetFieldListParams{ParentTableID: parentTableID})
	if indexErr != nil {
		return nil, fmt.Errorf("GetCellUpdateTimelineInfo: Unable to retrieve fields list for table: tableID=%v, error=%v ",
			parentTableID, indexErr)
	}

	allFieldValChanges := []FieldValTimelineChangeInfo{}
	for _, currUpdate := range fieldCellUpdates {

		fieldInfo, fieldErr := fieldRefIndex.GetFieldRefByID(currUpdate.FieldID)
		if fieldErr != nil {
			return nil, fmt.Errorf(
				"GetCellUpdateTimelineInfo: Unable to retrieve field information for field ID = %v: tableID=%v: error=%v ",
				currUpdate.FieldID, parentTableID, fieldErr)
		}

		decodedCellVal, decodeErr := DecodeCellValue(fieldInfo.Type, currUpdate.CellValue)
		if decodeErr != nil {
			return nil, fmt.Errorf(
				"NewUpdateFieldValueIndex: Unable to cell value for field ID = %v: tableID=%v: error=%v ",
				currUpdate.FieldID, parentTableID, decodeErr)

		}

		fieldValChangeInfo := FieldValTimelineChangeInfo{
			UpdateTimeStamp: currUpdate.UpdateTimeStamp,
			UpdatedValue:    decodedCellVal}

		allFieldValChanges = append(allFieldValChanges, fieldValChangeInfo)

	}
	sort.Sort(ByTimelineUpdateTime(allFieldValChanges))

	return allFieldValChanges, nil
}
