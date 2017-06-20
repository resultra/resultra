package attachment

import (
	"fmt"
	"log"
	"resultra/datasheet/server/displayTable/columns/common"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/uniqueID"
)

const attachmentEntityKind string = "attachment"

type Attachment struct {
	ParentTableID string               `json:"parentTableID"`
	AttachmentID  string               `json:"imageID"`
	ColumnID      string               `json:"columnID"`
	ColType       string               `json:"colType"`
	Properties    AttachmentProperties `json:"properties"`
}

type NewAttachmentParams struct {
	ParentTableID string `json:"parentTableID"`
	FieldID       string `json:"fieldID"`
}

func validAttachmentFieldType(fieldType string) bool {
	if fieldType == field.FieldTypeFile {
		return true
	} else {
		return false
	}
}

func saveAttachment(newAttachment Attachment) error {

	if saveErr := common.SaveNewTableColumn(attachmentEntityKind,
		newAttachment.ParentTableID, newAttachment.AttachmentID, newAttachment.Properties); saveErr != nil {
		return fmt.Errorf("saveNewAttachment: Unable to save image form component: error = %v", saveErr)
	}
	return nil

}

func saveNewAttachment(params NewAttachmentParams) (*Attachment, error) {

	if fieldErr := field.ValidateField(params.FieldID, validAttachmentFieldType); fieldErr != nil {
		return nil, fmt.Errorf("saveNewTextBox: %v", fieldErr)
	}

	properties := newDefaultAttachmentProperties()
	properties.FieldID = params.FieldID

	attachmentID := uniqueID.GenerateSnowflakeID()
	newAttachment := Attachment{ParentTableID: params.ParentTableID,
		AttachmentID: attachmentID,
		ColumnID:     attachmentID,
		ColType:      attachmentEntityKind,
		Properties:   properties}

	if saveErr := saveAttachment(newAttachment); saveErr != nil {
		return nil, fmt.Errorf("saveNewAttachment: Unable to save image form component with params=%+v: error = %v", params, saveErr)
	}

	log.Printf("INFO: API: saveNewAttachment: Created new image component: %+v", newAttachment)

	return &newAttachment, nil

}

func getAttachment(parentTableID string, attachmentID string) (*Attachment, error) {

	props := newDefaultAttachmentProperties()
	if getErr := common.GetTableColumn(attachmentEntityKind, parentTableID, attachmentID, &props); getErr != nil {
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

func GetAttachments(parentTableID string) ([]Attachment, error) {

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
	if getErr := common.GetTableColumns(attachmentEntityKind, parentTableID, addAttachment); getErr != nil {
		return nil, fmt.Errorf("GetCheckBoxes: Can't get image form components: %v")
	}

	return attachments, nil

}

func CloneAttachments(remappedIDs uniqueID.UniqueIDRemapper, parentTableID string) error {

	srcAttachments, err := GetAttachments(parentTableID)
	if err != nil {
		return fmt.Errorf("CloneAttachments: %v", err)
	}

	for _, srcAttachment := range srcAttachments {
		remappedAttachmentID := remappedIDs.AllocNewOrGetExistingRemappedID(srcAttachment.AttachmentID)
		remappedFormID, err := remappedIDs.GetExistingRemappedID(srcAttachment.ParentTableID)
		if err != nil {
			return fmt.Errorf("CloneAttachments: %v", err)
		}
		destProperties, err := srcAttachment.Properties.Clone(remappedIDs)
		if err != nil {
			return fmt.Errorf("CloneAttachments: %v", err)
		}
		destAttachment := Attachment{
			ParentTableID: remappedFormID,
			AttachmentID:  remappedAttachmentID,
			ColumnID:      remappedAttachmentID,
			ColType:       attachmentEntityKind,
			Properties:    *destProperties}
		if err := saveAttachment(destAttachment); err != nil {
			return fmt.Errorf("CloneAttachments: %v", err)
		}
	}

	return nil
}

func updateExistingAttachment(attachmentID string, updatedAttachment *Attachment) (*Attachment, error) {

	if updateErr := common.UpdateTableColumn(attachmentEntityKind, updatedAttachment.ParentTableID,
		updatedAttachment.AttachmentID, updatedAttachment.Properties); updateErr != nil {
		return nil, fmt.Errorf("updateExistingAttachment: error updating existing image component: %v", updateErr)
	}

	return updatedAttachment, nil

}
