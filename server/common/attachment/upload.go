package attachment

import (
	"fmt"
	"net/http"
	"path"
	"resultra/tracker/server/common/databaseWrapper"
	"resultra/tracker/server/common/userAuth"
	"resultra/tracker/server/generic/api"
	"resultra/tracker/server/generic/uniqueID"
)

func UniqueAttachmentFileNameFromUserFileName(userFileName string) string {

	uniqueIDStr := uniqueID.GenerateUniqueIDWithTime()

	fileExt := path.Ext(userFileName)

	cloudFileName := uniqueIDStr + fileExt

	return cloudFileName
}

type UploadedAttachment struct {
	Name           string         `json:"name"`
	Size           int            `json:"size"`
	Error          string         `json"error,omitempty"`
	Url            string         `json:"url"`
	AttachmentInfo AttachmentInfo `json:"attachmentInfo"`
}

type UploadedAttachmentResponse struct {
	// Even though only a single file is uploaded at once, the jQuery File Upload plugin
	// requires the return of upload information in an array.
	Files []UploadedAttachment `json:"files"`
}

func uploadAttachment(w http.ResponseWriter, req *http.Request) (*UploadedAttachmentResponse, error) {

	// The string "uploadFile" matches the parameter name used in clients.
	uploadInfo, uploadErr := api.ReadUploadFile(w, req, "uploadFile")
	if uploadErr != nil {
		return nil, fmt.Errorf("uploadAttachment: Unable to read file contents: %v", uploadErr)
	}

	// This call to req.FormValue needs to come after the reading the upload file, since
	// req.FormValue will in turn call http.ParseMultipartForm. However, api.ReadUploadFile
	// needs to do error checking on the size of the requested upload through a call
	// to http.ParseMultipartForm.
	parentDatabaseID := req.FormValue("parentDatabaseID")
	cloudFileName := UniqueAttachmentFileNameFromUserFileName(uploadInfo.FileName)

	saveParams := databaseWrapper.SaveAttachmentParams{
		CloudFileName:    cloudFileName,
		ParentDatabaseID: parentDatabaseID,
		HTTPReq:          req,
		FileData:         uploadInfo.FileData}

	if saveErr := databaseWrapper.SaveAttachment(saveParams); saveErr != nil {
		return nil, fmt.Errorf("uploadFile: Unable to save file to cloud storage: %v", saveErr)
	}

	// Generate an URL for the newly saved file
	fileURL := GetAttachmentURL(parentDatabaseID, cloudFileName)

	currUserID, userErr := userAuth.GetCurrentUserID(req)
	if userErr != nil {
		return nil, fmt.Errorf("uploadFile: Unable to get current user information: %v", userErr)
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(req)
	if dbErr != nil {
		return nil, dbErr
	}

	attachInfo := newAttachmentInfo(parentDatabaseID, currUserID, attachTypeFile, uploadInfo.FileName, cloudFileName)
	if saveErr := saveAttachmentInfo(trackerDBHandle, attachInfo); saveErr != nil {
		return nil, fmt.Errorf("uploadFile: unable to save attachment information/metadata: %v", saveErr)
	}

	uploadedAttachment := UploadedAttachment{
		Name:           cloudFileName,
		Size:           uploadInfo.FileLength,
		Error:          "",
		Url:            fileURL,
		AttachmentInfo: attachInfo}
	uploadResponse := UploadedAttachmentResponse{Files: []UploadedAttachment{uploadedAttachment}}

	return &uploadResponse, nil

}
