package record

import (
	"appengine"
	"fmt"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/generic/cloudStorageWrapper"
	"resultra/datasheet/server/runtimeConfig"
)

type GetFieldValUrlParams struct {
	RecordID string `json:"recordID"`
	FieldID  string `json:"fieldID"`
}

type RecordFileFieldURLResponse struct {
	Url string `json:"url"`
}

func getFieldValUrl(appEngCntxt appengine.Context, params GetFieldValUrlParams) (*RecordFileFieldURLResponse, error) {

	record, getErr := GetRecord(appEngCntxt, params.RecordID)
	if getErr != nil {
		return nil, fmt.Errorf("getFieldValUrl: Unabled to get record: id = %v: get err=%v", params.RecordID, getErr)
	}

	theField, fieldErr := field.GetField(appEngCntxt, params.FieldID)
	if fieldErr != nil {
		return nil, fmt.Errorf("getFieldValUrl: Unabled to get field: id = %+v: get err=%v", params.FieldID, fieldErr)
	}

	if theField.Type != field.FieldTypeFile {
		return nil, fmt.Errorf("getFieldValUrl: Unexpected field type, expecting file, got %v", theField.Type)
	}

	// TODO check both record and field have same parent table ID

	cloudFileName, fileNameErr := record.GetTextFieldValue(params.FieldID)
	if fileNameErr != nil {
		return nil, fmt.Errorf(
			"getFieldValUrl: Unabled to get record's value for field: recordID = %v field id = %v: get err=%v",
			params.RecordID, params.FieldID, fileNameErr)
	}
	if len(cloudFileName) == 0 {
		return nil, fmt.Errorf(
			"getFieldValUrl: Unabled to get record's value for field: recordID = %v field id = %v: unexpected 0 length file name",
			params.RecordID, params.FieldID)
	}

	signedURL, urlErr := cloudStorageWrapper.GetSignedURL(runtimeConfig.CloudStorageBucketName,
		cloudFileName, runtimeConfig.CloudStorageAuthConfig, 60)
	if urlErr != nil {
		return nil, fmt.Errorf("uploadFile: Unable to create signed URL for newly uploaded file: %v", urlErr)
	}

	return &RecordFileFieldURLResponse{Url: signedURL}, nil

}
