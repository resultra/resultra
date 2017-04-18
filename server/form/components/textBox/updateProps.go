package textBox

import (
	"fmt"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/form/components/common"
)

type TextBoxIDInterface interface {
	getTextBoxID() string
	getParentFormID() string
}

type TextBoxIDHeader struct {
	ParentFormID string `json:"parentFormID"`
	TextBoxID    string `json:"textBoxID"`
}

func (idHeader TextBoxIDHeader) getTextBoxID() string {
	return idHeader.TextBoxID
}

func (idHeader TextBoxIDHeader) getParentFormID() string {
	return idHeader.ParentFormID
}

type TextBoxPropUpdater interface {
	TextBoxIDInterface
	updateProps(textBox *TextBox) error
}

func updateTextBoxProps(propUpdater TextBoxPropUpdater) (*TextBox, error) {

	// Retrieve the bar chart from the data store
	textBoxForUpdate, getErr := getTextBox(propUpdater.getParentFormID(), propUpdater.getTextBoxID())
	if getErr != nil {
		return nil, fmt.Errorf("UpdateTextBoxProps: Unable to get existing text box: %v", getErr)
	}

	if propUpdateErr := propUpdater.updateProps(textBoxForUpdate); propUpdateErr != nil {
		return nil, fmt.Errorf("UpdateTextBoxProps: Unable to update existing text box properties: %v", propUpdateErr)
	}

	textBox, updateErr := updateExistingTextBox(propUpdater.getTextBoxID(), textBoxForUpdate)
	if updateErr != nil {
		return nil, fmt.Errorf("UpdateTextBoxProps: Unable to update existing text box properties: datastore update error =  %v", updateErr)
	}

	return textBox, nil
}

type TextBoxResizeParams struct {
	TextBoxIDHeader
	Geometry componentLayout.LayoutGeometry `json:"geometry"`
}

func (updateParams TextBoxResizeParams) updateProps(textBox *TextBox) error {

	if !componentLayout.ValidGeometry(updateParams.Geometry) {
		return fmt.Errorf("set text box dimensions: Invalid geometry: %+v", updateParams.Geometry)
	}

	textBox.Properties.Geometry = updateParams.Geometry

	return nil
}

type TextBoxLabelFormatParams struct {
	TextBoxIDHeader
	LabelFormat common.ComponentLabelFormatProperties `json:"labelFormat"`
}

func (updateParams TextBoxLabelFormatParams) updateProps(textBox *TextBox) error {

	// TODO - Validate format is well-formed.

	textBox.Properties.LabelFormat = updateParams.LabelFormat

	return nil
}

type TextBoxVisibilityParams struct {
	TextBoxIDHeader
	common.ComponentVisibilityProperties
}

func (updateParams TextBoxVisibilityParams) updateProps(textBox *TextBox) error {

	// TODO - Validate conditions

	textBox.Properties.VisibilityConditions = updateParams.VisibilityConditions

	return nil
}

type TextBoxPermissionParams struct {
	TextBoxIDHeader
	Permissions common.ComponentValuePermissionsProperties `json:"permissions"`
}

func (updateParams TextBoxPermissionParams) updateProps(textBox *TextBox) error {

	textBox.Properties.Permissions = updateParams.Permissions

	return nil
}

type TextBoxValueListParams struct {
	TextBoxIDHeader
	ValueListID *string `json:"valueListID"`
}

func (updateParams TextBoxValueListParams) updateProps(textBox *TextBox) error {

	if updateParams.ValueListID != nil {
		textBox.Properties.ValueListID = updateParams.ValueListID
	} else {
		textBox.Properties.ValueListID = nil

	}

	return nil
}
