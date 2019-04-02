// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package urlLink

import (
	"database/sql"
	"fmt"
	"github.com/resultra/resultra/server/common/componentLayout"
	"github.com/resultra/resultra/server/form/components/common"
)

type UrlLinkIDInterface interface {
	getUrlLinkID() string
	getParentFormID() string
}

type UrlLinkIDHeader struct {
	ParentFormID string `json:"parentFormID"`
	UrlLinkID    string `json:"urlLinkID"`
}

func (idHeader UrlLinkIDHeader) getUrlLinkID() string {
	return idHeader.UrlLinkID
}

func (idHeader UrlLinkIDHeader) getParentFormID() string {
	return idHeader.ParentFormID
}

type UrlLinkPropUpdater interface {
	UrlLinkIDInterface
	updateProps(urlLink *UrlLink) error
}

func updateUrlLinkProps(trackerDBHandle *sql.DB, propUpdater UrlLinkPropUpdater) (*UrlLink, error) {

	// Retrieve the bar chart from the data store
	urlLinkForUpdate, getErr := getUrlLink(trackerDBHandle, propUpdater.getParentFormID(), propUpdater.getUrlLinkID())
	if getErr != nil {
		return nil, fmt.Errorf("UpdateUrlLinkProps: Unable to get existing text box: %v", getErr)
	}

	if propUpdateErr := propUpdater.updateProps(urlLinkForUpdate); propUpdateErr != nil {
		return nil, fmt.Errorf("UpdateUrlLinkProps: Unable to update existing text box properties: %v", propUpdateErr)
	}

	urlLink, updateErr := updateExistingUrlLink(trackerDBHandle, propUpdater.getUrlLinkID(), urlLinkForUpdate)
	if updateErr != nil {
		return nil, fmt.Errorf("UpdateUrlLinkProps: Unable to update existing text box properties: datastore update error =  %v", updateErr)
	}

	return urlLink, nil
}

type UrlLinkResizeParams struct {
	UrlLinkIDHeader
	Geometry componentLayout.LayoutGeometry `json:"geometry"`
}

func (updateParams UrlLinkResizeParams) updateProps(urlLink *UrlLink) error {

	if !componentLayout.ValidGeometry(updateParams.Geometry) {
		return fmt.Errorf("set text box dimensions: Invalid geometry: %+v", updateParams.Geometry)
	}

	urlLink.Properties.Geometry = updateParams.Geometry

	return nil
}

type UrlLinkLabelFormatParams struct {
	UrlLinkIDHeader
	LabelFormat common.ComponentLabelFormatProperties `json:"labelFormat"`
}

func (updateParams UrlLinkLabelFormatParams) updateProps(urlLink *UrlLink) error {

	// TODO - Validate format is well-formed.

	urlLink.Properties.LabelFormat = updateParams.LabelFormat

	return nil
}

type UrlLinkVisibilityParams struct {
	UrlLinkIDHeader
	common.ComponentVisibilityProperties
}

func (updateParams UrlLinkVisibilityParams) updateProps(urlLink *UrlLink) error {

	// TODO - Validate conditions

	urlLink.Properties.VisibilityConditions = updateParams.VisibilityConditions

	return nil
}

type UrlLinkPermissionParams struct {
	UrlLinkIDHeader
	Permissions common.ComponentValuePermissionsProperties `json:"permissions"`
}

func (updateParams UrlLinkPermissionParams) updateProps(urlLink *UrlLink) error {

	urlLink.Properties.Permissions = updateParams.Permissions

	return nil
}

type UrlLinkValidationParams struct {
	UrlLinkIDHeader
	Validation UrlLinkValidationProperties `json:"validation"`
}

func (updateParams UrlLinkValidationParams) updateProps(urlLink *UrlLink) error {

	urlLink.Properties.Validation = updateParams.Validation

	return nil
}

type UrlLinkClearValueSupportedParams struct {
	UrlLinkIDHeader
	ClearValueSupported bool `json:"clearValueSupported"`
}

func (updateParams UrlLinkClearValueSupportedParams) updateProps(urlLink *UrlLink) error {

	urlLink.Properties.ClearValueSupported = updateParams.ClearValueSupported

	return nil
}

type HelpPopupMsgParams struct {
	UrlLinkIDHeader
	PopupMsg string `json:"popupMsg"`
}

func (updateParams HelpPopupMsgParams) updateProps(urlLink *UrlLink) error {

	urlLink.Properties.HelpPopupMsg = updateParams.PopupMsg

	return nil
}
