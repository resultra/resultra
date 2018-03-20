package attachment

import (
	"database/sql"
	"fmt"
	"net/http"
	"path/filepath"
	"resultra/datasheet/server/common/databaseWrapper"
	"resultra/datasheet/server/common/userAuth"
	"resultra/datasheet/server/generic/timestamp"
	"resultra/datasheet/server/generic/uniqueID"
	"time"
)

const attachTypeURL string = "url"
const attachTypeFile string = "file"
const cloudFileNameNone string = ""

type AttachmentInfo struct {
	AttachmentID       string    `json:"attachmentID"`
	ParentDatabaseID   string    `json:"parentDatabaseID"`
	UserID             string    `json:"userID"`
	CreateTimestampUTC time.Time `json:"createTimestampUTC"`
	Type               string    `json:"type"`
	CloudFileName      string    `json:"cloudFileName"`
	OrigFileName       string    `json:"origFileName"`
	Title              string    `json:"title"`
	Caption            string    `json:"caption"`
}

func saveAttachmentInfo(trackerDBHandle *sql.DB, attachInfo AttachmentInfo) error {

	if _, insertErr := trackerDBHandle.Exec(
		`INSERT INTO attachments (attachment_id, database_id, user_id, 
						create_timestamp_utc, type, orig_file_name,cloud_file_name,title,caption) 
			 VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)`,
		attachInfo.AttachmentID,
		attachInfo.ParentDatabaseID,
		attachInfo.UserID,
		attachInfo.CreateTimestampUTC,
		attachInfo.Type,
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

func newAttachmentInfo(parentDatabaseID string, userID string, attachmentType string, origFileName string, cloudFileName string) AttachmentInfo {

	attachInfo := AttachmentInfo{
		AttachmentID:       uniqueID.GenerateSnowflakeID(),
		ParentDatabaseID:   parentDatabaseID,
		UserID:             userID,
		Type:               attachmentType,
		CloudFileName:      cloudFileName,
		OrigFileName:       origFileName,
		CreateTimestampUTC: timestamp.CurrentTimestampUTC(),
		Title:              defaultTitle(origFileName),
		Caption:            ""}

	return attachInfo

}

func GetAttachmentInfo(trackerDBHandle *sql.DB, attachmentID string) (*AttachmentInfo, error) {

	attachInfo := AttachmentInfo{}
	getErr := trackerDBHandle.QueryRow(
		`SELECT attachment_id, database_id, user_id, create_timestamp_utc, type,orig_file_name,cloud_file_name,title,caption
		 FROM attachments
		 WHERE attachment_id=$1 LIMIT 1`, attachmentID).Scan(
		&attachInfo.AttachmentID,
		&attachInfo.ParentDatabaseID,
		&attachInfo.UserID,
		&attachInfo.CreateTimestampUTC,
		&attachInfo.Type,
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

func getOrigFilenameFromCloudFileName(trackerDBHandle *sql.DB, cloudFileName string) (string, error) {

	origFileName := ""

	getErr := trackerDBHandle.QueryRow(
		`SELECT orig_file_name
		 FROM attachments
		 WHERE cloud_file_name=$1 LIMIT 1`, cloudFileName).Scan(
		&origFileName)
	if getErr != nil {
		return "", fmt.Errorf("GetOrigFilenameFromCloudFileName: cloud file name = %v: datastore err=%v",
			cloudFileName, getErr)
	}
	return origFileName, nil

}

type SetCaptionParams struct {
	AttachmentID string `json:"attachmentID"`
	Caption      string `json:"caption"`
}

func setCaption(trackerDBHandle *sql.DB, params SetCaptionParams) (*AttachmentInfo, error) {

	if _, updateErr := trackerDBHandle.Exec(`UPDATE attachments 
			SET caption=$1 
			WHERE attachment_id=$2`,
		params.Caption,
		params.AttachmentID); updateErr != nil {
		return nil, fmt.Errorf("setCaption: Error updating caption: error = %v", updateErr)
	}

	return GetAttachmentInfo(trackerDBHandle, params.AttachmentID)
}

type SetTitleParams struct {
	AttachmentID string `json:"attachmentID"`
	Title        string `json:"title"`
}

func setTitle(trackerDBHandle *sql.DB, params SetTitleParams) (*AttachmentInfo, error) {

	if _, updateErr := trackerDBHandle.Exec(`UPDATE attachments 
			SET title=$1 
			WHERE attachment_id=$2`,
		params.Title,
		params.AttachmentID); updateErr != nil {
		return nil, fmt.Errorf("setTitle: Error updating title: error = %v", updateErr)
	}

	return GetAttachmentInfo(trackerDBHandle, params.AttachmentID)
}

type SaveURLParams struct {
	ParentDatabaseID string `json:"parentDatabaseID"`
	URL              string `json:"url"`
	Title            string `json:"title"`
	Caption          string `json:"caption"`
}

func saveURL(req *http.Request, params SaveURLParams) (*AttachmentInfo, error) {

	currUserID, userErr := userAuth.GetCurrentUserID(req)
	if userErr != nil {
		return nil, fmt.Errorf("uploadFile: Unable to get current user information: %v", userErr)
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(req)
	if dbErr != nil {
		return nil, dbErr
	}

	attachInfo := newAttachmentInfo(params.ParentDatabaseID, currUserID, attachTypeURL, params.URL, cloudFileNameNone)
	attachInfo.Title = params.Title
	attachInfo.Caption = params.Caption

	if saveErr := saveAttachmentInfo(trackerDBHandle, attachInfo); saveErr != nil {
		return nil, fmt.Errorf("uploadFile: unable to save attachment information/metadata: %v", saveErr)
	}

	return &attachInfo, nil

}
