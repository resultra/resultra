package form

import (
	"fmt"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/common/recordSortDataModel"
	"resultra/datasheet/server/generic/uniqueID"
	"resultra/datasheet/server/recordFilter"
)

type FormProperties struct {
	Layout                 componentLayout.ComponentLayout      `json:"layout"`
	DefaultRecordSortRules []recordSortDataModel.RecordSortRule `json:"defaultRecordSortRules"`
	DefaultFilterRules     []recordFilter.RecordFilterRule      `json:"defaultFilterRules"`
}

func (srcProps FormProperties) Clone(remappedIDs uniqueID.UniqueIDRemapper) (*FormProperties, error) {

	destFilterRules, err := recordFilter.CloneFilterRules(remappedIDs, srcProps.DefaultFilterRules)
	if err != nil {
		return nil, fmt.Errorf("FormProperties.Clone: %v")
	}

	destSortRules, err := recordSortDataModel.CloneSortRules(remappedIDs, srcProps.DefaultRecordSortRules)
	if err != nil {
		return nil, fmt.Errorf("FormProperties.Clone: %v")
	}

	destProps := FormProperties{
		Layout:                 srcProps.Layout.Clone(remappedIDs),
		DefaultRecordSortRules: destSortRules,
		DefaultFilterRules:     destFilterRules}

	return &destProps, nil
}

func newDefaultFormProperties() FormProperties {
	defaultProps := FormProperties{
		Layout:                 componentLayout.ComponentLayout{},
		DefaultRecordSortRules: []recordSortDataModel.RecordSortRule{},
		DefaultFilterRules:     []recordFilter.RecordFilterRule{}}

	return defaultProps
}
