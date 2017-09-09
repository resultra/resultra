package file

import (
	"fmt"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/form/components/common"
	"resultra/datasheet/server/generic/uniqueID"
)

type FileValidationProperties struct {
	ValueRequired bool `json:"valueRequired"`
}

func newDefaultFileValidationProperties() FileValidationProperties {
	return FileValidationProperties{false}
}

type FileProperties struct {
	FieldID     string                                     `json:"fieldID"`
	Geometry    componentLayout.LayoutGeometry             `json:"geometry"`
	LabelFormat common.ComponentLabelFormatProperties      `json:"labelFormat"`
	Permissions common.ComponentValuePermissionsProperties `json:"permissions"`
	common.ComponentVisibilityProperties
	Validation          FileValidationProperties `json:"validation"`
	ClearValueSupported bool                          `json:"clearValueSupported"`
	HelpPopupMsg        string                        `json:"helpPopupMsg"`
}

func (srcProps FileProperties) Clone(remappedIDs uniqueID.UniqueIDRemapper) (*FileProperties, error) {

	destProps := srcProps

	remappedFieldID, err := remappedIDs.GetExistingRemappedID(srcProps.FieldID)
	if err != nil {
		return nil, fmt.Errorf("Clone: %v", err)
	}
	destProps.FieldID = remappedFieldID

	destVisibilityConditions, err := srcProps.VisibilityConditions.Clone(remappedIDs)
	if err != nil {
		return nil, fmt.Errorf("CaptionProperties.Clone: %v")
	}
	destProps.VisibilityConditions = *destVisibilityConditions

	return &destProps, nil
}

func newDefaultFileProperties() FileProperties {
	props := FileProperties{
		ComponentVisibilityProperties: common.NewDefaultComponentVisibilityProperties(),
		LabelFormat:                   common.NewDefaultLabelFormatProperties(),
		Permissions:                   common.NewDefaultComponentValuePermissionsProperties(),
		Validation:                    newDefaultFileValidationProperties(),
		ClearValueSupported:           false,
		HelpPopupMsg:                  ""}
	return props
}
