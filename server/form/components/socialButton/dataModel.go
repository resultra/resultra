package socialButton

import (
	"fmt"
	"log"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/form/components/common"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/uniqueID"
)

const socialButtonEntityKind string = "socialButton"

type SocialButton struct {
	ParentFormID   string                 `json:"parentFormID"`
	SocialButtonID string                 `json:"socialButtonID"`
	Properties     SocialButtonProperties `json:"properties"`
}

type NewSocialButtonParams struct {
	ParentFormID string                         `json:"parentFormID"`
	FieldID      string                         `json:"fieldID"`
	Geometry     componentLayout.LayoutGeometry `json:"geometry"`
}

func validateSocialButtonFieldType(fieldType string) bool {
	if fieldType == field.FieldTypeUser {
		return true
	} else {
		return false
	}
}

func saveSocialButton(newSocialButton SocialButton) error {

	if saveErr := common.SaveNewFormComponent(socialButtonEntityKind,
		newSocialButton.ParentFormID, newSocialButton.SocialButtonID, newSocialButton.Properties); saveErr != nil {
		return fmt.Errorf("saveSocialButton: Unable to save socialButton: error = %v", saveErr)
	}

	return nil
}

func saveNewSocialButton(params NewSocialButtonParams) (*SocialButton, error) {

	if !componentLayout.ValidGeometry(params.Geometry) {
		return nil, fmt.Errorf("Invalid layout container parameters: %+v", params)
	}

	if fieldErr := field.ValidateField(params.FieldID, validateSocialButtonFieldType); fieldErr != nil {
		return nil, fmt.Errorf("saveNewSocialButton: %v", fieldErr)
	}

	properties := newDefaultSocialButtonProperties()
	properties.FieldID = params.FieldID
	properties.Geometry = params.Geometry

	newSocialButton := SocialButton{ParentFormID: params.ParentFormID,
		SocialButtonID: uniqueID.GenerateSnowflakeID(),
		Properties:     properties}

	if saveErr := saveSocialButton(newSocialButton); saveErr != nil {
		return nil, fmt.Errorf("saveNewSocialButton: Unable to save socialButton with params=%+v: error = %v", params, saveErr)
	}

	log.Printf("INFO: API: New SocialButton: Created new socialButton component:  %+v", newSocialButton)

	return &newSocialButton, nil

}

func getSocialButton(parentFormID string, socialButtonID string) (*SocialButton, error) {

	socialButtonProps := newDefaultSocialButtonProperties()
	if getErr := common.GetFormComponent(socialButtonEntityKind, parentFormID, socialButtonID, &socialButtonProps); getErr != nil {
		return nil, fmt.Errorf("getSocialButton: Unable to retrieve socialButton: %v", getErr)
	}

	socialButton := SocialButton{
		ParentFormID:   parentFormID,
		SocialButtonID: socialButtonID,
		Properties:     socialButtonProps}

	return &socialButton, nil
}

func GetSocialButtons(parentFormID string) ([]SocialButton, error) {

	socialButtons := []SocialButton{}
	addSocialButton := func(socialButtonID string, encodedProps string) error {

		socialButtonProps := newDefaultSocialButtonProperties()
		socialButtonProps.Tooltips = []string{} // Default to empty set of tooltips
		if decodeErr := generic.DecodeJSONString(encodedProps, &socialButtonProps); decodeErr != nil {
			return fmt.Errorf("GetSocialButtons: can't decode properties: %v", encodedProps)
		}

		currSocialButton := SocialButton{
			ParentFormID:   parentFormID,
			SocialButtonID: socialButtonID,
			Properties:     socialButtonProps}
		socialButtons = append(socialButtons, currSocialButton)

		return nil
	}
	if getErr := common.GetFormComponents(socialButtonEntityKind, parentFormID, addSocialButton); getErr != nil {
		return nil, fmt.Errorf("GetSocialButtons: Can't get socialButtons: %v")
	}

	return socialButtons, nil
}

func CloneSocialButtons(remappedIDs uniqueID.UniqueIDRemapper, parentFormID string) error {

	srcSocialButtons, err := GetSocialButtons(parentFormID)
	if err != nil {
		return fmt.Errorf("CloneSocialButtons: %v", err)
	}

	for _, srcSocialButton := range srcSocialButtons {
		remappedSocialButtonID := remappedIDs.AllocNewOrGetExistingRemappedID(srcSocialButton.SocialButtonID)
		remappedFormID, err := remappedIDs.GetExistingRemappedID(srcSocialButton.ParentFormID)
		if err != nil {
			return fmt.Errorf("CloneSocialButtons: %v", err)
		}
		destProperties, err := srcSocialButton.Properties.Clone(remappedIDs)
		if err != nil {
			return fmt.Errorf("CloneSocialButtons: %v", err)
		}
		destSocialButton := SocialButton{
			ParentFormID:   remappedFormID,
			SocialButtonID: remappedSocialButtonID,
			Properties:     *destProperties}
		if err := saveSocialButton(destSocialButton); err != nil {
			return fmt.Errorf("CloneSocialButtons: %v", err)
		}
	}

	return nil
}

func updateExistingSocialButton(updatedSocialButton *SocialButton) (*SocialButton, error) {

	if updateErr := common.UpdateFormComponent(socialButtonEntityKind, updatedSocialButton.ParentFormID,
		updatedSocialButton.SocialButtonID, updatedSocialButton.Properties); updateErr != nil {
		return nil, fmt.Errorf("updateExistingSocialButton: failure updating socialButton: %v", updateErr)
	}
	return updatedSocialButton, nil

}