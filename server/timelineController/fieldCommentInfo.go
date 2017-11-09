package timelineController

import (
	"database/sql"
	"fmt"
	"net/http"
	"resultra/datasheet/server/common/databaseWrapper"
	"resultra/datasheet/server/generic/userAuth"
	"time"
)

type TimelineCommentInfo struct {
	UserName      string    `json:"userName"`
	IsCurrentUser bool      `json:"isCurrentUser"`
	CommentID     string    `json:"commentID"`
	Comment       string    `json:"comment"`
	CommentDate   time.Time `json:"commentDate"`
}

type GetFieldRecordCommentInfoParams struct {
	RecordID string `json:"recordID"`
	FieldID  string `json:"fieldID"`
}

func newTimelineCommentInfo(trackerDBHandle *sql.DB,
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

func saveTimelineComment(req *http.Request, params SaveFieldCommentParams) (*TimelineCommentInfo, error) {

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(req)
	if dbErr != nil {
		return nil, dbErr
	}

	newComment, err := saveFieldComment(trackerDBHandle, req, params)
	if err != nil {
		return nil, fmt.Errorf("saveTimelineComment: %v", err)
	}
	currUserID, err := userAuth.GetCurrentUserID(req)
	if err != nil {
		return nil, fmt.Errorf("saveTimelineComment: %v", err)
	}

	return newTimelineCommentInfo(trackerDBHandle, currUserID, *newComment)
}

func getFieldRecordTimelineCommentInfo(req *http.Request, params GetFieldRecordCommentInfoParams) ([]TimelineCommentInfo, error) {

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(req)
	if dbErr != nil {
		return nil, dbErr
	}

	currUserID, err := userAuth.GetCurrentUserID(req)
	if err != nil {
		return nil, fmt.Errorf("getFieldRecordTimelineCommentInfo: %v", err)
	}

	commentParams := GetFieldCommentsParams{
		RecordID: params.RecordID,
		FieldID:  params.FieldID}
	comments, commentErr := GetFieldComments(trackerDBHandle, commentParams)
	if commentErr != nil {
		return nil, fmt.Errorf("getFieldRecordTimelineCommentInfo: %v", commentErr)
	}

	timelineCommentsInfo := []TimelineCommentInfo{}
	for _, currComment := range comments {

		commentInfo, err := newTimelineCommentInfo(trackerDBHandle, currUserID, currComment)
		if err != nil {
			return nil, fmt.Errorf("getFieldRecordTimelineCommentInfo: %v", err)
		}

		timelineCommentsInfo = append(timelineCommentsInfo, *commentInfo)
	}

	return timelineCommentsInfo, nil

}
