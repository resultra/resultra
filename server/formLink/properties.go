package formLink

import (
	"fmt"
	"resultra/datasheet/server/generic/uniqueID"
	"resultra/datasheet/server/record"
)

type FormLinkProperties struct {
	DefaultValues []record.DefaultFieldValue `json:"defaultValues"`
}

func newDefaultNewItemProperties() FormLinkProperties {
	defaultProps := FormLinkProperties{
		DefaultValues: []record.DefaultFieldValue{}}
	return defaultProps
}

func (srcProps FormLinkProperties) Clone(remappedIDs uniqueID.UniqueIDRemapper) (*FormLinkProperties, error) {

	destProps := srcProps

	destDefaultVals, cloneErr := record.CloneDefaultFieldValues(remappedIDs, srcProps.DefaultValues)
	if cloneErr != nil {
		return nil, fmt.Errorf("FormLinkProperties.Clone: %v", cloneErr)
	}

	destProps.DefaultValues = destDefaultVals

	return &destProps, nil

}
