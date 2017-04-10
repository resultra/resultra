package attachment

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"resultra/datasheet/server/common/runtimeConfig"
	"resultra/datasheet/server/generic/api"
	"resultra/datasheet/server/generic/uniqueID"
	"resultra/datasheet/server/generic/userAuth"
)

const permsOwnerReadWriteOnly os.FileMode = 0700

func fullyQualifiedAttachmentFileName(fileName string) string {
	return runtimeConfig.CurrRuntimeConfig.AttachmentBasePath + fileName
}

func InitAttachmentBasePath() error {
	err := os.MkdirAll(runtimeConfig.CurrRuntimeConfig.AttachmentBasePath, permsOwnerReadWriteOnly)
	if err != nil {
		return fmt.Errorf("Error initializing attachment directory %v: %v",
			runtimeConfig.CurrRuntimeConfig.AttachmentBasePath, err)
	}
	return nil
}

func SaveAttachmentFile(fileName string, fileData []byte) error {

	writeErr := ioutil.WriteFile(runtimeConfig.CurrRuntimeConfig.AttachmentBasePath+fileName, fileData, permsOwnerReadWriteOnly)
	if writeErr != nil {
		return fmt.Errorf("SaveCloudFile: Error writing file: filename = %v, error = %v", fileName, writeErr)
	}

	log.Printf("uploadFile: ... done uloading file to cloud: file name = %v", fileName)

	return nil

}

func UniqueAttachmentFileNameFromUserFileName(userFileName string) string {

	uniqueIDStr := uniqueID.GenerateUniqueID()

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

func uploadAttachment(req *http.Request) (*UploadedAttachmentResponse, error) {

	// The string "uploadFile" matches the parameter name used in clients.
	uploadInfo, uploadErr := api.ReadUploadFile(req, "uploadFile")
	if uploadErr != nil {
		return nil, fmt.Errorf("uploadAttachment: Unable to read file contents: %v", uploadErr)
	}

	cloudFileName := UniqueAttachmentFileNameFromUserFileName(uploadInfo.FileName)
	if saveErr := SaveAttachmentFile(cloudFileName, uploadInfo.FileData); saveErr != nil {
		return nil, fmt.Errorf("uploadFile: Unable to save file to cloud storage: %v", saveErr)
	}

	// Generate an URL for the newly saved file
	fileURL := GetAttachmentURL(cloudFileName)
	parentDatabaseID := req.FormValue("parentDatabaseID")

	currUserID, userErr := userAuth.GetCurrentUserID(req)
	if userErr != nil {
		return nil, fmt.Errorf("uploadFile: Unable to get current user information: %v", userErr)
	}

	attachInfo := newAttachmentInfo(parentDatabaseID, currUserID, attachTypeFile, uploadInfo.FileName, cloudFileName)
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
