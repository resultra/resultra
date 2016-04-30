package textBox

import (
	"appengine"
	"fmt"
	"resultra/datasheet/server/common"
)

type TextBoxIDInterface interface {
	getTextBoxID() string
}

type TextBoxIDHeader struct {
	TextBoxID string `json:"textBoxID"`
}

func (idHeader TextBoxIDHeader) getTextBoxID() string {
	return idHeader.TextBoxID
}

type TextBoxPropUpdater interface {
	TextBoxIDInterface
	updateProps(textBox *TextBox) error
}

func updateTextBoxProps(appEngContext appengine.Context, propUpdater TextBoxPropUpdater) (*TextBoxRef, error) {

	// Retrieve the bar chart from the data store
	textBoxForUpdate, getErr := getTextBox(appEngContext, propUpdater.getTextBoxID())
	if getErr != nil {
		return nil, fmt.Errorf("UpdateTextBoxProps: Unable to get existing text box: %v", getErr)
	}

	if propUpdateErr := propUpdater.updateProps(textBoxForUpdate); propUpdateErr != nil {
		return nil, fmt.Errorf("UpdateTextBoxProps: Unable to update existing text box properties: %v", propUpdateErr)
	}

	textBoxRef, updateErr := updateExistingTextBox(appEngContext, propUpdater.getTextBoxID(), textBoxForUpdate)
	if updateErr != nil {
		return nil, fmt.Errorf("UpdateTextBoxProps: Unable to update existing text box properties: datastore update error =  %v", updateErr)
	}

	return textBoxRef, nil
}

type TextBoxResizeParams struct {
	TextBoxIDHeader
	Geometry common.LayoutGeometry `json:"geometry"`
}

func (updateParams TextBoxResizeParams) updateProps(textBox *TextBox) error {

	if !common.ValidGeometry(updateParams.Geometry) {
		return fmt.Errorf("set text box dimensions: Invalid geometry: %+v", updateParams.Geometry)
	}

	textBox.Geometry = updateParams.Geometry

	return nil
}

type TextBoxRepositionParams struct {
	TextBoxIDHeader
	Position common.LayoutPosition `json:"position"`
}

func (updateParams TextBoxRepositionParams) updateProps(textBox *TextBox) error {

	if err := textBox.Geometry.SetPosition(updateParams.Position); err != nil {
		return fmt.Errorf("Error setting position for text box: Invalid geometry: %v", err)
	}

	return nil
}
