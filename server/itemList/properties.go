// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package itemList

import (
	"fmt"
	"resultra/tracker/server/common/recordSortDataModel"
	"resultra/tracker/server/generic/uniqueID"
	"resultra/tracker/server/recordFilter"
	"resultra/tracker/server/trackerDatabase"
)

type ItemListViewProperties struct {
	FormID   *string `json:"formID,omitempty"`
	TableID  *string `json:"tableID,omitempty"`
	PageSize int     `json:"pageSize"`
}

func (srcProps ItemListViewProperties) Clone(remappedIDs uniqueID.UniqueIDRemapper) ItemListViewProperties {

	destProps := srcProps

	if srcProps.FormID != nil {
		destFormID := remappedIDs.AllocNewOrGetExistingRemappedID(*srcProps.FormID)
		destProps.FormID = &destFormID
	} else if srcProps.TableID != nil {
		destTableID := remappedIDs.AllocNewOrGetExistingRemappedID(*srcProps.TableID)
		destProps.TableID = &destTableID
	}
	return destProps
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
	IncludeInSidebar       bool                                 `json:"includeInSidebar"`
}

func (srcProps ItemListProperties) Clone(cloneParams *trackerDatabase.CloneDatabaseParams) (*ItemListProperties, error) {

	destFilterRules, err := srcProps.DefaultFilterRules.Clone(cloneParams)
	if err != nil {
		return nil, fmt.Errorf("FormProperties.Clone: %v")
	}

	destPreFilterRules, err := srcProps.PreFilterRules.Clone(cloneParams)
	if err != nil {
		return nil, fmt.Errorf("FormProperties.Clone: %v")
	}

	destFilterFields := uniqueID.CloneIDList(cloneParams.IDRemapper, srcProps.DefaultFilterFields)

	destSortRules, err := recordSortDataModel.CloneSortRules(cloneParams.IDRemapper, srcProps.DefaultRecordSortRules)
	if err != nil {
		return nil, fmt.Errorf("FormProperties.Clone: %v")
	}
	destSortFields := uniqueID.CloneIDList(cloneParams.IDRemapper, srcProps.DefaultSortFields)

	destAltViews := []ItemListViewProperties{}
	for _, srcAltView := range srcProps.AlternateViews {
		destAltView := ItemListViewProperties{}
		destAltView.PageSize = srcAltView.PageSize
		if srcAltView.FormID != nil {
			remappedID := cloneParams.IDRemapper.AllocNewOrGetExistingRemappedID(*srcAltView.FormID)
			destAltView.FormID = &remappedID
		} else if srcAltView.TableID != nil {
			remappedID := cloneParams.IDRemapper.AllocNewOrGetExistingRemappedID(*srcAltView.TableID)
			destAltView.TableID = &remappedID
		}
		destAltViews = append(destAltViews, destAltView)
	}

	destDefaultView := srcProps.DefaultView.Clone(cloneParams.IDRemapper)

	destProps := ItemListProperties{
		DefaultView:            destDefaultView,
		DefaultRecordSortRules: destSortRules,
		DefaultFilterRules:     *destFilterRules,
		DefaultFilterFields:    destFilterFields,
		DefaultSortFields:      destSortFields,
		PreFilterRules:         *destPreFilterRules,
		AlternateViews:         destAltViews,
		IncludeInSidebar:       srcProps.IncludeInSidebar}

	return &destProps, nil
}

func newDefaultItemListProperties() ItemListProperties {
	defaultProps := ItemListProperties{
		DefaultRecordSortRules: []recordSortDataModel.RecordSortRule{},
		DefaultFilterRules:     recordFilter.NewDefaultRecordFilterRuleSet(),
		DefaultFilterFields:    []string{},
		PreFilterRules:         recordFilter.NewDefaultRecordFilterRuleSet(),
		AlternateViews:         []ItemListViewProperties{},
		DefaultSortFields:      []string{},
		IncludeInSidebar:       true}

	return defaultProps
}
