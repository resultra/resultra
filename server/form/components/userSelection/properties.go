package userSelection

import (
	"fmt"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/form/components/common"
	"resultra/datasheet/server/generic/uniqueID"
)

type UserSelectionProperties struct {
	FieldID     string                                `json:"fieldID"`
	Geometry    componentLayout.LayoutGeometry        `json:"geometry"`
	LabelFormat common.ComponentLabelFormatProperties `json:"labelFormat"`
}

func (srcProps UserSelectionProperties) Clone(remappedIDs uniqueID.UniqueIDRemapper) (*UserSelectionProperties, error) {

	destProps := srcProps

	remappedFieldID, err := remappedIDs.GetExistingRemappedID(srcProps.FieldID)
	if err != nil {
		return nil, fmt.Errorf("Clone: %v", err)
	}
	destProps.FieldID = remappedFieldID

	return &destProps, nil
}

func newDefaultUserSelectionProperties() UserSelectionProperties {
	props := UserSelectionProperties{
		LabelFormat: common.NewDefaultLabelFormatProperties()}
	return props
}
