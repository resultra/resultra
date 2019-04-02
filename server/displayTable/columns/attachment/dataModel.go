// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package attachment

import (
	"database/sql"
	"fmt"
	"log"
	"github.com/resultra/resultra/server/displayTable/columns/common"
	"github.com/resultra/resultra/server/field"
	"github.com/resultra/resultra/server/generic"
	"github.com/resultra/resultra/server/generic/uniqueID"
	"github.com/resultra/resultra/server/trackerDatabase"
)

const attachmentEntityKind string = "attachment"

type Attachment struct {
	ParentTableID string               `json:"parentTableID"`
	AttachmentID  string               `json:"attachmentID"`
	ColumnID      string               `json:"columnID"`
	ColType       string               `json:"colType"`
	Properties    AttachmentProperties `json:"properties"`
}

type NewAttachmentParams struct {
	ParentTableID string `json:"parentTableID"`
	FieldID       string `json:"fieldID"`
}

func validAttachmentFieldType(fieldType string) bool {
	if fieldType == field.FieldTypeAttachment {
		return true
	} else {
		return false
	}
}

func saveAttachment(destDBHandle *sql.DB, newAttachment Attachment) error {

	if saveErr := common.SaveNewTableColumn(destDBHandle, attachmentEntityKind,
		newAttachment.ParentTableID, newAttachment.AttachmentID, newAttachment.Properties); saveErr != nil {
		return fmt.Errorf("saveNewAttachment: Unable to save image form component: error = %v", saveErr)
	}
	return nil

}

func saveNewAttachment(trackerDBHandle *sql.DB, params NewAttachmentParams) (*Attachment, error) {

	if fieldErr := field.ValidateField(trackerDBHandle, params.FieldID, validAttachmentFieldType); fieldErr != nil {
		return nil, fmt.Errorf("saveNewTextBox: %v", fieldErr)
	}

	properties := newDefaultAttachmentProperties()
	properties.FieldID = params.FieldID

	attachmentID := uniqueID.GenerateUniqueID()
	newAttachment := Attachment{ParentTableID: params.ParentTableID,
		AttachmentID: attachmentID,
		ColumnID:     attachmentID,
		ColType:      attachmentEntityKind,
		Properties:   properties}

	if saveErr := saveAttachment(trackerDBHandle, newAttachment); saveErr != nil {
		return nil, fmt.Errorf("saveNewAttachment: Unable to save image form component with params=%+v: error = %v", params, saveErr)
	}

	log.Printf("INFO: API: saveNewAttachment: Created new image component: %+v", newAttachment)

	return &newAttachment, nil

}

func getAttachment(trackerDBHandle *sql.DB, parentTableID string, attachmentID string) (*Attachment, error) {

	props := newDefaultAttachmentProperties()
	if getErr := common.GetTableColumn(trackerDBHandle, attachmentEntityKind, parentTableID, attachmentID, &props); getErr != nil {
		return nil, fmt.Errorf("getAttachment: Unable to retrieve image form component: %v", getErr)
	}

	attachment := Attachment{
		ParentTableID: parentTableID,
		AttachmentID:  attachmentID,
		ColumnID:      attachmentID,
		ColType:       attachmentEntityKind,
		Properties:    props}

	return &attachment, nil
}

func getAttachmentsFromSrc(srcDBHandle *sql.DB, parentTableID string) ([]Attachment, error) {

	attachments := []Attachment{}
	addAttachment := func(attachmentID string, encodedProps string) error {

		props := newDefaultAttachmentProperties()
		if decodeErr := generic.DecodeJSONString(encodedProps, &props); decodeErr != nil {
			return fmt.Errorf("GetAttachments: can't decode properties: %v", encodedProps)
		}

		currAttachment := Attachment{
			ParentTableID: parentTableID,
			AttachmentID:  attachmentID,
			ColumnID:      attachmentID,
			ColType:       attachmentEntityKind,
			Properties:    props}
		attachments = append(attachments, currAttachment)

		return nil
	}
	if getErr := common.GetTableColumns(srcDBHandle, attachmentEntityKind, parentTableID, addAttachment); getErr != nil {
		return nil, fmt.Errorf("GetCheckBoxes: Can't get image form components: %v")
	}

	return attachments, nil

}

func GetAttachments(trackerDBHandle *sql.DB, parentTableID string) ([]Attachment, error) {
	return getAttachmentsFromSrc(trackerDBHandle, parentTableID)
}

func CloneAttachments(cloneParams *trackerDatabase.CloneDatabaseParams, parentTableID string) error {

	srcAttachments, err := getAttachmentsFromSrc(cloneParams.SrcDBHandle, parentTableID)
	if err != nil {
		return fmt.Errorf("CloneAttachments: %v", err)
	}

	for _, srcAttachment := range srcAttachments {
		remappedAttachmentID := cloneParams.IDRemapper.AllocNewOrGetExistingRemappedID(srcAttachment.AttachmentID)
		remappedFormID, err := cloneParams.IDRemapper.GetExistingRemappedID(srcAttachment.ParentTableID)
		if err != nil {
			return fmt.Errorf("CloneAttachments: %v", err)
		}
		destProperties, err := srcAttachment.Properties.Clone(cloneParams)
		if err != nil {
			return fmt.Errorf("CloneAttachments: %v", err)
		}
		destAttachment := Attachment{
			ParentTableID: remappedFormID,
			AttachmentID:  remappedAttachmentID,
			ColumnID:      remappedAttachmentID,
			ColType:       attachmentEntityKind,
			Properties:    *destProperties}
		if err := saveAttachment(cloneParams.DestDBHandle, destAttachment); err != nil {
			return fmt.Errorf("CloneAttachments: %v", err)
		}
	}

	return nil
}

func updateExistingAttachment(trackerDBHandle *sql.DB, attachmentID string, updatedAttachment *Attachment) (*Attachment, error) {

	if updateErr := common.UpdateTableColumn(trackerDBHandle, attachmentEntityKind, updatedAttachment.ParentTableID,
		updatedAttachment.AttachmentID, updatedAttachment.Properties); updateErr != nil {
		return nil, fmt.Errorf("updateExistingAttachment: error updating existing image component: %v", updateErr)
	}

	return updatedAttachment, nil

}
