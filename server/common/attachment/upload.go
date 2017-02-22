package attachment

import (
	"fmt"
	"net/http"
	"resultra/datasheet/server/generic/api"
	"resultra/datasheet/server/generic/cloudStorageWrapper"
	"resultra/datasheet/server/generic/userAuth"
)

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

	currUserID, userErr := userAuth.GetCurrentUserID(req)
	if userErr != nil {
		return nil, fmt.Errorf("uploadFile: Unable to get current user information: %v", userErr)
	}

	attachInfo := newAttachmentInfo(parentDatabaseID, currUserID, uploadInfo.FileName, cloudFileName)
	if saveErr := saveAttachmentInfo(attachInfo); saveErr != nil {
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
