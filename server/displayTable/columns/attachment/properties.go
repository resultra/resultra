package attachment

import (
	"fmt"
	"resultra/datasheet/server/form/components/common"
	"resultra/datasheet/server/generic/uniqueID"
)

type ValidationProperties struct {
	ValueRequired bool `json:"valueRequired"`
}

func newDefaultValidationProperties() ValidationProperties {
	return ValidationProperties{
		ValueRequired: true}
}

type AttachmentProperties struct {
	FieldID      string                                     `json:"fieldID"`
	LabelFormat  common.ComponentLabelFormatProperties      `json:"labelFormat"`
	Permissions  common.ComponentValuePermissionsProperties `json:"permissions"`
	Validation   ValidationProperties                       `json:"validation"`
	HelpPopupMsg string                                     `json:"helpPopupMsg"`
}

func (srcProps AttachmentProperties) Clone(remappedIDs uniqueID.UniqueIDRemapper) (*AttachmentProperties, error) {

	destProps := srcProps

	remappedFieldID, err := remappedIDs.GetExistingRemappedID(srcProps.FieldID)
	if err != nil {
		return nil, fmt.Errorf("Clone: %v", err)
	}
	destProps.FieldID = remappedFieldID

	return &destProps, nil
}

func newDefaultAttachmentProperties() AttachmentProperties {
	props := AttachmentProperties{
		LabelFormat:  common.NewDefaultLabelFormatProperties(),
		Permissions:  common.NewDefaultComponentValuePermissionsProperties(),
		Validation:   newDefaultValidationProperties(),
		HelpPopupMsg: ""}
	return props
}
