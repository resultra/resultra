package textBox

import (
	"fmt"
	"resultra/datasheet/server/common/componentLayout"
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

type TextBoxValueFormatParams struct {
	TextBoxIDHeader
	ValueFormat TextBoxValueFormatProperties `json:"valueFormat"`
}

func (updateParams TextBoxValueFormatParams) updateProps(textBox *TextBox) error {

	textBox.Properties.ValueFormat = updateParams.ValueFormat

	return nil
}
