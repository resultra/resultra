package htmlEditor

import (
	"fmt"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/form/components/common"
	"resultra/datasheet/server/generic/uniqueID"
	"resultra/datasheet/server/recordFilter"
)

type HtmlEditorProperties struct {
	FieldID     string                                     `json:"fieldID"`
	Geometry    componentLayout.LayoutGeometry             `json:"geometry"`
	LabelFormat common.ComponentLabelFormatProperties      `json:"labelFormat"`
	Permissions common.ComponentValuePermissionsProperties `json:"permissions"`
	common.ComponentVisibilityProperties
}

func (srcProps HtmlEditorProperties) Clone(remappedIDs uniqueID.UniqueIDRemapper) (*HtmlEditorProperties, error) {

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

func newDefaultEditorProperties() HtmlEditorProperties {
	props := HtmlEditorProperties{
		ComponentVisibilityProperties: common.NewDefaultComponentVisibilityProperties(),
		LabelFormat:                   common.NewDefaultLabelFormatProperties(),
		Permissions:                   common.NewDefaultComponentValuePermissionsProperties()}
	return props
}
