package comment

import (
	"fmt"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/generic/uniqueID"
)

type CommentProperties struct {
	FieldID  string                         `json:"fieldID"`
	Geometry componentLayout.LayoutGeometry `json:"geometry"`
}

func (srcProps CommentProperties) Clone(remappedIDs uniqueID.UniqueIDRemapper) (*CommentProperties, error) {

	destProps := srcProps

	remappedFieldID, err := remappedIDs.GetExistingRemappedID(srcProps.FieldID)
	if err != nil {
		return nil, fmt.Errorf("CommentProperties.Clone: %v", err)
	}
	destProps.FieldID = remappedFieldID

	return &destProps, nil
}
