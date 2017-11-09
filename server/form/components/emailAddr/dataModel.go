package emailAddr

import (
	"database/sql"
	"fmt"
	"log"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/form/components/common"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/uniqueID"
	"resultra/datasheet/server/trackerDatabase"
)

const emailAddrEntityKind string = "emailAddr"

type EmailAddr struct {
	ParentFormID string              `json:"parentFormID"`
	EmailAddrID  string              `json:"emailAddrID"`
	Properties   EmailAddrProperties `json:"properties"`
}

type NewEmailAddrParams struct {
	ParentFormID string                         `json:"parentFormID"`
	FieldID      string                         `json:"fieldID"`
	Geometry     componentLayout.LayoutGeometry `json:"geometry"`
}

func validEmailAddrFieldType(fieldType string) bool {
	if fieldType == field.FieldTypeEmail {
		return true
	} else {
		return false
	}
}

func saveEmailAddr(destDBHandle *sql.DB, newEmailAddr EmailAddr) error {
	if saveErr := common.SaveNewFormComponent(destDBHandle, emailAddrEntityKind,
		newEmailAddr.ParentFormID, newEmailAddr.EmailAddrID, newEmailAddr.Properties); saveErr != nil {
		return fmt.Errorf("saveEmailAddr: Unable to save text box: %v", saveErr)
	}
	return nil

}

func saveNewEmailAddr(trackerDBHandle *sql.DB, params NewEmailAddrParams) (*EmailAddr, error) {

	if !componentLayout.ValidGeometry(params.Geometry) {
		return nil, fmt.Errorf("Invalid layout container parameters: %+v", params)
	}

	if fieldErr := field.ValidateField(trackerDBHandle, params.FieldID, validEmailAddrFieldType); fieldErr != nil {
		return nil, fmt.Errorf("saveNewEmailAddr: %v", fieldErr)
	}

	properties := newDefaultEmailAddrProperties()
	properties.Geometry = params.Geometry
	properties.FieldID = params.FieldID

	newEmailAddr := EmailAddr{ParentFormID: params.ParentFormID,
		EmailAddrID: uniqueID.GenerateSnowflakeID(),
		Properties:  properties}

	if err := saveEmailAddr(trackerDBHandle, newEmailAddr); err != nil {
		return nil, fmt.Errorf("saveNewEmailAddr: Unable to save text box with params=%+v: error = %v", params, err)
	}

	log.Printf("INFO: API: NewLayout: Created new Layout container: %+v", newEmailAddr)

	return &newEmailAddr, nil

}

func getEmailAddr(trackerDBHandle *sql.DB, parentFormID string, emailAddrID string) (*EmailAddr, error) {

	emailAddrProps := newDefaultEmailAddrProperties()
	if getErr := common.GetFormComponent(trackerDBHandle, emailAddrEntityKind,
		parentFormID, emailAddrID, &emailAddrProps); getErr != nil {
		return nil, fmt.Errorf("getCheckBox: Unable to retrieve text box: %v", getErr)
	}

	emailAddr := EmailAddr{
		ParentFormID: parentFormID,
		EmailAddrID:  emailAddrID,
		Properties:   emailAddrProps}

	return &emailAddr, nil
}

func getEmailAddrsFromSrc(srcDBHandle *sql.DB, parentFormID string) ([]EmailAddr, error) {

	emailAddres := []EmailAddr{}
	addEmailAddr := func(emailAddrID string, encodedProps string) error {

		emailAddrProps := newDefaultEmailAddrProperties()
		if decodeErr := generic.DecodeJSONString(encodedProps, &emailAddrProps); decodeErr != nil {
			return fmt.Errorf("GetEmailAddr: can't decode properties: %v", encodedProps)
		}

		currEmailAddr := EmailAddr{
			ParentFormID: parentFormID,
			EmailAddrID:  emailAddrID,
			Properties:   emailAddrProps}
		emailAddres = append(emailAddres, currEmailAddr)

		return nil
	}
	if getErr := common.GetFormComponents(srcDBHandle, emailAddrEntityKind, parentFormID, addEmailAddr); getErr != nil {
		return nil, fmt.Errorf("GetCheckBoxes: Can't get text boxes: %v")
	}

	return emailAddres, nil

}

func GetEmailAddrs(trackerDBHandle *sql.DB, parentFormID string) ([]EmailAddr, error) {
	return getEmailAddrsFromSrc(trackerDBHandle, parentFormID)
}

func CloneEmailAddrs(cloneParams *trackerDatabase.CloneDatabaseParams, parentFormID string) error {

	srcEmailAddr, err := getEmailAddrsFromSrc(cloneParams.SrcDBHandle, parentFormID)
	if err != nil {
		return fmt.Errorf("CloneEmailAddr: %v", err)
	}

	for _, srcEmailAddr := range srcEmailAddr {
		remappedEmailAddrID := cloneParams.IDRemapper.AllocNewOrGetExistingRemappedID(srcEmailAddr.EmailAddrID)
		remappedFormID, err := cloneParams.IDRemapper.GetExistingRemappedID(srcEmailAddr.ParentFormID)
		if err != nil {
			return fmt.Errorf("CloneEmailAddr: %v", err)
		}
		destProperties, err := srcEmailAddr.Properties.Clone(cloneParams)
		if err != nil {
			return fmt.Errorf("CloneEmailAddr: %v", err)
		}
		destEmailAddr := EmailAddr{
			ParentFormID: remappedFormID,
			EmailAddrID:  remappedEmailAddrID,
			Properties:   *destProperties}
		if err := saveEmailAddr(cloneParams.DestDBHandle, destEmailAddr); err != nil {
			return fmt.Errorf("CloneEmailAddr: %v", err)
		}
	}

	return nil
}

func updateExistingEmailAddr(trackerDBHandle *sql.DB, emailAddrID string, updatedEmailAddr *EmailAddr) (*EmailAddr, error) {

	if updateErr := common.UpdateFormComponent(trackerDBHandle, emailAddrEntityKind, updatedEmailAddr.ParentFormID,
		updatedEmailAddr.EmailAddrID, updatedEmailAddr.Properties); updateErr != nil {
		return nil, fmt.Errorf("updateExistingEmailAddr: error updating existing text box component: %v", updateErr)
	}

	return updatedEmailAddr, nil

}
