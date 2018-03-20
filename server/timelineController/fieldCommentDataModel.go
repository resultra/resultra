package timelineController

import (
	"database/sql"
	"fmt"
	"net/http"
	"resultra/datasheet/server/common/userAuth"
	"resultra/datasheet/server/generic/timestamp"
	"resultra/datasheet/server/generic/uniqueID"
	"time"
)

type SaveFieldCommentParams struct {
	FieldID  string `json:"fieldID"`
	RecordID string `json:"recordID"`
	Comment  string `json:"comment"`
}

type FieldComment struct {
	UserID          string    `json:"userID"`
	CommentID       string    `json:"commentID"`
	Comment         string    `json:"comment"`
	RecordID        string    `json:"recordID"`
	FieldID         string    `json:"fieldID"`
	UpdateTimestamp time.Time `json:"updateTimestamp"`
	CreateTimestamp time.Time `json:"createTimestamp"`
}

func saveFieldComment(trackerDBHandle *sql.DB, req *http.Request, params SaveFieldCommentParams) (*FieldComment, error) {

	commentTimestamp := timestamp.CurrentTimestampUTC()
	commentID := uniqueID.GenerateUniqueID()
	currUserID, userErr := userAuth.GetCurrentUserID(req)
	if userErr != nil {
		return nil, fmt.Errorf("saveFieldComment: can't get current user user: %v", userErr)
	}

	newComment := FieldComment{
		UserID:          currUserID,
		CommentID:       commentID,
		Comment:         params.Comment,
		RecordID:        params.RecordID,
		FieldID:         params.FieldID,
		CreateTimestamp: commentTimestamp,
		UpdateTimestamp: commentTimestamp}

	if _, insertErr := trackerDBHandle.Exec(
		`INSERT INTO field_comments (user_id,comment_id,comment,record_id,field_id,create_timestamp_utc,update_timestamp_utc) 
					VALUES ($1,$2,$3,$4,$5,$6,$7)`,
		newComment.UserID,
		newComment.CommentID,
		newComment.Comment,
		newComment.RecordID,
		newComment.FieldID,
		newComment.CreateTimestamp,
		newComment.UpdateTimestamp); insertErr != nil {
		return nil, fmt.Errorf("saveFieldComment: Can't create comment: error = %v", insertErr)
	}

	return &newComment, nil

}

type GetFieldCommentsParams struct {
	RecordID string `json:"recordID"`
	FieldID  string `json:"fieldID"`
}

func GetFieldComments(trackerDBHandle *sql.DB, params GetFieldCommentsParams) ([]FieldComment, error) {

	rows, queryErr := trackerDBHandle.Query(
		`SELECT user_id,comment_id,comment,record_id,field_id,create_timestamp_utc,update_timestamp_utc 
		FROM field_comments 
		WHERE record_id=$1 AND field_id=$2`, params.RecordID, params.FieldID)
	if queryErr != nil {
		return nil, fmt.Errorf("getTableList: Failure querying database: %v", queryErr)
	}
	defer rows.Close()

	allComments := []FieldComment{}
	for rows.Next() {
		var currComment FieldComment
		if scanErr := rows.Scan(&currComment.UserID,
			&currComment.CommentID,
			&currComment.Comment,
			&currComment.RecordID,
			&currComment.FieldID,
			&currComment.CreateTimestamp,
			&currComment.UpdateTimestamp); scanErr != nil {
			return nil, fmt.Errorf("GetAllFieldFieldComments: Failure querying database: %v", scanErr)

		}
		allComments = append(allComments, currComment)
	}

	return allComments, nil
}
