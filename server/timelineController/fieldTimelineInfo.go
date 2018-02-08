package timelineController

import (
	"database/sql"
	"fmt"
	"net/http"
	"resultra/datasheet/server/common/databaseWrapper"
	"resultra/datasheet/server/common/userAuth"
	"resultra/datasheet/server/record"
	"sort"
	"time"
)

type FieldTimelineInfo struct {
	UpdateTime         time.Time                          `json:"updateTime"`
	CommentInfo        *TimelineCommentInfo               `json:"commentInfo,omitempty"`
	FieldValChangeInfo *record.FieldValTimelineChangeInfo `json:"fieldValChangeInfo,omitempty"`
}

// Custom sort function for the FieldValTimelineChangeInfo
type ByFieldTimelineInfoUpdateTime []FieldTimelineInfo

func (s ByFieldTimelineInfoUpdateTime) Len() int {
	return len(s)
}
func (s ByFieldTimelineInfoUpdateTime) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Sort in reverse chronological order; i.e. the most recent dates come first.
func (s ByFieldTimelineInfoUpdateTime) Less(i, j int) bool {
	return s[i].UpdateTime.After(s[j].UpdateTime)
}

type FieldValChangeTimelineInfo struct {
}

func newFieldValChangeTimelineInfo(trackerDBHandle *sql.DB,
	currUserID string, comment FieldComment) (*TimelineCommentInfo, error) {

	isCurrentUser := false
	if currUserID == comment.UserID {
		isCurrentUser = true
	}

	commentUserInfo, err := userAuth.GetUserInfoByID(trackerDBHandle, comment.UserID)
	if err != nil {
		return nil, fmt.Errorf("getFieldRecordTimelineCommentInfo: %v", err)
	}

	commentInfo := TimelineCommentInfo{
		UserName:      commentUserInfo.UserName,
		IsCurrentUser: isCurrentUser,
		CommentID:     comment.CommentID,
		Comment:       comment.Comment,
		CommentDate:   comment.CreateTimestamp}

	return &commentInfo, nil

}

type GetFieldTimelineInfoParams struct {
	RecordID string `json:"recordID"`
	FieldID  string `json:"fieldID"`
}

func getFieldTimelineInfo(req *http.Request, params GetFieldTimelineInfoParams) ([]FieldTimelineInfo, error) {

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(req)
	if dbErr != nil {
		return nil, dbErr
	}

	commentParams := GetFieldRecordCommentInfoParams{
		RecordID: params.RecordID,
		FieldID:  params.FieldID}
	timelineComments, err := getFieldRecordTimelineCommentInfo(req, commentParams)
	if err != nil {
		return nil, fmt.Errorf("getFieldTimelineInfo: Error retrieving timeline comments: %+v, error = %v", params, err)
	}

	timelineInfoItems := []FieldTimelineInfo{}

	for _, currComment := range timelineComments {

		comment := currComment

		timelineInfo := FieldTimelineInfo{
			UpdateTime:  comment.CommentDate,
			CommentInfo: &comment}
		timelineInfoItems = append(timelineInfoItems, timelineInfo)
	}

	currUserID, err := userAuth.GetCurrentUserID(req)
	if err != nil {
		return nil, fmt.Errorf("getFieldTimelineInfo: %v", err)
	}

	fieldValTimelineChanges, err := record.GetFieldValUpdateTimelineInfo(trackerDBHandle, currUserID,
		params.RecordID, params.FieldID, record.FullyCommittedCellUpdatesChangeSetID)
	if err != nil {
		return nil, fmt.Errorf("getFieldTimelineInfo: Error retrieving timeline field value changes: %+v, error = %v", params, err)
	}

	for _, currFieldValChange := range fieldValTimelineChanges {

		valChange := currFieldValChange

		timelineInfo := FieldTimelineInfo{
			UpdateTime:         valChange.UpdateTimeStamp,
			FieldValChangeInfo: &valChange}
		timelineInfoItems = append(timelineInfoItems, timelineInfo)
	}

	sort.Sort(ByFieldTimelineInfoUpdateTime(timelineInfoItems))

	return timelineInfoItems, nil
}
