package checkBox

import (
	"fmt"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/form/components/common"
)

type CheckboxIDInterface interface {
	getCheckBoxID() string
	getParentFormID() string
}

type CheckboxIDHeader struct {
	ParentFormID string `json:"parentFormID"`
	CheckBoxID   string `json:"checkBoxID"`
}

func (idHeader CheckboxIDHeader) getCheckBoxID() string {
	return idHeader.CheckBoxID
}

func (idHeader CheckboxIDHeader) getParentFormID() string {
	return idHeader.ParentFormID
}

type CheckBoxPropUpdater interface {
	CheckboxIDInterface
	updateProps(checkBox *CheckBox) error
}

func updateCheckBoxProps(propUpdater CheckBoxPropUpdater) (*CheckBox, error) {

	// Retrieve the bar chart from the data store
	checkBoxForUpdate, getErr := getCheckBox(propUpdater.getParentFormID(), propUpdater.getCheckBoxID())
	if getErr != nil {
		return nil, fmt.Errorf("updateCheckBoxProps: Unable to get existing check box: %v", getErr)
	}

	if propUpdateErr := propUpdater.updateProps(checkBoxForUpdate); propUpdateErr != nil {
		return nil, fmt.Errorf("updateCheckBoxProps: Unable to update existing check box properties: %v", propUpdateErr)
	}

	checkBox, updateErr := updateExistingCheckBox(checkBoxForUpdate)
	if updateErr != nil {
		return nil, fmt.Errorf("updateCheckBoxProps: Unable to update existing check box properties: datastore update error =  %v", updateErr)
	}

	return checkBox, nil
}

type CheckBoxResizeParams struct {
	CheckboxIDHeader
	Geometry componentLayout.LayoutGeometry `json:"geometry"`
}

func (updateParams CheckBoxResizeParams) updateProps(checkBox *CheckBox) error {

	if !componentLayout.ValidGeometry(updateParams.Geometry) {
		return fmt.Errorf("set check box dimensions: Invalid geometry: %+v", updateParams.Geometry)
	}

	checkBox.Properties.Geometry = updateParams.Geometry

	return nil
}

type CheckBoxColorSchemeParams struct {
	CheckboxIDHeader
	ColorScheme string `json:"colorScheme"`
}

func (updateParams CheckBoxColorSchemeParams) updateProps(checkBox *CheckBox) error {

	// TODO - Validate against list of valid color schemes

	checkBox.Properties.ColorScheme = updateParams.ColorScheme

	return nil
}

type CheckBoxStrikethroughParams struct {
	CheckboxIDHeader
	StrikethroughCompleted bool `json:"strikethroughCompleted"`
}

func (updateParams CheckBoxStrikethroughParams) updateProps(checkBox *CheckBox) error {

	// TODO - Validate against list of valid color schemes

	checkBox.Properties.StrikethroughCompleted = updateParams.StrikethroughCompleted

	return nil
}

type CheckBoxLabelFormatParams struct {
	CheckboxIDHeader
	LabelFormat common.ComponentLabelFormatProperties `json:"labelFormat"`
}

func (updateParams CheckBoxLabelFormatParams) updateProps(checkBox *CheckBox) error {

	// TODO - Validate format is well-formed.

	checkBox.Properties.LabelFormat = updateParams.LabelFormat

	return nil
}
