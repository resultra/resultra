package selection

import (
	"fmt"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/form/components/common"
	"resultra/datasheet/server/generic/uniqueID"
)

type SelectionProperties struct {
	ComponentLink  common.ComponentLink           `json:"componentLink"`
	Geometry       componentLayout.LayoutGeometry `json:"geometry"`
	SelectableVals []SelectionSelectableVal       `json:"selectableVals"`
}

func (srcProps SelectionProperties) Clone(remappedIDs uniqueID.UniqueIDRemapper) (*SelectionProperties, error) {

	destProps := srcProps

	destLink, err := srcProps.ComponentLink.Clone(remappedIDs)
	if err != nil {
		return nil, fmt.Errorf("SelectionProperties.Clone: %v", err)
	}
	destProps.ComponentLink = *destLink

	return &destProps, nil
}
