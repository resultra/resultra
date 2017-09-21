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
	DefaultFilterFields    []string                             `json:"defaultFilterFields"`
	DefaultSortFields      []string                             `json:"defaultSortFields"`
	PreFilterRules         recordFilter.RecordFilterRuleSet     `json:"preFilterRules"`
	DefaultView            ItemListViewProperties               `json:"defaultView"`
	AlternateViews         []ItemListViewProperties             `json:"alternateViews"`
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

	destFilterFields := uniqueID.CloneIDList(remappedIDs, srcProps.DefaultFilterFields)

	destSortRules, err := recordSortDataModel.CloneSortRules(remappedIDs, srcProps.DefaultRecordSortRules)
	if err != nil {
		return nil, fmt.Errorf("FormProperties.Clone: %v")
	}
	destSortFields := uniqueID.CloneIDList(remappedIDs, srcProps.DefaultSortFields)

	destAltViews := []ItemListViewProperties{}
	for _, srcAltView := range srcProps.AlternateViews {
		destAltView := ItemListViewProperties{}
		destAltView.PageSize = srcAltView.PageSize
		if srcAltView.FormID != nil {
			remappedID := remappedIDs.AllocNewOrGetExistingRemappedID(*srcAltView.FormID)
			destAltView.FormID = &remappedID
		} else if srcAltView.TableID != nil {
			remappedID := remappedIDs.AllocNewOrGetExistingRemappedID(*srcAltView.TableID)
			destAltView.TableID = &remappedID
		}
		destAltViews = append(destAltViews, destAltView)
	}

	destProps := ItemListProperties{
		DefaultRecordSortRules: destSortRules,
		DefaultFilterRules:     *destFilterRules,
		DefaultFilterFields:    destFilterFields,
		DefaultSortFields:      destSortFields,
		PreFilterRules:         *destPreFilterRules,
		AlternateViews:         destAltViews}

	return &destProps, nil
}

func newDefaultItemListProperties() ItemListProperties {
	defaultProps := ItemListProperties{
		DefaultRecordSortRules: []recordSortDataModel.RecordSortRule{},
		DefaultFilterRules:     recordFilter.NewDefaultRecordFilterRuleSet(),
		DefaultFilterFields:    []string{},
		PreFilterRules:         recordFilter.NewDefaultRecordFilterRuleSet(),
		AlternateViews:         []ItemListViewProperties{},
		DefaultSortFields:      []string{}}

	return defaultProps
}
