package form

import (
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/generic/uniqueID"
)

type FormProperties struct {
	Layout             componentLayout.ComponentLayout `json:"layout"`
	AddNewItemFromForm bool                            `json:"addNewItemFromForm"`
}

func (srcProps FormProperties) Clone(remappedIDs uniqueID.UniqueIDRemapper) (*FormProperties, error) {

	destProps := FormProperties{
		Layout: srcProps.Layout.Clone(remappedIDs)}

	return &destProps, nil
}

func newDefaultFormProperties() FormProperties {
	defaultProps := FormProperties{
		Layout:             componentLayout.ComponentLayout{},
		AddNewItemFromForm: false}

	return defaultProps
}
