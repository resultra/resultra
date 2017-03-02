package attachment

import (
	"fmt"
	"path"
	"resultra/datasheet/server/common/runtimeConfig"
	"strings"
)

func GetAttachmentURL(cloudFileName string) string {
	fileURLSubPath := "api/record/getFile/" + cloudFileName
	return runtimeConfig.GetSiteResourceURL(fileURLSubPath)
}

type AttachmentReference struct {
	AttachmentInfo AttachmentInfo `json:"attachmentInfo"`
	URL            string         `json:"url"`
	DataType       string         `json:"dataType"`
	Extension      string         `json:"extension"`
}

const extGIF string = ".gif"
const extJPG string = ".jpg"
const extPNG string = ".png"

const dataTypeImage string = "image"
const dataTypeFile string = "file"

func getAttachmentDataType(attachInfo *AttachmentInfo) string {

	ext := path.Ext(attachInfo.OrigFileName)
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
	ext := strings.TrimPrefix(path.Ext(attachInfo.OrigFileName), ".")

	attachRef := AttachmentReference{
		AttachmentInfo: *attachInfo,
		URL:            url,
		DataType:       dataType,
		Extension:      ext}

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
