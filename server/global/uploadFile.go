package global

import (
	"fmt"
	"log"
	"net/http"
	"resultra/datasheet/server/common/attachment"
	"resultra/datasheet/server/common/databaseWrapper"
	"resultra/datasheet/server/generic/api"
)

type UploadFile struct {
	Name            string          `json:"name"`
	Size            int             `json:"size"`
	Error           string          `json"error,omitempty"`
	Url             string          `json:"url"`
	GlobalValUpdate GlobalValUpdate `json:"globalValUpdate"`
}

type UploadFileResponse struct {
	Files []UploadFile `json:"files"`
}

func uploadFile(req *http.Request) (*UploadFileResponse, error) {

	globalID := req.FormValue("globalID")
	parentDatabaseID := req.FormValue("parentDatabaseID")

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(req)
	if dbErr != nil {
		return nil, dbErr
	}

	global, getErr := getGlobal(trackerDBHandle, globalID)
	if getErr != nil {
		return nil, fmt.Errorf("uploadFile: Unable to retrieve global info for global ID = %v: %v",
			globalID, getErr)
	}
	if global.ParentDatabaseID != parentDatabaseID {
		return nil, fmt.Errorf("uploadFile: database and global info mismatch (database IDs don't match) info = %+v: database ID %v",
			global, parentDatabaseID)

	}

	// The string "uploadFile" matches the parameter name used in clients.
	uploadInfo, uploadErr := api.ReadUploadFile(req, "uploadFile")
	if uploadErr != nil {
		return nil, fmt.Errorf("uploadFile: Unable to read file contents: %v", uploadErr)
	}

	cloudFileName := attachment.UniqueAttachmentFileNameFromUserFileName(uploadInfo.FileName)
	if saveErr := attachment.SaveAttachmentFile(global.ParentDatabaseID, cloudFileName, uploadInfo.FileData); saveErr != nil {
		return nil, fmt.Errorf("uploadFile: Unable to save file to cloud storage: %v", saveErr)
	}

	// Generate an URL for the newly saved file
	fileURL := GetFileURL(trackerDBHandle, cloudFileName)

	// setRecordFileNameFieldValue. Although the parameters for a record update with a filename aren't passed through the http request,
	// the standard record updating mechanism can be used to update the field with the filename.
	updateGlobalHeader := GlobalValUpdateHeader{
		ParentDatabaseID: parentDatabaseID,
		GlobalID:         globalID}
	updateGlobalParams := SetImageGlobalValueParams{
		GlobalValUpdateHeader: updateGlobalHeader,
		OrigFileName:          uploadInfo.FileName,
		CloudFileName:         cloudFileName}

	valUpdate, updateErr := updateGlobalValue(trackerDBHandle, updateGlobalParams)
	if updateErr != nil {
		return nil, fmt.Errorf("uploadFile: Unable to update record for newly uploaded file: %v", updateErr)
	}

	log.Printf("uploadFile: Done uploading file: updated global = %+v", valUpdate)

	uploadFile := UploadFile{
		Name:            cloudFileName,
		Size:            uploadInfo.FileLength,
		Error:           "",
		Url:             fileURL,
		GlobalValUpdate: *valUpdate}
	uploadResponse := UploadFileResponse{Files: []UploadFile{uploadFile}}

	return &uploadResponse, nil
}
