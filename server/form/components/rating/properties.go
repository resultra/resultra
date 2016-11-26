package rating

import (
	"fmt"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/form/components/common"
	"resultra/datasheet/server/generic/uniqueID"
)

type RatingProperties struct {
	ComponentLink common.ComponentLink           `json:"componentLink"`
	Geometry      componentLayout.LayoutGeometry `json:"geometry"`
	Tooltips      []string                       `json:"tooltips"`
}

func (srcProps RatingProperties) Clone(remappedIDs uniqueID.UniqueIDRemapper) (*RatingProperties, error) {

	destProps := srcProps

	destLink, err := srcProps.ComponentLink.Clone(remappedIDs)
	if err != nil {
		return nil, fmt.Errorf("RatingProperties.Clone: %v", err)
	}
	destProps.ComponentLink = *destLink

	return &destProps, nil
}
