package emailAddr

import (
	"database/sql"
	"fmt"
	"log"
	"resultra/tracker/server/displayTable/columns/common"
	"resultra/tracker/server/field"
	"resultra/tracker/server/generic"
	"resultra/tracker/server/generic/uniqueID"
	"resultra/tracker/server/trackerDatabase"
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

func saveEmailAddr(destDBHandle *sql.DB, newEmailAddr EmailAddr) error {
	if saveErr := common.SaveNewTableColumn(destDBHandle, emailAddrEntityKind,
		newEmailAddr.ParentTableID, newEmailAddr.EmailAddrID, newEmailAddr.Properties); saveErr != nil {
		return fmt.Errorf("saveEmailAddr: Unable to save text box: %v", saveErr)
	}
	return nil

}

func saveNewEmailAddr(trackerDBHandle *sql.DB, params NewEmailAddrParams) (*EmailAddr, error) {

	if fieldErr := field.ValidateField(trackerDBHandle, params.FieldID, validEmailAddrFieldType); fieldErr != nil {
		return nil, fmt.Errorf("saveNewEmailAddr: %v", fieldErr)
	}

	properties := newDefaultEmailAddrProperties()
	properties.FieldID = params.FieldID

	emailAddrID := uniqueID.GenerateUniqueID()
	newEmailAddr := EmailAddr{ParentTableID: params.ParentTableID,
		EmailAddrID: emailAddrID,
		ColumnID:    emailAddrID,
		Properties:  properties,
		ColType:     emailAddrEntityKind}

	if err := saveEmailAddr(trackerDBHandle, newEmailAddr); err != nil {
		return nil, fmt.Errorf("saveNewEmailAddr: Unable to save text box with params=%+v: error = %v", params, err)
	}

	log.Printf("INFO: API: NewLayout: Created new Layout container: %+v", newEmailAddr)

	return &newEmailAddr, nil

}

func getEmailAddr(trackerDBHandle *sql.DB, parentTableID string, emailAddrID string) (*EmailAddr, error) {

	emailAddrProps := newDefaultEmailAddrProperties()
	if getErr := common.GetTableColumn(trackerDBHandle, emailAddrEntityKind, parentTableID, emailAddrID, &emailAddrProps); getErr != nil {
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

func getEmailAddrsFromSrc(srcDBHandle *sql.DB, parentTableID string) ([]EmailAddr, error) {

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
	if getErr := common.GetTableColumns(srcDBHandle, emailAddrEntityKind, parentTableID, addEmailAddr); getErr != nil {
		return nil, fmt.Errorf("GetEmailAddrs: Can't get text boxes: %v")
	}

	return emailAddrs, nil

}

func GetEmailAddrs(trackerDBHandle *sql.DB, parentTableID string) ([]EmailAddr, error) {
	return getEmailAddrsFromSrc(trackerDBHandle, parentTableID)
}

func CloneEmailAddrs(cloneParams *trackerDatabase.CloneDatabaseParams, parentFormID string) error {

	srcEmailAddres, err := getEmailAddrsFromSrc(cloneParams.SrcDBHandle, parentFormID)
	if err != nil {
		return fmt.Errorf("CloneEmailAddres: %v", err)
	}

	for _, srcEmailAddr := range srcEmailAddres {
		remappedEmailAddrID := cloneParams.IDRemapper.AllocNewOrGetExistingRemappedID(srcEmailAddr.EmailAddrID)
		remappedFormID, err := cloneParams.IDRemapper.GetExistingRemappedID(srcEmailAddr.ParentTableID)
		if err != nil {
			return fmt.Errorf("CloneEmailAddrs: %v", err)
		}
		destProperties, err := srcEmailAddr.Properties.Clone(cloneParams)
		if err != nil {
			return fmt.Errorf("CloneEmailAddrs: %v", err)
		}
		destEmailAddr := EmailAddr{
			ParentTableID: remappedFormID,
			EmailAddrID:   remappedEmailAddrID,
			ColumnID:      remappedEmailAddrID,
			Properties:    *destProperties,
			ColType:       emailAddrEntityKind}
		if err := saveEmailAddr(cloneParams.DestDBHandle, destEmailAddr); err != nil {
			return fmt.Errorf("CloneEmailAddrs: %v", err)
		}
	}

	return nil
}

func updateExistingEmailAddr(trackerDBHandle *sql.DB, emailAddrID string, updatedEmailAddr *EmailAddr) (*EmailAddr, error) {

	if updateErr := common.UpdateTableColumn(trackerDBHandle, emailAddrEntityKind, updatedEmailAddr.ParentTableID,
		updatedEmailAddr.EmailAddrID, updatedEmailAddr.Properties); updateErr != nil {
		return nil, fmt.Errorf("updateExistingEmailAddr: error updating existing text box component: %v", updateErr)
	}

	return updatedEmailAddr, nil

}
