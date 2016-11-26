package textBox

import (
	"fmt"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/form/components/common"
	"resultra/datasheet/server/generic/uniqueID"
)

type TextBoxValueFormatProperties struct {
	Format string `json:"format"`
}

type TextBoxProperties struct {
	ComponentLink common.ComponentLink           `json:"componentLink"`
	Geometry      componentLayout.LayoutGeometry `json:"geometry"`
	ValueFormat   TextBoxValueFormatProperties   `json:"valueFormat"`
}

func (srcProps TextBoxProperties) Clone(remappedIDs uniqueID.UniqueIDRemapper) (*TextBoxProperties, error) {

	destProps := srcProps

	destLink, err := srcProps.ComponentLink.Clone(remappedIDs)
	if err != nil {
		return nil, fmt.Errorf("TextBoxProperties.Clone: %v", err)
	}
	destProps.ComponentLink = *destLink

	return &destProps, nil
}
