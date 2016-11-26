package image

import (
	"fmt"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/form/components/common"
	"resultra/datasheet/server/generic/uniqueID"
)

type ImageProperties struct {
	ComponentLink common.ComponentLink           `json:"componentLink"`
	Geometry      componentLayout.LayoutGeometry `json:"geometry"`
}

func (srcProps ImageProperties) Clone(remappedIDs uniqueID.UniqueIDRemapper) (*ImageProperties, error) {

	destProps := srcProps

	destLink, err := srcProps.ComponentLink.Clone(remappedIDs)
	if err != nil {
		return nil, fmt.Errorf("ImageProperties.Clone: %v", err)
	}
	destProps.ComponentLink = *destLink

	return &destProps, nil
}
