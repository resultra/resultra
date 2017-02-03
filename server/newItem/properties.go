package newItem

import (
	"fmt"
	"resultra/datasheet/server/generic/uniqueID"
	"resultra/datasheet/server/record"
)

type NewItemPresetProperties struct {
	DefaultValues []record.DefaultFieldValue `json:"defaultValues"`
}

func newDefaultNewItemProperties() NewItemPresetProperties {
	defaultProps := NewItemPresetProperties{
		DefaultValues: []record.DefaultFieldValue{}}
	return defaultProps
}

func (srcProps NewItemPresetProperties) Clone(remappedIDs uniqueID.UniqueIDRemapper) (*NewItemPresetProperties, error) {

	destProps := srcProps

	destDefaultVals, cloneErr := record.CloneDefaultFieldValues(remappedIDs, srcProps.DefaultValues)
	if cloneErr != nil {
		return nil, fmt.Errorf("NewItemPresetProperties.Clone: %v", cloneErr)
	}

	destProps.DefaultValues = destDefaultVals

	return &destProps, nil

}
