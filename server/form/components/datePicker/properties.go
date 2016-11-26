package datePicker

import (
	"fmt"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/form/components/common"
	"resultra/datasheet/server/generic/uniqueID"
)

type DatePickerProperties struct {
	ComponentLink common.ComponentLink           `json:"componentLink"`
	Geometry      componentLayout.LayoutGeometry `json:"geometry"`
}

func (srcProps DatePickerProperties) Clone(remappedIDs uniqueID.UniqueIDRemapper) (*DatePickerProperties, error) {

	destProps := srcProps

	destLink, err := srcProps.ComponentLink.Clone(remappedIDs)
	if err != nil {
		return nil, fmt.Errorf("DatePickerProperties.Clone: %v", err)
	}
	destProps.ComponentLink = *destLink

	return &destProps, nil
}
