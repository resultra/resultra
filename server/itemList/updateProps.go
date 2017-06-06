package itemList

import (
	"fmt"
	"resultra/datasheet/server/common/recordSortDataModel"
	"resultra/datasheet/server/recordFilter"
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

func updateItemListProps(propUpdater ItemListPropUpdater) (*ItemList, error) {

	// Retrieve the bar chart from the data store
	listForUpdate, getErr := GetItemList(propUpdater.getListID())
	if getErr != nil {
		return nil, fmt.Errorf("updateItemListProps: Unable to get existing list: %v", getErr)
	}

	if propUpdateErr := propUpdater.updateProps(listForUpdate); propUpdateErr != nil {
		return nil, fmt.Errorf("updateItemListProps: Unable to update existing form properties: %v", propUpdateErr)
	}

	updatedItemList, updateErr := updateExistingItemList(propUpdater.getListID(), listForUpdate)
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

type SetPageSizeParams struct {
	ItemListIDHeader
	PageSize int `json:"pageSize"`
}

func (updateParams SetPageSizeParams) updateProps(itemList *ItemList) error {

	itemList.Properties.DefaultPageSize = updateParams.PageSize

	return nil
}

type SetAlternateFormsParams struct {
	ItemListIDHeader
	AlternateForms []string `json:"alternateForms"`
}

func (updateParams SetAlternateFormsParams) updateProps(itemList *ItemList) error {

	itemList.Properties.AlternateForms = updateParams.AlternateForms

	return nil
}
