package checkBox

import (
	"appengine"
	"fmt"
	"resultra/datasheet/server/common"
)

type CheckboxIDInterface interface {
	getCheckBoxID() string
}

type CheckboxIDHeader struct {
	CheckBoxID string `json:"checkBoxID"`
}

func (idHeader CheckboxIDHeader) getCheckBoxID() string {
	return idHeader.CheckBoxID
}

type CheckBoxPropUpdater interface {
	CheckboxIDInterface
	updateProps(checkBox *CheckBox) error
}

func updateCheckBoxProps(appEngContext appengine.Context, propUpdater CheckBoxPropUpdater) (*CheckBox, error) {

	// Retrieve the bar chart from the data store
	checkBoxForUpdate, getErr := getCheckBox(appEngContext, propUpdater.getCheckBoxID())
	if getErr != nil {
		return nil, fmt.Errorf("updateCheckBoxProps: Unable to get existing check box: %v", getErr)
	}

	if propUpdateErr := propUpdater.updateProps(checkBoxForUpdate); propUpdateErr != nil {
		return nil, fmt.Errorf("updateCheckBoxProps: Unable to update existing check box properties: %v", propUpdateErr)
	}

	checkBox, updateErr := updateExistingCheckBox(appEngContext, propUpdater.getCheckBoxID(), checkBoxForUpdate)
	if updateErr != nil {
		return nil, fmt.Errorf("updateCheckBoxProps: Unable to update existing check box properties: datastore update error =  %v", updateErr)
	}

	return checkBox, nil
}

type CheckBoxResizeParams struct {
	CheckboxIDHeader
	Geometry common.LayoutGeometry `json:"geometry"`
}

func (updateParams CheckBoxResizeParams) updateProps(checkBox *CheckBox) error {

	if !common.ValidGeometry(updateParams.Geometry) {
		return fmt.Errorf("set check box dimensions: Invalid geometry: %+v", updateParams.Geometry)
	}

	checkBox.Geometry = updateParams.Geometry

	return nil
}

type CheckBoxRepositionParams struct {
	CheckboxIDHeader
	Position common.LayoutPosition `json:"position"`
}

func (updateParams CheckBoxRepositionParams) updateProps(checkBox *CheckBox) error {

	if err := checkBox.Geometry.SetPosition(updateParams.Position); err != nil {
		return fmt.Errorf("Error setting position for check box: Invalid geometry: %v", err)
	}

	return nil
}
