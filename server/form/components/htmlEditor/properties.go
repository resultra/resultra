package htmlEditor

import (
	"fmt"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/form/components/common"
	"resultra/datasheet/server/generic/uniqueID"
)

type HtmlEditorProperties struct {
	FieldID     string                                `json:"fieldID"`
	Geometry    componentLayout.LayoutGeometry        `json:"geometry"`
	LabelFormat common.ComponentLabelFormatProperties `json:"labelFormat"`
}

func (srcProps HtmlEditorProperties) Clone(remappedIDs uniqueID.UniqueIDRemapper) (*HtmlEditorProperties, error) {

	destProps := srcProps

	remappedFieldID, err := remappedIDs.GetExistingRemappedID(srcProps.FieldID)
	if err != nil {
		return nil, fmt.Errorf("Clone: %v", err)
	}
	destProps.FieldID = remappedFieldID

	return &destProps, nil
}

func newDefaultEditorProperties() HtmlEditorProperties {
	props := HtmlEditorProperties{
		LabelFormat: common.NewDefaultLabelFormatProperties()}
	return props
}
