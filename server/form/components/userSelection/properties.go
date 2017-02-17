package userSelection

import (
	"fmt"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/generic/uniqueID"
)

type UserSelectionProperties struct {
	FieldID  string                         `json:"fieldID"`
	Geometry componentLayout.LayoutGeometry `json:"geometry"`
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
