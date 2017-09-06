package urlLink

import (
	"fmt"
	"log"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/form/components/common"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/uniqueID"
)

const urlLinkEntityKind string = "urlLink"

type UrlLink struct {
	ParentFormID string            `json:"parentFormID"`
	UrlLinkID    string            `json:"urlLinkID"`
	Properties   UrlLinkProperties `json:"properties"`
}

type NewUrlLinkParams struct {
	ParentFormID string                         `json:"parentFormID"`
	FieldID      string                         `json:"fieldID"`
	Geometry     componentLayout.LayoutGeometry `json:"geometry"`
}

func validUrlLinkFieldType(fieldType string) bool {
	if fieldType == field.FieldTypeURL {
		return true
	} else {
		return false
	}
}

func saveUrlLink(newUrlLink UrlLink) error {
	if saveErr := common.SaveNewFormComponent(urlLinkEntityKind,
		newUrlLink.ParentFormID, newUrlLink.UrlLinkID, newUrlLink.Properties); saveErr != nil {
		return fmt.Errorf("saveUrlLink: Unable to save text box: %v", saveErr)
	}
	return nil

}

func saveNewUrlLink(params NewUrlLinkParams) (*UrlLink, error) {

	if !componentLayout.ValidGeometry(params.Geometry) {
		return nil, fmt.Errorf("Invalid layout container parameters: %+v", params)
	}

	if fieldErr := field.ValidateField(params.FieldID, validUrlLinkFieldType); fieldErr != nil {
		return nil, fmt.Errorf("saveNewUrlLink: %v", fieldErr)
	}

	properties := newDefaultUrlLinkProperties()
	properties.Geometry = params.Geometry
	properties.FieldID = params.FieldID

	newUrlLink := UrlLink{ParentFormID: params.ParentFormID,
		UrlLinkID:  uniqueID.GenerateSnowflakeID(),
		Properties: properties}

	if err := saveUrlLink(newUrlLink); err != nil {
		return nil, fmt.Errorf("saveNewUrlLink: Unable to save text box with params=%+v: error = %v", params, err)
	}

	log.Printf("INFO: API: NewLayout: Created new Layout container: %+v", newUrlLink)

	return &newUrlLink, nil

}

func getUrlLink(parentFormID string, urlLinkID string) (*UrlLink, error) {

	urlLinkProps := newDefaultUrlLinkProperties()
	if getErr := common.GetFormComponent(urlLinkEntityKind, parentFormID, urlLinkID, &urlLinkProps); getErr != nil {
		return nil, fmt.Errorf("getCheckBox: Unable to retrieve text box: %v", getErr)
	}

	urlLink := UrlLink{
		ParentFormID: parentFormID,
		UrlLinkID:    urlLinkID,
		Properties:   urlLinkProps}

	return &urlLink, nil
}

func GetUrlLinks(parentFormID string) ([]UrlLink, error) {

	urlLinkes := []UrlLink{}
	addUrlLink := func(urlLinkID string, encodedProps string) error {

		urlLinkProps := newDefaultUrlLinkProperties()
		if decodeErr := generic.DecodeJSONString(encodedProps, &urlLinkProps); decodeErr != nil {
			return fmt.Errorf("GetUrlLink: can't decode properties: %v", encodedProps)
		}

		currUrlLink := UrlLink{
			ParentFormID: parentFormID,
			UrlLinkID:    urlLinkID,
			Properties:   urlLinkProps}
		urlLinkes = append(urlLinkes, currUrlLink)

		return nil
	}
	if getErr := common.GetFormComponents(urlLinkEntityKind, parentFormID, addUrlLink); getErr != nil {
		return nil, fmt.Errorf("GetCheckBoxes: Can't get text boxes: %v")
	}

	return urlLinkes, nil

}

func CloneUrlLinks(remappedIDs uniqueID.UniqueIDRemapper, parentFormID string) error {

	srcUrlLink, err := GetUrlLinks(parentFormID)
	if err != nil {
		return fmt.Errorf("CloneUrlLink: %v", err)
	}

	for _, srcUrlLink := range srcUrlLink {
		remappedUrlLinkID := remappedIDs.AllocNewOrGetExistingRemappedID(srcUrlLink.UrlLinkID)
		remappedFormID, err := remappedIDs.GetExistingRemappedID(srcUrlLink.ParentFormID)
		if err != nil {
			return fmt.Errorf("CloneUrlLink: %v", err)
		}
		destProperties, err := srcUrlLink.Properties.Clone(remappedIDs)
		if err != nil {
			return fmt.Errorf("CloneUrlLink: %v", err)
		}
		destUrlLink := UrlLink{
			ParentFormID: remappedFormID,
			UrlLinkID:    remappedUrlLinkID,
			Properties:   *destProperties}
		if err := saveUrlLink(destUrlLink); err != nil {
			return fmt.Errorf("CloneUrlLink: %v", err)
		}
	}

	return nil
}

func updateExistingUrlLink(urlLinkID string, updatedUrlLink *UrlLink) (*UrlLink, error) {

	if updateErr := common.UpdateFormComponent(urlLinkEntityKind, updatedUrlLink.ParentFormID,
		updatedUrlLink.UrlLinkID, updatedUrlLink.Properties); updateErr != nil {
		return nil, fmt.Errorf("updateExistingUrlLink: error updating existing text box component: %v", updateErr)
	}

	return updatedUrlLink, nil

}
