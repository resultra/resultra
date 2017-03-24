package image

import (
	"fmt"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/form/components/common"
	"resultra/datasheet/server/generic/uniqueID"
)

type ImageProperties struct {
	FieldID     string                                     `json:"fieldID"`
	Geometry    componentLayout.LayoutGeometry             `json:"geometry"`
	LabelFormat common.ComponentLabelFormatProperties      `json:"labelFormat"`
	Permissions common.ComponentValuePermissionsProperties `json:"permissions"`
}

func (srcProps ImageProperties) Clone(remappedIDs uniqueID.UniqueIDRemapper) (*ImageProperties, error) {

	destProps := srcProps

	remappedFieldID, err := remappedIDs.GetExistingRemappedID(srcProps.FieldID)
	if err != nil {
		return nil, fmt.Errorf("Clone: %v", err)
	}
	destProps.FieldID = remappedFieldID

	return &destProps, nil
}

func newDefaultAttachmentProperties() ImageProperties {
	props := ImageProperties{
		LabelFormat: common.NewDefaultLabelFormatProperties(),
		Permissions: common.NewDefaultComponentValuePermissionsProperties()}
	return props
}
