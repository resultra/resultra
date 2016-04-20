package checkBox

import (
	"appengine"
	"fmt"
	"resultra/datasheet/server/common"
	"resultra/datasheet/server/generic/datastoreWrapper"
)

type CheckBoxPropUpdater interface {
	datastoreWrapper.UniqueIDInterface
	updateProps(checkBox *CheckBox) error
}

func updateCheckBoxProps(appEngContext appengine.Context, propUpdater CheckBoxPropUpdater) (*CheckBoxRef, error) {

	// Retrieve the bar chart from the data store
	checkBoxForUpdate, getErr := getCheckBox(appEngContext, propUpdater.GetUniqueID())
	if getErr != nil {
		return nil, fmt.Errorf("updateCheckBoxProps: Unable to get existing check box: %v", getErr)
	}

	if propUpdateErr := propUpdater.updateProps(checkBoxForUpdate); propUpdateErr != nil {
		return nil, fmt.Errorf("updateCheckBoxProps: Unable to update existing check box properties: %v", propUpdateErr)
	}

	checkBoxRef, updateErr := updateExistingCheckBox(appEngContext, propUpdater.GetUniqueID(), checkBoxForUpdate)
	if updateErr != nil {
		return nil, fmt.Errorf("updateCheckBoxProps: Unable to update existing check box properties: datastore update error =  %v", updateErr)
	}

	return checkBoxRef, nil
}

type CheckBoxResizeParams common.ObjectDimensionsParams

func (updateParams CheckBoxResizeParams) updateProps(checkBox *CheckBox) error {

	if !common.ValidGeometry(updateParams.Geometry) {
		return fmt.Errorf("set check box dimensions: Invalid geometry: %+v", updateParams.Geometry)
	}

	checkBox.Geometry = updateParams.Geometry

	return nil
}

type CheckBoxRepositionParams common.ObjectRepositionParams

func (updateParams CheckBoxRepositionParams) updateProps(checkBox *CheckBox) error {

	if err := checkBox.Geometry.SetPosition(updateParams.Position); err != nil {
		return fmt.Errorf("Error setting position for check box: Invalid geometry: %v", err)
	}

	return nil
}
