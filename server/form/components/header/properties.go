package header

import (
	"fmt"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/form/components/common"
	"resultra/datasheet/server/generic/uniqueID"
	"resultra/datasheet/server/recordFilter"
)

const headerSizeMedium string = "medium"

type HeaderProperties struct {
	Label      string                         `json:"label"`
	Geometry   componentLayout.LayoutGeometry `json:"geometry"`
	HeaderSize string                         `json:"headerSize"`
	Underlined bool                           `json:"underlined"`
	common.ComponentVisibilityProperties
}

func (srcProps HeaderProperties) Clone(remappedIDs uniqueID.UniqueIDRemapper) (*HeaderProperties, error) {

	destProps := srcProps

	destVisibilityConditions, err := recordFilter.CloneFilterRules(remappedIDs, srcProps.VisibilityConditions)
	if err != nil {
		return nil, fmt.Errorf("CaptionProperties.Clone: %v")
	}
	destProps.VisibilityConditions = destVisibilityConditions

	return &destProps, nil
}

func newDefaultHeaderProperties() HeaderProperties {
	props := HeaderProperties{
		ComponentVisibilityProperties: common.NewDefaultComponentVisibilityProperties(),
		HeaderSize:                    headerSizeMedium,
		Underlined:                    false}
	return props
}
