package record

import (
	"fmt"
	"resultra/datasheet/server/field"
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

func GetFileURL(cloudFileName string) string {
	// TODO - Replace localhost part with dynamically configured host name.
	fileURL := "http://localhost:8080/api/record/getFile/" + cloudFileName

	return fileURL
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

	fileURL := GetFileURL(params.CloudFileName)

	return &RecordFileFieldURLResponse{Url: fileURL}, nil

}
