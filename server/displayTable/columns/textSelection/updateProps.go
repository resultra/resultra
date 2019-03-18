package textSelection

import (
	"database/sql"
	"fmt"
	"resultra/tracker/server/form/components/common"
)

type TextSelectionIDInterface interface {
	getSelectionID() string
	getParentTableID() string
}

type TextSelectionIDHeader struct {
	ParentTableID string `json:"parentTableID"`
	SelectionID   string `json:"selectionID"`
}

func (idHeader TextSelectionIDHeader) getSelectionID() string {
	return idHeader.SelectionID
}

func (idHeader TextSelectionIDHeader) getParentTableID() string {
	return idHeader.ParentTableID
}

type TextSelectionPropUpdater interface {
	TextSelectionIDInterface
	updateProps(textSelection *TextSelection) error
}

func updateTextSelectionProps(trackerDBHandle *sql.DB, propUpdater TextSelectionPropUpdater) (*TextSelection, error) {

	// Retrieve the bar chart from the data store
	textSelectionForUpdate, getErr := getTextSelection(trackerDBHandle, propUpdater.getParentTableID(), propUpdater.getSelectionID())
	if getErr != nil {
		return nil, fmt.Errorf("UpdateTextInputProps: Unable to get existing text box: %v", getErr)
	}

	if propUpdateErr := propUpdater.updateProps(textSelectionForUpdate); propUpdateErr != nil {
		return nil, fmt.Errorf("UpdateTextInputProps: Unable to update existing text box properties: %v", propUpdateErr)
	}

	textSelection, updateErr := updateExistingTextSelection(trackerDBHandle, propUpdater.getSelectionID(), textSelectionForUpdate)
	if updateErr != nil {
		return nil, fmt.Errorf("UpdateTextInputProps: Unable to update existing text box properties: datastore update error =  %v", updateErr)
	}

	return textSelection, nil
}

type TextSelectionLabelFormatParams struct {
	TextSelectionIDHeader
	LabelFormat common.ComponentLabelFormatProperties `json:"labelFormat"`
}

func (updateParams TextSelectionLabelFormatParams) updateProps(textSelection *TextSelection) error {

	// TODO - Validate format is well-formed.

	textSelection.Properties.LabelFormat = updateParams.LabelFormat

	return nil
}

type TextSelectionPermissionParams struct {
	TextSelectionIDHeader
	Permissions common.ComponentValuePermissionsProperties `json:"permissions"`
}

func (updateParams TextSelectionPermissionParams) updateProps(textSelection *TextSelection) error {

	textSelection.Properties.Permissions = updateParams.Permissions

	return nil
}

type TextSelectionValueListParams struct {
	TextSelectionIDHeader
	ValueListID *string `json:"valueListID"`
}

func (updateParams TextSelectionValueListParams) updateProps(textSelection *TextSelection) error {

	if updateParams.ValueListID != nil {
		textSelection.Properties.ValueListID = updateParams.ValueListID
	} else {
		textSelection.Properties.ValueListID = nil

	}

	return nil
}

type TextSelectionValidationParams struct {
	TextSelectionIDHeader
	Validation TextSelectionValidationProperties `json:"validation"`
}

func (updateParams TextSelectionValidationParams) updateProps(textSelection *TextSelection) error {

	textSelection.Properties.Validation = updateParams.Validation

	return nil
}

type TextSelectionClearValueSupportedParams struct {
	TextSelectionIDHeader
	ClearValueSupported bool `json:"clearValueSupported"`
}

func (updateParams TextSelectionClearValueSupportedParams) updateProps(textSelection *TextSelection) error {

	textSelection.Properties.ClearValueSupported = updateParams.ClearValueSupported

	return nil
}

type HelpPopupMsgParams struct {
	TextSelectionIDHeader
	PopupMsg string `json:"popupMsg"`
}

func (updateParams HelpPopupMsgParams) updateProps(textSelection *TextSelection) error {

	textSelection.Properties.HelpPopupMsg = updateParams.PopupMsg

	return nil
}
