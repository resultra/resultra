package formButton

import (
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/generic/uniqueID"
)

type ButtonProperties struct {
	Geometry componentLayout.LayoutGeometry `json:"geometry"`
}

func (srcProps ButtonProperties) Clone(remappedIDs uniqueID.UniqueIDRemapper) (*ButtonProperties, error) {

	destProps := srcProps

	return &destProps, nil
}
