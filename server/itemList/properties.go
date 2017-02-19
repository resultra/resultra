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
	PreFilterRules         []recordFilter.RecordFilterRule      `json:"preFilterRules"`
	DefaultPageSize        int                                  `json:"defaultPageSize"`
	AlternateForms         []string                             `json:"alternateForms"`
}

func (srcProps ItemListProperties) Clone(remappedIDs uniqueID.UniqueIDRemapper) (*ItemListProperties, error) {

	destFilterRules, err := recordFilter.CloneFilterRules(remappedIDs, srcProps.DefaultFilterRules)
	if err != nil {
		return nil, fmt.Errorf("FormProperties.Clone: %v")
	}

	destPreFilterRules, err := recordFilter.CloneFilterRules(remappedIDs, srcProps.PreFilterRules)
	if err != nil {
		return nil, fmt.Errorf("FormProperties.Clone: %v")
	}

	destSortRules, err := recordSortDataModel.CloneSortRules(remappedIDs, srcProps.DefaultRecordSortRules)
	if err != nil {
		return nil, fmt.Errorf("FormProperties.Clone: %v")
	}

	destAlternateForms := uniqueID.CloneIDList(remappedIDs, srcProps.AlternateForms)

	destProps := ItemListProperties{
		DefaultRecordSortRules: destSortRules,
		DefaultFilterRules:     destFilterRules,
		PreFilterRules:         destPreFilterRules,
		DefaultPageSize:        srcProps.DefaultPageSize,
		AlternateForms:         destAlternateForms}

	return &destProps, nil
}

func newDefaultItemListProperties() ItemListProperties {
	defaultProps := ItemListProperties{
		DefaultRecordSortRules: []recordSortDataModel.RecordSortRule{},
		DefaultFilterRules:     []recordFilter.RecordFilterRule{},
		PreFilterRules:         []recordFilter.RecordFilterRule{},
		DefaultPageSize:        1,
		AlternateForms:         []string{}}

	return defaultProps
}
