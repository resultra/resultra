package attachment

import (
	"fmt"
	"resultra/tracker/server/form/components/common"
	"resultra/tracker/server/trackerDatabase"
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

func (srcProps AttachmentProperties) Clone(cloneParams *trackerDatabase.CloneDatabaseParams) (*AttachmentProperties, error) {

	destProps := srcProps

	remappedFieldID, err := cloneParams.IDRemapper.GetExistingRemappedID(srcProps.FieldID)
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
