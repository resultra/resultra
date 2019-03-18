package itemList

import (
	"database/sql"
	"fmt"
	"resultra/tracker/server/common/recordSortDataModel"
	"resultra/tracker/server/recordFilter"
)

type ItemListIDInterface interface {
	getListID() string
}

type ItemListIDHeader struct {
	ListID string `json:"listID"`
}

func (idHeader ItemListIDHeader) getListID() string {
	return idHeader.ListID
}

type ItemListPropUpdater interface {
	ItemListIDInterface
	updateProps(itemList *ItemList) error
}

func updateItemListProps(trackerDBHandle *sql.DB, propUpdater ItemListPropUpdater) (*ItemList, error) {

	// Retrieve the bar chart from the data store
	listForUpdate, getErr := GetItemList(trackerDBHandle, propUpdater.getListID())
	if getErr != nil {
		return nil, fmt.Errorf("updateItemListProps: Unable to get existing list: %v", getErr)
	}

	if propUpdateErr := propUpdater.updateProps(listForUpdate); propUpdateErr != nil {
		return nil, fmt.Errorf("updateItemListProps: Unable to update existing form properties: %v", propUpdateErr)
	}

	updatedItemList, updateErr := updateExistingItemList(trackerDBHandle, propUpdater.getListID(), listForUpdate)
	if updateErr != nil {
		return nil, fmt.Errorf("updateItemListProps: Unable to update existing list properties: datastore update error =  %v", updateErr)
	}

	return updatedItemList, nil
}

type SetItemListNameParams struct {
	ItemListIDHeader
	NewListName string `json:"newListName"`
}

func (updateParams SetItemListNameParams) updateProps(itemList *ItemList) error {

	// TODO - Validate name

	itemList.Name = updateParams.NewListName

	return nil
}

type SetFilterRulesParams struct {
	ItemListIDHeader
	FilterRules recordFilter.RecordFilterRuleSet `json:"filterRules"`
}

func (updateParams SetFilterRulesParams) updateProps(itemList *ItemList) error {

	// TODO - Validate filter rules before saving
	itemList.Properties.DefaultFilterRules = updateParams.FilterRules

	return nil
}

type SetPreFilterRulesParams struct {
	ItemListIDHeader
	FilterRules recordFilter.RecordFilterRuleSet `json:"filterRules"`
}

func (updateParams SetPreFilterRulesParams) updateProps(itemList *ItemList) error {

	// TODO - Validate filter rules before saving
	itemList.Properties.PreFilterRules = updateParams.FilterRules

	return nil
}

type SetDefaultSortRulesParams struct {
	ItemListIDHeader
	SortRules []recordSortDataModel.RecordSortRule `json:"sortRules"`
}

func (updateParams SetDefaultSortRulesParams) updateProps(itemList *ItemList) error {

	// TODO - Validate sort rules before saving
	itemList.Properties.DefaultRecordSortRules = updateParams.SortRules

	return nil
}

type SetDefaultViewParams struct {
	ItemListIDHeader
	View ItemListViewProperties `json:"view"`
}

func (updateParams SetDefaultViewParams) updateProps(itemList *ItemList) error {

	// TODO - Validate  before saving

	if err := updateParams.View.validate(); err != nil {
		return err
	}

	itemList.Properties.DefaultView = updateParams.View

	return nil
}

type SetAlternateViewsParams struct {
	ItemListIDHeader
	AlternateViews []ItemListViewProperties `json:"alternateViews"`
}

func (updateParams SetAlternateViewsParams) updateProps(itemList *ItemList) error {

	for _, altView := range updateParams.AlternateViews {
		if err := altView.validate(); err != nil {
			return fmt.Errorf("updateProps (item list): %v", err)
		}
	}

	itemList.Properties.AlternateViews = updateParams.AlternateViews

	return nil
}

type SetDefaultFilterFieldsParams struct {
	ItemListIDHeader
	DefaultFieldFields []string `json:"defaultFilterFields"`
}

func (updateParams SetDefaultFilterFieldsParams) updateProps(itemList *ItemList) error {

	// TODO - Validate fields

	itemList.Properties.DefaultFilterFields = updateParams.DefaultFieldFields

	return nil
}

type SetDefaultSortFieldsParams struct {
	ItemListIDHeader
	DefaultSortFields []string `json:"defaultSortFields"`
}

func (updateParams SetDefaultSortFieldsParams) updateProps(itemList *ItemList) error {

	// TODO - Validate fields

	itemList.Properties.DefaultSortFields = updateParams.DefaultSortFields

	return nil
}

type SetIncludeInSidebarParams struct {
	ItemListIDHeader
	IncludeInSidebar bool `json:"includeInSidebar"`
}

func (updateParams SetIncludeInSidebarParams) updateProps(itemList *ItemList) error {

	// TODO - Validate fields

	itemList.Properties.IncludeInSidebar = updateParams.IncludeInSidebar

	return nil
}
