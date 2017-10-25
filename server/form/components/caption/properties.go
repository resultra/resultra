package caption

import (
	"fmt"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/form/components/common"
	"resultra/datasheet/server/trackerDatabase"
)

const colorSchemeDefault string = "default"

type CaptionProperties struct {
	common.ComponentVisibilityProperties
	Label       string                         `json:"label"`
	Caption     string                         `json:"caption"`
	Geometry    componentLayout.LayoutGeometry `json:"geometry"`
	ColorScheme string                         `json:"colorScheme"`
}

func (srcProps CaptionProperties) Clone(cloneParams *trackerDatabase.CloneDatabaseParams) (*CaptionProperties, error) {

	destProps := srcProps

	destVisibilityConditions, err := srcProps.VisibilityConditions.Clone(cloneParams)
	if err != nil {
		return nil, fmt.Errorf("CaptionProperties.Clone: %v")
	}
	destProps.VisibilityConditions = *destVisibilityConditions

	return &destProps, nil
}

func newDefaultCaptionProperties() CaptionProperties {
	props := CaptionProperties{
		ComponentVisibilityProperties: common.NewDefaultComponentVisibilityProperties(),
		ColorScheme:                   colorSchemeDefault}
	return props
}
