package toggle

import (
	"fmt"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/form/components/common"
)

type ToggleIDInterface interface {
	getToggleID() string
	getParentFormID() string
}

type ToggleIDHeader struct {
	ParentFormID string `json:"parentFormID"`
	ToggleID     string `json:"toggleID"`
}

func (idHeader ToggleIDHeader) getToggleID() string {
	return idHeader.ToggleID
}

func (idHeader ToggleIDHeader) getParentFormID() string {
	return idHeader.ParentFormID
}

type TogglePropUpdater interface {
	ToggleIDInterface
	updateProps(toggle *Toggle) error
}

func updateToggleProps(propUpdater TogglePropUpdater) (*Toggle, error) {

	// Retrieve the bar chart from the data store
	toggleForUpdate, getErr := getToggle(propUpdater.getParentFormID(), propUpdater.getToggleID())
	if getErr != nil {
		return nil, fmt.Errorf("updateToggleProps: Unable to get existing check box: %v", getErr)
	}

	if propUpdateErr := propUpdater.updateProps(toggleForUpdate); propUpdateErr != nil {
		return nil, fmt.Errorf("updateToggleProps: Unable to update existing check box properties: %v", propUpdateErr)
	}

	toggle, updateErr := updateExistingToggle(toggleForUpdate)
	if updateErr != nil {
		return nil, fmt.Errorf("updateToggleProps: Unable to update existing check box properties: datastore update error =  %v", updateErr)
	}

	return toggle, nil
}

type ToggleResizeParams struct {
	ToggleIDHeader
	Geometry componentLayout.LayoutGeometry `json:"geometry"`
}

func (updateParams ToggleResizeParams) updateProps(toggle *Toggle) error {

	if !componentLayout.ValidGeometry(updateParams.Geometry) {
		return fmt.Errorf("set check box dimensions: Invalid geometry: %+v", updateParams.Geometry)
	}

	toggle.Properties.Geometry = updateParams.Geometry

	return nil
}

type ToggleOnColorSchemeParams struct {
	ToggleIDHeader
	ColorScheme string `json:"colorScheme"`
}

func (updateParams ToggleOnColorSchemeParams) updateProps(toggle *Toggle) error {

	// TODO - Validate against list of valid color schemes

	toggle.Properties.OnColorScheme = updateParams.ColorScheme

	return nil
}

type ToggleOffColorSchemeParams struct {
	ToggleIDHeader
	ColorScheme string `json:"colorScheme"`
}

func (updateParams ToggleOffColorSchemeParams) updateProps(toggle *Toggle) error {

	// TODO - Validate against list of valid color schemes

	toggle.Properties.OffColorScheme = updateParams.ColorScheme

	return nil
}

type ToggleLabelFormatParams struct {
	ToggleIDHeader
	LabelFormat common.ComponentLabelFormatProperties `json:"labelFormat"`
}

func (updateParams ToggleLabelFormatParams) updateProps(toggle *Toggle) error {

	// TODO - Validate format is well-formed.

	toggle.Properties.LabelFormat = updateParams.LabelFormat

	return nil
}

type ToggleVisibilityParams struct {
	ToggleIDHeader
	common.ComponentVisibilityProperties
}

func (updateParams ToggleVisibilityParams) updateProps(toggle *Toggle) error {

	// TODO - Validate conditions

	toggle.Properties.VisibilityConditions = updateParams.VisibilityConditions

	return nil
}

type TogglePermissionParams struct {
	ToggleIDHeader
	Permissions common.ComponentValuePermissionsProperties `json:"permissions"`
}

func (updateParams TogglePermissionParams) updateProps(toggle *Toggle) error {

	// TODO - Validate conditions

	toggle.Properties.Permissions = updateParams.Permissions

	return nil
}

type ToggleValidationParams struct {
	ToggleIDHeader
	Validation ToggleValidationProperties `json:"validation"`
}

func (updateParams ToggleValidationParams) updateProps(toggle *Toggle) error {

	toggle.Properties.Validation = updateParams.Validation

	return nil
}
