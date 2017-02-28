package progress

import (
	"fmt"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/generic/uniqueID"
)

type ProgressProperties struct {
	FieldID  string                         `json:"fieldID"`
	Geometry componentLayout.LayoutGeometry `json:"geometry"`
	MinVal   float64                        `json:"minVal"`
	MaxVal   float64                        `json:"maxVal"`
}

func newDefaultProgressProperties() ProgressProperties {
	props := ProgressProperties{
		FieldID: "",
		MinVal:  0.0,
		MaxVal:  100.0}
	return props

}

func (srcProps ProgressProperties) Clone(remappedIDs uniqueID.UniqueIDRemapper) (*ProgressProperties, error) {

	destProps := srcProps

	remappedFieldID, err := remappedIDs.GetExistingRemappedID(srcProps.FieldID)
	if err != nil {
		return nil, fmt.Errorf("Clone: %v", err)
	}
	destProps.FieldID = remappedFieldID

	return &destProps, nil
}
