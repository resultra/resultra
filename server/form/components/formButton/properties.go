package formButton

import (
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/generic/uniqueID"
)

type ButtonProperties struct {
	Geometry     componentLayout.LayoutGeometry `json:"geometry"`
	LinkedFormID string                         `json:"linkedFormID"`
}

func (srcProps ButtonProperties) Clone(remappedIDs uniqueID.UniqueIDRemapper) (*ButtonProperties, error) {

	destProps := srcProps

	destProps.LinkedFormID = remappedIDs.AllocNewOrGetExistingRemappedID(srcProps.LinkedFormID)

	return &destProps, nil
}
