package selection

import (
	"fmt"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/form/components/common"
)

type SelectionIDInterface interface {
	getSelectionID() string
	getParentFormID() string
}

type SelectionIDHeader struct {
	ParentFormID string `json:"parentFormID"`
	SelectionID  string `json:"selectionID"`
}

func (idHeader SelectionIDHeader) getSelectionID() string {
	return idHeader.SelectionID
}

func (idHeader SelectionIDHeader) getParentFormID() string {
	return idHeader.ParentFormID
}

type SelectionPropUpdater interface {
	SelectionIDInterface
	updateProps(selection *Selection) error
}

func updateSelectionProps(propUpdater SelectionPropUpdater) (*Selection, error) {

	// Retrieve the bar chart from the data store
	selectionForUpdate, getErr := getSelection(propUpdater.getParentFormID(), propUpdater.getSelectionID())
	if getErr != nil {
		return nil, fmt.Errorf("UpdateSelectionProps: Unable to get existing text box: %v", getErr)
	}

	if propUpdateErr := propUpdater.updateProps(selectionForUpdate); propUpdateErr != nil {
		return nil, fmt.Errorf("UpdateSelectionProps: Unable to update existing selection properties: %v", propUpdateErr)
	}

	selection, updateErr := updateExistingSelection(propUpdater.getSelectionID(), selectionForUpdate)
	if updateErr != nil {
		return nil, fmt.Errorf("UpdateSelectionProps: Unable to update existing selection properties: datastore update error =  %v", updateErr)
	}

	return selection, nil
}

type SelectionResizeParams struct {
	SelectionIDHeader
	Geometry componentLayout.LayoutGeometry `json:"geometry"`
}

func (updateParams SelectionResizeParams) updateProps(selection *Selection) error {

	if !componentLayout.ValidGeometry(updateParams.Geometry) {
		return fmt.Errorf("set text box dimensions: Invalid geometry: %+v", updateParams.Geometry)
	}

	selection.Properties.Geometry = updateParams.Geometry

	return nil
}

type SelectionSelectableValsParams struct {
	SelectionIDHeader
	SelectableVals []SelectionSelectableVal `json:"selectableVals"`
}

func (updateParams SelectionSelectableValsParams) updateProps(selection *Selection) error {

	selection.Properties.SelectableVals = updateParams.SelectableVals

	return nil
}

type SelectionLabelFormatParams struct {
	SelectionIDHeader
	LabelFormat common.ComponentLabelFormatProperties `json:"labelFormat"`
}

func (updateParams SelectionLabelFormatParams) updateProps(selection *Selection) error {

	// TODO - Validate format is well-formed.

	selection.Properties.LabelFormat = updateParams.LabelFormat

	return nil
}

type SelectionVisibilityParams struct {
	SelectionIDHeader
	common.ComponentVisibilityProperties
}

func (updateParams SelectionVisibilityParams) updateProps(selection *Selection) error {

	// TODO - Validate conditions

	selection.Properties.VisibilityConditions = updateParams.VisibilityConditions

	return nil
}
