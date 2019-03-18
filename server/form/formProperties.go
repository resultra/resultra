package form

import (
	"resultra/tracker/server/common/componentLayout"
	"resultra/tracker/server/trackerDatabase"
)

type FormProperties struct {
	Layout componentLayout.ComponentLayout `json:"layout"`
}

func (srcProps FormProperties) Clone(cloneParams *trackerDatabase.CloneDatabaseParams) (*FormProperties, error) {

	destProps := FormProperties{
		Layout: srcProps.Layout.Clone(cloneParams.IDRemapper)}

	return &destProps, nil
}

func newDefaultFormProperties() FormProperties {
	defaultProps := FormProperties{
		Layout: componentLayout.ComponentLayout{}}

	return defaultProps
}
