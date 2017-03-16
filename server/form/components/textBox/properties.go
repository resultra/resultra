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

func defaultValueFormat() TextBoxValueFormatProperties {
	return TextBoxValueFormatProperties{Format: "general"}
}

type TextBoxProperties struct {
	FieldID     string                                `json:"fieldID"`
	Geometry    componentLayout.LayoutGeometry        `json:"geometry"`
	ValueFormat TextBoxValueFormatProperties          `json:"valueFormat"`
	LabelFormat common.ComponentLabelFormatProperties `json:"labelFormat"`
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

func newDefaultTextBoxProperties() TextBoxProperties {
	props := TextBoxProperties{
		LabelFormat: common.NewDefaultLabelFormatProperties(),
		ValueFormat: defaultValueFormat()}
	return props
}
