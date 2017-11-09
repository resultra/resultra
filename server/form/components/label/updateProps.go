package label

import (
	"database/sql"
	"fmt"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/form/components/common"
)

type LabelIDInterface interface {
	getLabelID() string
	getParentFormID() string
}

type LabelIDHeader struct {
	ParentFormID string `json:"parentFormID"`
	LabelID      string `json:"labelID"`
}

func (idHeader LabelIDHeader) getLabelID() string {
	return idHeader.LabelID
}

func (idHeader LabelIDHeader) getParentFormID() string {
	return idHeader.ParentFormID
}

type LabelPropUpdater interface {
	LabelIDInterface
	updateProps(label *Label) error
}

func updateLabelProps(trackerDBHandle *sql.DB, propUpdater LabelPropUpdater) (*Label, error) {

	// Retrieve the bar chart from the data store
	labelForUpdate, getErr := getLabel(trackerDBHandle, propUpdater.getParentFormID(), propUpdater.getLabelID())
	if getErr != nil {
		return nil, fmt.Errorf("updateLabelProps: Unable to get existing label: %v", getErr)
	}

	if propUpdateErr := propUpdater.updateProps(labelForUpdate); propUpdateErr != nil {
		return nil, fmt.Errorf("updateLabelProps: Unable to update existing label properties: %v", propUpdateErr)
	}

	updatedLabel, updateErr := updateExistingLabel(trackerDBHandle, labelForUpdate)
	if updateErr != nil {
		return nil, fmt.Errorf("updateLabelProps: Unable to update existing label properties: datastore update error =  %v", updateErr)
	}

	return updatedLabel, nil
}

type LabelResizeParams struct {
	LabelIDHeader
	Geometry componentLayout.LayoutGeometry `json:"geometry"`
}

func (updateParams LabelResizeParams) updateProps(label *Label) error {

	if !componentLayout.ValidGeometry(updateParams.Geometry) {
		return fmt.Errorf("set check box dimensions: Invalid geometry: %+v", updateParams.Geometry)
	}

	label.Properties.Geometry = updateParams.Geometry

	return nil
}

type LabelLabelFormatParams struct {
	LabelIDHeader
	LabelFormat common.ComponentLabelFormatProperties `json:"labelFormat"`
}

func (updateParams LabelLabelFormatParams) updateProps(label *Label) error {

	// TODO - Validate format is well-formed.

	label.Properties.LabelFormat = updateParams.LabelFormat

	return nil
}

type LabelVisibilityParams struct {
	LabelIDHeader
	common.ComponentVisibilityProperties
}

func (updateParams LabelVisibilityParams) updateProps(label *Label) error {

	// TODO - Validate conditions

	label.Properties.VisibilityConditions = updateParams.VisibilityConditions

	return nil
}

type LabelPermissionParams struct {
	LabelIDHeader
	Permissions common.ComponentValuePermissionsProperties `json:"permissions"`
}

func (updateParams LabelPermissionParams) updateProps(label *Label) error {

	label.Properties.Permissions = updateParams.Permissions

	return nil
}

type LabelValidationParams struct {
	LabelIDHeader
	Validation ValidationProperties `json:"validation"`
}

func (updateParams LabelValidationParams) updateProps(label *Label) error {

	label.Properties.Validation = updateParams.Validation

	return nil
}

type HelpPopupMsgParams struct {
	LabelIDHeader
	PopupMsg string `json:"popupMsg"`
}

func (updateParams HelpPopupMsgParams) updateProps(label *Label) error {

	label.Properties.HelpPopupMsg = updateParams.PopupMsg

	return nil
}
