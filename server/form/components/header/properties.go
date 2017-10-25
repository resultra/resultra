package header

import (
	"fmt"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/form/components/common"
	"resultra/datasheet/server/trackerDatabase"
)

const headerSizeMedium string = "medium"

type HeaderProperties struct {
	Label      string                         `json:"label"`
	Geometry   componentLayout.LayoutGeometry `json:"geometry"`
	HeaderSize string                         `json:"headerSize"`
	Underlined bool                           `json:"underlined"`
	common.ComponentVisibilityProperties
}

func (srcProps HeaderProperties) Clone(cloneParams *trackerDatabase.CloneDatabaseParams) (*HeaderProperties, error) {

	destProps := srcProps

	destVisibilityConditions, err := srcProps.VisibilityConditions.Clone(cloneParams)
	if err != nil {
		return nil, fmt.Errorf("CaptionProperties.Clone: %v")
	}
	destProps.VisibilityConditions = *destVisibilityConditions

	return &destProps, nil
}

func newDefaultHeaderProperties() HeaderProperties {
	props := HeaderProperties{
		ComponentVisibilityProperties: common.NewDefaultComponentVisibilityProperties(),
		HeaderSize:                    headerSizeMedium,
		Underlined:                    false}
	return props
}
