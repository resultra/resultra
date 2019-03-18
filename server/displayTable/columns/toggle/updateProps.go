package toggle

import (
	"database/sql"
	"fmt"
	"resultra/tracker/server/form/components/common"
)

type ToggleIDInterface interface {
	getToggleID() string
	getParentTableID() string
}

type ToggleIDHeader struct {
	ParentTableID string `json:"parentTableID"`
	ToggleID      string `json:"toggleID"`
}

func (idHeader ToggleIDHeader) getToggleID() string {
	return idHeader.ToggleID
}

func (idHeader ToggleIDHeader) getParentTableID() string {
	return idHeader.ParentTableID
}

type TogglePropUpdater interface {
	ToggleIDInterface
	updateProps(toggle *Toggle) error
}

func updateToggleProps(trackerDBHandle *sql.DB, propUpdater TogglePropUpdater) (*Toggle, error) {

	// Retrieve the bar chart from the data store
	toggleForUpdate, getErr := getToggle(trackerDBHandle, propUpdater.getParentTableID(), propUpdater.getToggleID())
	if getErr != nil {
		return nil, fmt.Errorf("updateToggleProps: Unable to get existing check box: %v", getErr)
	}

	if propUpdateErr := propUpdater.updateProps(toggleForUpdate); propUpdateErr != nil {
		return nil, fmt.Errorf("updateToggleProps: Unable to update existing check box properties: %v", propUpdateErr)
	}

	toggle, updateErr := updateExistingToggle(trackerDBHandle, toggleForUpdate)
	if updateErr != nil {
		return nil, fmt.Errorf("updateToggleProps: Unable to update existing check box properties: datastore update error =  %v", updateErr)
	}

	return toggle, nil
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

type ToggleOffLabelParams struct {
	ToggleIDHeader
	Label string `json:"label"`
}

func (updateParams ToggleOffLabelParams) updateProps(toggle *Toggle) error {

	// TODO - Validate against list of valid color schemes

	toggle.Properties.OffLabel = updateParams.Label

	return nil
}

type ToggleOnLabelParams struct {
	ToggleIDHeader
	Label string `json:"label"`
}

func (updateParams ToggleOnLabelParams) updateProps(toggle *Toggle) error {

	// TODO - Validate against list of valid color schemes

	toggle.Properties.OnLabel = updateParams.Label

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

type ToggleClearValueSupportedParams struct {
	ToggleIDHeader
	ClearValueSupported bool `json:"clearValueSupported"`
}

func (updateParams ToggleClearValueSupportedParams) updateProps(toggle *Toggle) error {

	toggle.Properties.ClearValueSupported = updateParams.ClearValueSupported

	return nil
}

type HelpPopupMsgParams struct {
	ToggleIDHeader
	PopupMsg string `json:"popupMsg"`
}

func (updateParams HelpPopupMsgParams) updateProps(toggle *Toggle) error {

	toggle.Properties.HelpPopupMsg = updateParams.PopupMsg

	return nil
}
