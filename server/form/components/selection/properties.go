package selection

import (
	"fmt"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/form/components/common"
	"resultra/datasheet/server/generic/uniqueID"
)

type SelectionProperties struct {
	FieldID        string                                     `json:"fieldID"`
	Geometry       componentLayout.LayoutGeometry             `json:"geometry"`
	SelectableVals []SelectionSelectableVal                   `json:"selectableVals"`
	LabelFormat    common.ComponentLabelFormatProperties      `json:"labelFormat"`
	Permissions    common.ComponentValuePermissionsProperties `json:"permissions"`
	common.ComponentVisibilityProperties
}

func (srcProps SelectionProperties) Clone(remappedIDs uniqueID.UniqueIDRemapper) (*SelectionProperties, error) {

	destProps := srcProps

	remappedFieldID, err := remappedIDs.GetExistingRemappedID(srcProps.FieldID)
	if err != nil {
		return nil, fmt.Errorf("Clone: %v", err)
	}
	destProps.FieldID = remappedFieldID

	destVisibilityConditions, err := srcProps.VisibilityConditions.Clone(remappedIDs)
	if err != nil {
		return nil, fmt.Errorf("CaptionProperties.Clone: %v")
	}
	destProps.VisibilityConditions = *destVisibilityConditions

	return &destProps, nil
}

func newDefaultSelectionProperties() SelectionProperties {
	props := SelectionProperties{
		ComponentVisibilityProperties: common.NewDefaultComponentVisibilityProperties(),
		SelectableVals:                []SelectionSelectableVal{},
		LabelFormat:                   common.NewDefaultLabelFormatProperties(),
		Permissions:                   common.NewDefaultComponentValuePermissionsProperties()}
	return props
}
