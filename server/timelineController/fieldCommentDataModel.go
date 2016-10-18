package timelineController

import (
	"fmt"
	"net/http"
	"resultra/datasheet/server/generic/databaseWrapper"
	"resultra/datasheet/server/generic/uniqueID"
	"resultra/datasheet/server/generic/userAuth"
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

func saveFieldComment(req *http.Request, params SaveFieldCommentParams) (*FieldComment, error) {

	commentTimestamp := time.Now().UTC()
	commentID := uniqueID.GenerateSnowflakeID()
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

	if _, insertErr := databaseWrapper.DBHandle().Exec(
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

func GetFieldComments(params GetFieldCommentsParams) ([]FieldComment, error) {

	rows, queryErr := databaseWrapper.DBHandle().Query(
		`SELECT user_id,comment_id,comment,record_id,field_id,create_timestamp_utc,update_timestamp_utc 
		FROM field_comments 
		WHERE record_id=$1 AND field_id=$2`, params.RecordID, params.FieldID)
	if queryErr != nil {
		return nil, fmt.Errorf("getTableList: Failure querying database: %v", queryErr)
	}
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
