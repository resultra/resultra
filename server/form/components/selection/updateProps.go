// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package selection

import (
	"database/sql"
	"fmt"
	"github.com/resultra/resultra/server/common/componentLayout"
	"github.com/resultra/resultra/server/form/components/common"
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

func updateSelectionProps(trackerDBHandle *sql.DB, propUpdater SelectionPropUpdater) (*Selection, error) {

	// Retrieve the bar chart from the data store
	selectionForUpdate, getErr := getSelection(trackerDBHandle, propUpdater.getParentFormID(), propUpdater.getSelectionID())
	if getErr != nil {
		return nil, fmt.Errorf("UpdateSelectionProps: Unable to get existing text box: %v", getErr)
	}

	if propUpdateErr := propUpdater.updateProps(selectionForUpdate); propUpdateErr != nil {
		return nil, fmt.Errorf("UpdateSelectionProps: Unable to update existing selection properties: %v", propUpdateErr)
	}

	selection, updateErr := updateExistingSelection(trackerDBHandle, propUpdater.getSelectionID(), selectionForUpdate)
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

type SelectionPermissionParams struct {
	SelectionIDHeader
	Permissions common.ComponentValuePermissionsProperties `json:"permissions"`
}

func (updateParams SelectionPermissionParams) updateProps(selection *Selection) error {

	selection.Properties.Permissions = updateParams.Permissions

	return nil
}

type SelectionClearValueSupportedParams struct {
	SelectionIDHeader
	ClearValueSupported bool `json:"clearValueSupported"`
}

func (updateParams SelectionClearValueSupportedParams) updateProps(selection *Selection) error {

	selection.Properties.ClearValueSupported = updateParams.ClearValueSupported

	return nil
}

type SelectionValueListParams struct {
	SelectionIDHeader
	ValueListID *string `json:"valueListID"`
}

func (updateParams SelectionValueListParams) updateProps(selection *Selection) error {

	if updateParams.ValueListID != nil {
		selection.Properties.ValueListID = updateParams.ValueListID
	} else {
		selection.Properties.ValueListID = nil

	}

	return nil
}
