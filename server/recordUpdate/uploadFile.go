package recordUpdate

import (
	"fmt"
	"log"
	"net/http"
	"resultra/datasheet/server/generic/api"
	"resultra/datasheet/server/generic/cloudStorageWrapper"
	"resultra/datasheet/server/record"
	"resultra/datasheet/server/recordValue"
)

type UploadFile struct {
	Name          string                         `json:"name"`
	Size          int                            `json:"size"`
	Error         string                         `json"error,omitempty"`
	Url           string                         `json:"url"`
	UpdatedRecord recordValue.RecordValueResults `json:"updatedRecord"`
}

type UploadFileResponse struct {
	Files []UploadFile `json:"files"`
}

func uploadFile(req *http.Request) (*UploadFileResponse, error) {

	// The string "uploadFile" matches the parameter name used in clients.
	uploadInfo, uploadErr := api.ReadUploadFile(req, "uploadFile")
	if uploadErr != nil {
		return nil, fmt.Errorf("uploadFile: Unable to read file contents: %v", uploadErr)
	}

	cloudFileName := cloudStorageWrapper.UniqueCloudFileNameFromUserFileName(uploadInfo.FileName)
	if saveErr := cloudStorageWrapper.SaveCloudFile(cloudFileName, uploadInfo.FileData); saveErr != nil {
		return nil, fmt.Errorf("uploadFile: Unable to save file to cloud storage: %v", saveErr)
	}

	// Generate an URL for the newly saved file
	fileURL := record.GetFileURL(cloudFileName)

	// setRecordFileNameFieldValue. Although the parameters for a record update with a filename aren't passed through the http request,
	// the standard record updating mechanism can be used to update the field with the filename.
	updateRecordHeader := record.RecordUpdateHeader{
		ParentTableID: req.FormValue("parentTableID"),
		RecordID:      req.FormValue("recordID"),
		FieldID:       req.FormValue("fieldID")}
	valueFormat := record.CellUpdateValueFormat{
		Context: req.FormValue("valueFormatContext"),
		Format:  req.FormValue("valueFormatFormat")}
	updateRecordParams := record.SetRecordFileValueParams{
		RecordUpdateHeader: updateRecordHeader,
		OrigFileName:       uploadInfo.FileName,
		CloudFileName:      cloudFileName,
		ValueFormat:        valueFormat}

	updatedRecord, updateErr := updateRecordValue(req, updateRecordParams)
	if updateErr != nil {
		return nil, fmt.Errorf("uploadFile: Unable to update record for newly uploaded file: %v", updateErr)
	}

	log.Printf("uploadFile: Done uploading file: updated record = %+v", updatedRecord)

	uploadFile := UploadFile{
		Name:          cloudFileName,
		Size:          uploadInfo.FileLength,
		Error:         "",
		Url:           fileURL,
		UpdatedRecord: *updatedRecord}
	uploadResponse := UploadFileResponse{Files: []UploadFile{uploadFile}}

	return &uploadResponse, nil
}
