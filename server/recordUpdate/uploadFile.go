package recordUpdate

import (
	"appengine"
	"fmt"
	"net/http"
	"resultra/datasheet/server/generic/api"
	"resultra/datasheet/server/generic/cloudStorageWrapper"
	"resultra/datasheet/server/record"
)

type UploadFile struct {
	Name             string           `json:"name"`
	Size             int              `json:"size"`
	Error            string           `json"error,omitempty"`
	Url              string           `json:"url"`
	UpdatedRecordRef record.RecordRef `json:"updatedRecordRef"`
}

type UploadFileResponse struct {
	Files []UploadFile `json:"files"`
}

const cloudStorageBucketName string = "resultra-db-dev"
const cloudStorageJSONAuthInfoFile string = "/Users/sroehling/Development/Datasheet-Dev-60167588e163.json"

func uploadFile(req *http.Request) (*UploadFileResponse, error) {

	uploadInfo, uploadErr := api.ReadUploadFile(req, "uploadFile")
	if uploadErr != nil {
		return nil, fmt.Errorf("uploadFile: Unable to read file contents: %v", uploadErr)
	}

	authConfig, configErr := cloudStorageWrapper.ReadServiceAuthConfig(cloudStorageJSONAuthInfoFile)
	if configErr != nil {
		return nil, fmt.Errorf("uploadFile: Unable to get credentials for cloud storage: %v", configErr)
	}
	cloudAppEngContext := cloudStorageWrapper.NewCloudStorageContext(req)

	cloudFileName := cloudStorageWrapper.UniqueCloudFileNameFromUserFileName(uploadInfo.FileName)
	if saveErr := cloudStorageWrapper.SaveCloudFile(cloudAppEngContext, authConfig,
		cloudStorageBucketName, cloudFileName, uploadInfo.FileData); saveErr != nil {
		return nil, fmt.Errorf("uploadFile: Unable to save file to cloud storage: %v", saveErr)
	}

	// Generate an URL for the newly saved file
	signedURL, urlErr := cloudStorageWrapper.GetSignedURL(cloudStorageBucketName, cloudFileName, authConfig, 60)
	if urlErr != nil {
		return nil, fmt.Errorf("uploadFile: Unable to create signed URL for newly uploaded file: %v", urlErr)
	}

	appEngContext := appengine.NewContext(req)
	updatedRecordRef, updateErr := setRecordFileNameFieldValue(appEngContext,
		req.FormValue("recordID"), req.FormValue("fieldID"), cloudFileName)
	if updateErr != nil {
		return nil, fmt.Errorf("uploadFile: Unable to update record for newly uploaded file: %v", updateErr)
	}

	uploadFile := UploadFile{
		Name:             cloudFileName,
		Size:             uploadInfo.FileLength,
		Error:            "",
		Url:              signedURL,
		UpdatedRecordRef: *updatedRecordRef}
	uploadResponse := UploadFileResponse{Files: []UploadFile{uploadFile}}

	return &uploadResponse, nil
}
