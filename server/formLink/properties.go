package formLink

import (
	"fmt"
	"resultra/tracker/server/record"
	"resultra/tracker/server/trackerDatabase"
)

type FormLinkProperties struct {
	DefaultValues []record.DefaultFieldValue `json:"defaultValues"`
}

func newDefaultNewItemProperties() FormLinkProperties {
	defaultProps := FormLinkProperties{
		DefaultValues: []record.DefaultFieldValue{}}
	return defaultProps
}

func (srcProps FormLinkProperties) Clone(cloneParams *trackerDatabase.CloneDatabaseParams) (*FormLinkProperties, error) {

	destProps := srcProps

	destDefaultVals, cloneErr := record.CloneDefaultFieldValues(cloneParams.IDRemapper, srcProps.DefaultValues)
	if cloneErr != nil {
		return nil, fmt.Errorf("FormLinkProperties.Clone: %v", cloneErr)
	}

	destProps.DefaultValues = destDefaultVals

	return &destProps, nil

}
