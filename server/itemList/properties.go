package itemList

import (
	"fmt"
	"resultra/datasheet/server/common/recordSortDataModel"
	"resultra/datasheet/server/generic/uniqueID"
	"resultra/datasheet/server/recordFilter"
)

type ItemListViewProperties struct {
	FormID   *string `json:"formID,omitempty"`
	TableID  *string `json:"tableID,omitempty"`
	PageSize int     `json:"pageSize"`
}

func (viewProps ItemListViewProperties) validate() error {
	// TODO - Retrieve the form or table to validate it exists
	if viewProps.TableID != nil {
		return nil
	} else if viewProps.FormID != nil && viewProps.PageSize > 0 {
		return nil
	}
	return fmt.Errorf("Invalid item list view properties: %+v", viewProps)
}

type ItemListProperties struct {
	DefaultRecordSortRules []recordSortDataModel.RecordSortRule `json:"defaultRecordSortRules"`
	DefaultFilterRules     recordFilter.RecordFilterRuleSet     `json:"defaultFilterRules"`
	PreFilterRules         recordFilter.RecordFilterRuleSet     `json:"preFilterRules"`
	AlternateForms         []string                             `json:"alternateForms"`
	DefaultView            ItemListViewProperties               `json:"defaultView"`
}

func (srcProps ItemListProperties) Clone(remappedIDs uniqueID.UniqueIDRemapper) (*ItemListProperties, error) {

	destFilterRules, err := srcProps.DefaultFilterRules.Clone(remappedIDs)
	if err != nil {
		return nil, fmt.Errorf("FormProperties.Clone: %v")
	}

	destPreFilterRules, err := srcProps.PreFilterRules.Clone(remappedIDs)
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
		DefaultFilterRules:     *destFilterRules,
		PreFilterRules:         *destPreFilterRules,
		AlternateForms:         destAlternateForms}

	return &destProps, nil
}

func newDefaultItemListProperties() ItemListProperties {
	defaultProps := ItemListProperties{
		DefaultRecordSortRules: []recordSortDataModel.RecordSortRule{},
		DefaultFilterRules:     recordFilter.NewDefaultRecordFilterRuleSet(),
		PreFilterRules:         recordFilter.NewDefaultRecordFilterRuleSet(),
		AlternateForms:         []string{}}

	return defaultProps
}
