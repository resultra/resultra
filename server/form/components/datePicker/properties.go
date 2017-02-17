package datePicker

import (
	"fmt"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/generic/uniqueID"
)

type DatePickerProperties struct {
	FieldID  string                         `json:"fieldID"`
	Geometry componentLayout.LayoutGeometry `json:"geometry"`
}

func (srcProps DatePickerProperties) Clone(remappedIDs uniqueID.UniqueIDRemapper) (*DatePickerProperties, error) {

	destProps := srcProps

	remappedFieldID, err := remappedIDs.GetExistingRemappedID(srcProps.FieldID)
	if err != nil {
		return nil, fmt.Errorf("Clone: %v", err)
	}
	destProps.FieldID = remappedFieldID

	return &destProps, nil
}
