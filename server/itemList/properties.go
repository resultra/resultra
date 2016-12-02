package itemList

import (
	"fmt"
	"resultra/datasheet/server/common/recordSortDataModel"
	"resultra/datasheet/server/generic/uniqueID"
	"resultra/datasheet/server/recordFilter"
)

type ItemListProperties struct {
	DefaultRecordSortRules []recordSortDataModel.RecordSortRule `json:"defaultRecordSortRules"`
	DefaultFilterRules     []recordFilter.RecordFilterRule      `json:"defaultFilterRules"`
}

func (srcProps ItemListProperties) Clone(remappedIDs uniqueID.UniqueIDRemapper) (*ItemListProperties, error) {

	destFilterRules, err := recordFilter.CloneFilterRules(remappedIDs, srcProps.DefaultFilterRules)
	if err != nil {
		return nil, fmt.Errorf("FormProperties.Clone: %v")
	}

	destSortRules, err := recordSortDataModel.CloneSortRules(remappedIDs, srcProps.DefaultRecordSortRules)
	if err != nil {
		return nil, fmt.Errorf("FormProperties.Clone: %v")
	}

	destProps := ItemListProperties{
		DefaultRecordSortRules: destSortRules,
		DefaultFilterRules:     destFilterRules}

	return &destProps, nil
}

func newDefaultItemListProperties() ItemListProperties {
	defaultProps := ItemListProperties{
		DefaultRecordSortRules: []recordSortDataModel.RecordSortRule{},
		DefaultFilterRules:     []recordFilter.RecordFilterRule{}}

	return defaultProps
}
