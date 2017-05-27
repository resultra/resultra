package checkBox

import (
	"fmt"
	"log"
	"resultra/datasheet/server/displayTable/columns/common"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/uniqueID"
)

const checkBoxEntityKind string = "checkbox"

type CheckBox struct {
	ParentTableID string             `json:"parentTableID"`
	CheckBoxID    string             `json:"checkBoxID"`
	Properties    CheckBoxProperties `json:"properties"`
}

type NewCheckBoxParams struct {
	ParentTableID string `json:"parentTableID"`
	FieldID       string `json:"fieldID"`
}

func validCheckBoxFieldType(fieldType string) bool {
	if fieldType == field.FieldTypeBool {
		return true
	} else {
		return false
	}
}

func saveCheckbox(newCheckBox CheckBox) error {
	if saveErr := common.SaveNewTableColumn(checkBoxEntityKind,
		newCheckBox.ParentTableID, newCheckBox.CheckBoxID, newCheckBox.Properties); saveErr != nil {
		return fmt.Errorf("saveCheckbox: Unable to save bar chart with error = %v", saveErr)
	}
	return nil
}

func saveNewCheckBox(params NewCheckBoxParams) (*CheckBox, error) {

	if fieldErr := field.ValidateField(params.FieldID, validCheckBoxFieldType); fieldErr != nil {
		return nil, fmt.Errorf("saveNewCheckBox: %v", fieldErr)
	}

	properties := newDefaultCheckBoxProperties()
	properties.FieldID = params.FieldID

	newCheckBox := CheckBox{ParentTableID: params.ParentTableID,
		CheckBoxID: uniqueID.GenerateSnowflakeID(),
		Properties: properties}

	if err := saveCheckbox(newCheckBox); err != nil {
		return nil, fmt.Errorf("saveNewCheckBox: Unable to save bar chart with params=%+v: error = %v", params, err)
	}

	log.Printf("INFO: API: New Checkbox: Created new check box container:  %+v", newCheckBox)

	return &newCheckBox, nil

}

func getCheckBox(parentTableID string, checkBoxID string) (*CheckBox, error) {

	checkBoxProps := newDefaultCheckBoxProperties()
	if getErr := common.GetTableColumn(checkBoxEntityKind, parentTableID, checkBoxID, &checkBoxProps); getErr != nil {
		return nil, fmt.Errorf("getCheckBox: Unable to retrieve check box: %v", getErr)
	}

	checkBox := CheckBox{
		ParentTableID: parentTableID,
		CheckBoxID:    checkBoxID,
		Properties:    checkBoxProps}

	return &checkBox, nil
}

func GetCheckBoxes(parentTableID string) ([]CheckBox, error) {

	checkBoxes := []CheckBox{}
	addCheckbox := func(checkboxID string, encodedProps string) error {

		checkBoxProps := newDefaultCheckBoxProperties()
		if decodeErr := generic.DecodeJSONString(encodedProps, &checkBoxProps); decodeErr != nil {
			return fmt.Errorf("GetCheckBoxes: can't decode properties: %v", encodedProps)
		}

		currCheckBox := CheckBox{
			ParentTableID: parentTableID,
			CheckBoxID:    checkboxID,
			Properties:    checkBoxProps}
		checkBoxes = append(checkBoxes, currCheckBox)

		return nil
	}
	if getErr := common.GetTableColumns(checkBoxEntityKind, parentTableID, addCheckbox); getErr != nil {
		return nil, fmt.Errorf("GetCheckBoxes: Can't get checkboxes: %v")
	}

	return checkBoxes, nil
}

func CloneCheckBoxes(remappedIDs uniqueID.UniqueIDRemapper, parentTableID string) error {

	srcCheckBoxes, err := GetCheckBoxes(parentTableID)
	if err != nil {
		return fmt.Errorf("CloneCheckBoxes: %v", err)
	}

	for _, srcCheckBox := range srcCheckBoxes {
		remappedCheckBoxID := remappedIDs.AllocNewOrGetExistingRemappedID(srcCheckBox.CheckBoxID)
		remappedFormID, err := remappedIDs.GetExistingRemappedID(srcCheckBox.ParentTableID)
		if err != nil {
			return fmt.Errorf("CloneCheckBoxes: %v", err)
		}
		destProperties, err := srcCheckBox.Properties.Clone(remappedIDs)
		if err != nil {
			return fmt.Errorf("CloneCheckBoxes: %v", err)
		}
		destCheckBox := CheckBox{
			ParentTableID: remappedFormID,
			CheckBoxID:    remappedCheckBoxID,
			Properties:    *destProperties}
		if err := saveCheckbox(destCheckBox); err != nil {
			return fmt.Errorf("CloneCheckBoxes: %v", err)
		}
	}

	return nil
}

func updateExistingCheckBox(updatedCheckBox *CheckBox) (*CheckBox, error) {

	if updateErr := common.UpdateTableColumn(checkBoxEntityKind, updatedCheckBox.ParentTableID,
		updatedCheckBox.CheckBoxID, updatedCheckBox.Properties); updateErr != nil {
		return nil, fmt.Errorf("updateExistingCheckBox: failure updating checkbox: %v", updateErr)
	}
	return updatedCheckBox, nil

}
