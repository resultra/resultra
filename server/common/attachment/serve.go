package attachment

import (
	"fmt"
)

func GetAttachmentURL(cloudFileName string) string {
	// TODO - Replace localhost part with dynamically configured host name.
	fileURL := "http://localhost:8080/api/record/getFile/" + cloudFileName

	return fileURL
}

type AttachmentReference struct {
	AttachmentInfo AttachmentInfo `json:"attachmentInfo"`
	URL            string         `json:"url"`
}

func GetAttachmentReference(attachmentID string) (*AttachmentReference, error) {

	attachInfo, err := GetAttachmentInfo(attachmentID)
	if err != nil {
		return nil, fmt.Errorf("GetAttachmentReference: %v", err)
	}

	url := GetAttachmentURL(attachInfo.CloudFileName)

	attachRef := AttachmentReference{
		AttachmentInfo: *attachInfo,
		URL:            url}

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
