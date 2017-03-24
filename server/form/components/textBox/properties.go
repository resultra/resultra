package textBox

import (
	"fmt"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/form/components/common"
	"resultra/datasheet/server/generic/numberFormat"
	"resultra/datasheet/server/generic/uniqueID"

	"resultra/datasheet/server/recordFilter"
)

type TextBoxProperties struct {
	FieldID     string                                     `json:"fieldID"`
	Geometry    componentLayout.LayoutGeometry             `json:"geometry"`
	ValueFormat numberFormat.NumberFormatProperties        `json:"valueFormat"`
	LabelFormat common.ComponentLabelFormatProperties      `json:"labelFormat"`
	Permissions common.ComponentValuePermissionsProperties `json:"permissions"`
	common.ComponentVisibilityProperties
}

func (srcProps TextBoxProperties) Clone(remappedIDs uniqueID.UniqueIDRemapper) (*TextBoxProperties, error) {

	destProps := srcProps

	remappedFieldID, err := remappedIDs.GetExistingRemappedID(srcProps.FieldID)
	if err != nil {
		return nil, fmt.Errorf("Clone: %v", err)
	}
	destProps.FieldID = remappedFieldID

	destVisibilityConditions, err := recordFilter.CloneFilterRules(remappedIDs, srcProps.VisibilityConditions)
	if err != nil {
		return nil, fmt.Errorf("CaptionProperties.Clone: %v")
	}
	destProps.VisibilityConditions = destVisibilityConditions

	return &destProps, nil
}

func newDefaultTextBoxProperties() TextBoxProperties {
	props := TextBoxProperties{
		ComponentVisibilityProperties: common.NewDefaultComponentVisibilityProperties(),
		LabelFormat:                   common.NewDefaultLabelFormatProperties(),
		ValueFormat:                   numberFormat.DefaultNumberFormatProperties(),
		Permissions:                   common.NewDefaultComponentValuePermissionsProperties()}
	return props
}
