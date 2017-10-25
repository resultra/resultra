package tag

import (
	"fmt"
	"resultra/datasheet/server/form/components/common"
	"resultra/datasheet/server/trackerDatabase"
)

type ValidationProperties struct {
	ValueRequired bool `json:"valueRequired"`
}

func newDefaultValidationProperties() ValidationProperties {
	return ValidationProperties{ValueRequired: true}
}

type TagProperties struct {
	FieldID      string                                     `json:"fieldID"`
	LabelFormat  common.ComponentLabelFormatProperties      `json:"labelFormat"`
	Permissions  common.ComponentValuePermissionsProperties `json:"permissions"`
	Validation   ValidationProperties                       `json:"validation"`
	HelpPopupMsg string                                     `json:"helpPopupMsg"`
}

func (srcProps TagProperties) Clone(cloneParams *trackerDatabase.CloneDatabaseParams) (*TagProperties, error) {

	destProps := srcProps

	remappedFieldID, err := cloneParams.IDRemapper.GetExistingRemappedID(srcProps.FieldID)
	if err != nil {
		return nil, fmt.Errorf("Clone: %v", err)
	}
	destProps.FieldID = remappedFieldID

	return &destProps, nil
}

func newDefaultTagProperties() TagProperties {
	props := TagProperties{
		LabelFormat:  common.NewDefaultLabelFormatProperties(),
		Permissions:  common.NewDefaultComponentValuePermissionsProperties(),
		Validation:   newDefaultValidationProperties(),
		HelpPopupMsg: ""}
	return props
}
