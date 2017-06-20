package comment

import (
	"fmt"
	"resultra/datasheet/server/form/components/common"
	"resultra/datasheet/server/generic/uniqueID"
)

type CommentProperties struct {
	FieldID     string                                     `json:"fieldID"`
	LabelFormat common.ComponentLabelFormatProperties      `json:"labelFormat"`
	Permissions common.ComponentValuePermissionsProperties `json:"permissions"`
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

func newDefaultCommentProperties() CommentProperties {
	props := CommentProperties{
		LabelFormat: common.NewDefaultLabelFormatProperties(),
		Permissions: common.NewDefaultComponentValuePermissionsProperties()}
	return props
}
