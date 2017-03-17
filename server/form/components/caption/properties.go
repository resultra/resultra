package caption

import (
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/form/components/common"
	"resultra/datasheet/server/generic/uniqueID"
)

const colorSchemeDefault string = "default"

type CaptionProperties struct {
	common.ComponentVisibilityProperties
	Label       string                         `json:"label"`
	Caption     string                         `json:"caption"`
	Geometry    componentLayout.LayoutGeometry `json:"geometry"`
	ColorScheme string                         `json:"colorScheme"`
}

func (srcProps CaptionProperties) Clone(remappedIDs uniqueID.UniqueIDRemapper) (*CaptionProperties, error) {

	destProps := srcProps

	return &destProps, nil
}

func newDefaultCaptionProperties() CaptionProperties {
	props := CaptionProperties{
		ComponentVisibilityProperties: common.NewDefaultComponentVisibilityProperties(),
		ColorScheme:                   colorSchemeDefault}
	return props
}
