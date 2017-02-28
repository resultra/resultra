package caption

import (
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/generic/uniqueID"
)

type CaptionProperties struct {
	Label    string                         `json:"label"`
	Geometry componentLayout.LayoutGeometry `json:"geometry"`
}

func (srcProps CaptionProperties) Clone(remappedIDs uniqueID.UniqueIDRemapper) (*CaptionProperties, error) {

	destProps := srcProps

	return &destProps, nil
}
