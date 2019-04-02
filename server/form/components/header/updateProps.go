// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package header

import (
	"database/sql"
	"fmt"
	"github.com/resultra/resultra/server/common/componentLayout"
	"github.com/resultra/resultra/server/form/components/common"
	"github.com/resultra/resultra/server/generic/stringValidation"
)

type HeaderIDInterface interface {
	getHeaderID() string
	getParentFormID() string
}

type HeaderIDHeader struct {
	ParentFormID string `json:"parentFormID"`
	HeaderID     string `json:"headerID"`
}

func (idHeader HeaderIDHeader) getHeaderID() string {
	return idHeader.HeaderID
}

func (idHeader HeaderIDHeader) getParentFormID() string {
	return idHeader.ParentFormID
}

type HeaderPropUpdater interface {
	HeaderIDInterface
	updateProps(header *Header) error
}

func updateHeaderProps(trackerDBHandle *sql.DB, propUpdater HeaderPropUpdater) (*Header, error) {

	// Retrieve the bar chart from the data store
	headerForUpdate, getErr := getHeader(trackerDBHandle, propUpdater.getParentFormID(), propUpdater.getHeaderID())
	if getErr != nil {
		return nil, fmt.Errorf("updateHeaderProps: Unable to get existing header: %v", getErr)
	}

	if propUpdateErr := propUpdater.updateProps(headerForUpdate); propUpdateErr != nil {
		return nil, fmt.Errorf("updateHeaderProps: Unable to update existing header properties: %v", propUpdateErr)
	}

	updatedHeader, updateErr := updateExistingHeader(trackerDBHandle, headerForUpdate)
	if updateErr != nil {
		return nil, fmt.Errorf("updateHeaderProps: Unable to update existing header properties: datastore update error =  %v", updateErr)
	}

	return updatedHeader, nil
}

type HeaderLabelParams struct {
	HeaderIDHeader
	Label string `json:"label"`
}

func (updateParams HeaderLabelParams) updateProps(headerForUpdate *Header) error {

	if !stringValidation.WellFormedItemLabel(updateParams.Label) {
		return fmt.Errorf("Update header label: invalid label: %v", updateParams.Label)
	}

	headerForUpdate.Properties.Label = updateParams.Label

	return nil
}

type HeaderResizeParams struct {
	HeaderIDHeader
	Geometry componentLayout.LayoutGeometry `json:"geometry"`
}

func (updateParams HeaderResizeParams) updateProps(headerForUpdate *Header) error {

	if !componentLayout.ValidGeometry(updateParams.Geometry) {
		return fmt.Errorf("set comment box dimensions: Invalid geometry: %+v", updateParams.Geometry)
	}

	headerForUpdate.Properties.Geometry = updateParams.Geometry

	return nil
}

type HeaderSizeParams struct {
	HeaderIDHeader
	Size string `json:"size"`
}

func (updateParams HeaderSizeParams) updateProps(headerForUpdate *Header) error {

	// TODO - Validate size

	headerForUpdate.Properties.HeaderSize = updateParams.Size

	return nil
}

type HeaderUnderlinedParams struct {
	HeaderIDHeader
	Underlined bool `json:"underlined"`
}

func (updateParams HeaderUnderlinedParams) updateProps(headerForUpdate *Header) error {

	headerForUpdate.Properties.Underlined = updateParams.Underlined

	return nil
}

type HeaderVisibilityParams struct {
	HeaderIDHeader
	common.ComponentVisibilityProperties
}

func (updateParams HeaderVisibilityParams) updateProps(headerForUpdate *Header) error {

	// TODO - Validate conditions

	headerForUpdate.Properties.VisibilityConditions = updateParams.VisibilityConditions

	return nil
}
