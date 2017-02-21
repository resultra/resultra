package attachment

import (
	"fmt"
	"net/http"
	"resultra/datasheet/server/generic/api"
	"resultra/datasheet/server/generic/cloudStorageWrapper"
)

type AttachmentInfo struct {
	ParentDatabaseID string `json:"parentDatabaseID"`
	CloudFileName    string `json:"cloudFileName"`
	OrigFileName     string `json:"origFileName"`
}

type UploadedAttachment struct {
	Name           string         `json:"name"`
	Size           int            `json:"size"`
	Error          string         `json"error,omitempty"`
	Url            string         `json:"url"`
	AttachmentInfo AttachmentInfo `json:"attachmentInfo"`
}

type UploadedAttachmentResponse struct {
	Files []UploadedAttachment `json:"files"`
}

func uploadAttachment(req *http.Request) (*UploadedAttachmentResponse, error) {

	// The string "uploadFile" matches the parameter name used in clients.
	uploadInfo, uploadErr := api.ReadUploadFile(req, "uploadFile")
	if uploadErr != nil {
		return nil, fmt.Errorf("uploadAttachment: Unable to read file contents: %v", uploadErr)
	}

	cloudFileName := cloudStorageWrapper.UniqueCloudFileNameFromUserFileName(uploadInfo.FileName)
	if saveErr := cloudStorageWrapper.SaveCloudFile(cloudFileName, uploadInfo.FileData); saveErr != nil {
		return nil, fmt.Errorf("uploadFile: Unable to save file to cloud storage: %v", saveErr)
	}

	// Generate an URL for the newly saved file
	fileURL := GetAttachmentURL(cloudFileName)
	parentDatabaseID := req.FormValue("parentDatabaseID")

	// TODO - Add timestamp, current user

	attachInfo := AttachmentInfo{
		ParentDatabaseID: parentDatabaseID,
		CloudFileName:    cloudFileName,
		OrigFileName:     uploadInfo.FileName}

	uploadedAttachment := UploadedAttachment{
		Name:           cloudFileName,
		Size:           uploadInfo.FileLength,
		Error:          "",
		Url:            fileURL,
		AttachmentInfo: attachInfo}
	uploadResponse := UploadedAttachmentResponse{Files: []UploadedAttachment{uploadedAttachment}}

	return &uploadResponse, nil

}
