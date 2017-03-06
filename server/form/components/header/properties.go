package header

import (
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/generic/uniqueID"
)

const headerSizeMedium string = "medium"

type HeaderProperties struct {
	Label      string                         `json:"label"`
	Geometry   componentLayout.LayoutGeometry `json:"geometry"`
	HeaderSize string                         `json:"headerSize"`
	Underlined bool                           `json:"underlined"`
}

func (srcProps HeaderProperties) Clone(remappedIDs uniqueID.UniqueIDRemapper) (*HeaderProperties, error) {

	destProps := srcProps

	return &destProps, nil
}

func newDefaultHeaderProperties() HeaderProperties {
	props := HeaderProperties{
		HeaderSize: headerSizeMedium,
		Underlined: false}
	return props
}
