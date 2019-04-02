// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package urlLink

import (
	"database/sql"
	"fmt"
	"log"
	"github.com/resultra/resultra/server/common/componentLayout"
	"github.com/resultra/resultra/server/field"
	"github.com/resultra/resultra/server/form/components/common"
	"github.com/resultra/resultra/server/generic"
	"github.com/resultra/resultra/server/generic/uniqueID"
	"github.com/resultra/resultra/server/trackerDatabase"
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

func saveUrlLink(destDBHandle *sql.DB, newUrlLink UrlLink) error {
	if saveErr := common.SaveNewFormComponent(destDBHandle, urlLinkEntityKind,
		newUrlLink.ParentFormID, newUrlLink.UrlLinkID, newUrlLink.Properties); saveErr != nil {
		return fmt.Errorf("saveUrlLink: Unable to save text box: %v", saveErr)
	}
	return nil

}

func saveNewUrlLink(trackerDBHandle *sql.DB, params NewUrlLinkParams) (*UrlLink, error) {

	if !componentLayout.ValidGeometry(params.Geometry) {
		return nil, fmt.Errorf("Invalid layout container parameters: %+v", params)
	}

	if fieldErr := field.ValidateField(trackerDBHandle, params.FieldID, validUrlLinkFieldType); fieldErr != nil {
		return nil, fmt.Errorf("saveNewUrlLink: %v", fieldErr)
	}

	properties := newDefaultUrlLinkProperties()
	properties.Geometry = params.Geometry
	properties.FieldID = params.FieldID

	newUrlLink := UrlLink{ParentFormID: params.ParentFormID,
		UrlLinkID:  uniqueID.GenerateUniqueID(),
		Properties: properties}

	if err := saveUrlLink(trackerDBHandle, newUrlLink); err != nil {
		return nil, fmt.Errorf("saveNewUrlLink: Unable to save text box with params=%+v: error = %v", params, err)
	}

	log.Printf("INFO: API: NewLayout: Created new Layout container: %+v", newUrlLink)

	return &newUrlLink, nil

}

func getUrlLink(trackerDBHandle *sql.DB, parentFormID string, urlLinkID string) (*UrlLink, error) {

	urlLinkProps := newDefaultUrlLinkProperties()
	if getErr := common.GetFormComponent(trackerDBHandle, urlLinkEntityKind, parentFormID, urlLinkID, &urlLinkProps); getErr != nil {
		return nil, fmt.Errorf("getCheckBox: Unable to retrieve text box: %v", getErr)
	}

	urlLink := UrlLink{
		ParentFormID: parentFormID,
		UrlLinkID:    urlLinkID,
		Properties:   urlLinkProps}

	return &urlLink, nil
}

func getUrlLinksFromSrc(srcDBHandle *sql.DB, parentFormID string) ([]UrlLink, error) {

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
	if getErr := common.GetFormComponents(srcDBHandle, urlLinkEntityKind, parentFormID, addUrlLink); getErr != nil {
		return nil, fmt.Errorf("GetCheckBoxes: Can't get text boxes: %v")
	}

	return urlLinkes, nil

}

func GetUrlLinks(trackerDBHandle *sql.DB, parentFormID string) ([]UrlLink, error) {
	return getUrlLinksFromSrc(trackerDBHandle, parentFormID)
}

func CloneUrlLinks(cloneParams *trackerDatabase.CloneDatabaseParams, parentFormID string) error {

	srcUrlLink, err := getUrlLinksFromSrc(cloneParams.SrcDBHandle, parentFormID)
	if err != nil {
		return fmt.Errorf("CloneUrlLink: %v", err)
	}

	for _, srcUrlLink := range srcUrlLink {
		remappedUrlLinkID := cloneParams.IDRemapper.AllocNewOrGetExistingRemappedID(srcUrlLink.UrlLinkID)
		remappedFormID, err := cloneParams.IDRemapper.GetExistingRemappedID(srcUrlLink.ParentFormID)
		if err != nil {
			return fmt.Errorf("CloneUrlLink: %v", err)
		}
		destProperties, err := srcUrlLink.Properties.Clone(cloneParams)
		if err != nil {
			return fmt.Errorf("CloneUrlLink: %v", err)
		}
		destUrlLink := UrlLink{
			ParentFormID: remappedFormID,
			UrlLinkID:    remappedUrlLinkID,
			Properties:   *destProperties}
		if err := saveUrlLink(cloneParams.DestDBHandle, destUrlLink); err != nil {
			return fmt.Errorf("CloneUrlLink: %v", err)
		}
	}

	return nil
}

func updateExistingUrlLink(trackerDBHandle *sql.DB, urlLinkID string, updatedUrlLink *UrlLink) (*UrlLink, error) {

	if updateErr := common.UpdateFormComponent(trackerDBHandle, urlLinkEntityKind, updatedUrlLink.ParentFormID,
		updatedUrlLink.UrlLinkID, updatedUrlLink.Properties); updateErr != nil {
		return nil, fmt.Errorf("updateExistingUrlLink: error updating existing text box component: %v", updateErr)
	}

	return updatedUrlLink, nil

}
