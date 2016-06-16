package record

import (
	"fmt"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/generic/cloudStorageWrapper"
	"resultra/datasheet/server/runtimeConfig"
)

type GetFieldValUrlParams struct {
	ParentTableID string `json:"parentTableID"`
	RecordID      string `json:"recordID"`
	FieldID       string `json:"fieldID"`
	CloudFileName string `json:"cloudFileName"`
}

type RecordFileFieldURLResponse struct {
	Url string `json:"url"`
}

func getFieldValUrl(params GetFieldValUrlParams) (*RecordFileFieldURLResponse, error) {

	_, getErr := GetRecord(params.ParentTableID, params.RecordID)
	if getErr != nil {
		return nil, fmt.Errorf("getFieldValUrl: Unabled to get record: id = %v: get err=%v", params.RecordID, getErr)
	}

	theField, fieldErr := field.GetField(params.ParentTableID, params.FieldID)
	if fieldErr != nil {
		return nil, fmt.Errorf("getFieldValUrl: Unabled to get field: id = %+v: get err=%v", params.FieldID, fieldErr)
	}

	if theField.Type != field.FieldTypeFile {
		return nil, fmt.Errorf("getFieldValUrl: Unexpected field type, expecting file, got %v", theField.Type)
	}

	// TODO check both record and field have same parent table ID

	if len(params.CloudFileName) == 0 {
		return nil, fmt.Errorf(
			"getFieldValUrl: Unabled to get record's value for field: recordID = %v field id = %v: unexpected 0 length file name",
			params.RecordID, params.FieldID)
	}

	signedURL, urlErr := cloudStorageWrapper.GetSignedURL(runtimeConfig.CloudStorageBucketName,
		params.CloudFileName, runtimeConfig.CloudStorageAuthConfig, 60)
	if urlErr != nil {
		return nil, fmt.Errorf("uploadFile: Unable to create signed URL for newly uploaded file: %v", urlErr)
	}

	return &RecordFileFieldURLResponse{Url: signedURL}, nil

}
