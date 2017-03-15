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
const extPDF string = ".pdf"
const extDOC string = ".doc"

const dataTypeImage string = "image"
const dataTypeFile string = "file"
const dataTypeLink string = "link"

func getAttachmentDataType(attachInfo *AttachmentInfo) string {
	ext := path.Ext(attachInfo.OrigFileName)
	if strings.EqualFold(ext, extGIF) || strings.EqualFold(ext, extJPG) || strings.EqualFold(ext, extPNG) {
		// If the original name (whether it be a file name or URL) ends with an image extension, always
		// reference the attachment as an image. This will cause the image to be loaded as a thumbnail.
		return dataTypeImage
	} else if strings.EqualFold(ext, extPDF) || strings.EqualFold(ext, extDOC) {
		// Always treat PDF and DOC as a file, even if the file was given as an URL. This will allow the client
		// to provide a download link, preview etc., rather than just a link.
		// TODO - Include other well-known file types which might be linked to from an URL.
		return dataTypeFile
	} else {
		// Finally, if the attachment was provided as an URL, and it's not otherwise a simple link to a file
		// or image, then reference the attachment as a link.
		if attachInfo.Type == attachTypeURL {
			return dataTypeLink
		}
		return dataTypeFile
	}
}

func GetAttachmentReference(attachmentID string) (*AttachmentReference, error) {

	attachInfo, err := GetAttachmentInfo(attachmentID)
	if err != nil {
		return nil, fmt.Errorf("GetAttachmentReference: %v", err)
	}

	var url string
	if attachInfo.Type == attachTypeFile {
		url = GetAttachmentURL(attachInfo.CloudFileName)
	} else {
		url = attachInfo.OrigFileName
	}

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
