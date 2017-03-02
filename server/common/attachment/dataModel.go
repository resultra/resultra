package attachment

import (
	"fmt"
	"path/filepath"
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
	Title              string    `json:"title"`
	Caption            string    `json:"caption"`
}

func saveAttachmentInfo(attachInfo AttachmentInfo) error {

	if _, insertErr := databaseWrapper.DBHandle().Exec(
		`INSERT INTO attachments (attachment_id, database_id, user_id, 
						create_timestamp_utc, orig_file_name,cloud_file_name,title,caption) 
			 VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`,
		attachInfo.AttachmentID,
		attachInfo.ParentDatabaseID,
		attachInfo.UserID,
		attachInfo.CreateTimestampUTC,
		attachInfo.OrigFileName,
		attachInfo.CloudFileName,
		attachInfo.Title,
		attachInfo.Caption); insertErr != nil {
		return fmt.Errorf("saveAttachmentInfo: insert failed: update = %+v, error = %v", attachInfo, insertErr)
	}
	return nil

}

/* Default title is the file name without the extension */
func defaultTitle(origFileName string) string {
	fileName := filepath.Base(origFileName)
	ext := filepath.Ext(origFileName)
	baseName := fileName[:len(fileName)-len(ext)]
	return baseName
}

func newAttachmentInfo(parentDatabaseID string, userID string, origFileName string, cloudFileName string) AttachmentInfo {

	attachInfo := AttachmentInfo{
		AttachmentID:       uniqueID.GenerateSnowflakeID(),
		ParentDatabaseID:   parentDatabaseID,
		UserID:             userID,
		CloudFileName:      cloudFileName,
		OrigFileName:       origFileName,
		CreateTimestampUTC: time.Now().UTC(),
		Title:              defaultTitle(origFileName),
		Caption:            ""}

	return attachInfo

}

func GetAttachmentInfo(attachmentID string) (*AttachmentInfo, error) {

	attachInfo := AttachmentInfo{}
	getErr := databaseWrapper.DBHandle().QueryRow(
		`SELECT attachment_id, database_id, user_id, create_timestamp_utc, orig_file_name,cloud_file_name,title,caption
		 FROM attachments
		 WHERE attachment_id=$1 LIMIT 1`, attachmentID).Scan(
		&attachInfo.AttachmentID,
		&attachInfo.ParentDatabaseID,
		&attachInfo.UserID,
		&attachInfo.CreateTimestampUTC,
		&attachInfo.OrigFileName,
		&attachInfo.CloudFileName,
		&attachInfo.Title,
		&attachInfo.Caption)
	if getErr != nil {
		return nil, fmt.Errorf("GetAttachmentInfo: Unabled to get attachment info: ID = %v: datastore err=%v",
			attachmentID, getErr)
	}
	return &attachInfo, nil
}

type SetCaptionParams struct {
	AttachmentID string `json:"attachmentID"`
	Caption      string `json:"caption"`
}

func setCaption(params SetCaptionParams) (*AttachmentInfo, error) {

	if _, updateErr := databaseWrapper.DBHandle().Exec(`UPDATE attachments 
			SET caption=$1 
			WHERE attachment_id=$2`,
		params.Caption,
		params.AttachmentID); updateErr != nil {
		return nil, fmt.Errorf("setCaption: Error updating caption: error = %v", updateErr)
	}

	return GetAttachmentInfo(params.AttachmentID)
}

type SetTitleParams struct {
	AttachmentID string `json:"attachmentID"`
	Title        string `json:"title"`
}

func setTitle(params SetTitleParams) (*AttachmentInfo, error) {

	if _, updateErr := databaseWrapper.DBHandle().Exec(`UPDATE attachments 
			SET title=$1 
			WHERE attachment_id=$2`,
		params.Title,
		params.AttachmentID); updateErr != nil {
		return nil, fmt.Errorf("setTitle: Error updating title: error = %v", updateErr)
	}

	return GetAttachmentInfo(params.AttachmentID)
}