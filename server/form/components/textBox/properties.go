package textBox

import (
	"fmt"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/generic/uniqueID"
)

type TextBoxValueFormatProperties struct {
	Format string `json:"format"`
}

type TextBoxProperties struct {
	FieldID     string                         `json:"fieldID"`
	Geometry    componentLayout.LayoutGeometry `json:"geometry"`
	ValueFormat TextBoxValueFormatProperties   `json:"valueFormat"`
}

func (srcProps TextBoxProperties) Clone(remappedIDs uniqueID.UniqueIDRemapper) (*TextBoxProperties, error) {

	destProps := srcProps

	remappedFieldID, err := remappedIDs.GetExistingRemappedID(srcProps.FieldID)
	if err != nil {
		return nil, fmt.Errorf("Clone: %v", err)
	}
	destProps.FieldID = remappedFieldID

	return &destProps, nil
}
