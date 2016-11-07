package record

import (
	"fmt"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/userAuth"
	"sort"
	"time"
)

type FieldValTimelineChangeInfo struct {
	UpdateTimeStamp time.Time             `json:"updateTime"`
	UserName        string                `json:"userName"`
	IsCurrentUser   bool                  `json:"isCurrentUser"`
	UpdatedValue    interface{}           `json:"updatedValue"` // Decoded value, type depends on field.
	ValueFormat     CellUpdateValueFormat `json:"valueFormat"`
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

type UserTimelineVal struct {
	UserName      string `json:"userName"`
	IsCurrentUser bool   `json:"isCurrentUser"`
}

type FileTimelineValue struct {
	CloudName string `json:"cloudName"`
	OrigName  string `json:"origName"`
	Url       string `json:"url"`
}

func DecodeTimelineCellValue(currUserID string, fieldType string, encodedVal string) (interface{}, error) {

	if fieldType == field.FieldTypeUser {
		var userVal UserCellValue
		if err := generic.DecodeJSONString(encodedVal, &userVal); err != nil {
			return nil, fmt.Errorf("DecodeCellValue: failure decoding user value: %v", err)
		}

		userInfo, err := userAuth.GetUserInfoByID(userVal.UserID)
		if err != nil {
			return nil, fmt.Errorf("DecodeTimelineCellValue: %v", err)
		}

		isCurrentUser := false
		if currUserID == userVal.UserID {
			isCurrentUser = true
		}

		userTimelineVal := UserTimelineVal{
			UserName:      userInfo.UserName,
			IsCurrentUser: isCurrentUser}

		return userTimelineVal, nil

	} else if fieldType == field.FieldTypeFile {

		var fileVal FileCellValue
		if err := generic.DecodeJSONString(encodedVal, &fileVal); err != nil {
			return nil, fmt.Errorf("DecodeCellValue: failure decoding file value: %v", err)
		}

		fileTimelineVal := FileTimelineValue{
			OrigName:  fileVal.OrigName,
			CloudName: fileVal.CloudName,
			Url:       GetFileURL(fileVal.CloudName)}

		return fileTimelineVal, nil

	} else {
		// If the field type is anything but a user or file field, the decoded value can be returned as is.
		// However, if the field type is a user, additional information needs to be retrieved for display.
		return DecodeCellValue(fieldType, encodedVal)
	}
}

func GetFieldValUpdateTimelineInfo(currUserID string,
	parentTableID string, recordID string, fieldID string) ([]FieldValTimelineChangeInfo, error) {

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

		decodedCellVal, decodeErr := DecodeTimelineCellValue(currUserID, fieldInfo.Type, currUpdate.CellValue)
		if decodeErr != nil {
			return nil, fmt.Errorf(
				"NewUpdateFieldValueIndex: Unable to cell value for field ID = %v: tableID=%v: error=%v ",
				currUpdate.FieldID, parentTableID, decodeErr)

		}

		updateUserInfo, err := userAuth.GetUserInfoByID(currUpdate.UserID)
		if err != nil {
			return nil, fmt.Errorf("GetFieldValUpdateTimelineInfo: %v", err)
		}

		isCurrentUser := false
		if currUserID == currUpdate.UserID {
			isCurrentUser = true
		}

		fieldValChangeInfo := FieldValTimelineChangeInfo{
			UpdateTimeStamp: currUpdate.UpdateTimeStamp,
			UserName:        updateUserInfo.UserName,
			IsCurrentUser:   isCurrentUser,
			UpdatedValue:    decodedCellVal,
			ValueFormat:     currUpdate.Properties.ValueFormat}

		allFieldValChanges = append(allFieldValChanges, fieldValChangeInfo)

	}
	sort.Sort(ByTimelineUpdateTime(allFieldValChanges))

	return allFieldValChanges, nil
}
