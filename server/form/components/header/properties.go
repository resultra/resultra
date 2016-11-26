package header

import (
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/generic/uniqueID"
)

type HeaderProperties struct {
	Label    string                         `json:"label"`
	Geometry componentLayout.LayoutGeometry `json:"geometry"`
}

func (srcProps HeaderProperties) Clone(remappedIDs uniqueID.UniqueIDRemapper) (*HeaderProperties, error) {

	destProps := srcProps

	return &destProps, nil
}
