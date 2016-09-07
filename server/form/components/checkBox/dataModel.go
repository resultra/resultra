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

type CheckBoxProperties struct {
	ComponentLink common.ComponentLink           `json:"componentLink"`
	Geometry      componentLayout.LayoutGeometry `json:"geometry"`
}

type CheckBox struct {
	ParentFormID string             `json:"parentID"`
	CheckBoxID   string             `json:"checkBoxID"`
	Properties   CheckBoxProperties `json:"properties"`
}

type NewCheckBoxParams struct {
	ParentFormID  string                         `json:"parentFormID"`
	ComponentLink common.ComponentLink           `json:"componentLink"`
	Geometry      componentLayout.LayoutGeometry `json:"geometry"`
}

func validCheckBoxFieldType(fieldType string) bool {
	if fieldType == field.FieldTypeBool {
		return true
	} else {
		return false
	}
}

func saveNewCheckBox(params NewCheckBoxParams) (*CheckBox, error) {

	if !componentLayout.ValidGeometry(params.Geometry) {
		return nil, fmt.Errorf("Invalid layout container parameters: %+v", params)
	}

	if compLinkErr := common.ValidateComponentLink(params.ComponentLink, validCheckBoxFieldType); compLinkErr != nil {
		return nil, fmt.Errorf("saveNewCheckBox: %v", compLinkErr)
	}

	properties := CheckBoxProperties{
		ComponentLink: params.ComponentLink,
		Geometry:      params.Geometry}

	newCheckBox := CheckBox{ParentFormID: params.ParentFormID,
		CheckBoxID: uniqueID.GenerateSnowflakeID(),
		Properties: properties}

	if saveErr := common.SaveNewFormComponent(checkBoxEntityKind,
		newCheckBox.ParentFormID, newCheckBox.CheckBoxID, newCheckBox.Properties); saveErr != nil {
		return nil, fmt.Errorf("saveNewCheckBox: Unable to save bar chart with params=%+v: error = %v", params, saveErr)
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

func updateExistingCheckBox(updatedCheckBox *CheckBox) (*CheckBox, error) {

	if updateErr := common.UpdateFormComponent(checkBoxEntityKind, updatedCheckBox.ParentFormID,
		updatedCheckBox.CheckBoxID, updatedCheckBox.Properties); updateErr != nil {
	}
	return updatedCheckBox, nil

}
