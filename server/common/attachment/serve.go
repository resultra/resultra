package attachment

import (
	"fmt"
	"path/filepath"
	"strings"
)

func GetAttachmentURL(cloudFileName string) string {
	// TODO - Replace localhost part with dynamically configured host name.
	fileURL := "http://localhost:8080/api/record/getFile/" + cloudFileName

	return fileURL
}

type AttachmentReference struct {
	AttachmentInfo AttachmentInfo `json:"attachmentInfo"`
	URL            string         `json:"url"`
	DataType       string         `json:"dataType"`
}

const extGIF string = "gif"
const extJPG string = "jpg"
const extPNG string = "png"

const dataTypeImage string = "image"
const dataTypeFile string = "file"

func getAttachmentDataType(attachInfo *AttachmentInfo) string {

	ext := filepath.Ext(attachInfo.OrigFileName)
	if strings.EqualFold(ext, extGIF) || strings.EqualFold(ext, extJPG) || strings.EqualFold(ext, extPNG) {
		return dataTypeImage
	} else {
		return dataTypeFile
	}
}

func GetAttachmentReference(attachmentID string) (*AttachmentReference, error) {

	attachInfo, err := GetAttachmentInfo(attachmentID)
	if err != nil {
		return nil, fmt.Errorf("GetAttachmentReference: %v", err)
	}

	url := GetAttachmentURL(attachInfo.CloudFileName)
	dataType := getAttachmentDataType(attachInfo)

	attachRef := AttachmentReference{
		AttachmentInfo: *attachInfo,
		URL:            url,
		DataType:       dataType}

	return &attachRef, nil

}

func getAttachmentReferences(attachmentIDs []string) ([]AttachmentReference, error) {

	refs := []AttachmentReference{}
	for _, attachmentID := range attachmentIDs {
		currRef, err := GetAttachmentReference(attachmentID)
		if err != nil {
			return nil, fmt.Errorf("getAttachmentReferences: error getting reference for attachment with ID = %v: %v", attachmentID, err)
		}
		refs = append(refs, *currRef)
	}

	return refs, nil

}
