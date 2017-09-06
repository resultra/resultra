package urlLink

import (
	"fmt"
	"log"
	"resultra/datasheet/server/displayTable/columns/common"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/uniqueID"
)

const urlLinkEntityKind string = "urlLink"

type UrlLink struct {
	ParentTableID string            `json:"parentTableID"`
	UrlLinkID     string            `json:"urlLinkID"`
	ColType       string            `json:"colType"`
	ColumnID      string            `json:"columnID"`
	Properties    UrlLinkProperties `json:"properties"`
}

type NewUrlLinkParams struct {
	ParentTableID string `json:"parentTableID"`
	FieldID       string `json:"fieldID"`
}

func validUrlLinkFieldType(fieldType string) bool {
	if fieldType == field.FieldTypeURL {
		return true
	} else {
		return false
	}
}

func saveUrlLink(newUrlLink UrlLink) error {
	if saveErr := common.SaveNewTableColumn(urlLinkEntityKind,
		newUrlLink.ParentTableID, newUrlLink.UrlLinkID, newUrlLink.Properties); saveErr != nil {
		return fmt.Errorf("saveUrlLink: Unable to save url link: %v", saveErr)
	}
	return nil

}

func saveNewUrlLink(params NewUrlLinkParams) (*UrlLink, error) {

	if fieldErr := field.ValidateField(params.FieldID, validUrlLinkFieldType); fieldErr != nil {
		return nil, fmt.Errorf("saveNewUrlLink: %v", fieldErr)
	}

	properties := newDefaultUrlLinkProperties()
	properties.FieldID = params.FieldID

	urlLinkID := uniqueID.GenerateSnowflakeID()
	newUrlLink := UrlLink{ParentTableID: params.ParentTableID,
		UrlLinkID:  urlLinkID,
		ColumnID:   urlLinkID,
		Properties: properties,
		ColType:    urlLinkEntityKind}

	if err := saveUrlLink(newUrlLink); err != nil {
		return nil, fmt.Errorf("saveNewUrlLink: Unable to save text box with params=%+v: error = %v", params, err)
	}

	log.Printf("INFO: API: NewLayout: Created new Layout container: %+v", newUrlLink)

	return &newUrlLink, nil

}

func getUrlLink(parentTableID string, urlLinkID string) (*UrlLink, error) {

	urlLinkProps := newDefaultUrlLinkProperties()
	if getErr := common.GetTableColumn(urlLinkEntityKind, parentTableID, urlLinkID, &urlLinkProps); getErr != nil {
		return nil, fmt.Errorf("getCheckBox: Unable to retrieve text box: %v", getErr)
	}

	urlLink := UrlLink{
		ParentTableID: parentTableID,
		UrlLinkID:     urlLinkID,
		ColumnID:      urlLinkID,
		Properties:    urlLinkProps,
		ColType:       urlLinkEntityKind}

	return &urlLink, nil
}

func GetUrlLinks(parentTableID string) ([]UrlLink, error) {

	urlLinks := []UrlLink{}
	addUrlLink := func(urlLinkID string, encodedProps string) error {

		urlLinkProps := newDefaultUrlLinkProperties()
		if decodeErr := generic.DecodeJSONString(encodedProps, &urlLinkProps); decodeErr != nil {
			return fmt.Errorf("GetUrlLinks: can't decode properties: %v", encodedProps)
		}

		currUrlLink := UrlLink{
			ParentTableID: parentTableID,
			UrlLinkID:     urlLinkID,
			ColumnID:      urlLinkID,
			Properties:    urlLinkProps,
			ColType:       urlLinkEntityKind}
		urlLinks = append(urlLinks, currUrlLink)

		return nil
	}
	if getErr := common.GetTableColumns(urlLinkEntityKind, parentTableID, addUrlLink); getErr != nil {
		return nil, fmt.Errorf("GetUrlLinks: Can't get text boxes: %v")
	}

	return urlLinks, nil

}

func CloneUrlLinks(remappedIDs uniqueID.UniqueIDRemapper, parentFormID string) error {

	srcUrlLinkes, err := GetUrlLinks(parentFormID)
	if err != nil {
		return fmt.Errorf("CloneUrlLinkes: %v", err)
	}

	for _, srcUrlLink := range srcUrlLinkes {
		remappedUrlLinkID := remappedIDs.AllocNewOrGetExistingRemappedID(srcUrlLink.UrlLinkID)
		remappedFormID, err := remappedIDs.GetExistingRemappedID(srcUrlLink.ParentTableID)
		if err != nil {
			return fmt.Errorf("CloneUrlLinks: %v", err)
		}
		destProperties, err := srcUrlLink.Properties.Clone(remappedIDs)
		if err != nil {
			return fmt.Errorf("CloneUrlLinks: %v", err)
		}
		destUrlLink := UrlLink{
			ParentTableID: remappedFormID,
			UrlLinkID:     remappedUrlLinkID,
			ColumnID:      remappedUrlLinkID,
			Properties:    *destProperties,
			ColType:       urlLinkEntityKind}
		if err := saveUrlLink(destUrlLink); err != nil {
			return fmt.Errorf("CloneUrlLinks: %v", err)
		}
	}

	return nil
}

func updateExistingUrlLink(urlLinkID string, updatedUrlLink *UrlLink) (*UrlLink, error) {

	if updateErr := common.UpdateTableColumn(urlLinkEntityKind, updatedUrlLink.ParentTableID,
		updatedUrlLink.UrlLinkID, updatedUrlLink.Properties); updateErr != nil {
		return nil, fmt.Errorf("updateExistingUrlLink: error updating existing text box component: %v", updateErr)
	}

	return updatedUrlLink, nil

}
