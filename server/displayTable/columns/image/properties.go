package image

import (
	"fmt"
	"resultra/tracker/server/form/components/common"
	"resultra/tracker/server/trackerDatabase"
)

type ImageValidationProperties struct {
	ValueRequired bool `json:"valueRequired"`
}

func newDefaultImageValidationProperties() ImageValidationProperties {
	return ImageValidationProperties{false}
}

type ImageProperties struct {
	FieldID             string                                     `json:"fieldID"`
	LabelFormat         common.ComponentLabelFormatProperties      `json:"labelFormat"`
	Permissions         common.ComponentValuePermissionsProperties `json:"permissions"`
	Validation          ImageValidationProperties                  `json:"validation"`
	ClearValueSupported bool                                       `json:"clearValueSupported"`
	HelpPopupMsg        string                                     `json:"helpPopupMsg"`
}

func (srcProps ImageProperties) Clone(cloneParams *trackerDatabase.CloneDatabaseParams) (*ImageProperties, error) {

	destProps := srcProps

	remappedFieldID, err := cloneParams.IDRemapper.GetExistingRemappedID(srcProps.FieldID)
	if err != nil {
		return nil, fmt.Errorf("Clone: %v", err)
	}
	destProps.FieldID = remappedFieldID

	return &destProps, nil
}

func newDefaultImageProperties() ImageProperties {
	props := ImageProperties{
		LabelFormat:         common.NewDefaultLabelFormatProperties(),
		Permissions:         common.NewDefaultComponentValuePermissionsProperties(),
		Validation:          newDefaultImageValidationProperties(),
		ClearValueSupported: false,
		HelpPopupMsg:        ""}
	return props
}
