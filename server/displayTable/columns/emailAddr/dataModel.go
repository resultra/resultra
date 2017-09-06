package emailAddr

import (
	"fmt"
	"log"
	"resultra/datasheet/server/displayTable/columns/common"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/uniqueID"
)

const emailAddrEntityKind string = "emailAddr"

type EmailAddr struct {
	ParentTableID string              `json:"parentTableID"`
	EmailAddrID   string              `json:"emailAddrID"`
	ColType       string              `json:"colType"`
	ColumnID      string              `json:"columnID"`
	Properties    EmailAddrProperties `json:"properties"`
}

type NewEmailAddrParams struct {
	ParentTableID string `json:"parentTableID"`
	FieldID       string `json:"fieldID"`
}

func validEmailAddrFieldType(fieldType string) bool {
	if fieldType == field.FieldTypeEmail {
		return true
	} else {
		return false
	}
}

func saveEmailAddr(newEmailAddr EmailAddr) error {
	if saveErr := common.SaveNewTableColumn(emailAddrEntityKind,
		newEmailAddr.ParentTableID, newEmailAddr.EmailAddrID, newEmailAddr.Properties); saveErr != nil {
		return fmt.Errorf("saveEmailAddr: Unable to save text box: %v", saveErr)
	}
	return nil

}

func saveNewEmailAddr(params NewEmailAddrParams) (*EmailAddr, error) {

	if fieldErr := field.ValidateField(params.FieldID, validEmailAddrFieldType); fieldErr != nil {
		return nil, fmt.Errorf("saveNewEmailAddr: %v", fieldErr)
	}

	properties := newDefaultEmailAddrProperties()
	properties.FieldID = params.FieldID

	emailAddrID := uniqueID.GenerateSnowflakeID()
	newEmailAddr := EmailAddr{ParentTableID: params.ParentTableID,
		EmailAddrID: emailAddrID,
		ColumnID:    emailAddrID,
		Properties:  properties,
		ColType:     emailAddrEntityKind}

	if err := saveEmailAddr(newEmailAddr); err != nil {
		return nil, fmt.Errorf("saveNewEmailAddr: Unable to save text box with params=%+v: error = %v", params, err)
	}

	log.Printf("INFO: API: NewLayout: Created new Layout container: %+v", newEmailAddr)

	return &newEmailAddr, nil

}

func getEmailAddr(parentTableID string, emailAddrID string) (*EmailAddr, error) {

	emailAddrProps := newDefaultEmailAddrProperties()
	if getErr := common.GetTableColumn(emailAddrEntityKind, parentTableID, emailAddrID, &emailAddrProps); getErr != nil {
		return nil, fmt.Errorf("getCheckBox: Unable to retrieve text box: %v", getErr)
	}

	emailAddr := EmailAddr{
		ParentTableID: parentTableID,
		EmailAddrID:   emailAddrID,
		ColumnID:      emailAddrID,
		Properties:    emailAddrProps,
		ColType:       emailAddrEntityKind}

	return &emailAddr, nil
}

func GetEmailAddrs(parentTableID string) ([]EmailAddr, error) {

	emailAddrs := []EmailAddr{}
	addEmailAddr := func(emailAddrID string, encodedProps string) error {

		emailAddrProps := newDefaultEmailAddrProperties()
		if decodeErr := generic.DecodeJSONString(encodedProps, &emailAddrProps); decodeErr != nil {
			return fmt.Errorf("GetEmailAddrs: can't decode properties: %v", encodedProps)
		}

		currEmailAddr := EmailAddr{
			ParentTableID: parentTableID,
			EmailAddrID:   emailAddrID,
			ColumnID:      emailAddrID,
			Properties:    emailAddrProps,
			ColType:       emailAddrEntityKind}
		emailAddrs = append(emailAddrs, currEmailAddr)

		return nil
	}
	if getErr := common.GetTableColumns(emailAddrEntityKind, parentTableID, addEmailAddr); getErr != nil {
		return nil, fmt.Errorf("GetEmailAddrs: Can't get text boxes: %v")
	}

	return emailAddrs, nil

}

func CloneEmailAddrs(remappedIDs uniqueID.UniqueIDRemapper, parentFormID string) error {

	srcEmailAddres, err := GetEmailAddrs(parentFormID)
	if err != nil {
		return fmt.Errorf("CloneEmailAddres: %v", err)
	}

	for _, srcEmailAddr := range srcEmailAddres {
		remappedEmailAddrID := remappedIDs.AllocNewOrGetExistingRemappedID(srcEmailAddr.EmailAddrID)
		remappedFormID, err := remappedIDs.GetExistingRemappedID(srcEmailAddr.ParentTableID)
		if err != nil {
			return fmt.Errorf("CloneEmailAddrs: %v", err)
		}
		destProperties, err := srcEmailAddr.Properties.Clone(remappedIDs)
		if err != nil {
			return fmt.Errorf("CloneEmailAddrs: %v", err)
		}
		destEmailAddr := EmailAddr{
			ParentTableID: remappedFormID,
			EmailAddrID:   remappedEmailAddrID,
			ColumnID:      remappedEmailAddrID,
			Properties:    *destProperties,
			ColType:       emailAddrEntityKind}
		if err := saveEmailAddr(destEmailAddr); err != nil {
			return fmt.Errorf("CloneEmailAddrs: %v", err)
		}
	}

	return nil
}

func updateExistingEmailAddr(emailAddrID string, updatedEmailAddr *EmailAddr) (*EmailAddr, error) {

	if updateErr := common.UpdateTableColumn(emailAddrEntityKind, updatedEmailAddr.ParentTableID,
		updatedEmailAddr.EmailAddrID, updatedEmailAddr.Properties); updateErr != nil {
		return nil, fmt.Errorf("updateExistingEmailAddr: error updating existing text box component: %v", updateErr)
	}

	return updatedEmailAddr, nil

}
