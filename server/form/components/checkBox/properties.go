package checkBox

import (
	"fmt"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/form/components/common"
	"resultra/datasheet/server/generic/uniqueID"
)

type CheckBoxProperties struct {
	ComponentLink common.ComponentLink           `json:"componentLink"`
	Geometry      componentLayout.LayoutGeometry `json:"geometry"`
}

func (srcProps CheckBoxProperties) Clone(remappedIDs uniqueID.UniqueIDRemapper) (*CheckBoxProperties, error) {

	destProps := srcProps

	destLink, err := srcProps.ComponentLink.Clone(remappedIDs)
	if err != nil {
		return nil, fmt.Errorf("CheckBoxProperties.Clone: %v", err)
	}
	destProps.ComponentLink = *destLink

	return &destProps, nil
}
