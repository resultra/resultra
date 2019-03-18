package socialButton

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

const socialButtonEntityKind string = "socialButton"

type SocialButton struct {
	ParentTableID  string                 `json:"parentTableID"`
	SocialButtonID string                 `json:"socialButtonID"`
	ColumnID       string                 `json:"columnID"`
	ColType        string                 `json:"colType"`
	Properties     SocialButtonProperties `json:"properties"`
}

type NewSocialButtonParams struct {
	ParentTableID string `json:"parentTableID"`
	FieldID       string `json:"fieldID"`
}

func validateSocialButtonFieldType(fieldType string) bool {
	if fieldType == field.FieldTypeUser {
		return true
	} else {
		return false
	}
}

func saveSocialButton(destDBHandle *sql.DB, newSocialButton SocialButton) error {

	if saveErr := common.SaveNewTableColumn(destDBHandle, socialButtonEntityKind,
		newSocialButton.ParentTableID, newSocialButton.SocialButtonID, newSocialButton.Properties); saveErr != nil {
		return fmt.Errorf("saveSocialButton: Unable to save socialButton: error = %v", saveErr)
	}

	return nil
}

func saveNewSocialButton(trackerDBHandle *sql.DB, params NewSocialButtonParams) (*SocialButton, error) {

	if fieldErr := field.ValidateField(trackerDBHandle, params.FieldID, validateSocialButtonFieldType); fieldErr != nil {
		return nil, fmt.Errorf("saveNewSocialButton: %v", fieldErr)
	}

	properties := newDefaultSocialButtonProperties()
	properties.FieldID = params.FieldID

	columnID := uniqueID.GenerateUniqueID()
	newSocialButton := SocialButton{ParentTableID: params.ParentTableID,
		SocialButtonID: columnID,
		ColumnID:       columnID,
		ColType:        socialButtonEntityKind,
		Properties:     properties}

	if saveErr := saveSocialButton(trackerDBHandle, newSocialButton); saveErr != nil {
		return nil, fmt.Errorf("saveNewSocialButton: Unable to save socialButton with params=%+v: error = %v", params, saveErr)
	}

	log.Printf("INFO: API: New SocialButton: Created new socialButton component:  %+v", newSocialButton)

	return &newSocialButton, nil

}

func getSocialButton(trackerDBHandle *sql.DB, parentTableID string, socialButtonID string) (*SocialButton, error) {

	socialButtonProps := newDefaultSocialButtonProperties()
	if getErr := common.GetTableColumn(trackerDBHandle, socialButtonEntityKind, parentTableID, socialButtonID, &socialButtonProps); getErr != nil {
		return nil, fmt.Errorf("getSocialButton: Unable to retrieve socialButton: %v", getErr)
	}

	socialButton := SocialButton{
		ParentTableID:  parentTableID,
		SocialButtonID: socialButtonID,
		ColumnID:       socialButtonID,
		ColType:        socialButtonEntityKind,
		Properties:     socialButtonProps}

	return &socialButton, nil
}

func getSocialButtonsFromSrc(srcDBHandle *sql.DB, parentTableID string) ([]SocialButton, error) {

	socialButtons := []SocialButton{}
	addSocialButton := func(socialButtonID string, encodedProps string) error {

		socialButtonProps := newDefaultSocialButtonProperties()
		if decodeErr := generic.DecodeJSONString(encodedProps, &socialButtonProps); decodeErr != nil {
			return fmt.Errorf("GetSocialButtons: can't decode properties: %v", encodedProps)
		}

		currSocialButton := SocialButton{
			ParentTableID:  parentTableID,
			SocialButtonID: socialButtonID,
			ColumnID:       socialButtonID,
			ColType:        socialButtonEntityKind,
			Properties:     socialButtonProps}
		socialButtons = append(socialButtons, currSocialButton)

		return nil
	}
	if getErr := common.GetTableColumns(srcDBHandle, socialButtonEntityKind, parentTableID, addSocialButton); getErr != nil {
		return nil, fmt.Errorf("GetSocialButtons: Can't get socialButtons: %v")
	}

	return socialButtons, nil
}

func GetSocialButtons(trackerDBHandle *sql.DB, parentTableID string) ([]SocialButton, error) {
	return getSocialButtonsFromSrc(trackerDBHandle, parentTableID)
}

func CloneSocialButtons(cloneParams *trackerDatabase.CloneDatabaseParams, parentTableID string) error {

	srcSocialButtons, err := getSocialButtonsFromSrc(cloneParams.SrcDBHandle, parentTableID)
	if err != nil {
		return fmt.Errorf("CloneSocialButtons: %v", err)
	}

	for _, srcSocialButton := range srcSocialButtons {
		remappedSocialButtonID := cloneParams.IDRemapper.AllocNewOrGetExistingRemappedID(srcSocialButton.SocialButtonID)
		remappedFormID, err := cloneParams.IDRemapper.GetExistingRemappedID(srcSocialButton.ParentTableID)
		if err != nil {
			return fmt.Errorf("CloneSocialButtons: %v", err)
		}
		destProperties, err := srcSocialButton.Properties.Clone(cloneParams)
		if err != nil {
			return fmt.Errorf("CloneSocialButtons: %v", err)
		}
		destSocialButton := SocialButton{
			ParentTableID:  remappedFormID,
			SocialButtonID: remappedSocialButtonID,
			ColumnID:       remappedSocialButtonID,
			ColType:        socialButtonEntityKind,
			Properties:     *destProperties}
		if err := saveSocialButton(cloneParams.DestDBHandle, destSocialButton); err != nil {
			return fmt.Errorf("CloneSocialButtons: %v", err)
		}
	}

	return nil
}

func updateExistingSocialButton(trackerDBHandle *sql.DB, updatedSocialButton *SocialButton) (*SocialButton, error) {

	if updateErr := common.UpdateTableColumn(trackerDBHandle, socialButtonEntityKind, updatedSocialButton.ParentTableID,
		updatedSocialButton.SocialButtonID, updatedSocialButton.Properties); updateErr != nil {
		return nil, fmt.Errorf("updateExistingSocialButton: failure updating socialButton: %v", updateErr)
	}
	return updatedSocialButton, nil

}
