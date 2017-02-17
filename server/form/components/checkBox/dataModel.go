package checkBox

import (
	"fmt"
	"log"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/form/components/common"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/uniqueID"
)

const checkBoxEntityKind string = "checkbox"

type CheckBox struct {
	ParentFormID string             `json:"parentID"`
	CheckBoxID   string             `json:"checkBoxID"`
	Properties   CheckBoxProperties `json:"properties"`
}

type NewCheckBoxParams struct {
	ParentFormID string                         `json:"parentFormID"`
	FieldID      string                         `json:"fieldID"`
	Geometry     componentLayout.LayoutGeometry `json:"geometry"`
}

func validCheckBoxFieldType(fieldType string) bool {
	if fieldType == field.FieldTypeBool {
		return true
	} else {
		return false
	}
}

func saveCheckbox(newCheckBox CheckBox) error {
	if saveErr := common.SaveNewFormComponent(checkBoxEntityKind,
		newCheckBox.ParentFormID, newCheckBox.CheckBoxID, newCheckBox.Properties); saveErr != nil {
		return fmt.Errorf("saveCheckbox: Unable to save bar chart with error = %v", saveErr)
	}
	return nil
}

func saveNewCheckBox(params NewCheckBoxParams) (*CheckBox, error) {

	if !componentLayout.ValidGeometry(params.Geometry) {
		return nil, fmt.Errorf("Invalid layout container parameters: %+v", params)
	}

	if fieldErr := common.ValidateField(params.FieldID, validCheckBoxFieldType); fieldErr != nil {
		return nil, fmt.Errorf("saveNewCheckBox: %v", fieldErr)
	}

	properties := CheckBoxProperties{
		FieldID:  params.FieldID,
		Geometry: params.Geometry}

	newCheckBox := CheckBox{ParentFormID: params.ParentFormID,
		CheckBoxID: uniqueID.GenerateSnowflakeID(),
		Properties: properties}

	if err := saveCheckbox(newCheckBox); err != nil {
		return nil, fmt.Errorf("saveNewCheckBox: Unable to save bar chart with params=%+v: error = %v", params, err)
	}

	log.Printf("INFO: API: New Checkbox: Created new check box container:  %+v", newCheckBox)

	return &newCheckBox, nil

}

func getCheckBox(parentFormID string, checkBoxID string) (*CheckBox, error) {

	checkBoxProps := CheckBoxProperties{}
	if getErr := common.GetFormComponent(checkBoxEntityKind, parentFormID, checkBoxID, &checkBoxProps); getErr != nil {
		return nil, fmt.Errorf("getCheckBox: Unable to retrieve check box: %v", getErr)
	}

	checkBox := CheckBox{
		ParentFormID: parentFormID,
		CheckBoxID:   checkBoxID,
		Properties:   checkBoxProps}

	return &checkBox, nil
}

func GetCheckBoxes(parentFormID string) ([]CheckBox, error) {

	checkBoxes := []CheckBox{}
	addCheckbox := func(checkboxID string, encodedProps string) error {

		var checkBoxProps CheckBoxProperties
		if decodeErr := generic.DecodeJSONString(encodedProps, &checkBoxProps); decodeErr != nil {
			return fmt.Errorf("GetCheckBoxes: can't decode properties: %v", encodedProps)
		}

		currCheckBox := CheckBox{
			ParentFormID: parentFormID,
			CheckBoxID:   checkboxID,
			Properties:   checkBoxProps}
		checkBoxes = append(checkBoxes, currCheckBox)

		return nil
	}
	if getErr := common.GetFormComponents(checkBoxEntityKind, parentFormID, addCheckbox); getErr != nil {
		return nil, fmt.Errorf("GetCheckBoxes: Can't get checkboxes: %v")
	}

	return checkBoxes, nil
}

func CloneCheckBoxes(remappedIDs uniqueID.UniqueIDRemapper, parentFormID string) error {

	srcCheckBoxes, err := GetCheckBoxes(parentFormID)
	if err != nil {
		return fmt.Errorf("CloneCheckBoxes: %v", err)
	}

	for _, srcCheckBox := range srcCheckBoxes {
		remappedCheckBoxID := remappedIDs.AllocNewOrGetExistingRemappedID(srcCheckBox.CheckBoxID)
		remappedFormID, err := remappedIDs.GetExistingRemappedID(srcCheckBox.ParentFormID)
		if err != nil {
			return fmt.Errorf("CloneCheckBoxes: %v", err)
		}
		destProperties, err := srcCheckBox.Properties.Clone(remappedIDs)
		if err != nil {
			return fmt.Errorf("CloneCheckBoxes: %v", err)
		}
		destCheckBox := CheckBox{
			ParentFormID: remappedFormID,
			CheckBoxID:   remappedCheckBoxID,
			Properties:   *destProperties}
		if err := saveCheckbox(destCheckBox); err != nil {
			return fmt.Errorf("CloneCheckBoxes: %v", err)
		}
	}

	return nil
}

func updateExistingCheckBox(updatedCheckBox *CheckBox) (*CheckBox, error) {

	if updateErr := common.UpdateFormComponent(checkBoxEntityKind, updatedCheckBox.ParentFormID,
		updatedCheckBox.CheckBoxID, updatedCheckBox.Properties); updateErr != nil {
		return nil, fmt.Errorf("updateExistingCheckBox: failure updating checkbox: %v", updateErr)
	}
	return updatedCheckBox, nil

}
