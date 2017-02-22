package attachment

import (
	"fmt"
	"resultra/datasheet/server/generic/databaseWrapper"
	"resultra/datasheet/server/generic/uniqueID"
	"time"
)

type AttachmentInfo struct {
	AttachmentID       string    `json:"attachmentID"`
	ParentDatabaseID   string    `json:"parentDatabaseID"`
	UserID             string    `json:"userID"`
	CreateTimestampUTC time.Time `json:"createTimestampUTC"`
	CloudFileName      string    `json:"cloudFileName"`
	OrigFileName       string    `json:"origFileName"`
}

func saveAttachmentInfo(attachInfo AttachmentInfo) error {

	if _, insertErr := databaseWrapper.DBHandle().Exec(
		`INSERT INTO attachments (attachment_id, database_id, user_id, create_timestamp_utc, orig_file_name,cloud_file_name) 
			 VALUES ($1,$2,$3,$4,$5,$6)`,
		attachInfo.AttachmentID,
		attachInfo.ParentDatabaseID,
		attachInfo.UserID,
		attachInfo.CreateTimestampUTC,
		attachInfo.OrigFileName,
		attachInfo.CloudFileName); insertErr != nil {
		return fmt.Errorf("saveAttachmentInfo: insert failed: update = %+v, error = %v", attachInfo, insertErr)
	}
	return nil

}

func newAttachmentInfo(parentDatabaseID string, userID string, origFileName string, cloudFileName string) AttachmentInfo {

	attachInfo := AttachmentInfo{
		AttachmentID:       uniqueID.GenerateSnowflakeID(),
		ParentDatabaseID:   parentDatabaseID,
		UserID:             userID,
		CloudFileName:      cloudFileName,
		OrigFileName:       origFileName,
		CreateTimestampUTC: time.Now().UTC()}

	return attachInfo

}

func GetAttachmentInfo(attachmentID string) (*AttachmentInfo, error) {

	attachInfo := AttachmentInfo{}
	getErr := databaseWrapper.DBHandle().QueryRow(
		`SELECT attachment_id, database_id, user_id, create_timestamp_utc, orig_file_name,cloud_file_name FROM attachments
		 WHERE attachment_id=$1 LIMIT 1`, attachmentID).Scan(
		&attachInfo.AttachmentID,
		&attachInfo.ParentDatabaseID,
		&attachInfo.UserID,
		&attachInfo.CreateTimestampUTC,
		&attachInfo.OrigFileName,
		&attachInfo.CloudFileName)
	if getErr != nil {
		return nil, fmt.Errorf("GetAttachmentInfo: Unabled to get attachment info: ID = %v: datastore err=%v",
			attachmentID, getErr)
	}
	return &attachInfo, nil
}
