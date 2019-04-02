// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package caption

import (
	"database/sql"
	"fmt"
	"github.com/resultra/resultra/server/common/componentLayout"
	"github.com/resultra/resultra/server/form/components/common"
	"github.com/resultra/resultra/server/generic/stringValidation"
)

type CaptionIDInterface interface {
	getCaptionID() string
	getParentFormID() string
}

type CaptionIDHeader struct {
	ParentFormID string `json:"parentFormID"`
	CaptionID    string `json:"captionID"`
}

func (idCaption CaptionIDHeader) getCaptionID() string {
	return idCaption.CaptionID
}

func (idCaption CaptionIDHeader) getParentFormID() string {
	return idCaption.ParentFormID
}

type CaptionPropUpdater interface {
	CaptionIDInterface
	updateProps(caption *Caption) error
}

func updateCaptionProps(trackerDBHandle *sql.DB, propUpdater CaptionPropUpdater) (*Caption, error) {

	// Retrieve the bar chart from the data store
	captionForUpdate, getErr := getCaption(trackerDBHandle, propUpdater.getParentFormID(), propUpdater.getCaptionID())
	if getErr != nil {
		return nil, fmt.Errorf("updateCaptionProps: Unable to get existing caption: %v", getErr)
	}

	if propUpdateErr := propUpdater.updateProps(captionForUpdate); propUpdateErr != nil {
		return nil, fmt.Errorf("updateCaptionProps: Unable to update existing caption properties: %v", propUpdateErr)
	}

	updatedCaption, updateErr := updateExistingCaption(trackerDBHandle, captionForUpdate)
	if updateErr != nil {
		return nil, fmt.Errorf("updateCaptionProps: Unable to update existing caption properties: datastore update error =  %v", updateErr)
	}

	return updatedCaption, nil
}

type CaptionLabelParams struct {
	CaptionIDHeader
	Label string `json:"label"`
}

func (updateParams CaptionLabelParams) updateProps(captionForUpdate *Caption) error {

	if !stringValidation.WellFormedItemLabel(updateParams.Label) {
		return fmt.Errorf("Update caption label: invalid label: %v", updateParams.Label)
	}

	captionForUpdate.Properties.Label = updateParams.Label

	return nil
}

type CaptionResizeParams struct {
	CaptionIDHeader
	Geometry componentLayout.LayoutGeometry `json:"geometry"`
}

func (updateParams CaptionResizeParams) updateProps(captionForUpdate *Caption) error {

	if !componentLayout.ValidGeometry(updateParams.Geometry) {
		return fmt.Errorf("set comment box dimensions: Invalid geometry: %+v", updateParams.Geometry)
	}

	captionForUpdate.Properties.Geometry = updateParams.Geometry

	return nil
}

type CaptionCaptionParams struct {
	CaptionIDHeader
	Caption string `json:"caption"`
}

func (updateParams CaptionCaptionParams) updateProps(captionForUpdate *Caption) error {

	captionForUpdate.Properties.Caption = updateParams.Caption

	return nil
}

type CaptionColorParams struct {
	CaptionIDHeader
	ColorScheme string `json:"colorScheme"`
}

func (updateParams CaptionColorParams) updateProps(captionForUpdate *Caption) error {

	// TODO - Validate color scheme

	captionForUpdate.Properties.ColorScheme = updateParams.ColorScheme

	return nil
}

type CaptionVisibilityParams struct {
	CaptionIDHeader
	common.ComponentVisibilityProperties
}

func (updateParams CaptionVisibilityParams) updateProps(captionForUpdate *Caption) error {

	// TODO - Validate conditions

	captionForUpdate.Properties.VisibilityConditions = updateParams.VisibilityConditions

	return nil
}
