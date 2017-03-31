package numberInput

import (
	"fmt"
	"log"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/form/components/common"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/uniqueID"
)

const numberInputEntityKind string = "numberInput"

type NumberInput struct {
	ParentFormID  string                `json:"parentFormID"`
	NumberInputID string                `json:"numberInputID"`
	Properties    NumberInputProperties `json:"properties"`
}

type NewNumberInputParams struct {
	ParentFormID string                         `json:"parentFormID"`
	FieldID      string                         `json:"fieldID"`
	Geometry     componentLayout.LayoutGeometry `json:"geometry"`
}

func validNumberInputFieldType(fieldType string) bool {
	if fieldType == field.FieldTypeText {
		return true
	} else if fieldType == field.FieldTypeNumber {
		return true
	} else {
		return false
	}
}

func saveNumberInput(newNumberInput NumberInput) error {
	if saveErr := common.SaveNewFormComponent(numberInputEntityKind,
		newNumberInput.ParentFormID, newNumberInput.NumberInputID, newNumberInput.Properties); saveErr != nil {
		return fmt.Errorf("saveNumberInput: Unable to save text box: %v", saveErr)
	}
	return nil

}

func saveNewNumberInput(params NewNumberInputParams) (*NumberInput, error) {

	if !componentLayout.ValidGeometry(params.Geometry) {
		return nil, fmt.Errorf("Invalid layout container parameters: %+v", params)
	}

	if fieldErr := common.ValidateField(params.FieldID, validNumberInputFieldType); fieldErr != nil {
		return nil, fmt.Errorf("saveNewNumberInput: %v", fieldErr)
	}

	properties := newDefaultNumberInputProperties()
	properties.Geometry = params.Geometry
	properties.FieldID = params.FieldID

	newNumberInput := NumberInput{ParentFormID: params.ParentFormID,
		NumberInputID: uniqueID.GenerateSnowflakeID(),
		Properties:    properties}

	if err := saveNumberInput(newNumberInput); err != nil {
		return nil, fmt.Errorf("saveNewNumberInput: Unable to save text box with params=%+v: error = %v", params, err)
	}

	log.Printf("INFO: API: NewLayout: Created new Layout container: %+v", newNumberInput)

	return &newNumberInput, nil

}

func getNumberInput(parentFormID string, numberInputID string) (*NumberInput, error) {

	numberInputProps := newDefaultNumberInputProperties()
	if getErr := common.GetFormComponent(numberInputEntityKind, parentFormID, numberInputID, &numberInputProps); getErr != nil {
		return nil, fmt.Errorf("getCheckBox: Unable to retrieve text box: %v", getErr)
	}

	numberInput := NumberInput{
		ParentFormID:  parentFormID,
		NumberInputID: numberInputID,
		Properties:    numberInputProps}

	return &numberInput, nil
}

func GetNumberInputs(parentFormID string) ([]NumberInput, error) {

	numberInputs := []NumberInput{}
	addNumberInput := func(numberInputID string, encodedProps string) error {

		numberInputProps := newDefaultNumberInputProperties()
		if decodeErr := generic.DecodeJSONString(encodedProps, &numberInputProps); decodeErr != nil {
			return fmt.Errorf("GetNumberInputes: can't decode properties: %v", encodedProps)
		}

		currNumberInput := NumberInput{
			ParentFormID:  parentFormID,
			NumberInputID: numberInputID,
			Properties:    numberInputProps}
		numberInputs = append(numberInputs, currNumberInput)

		return nil
	}
	if getErr := common.GetFormComponents(numberInputEntityKind, parentFormID, addNumberInput); getErr != nil {
		return nil, fmt.Errorf("GetNumberInputs: Can't get number inputs: %v")
	}

	return numberInputs, nil

}

func CloneNumberInputs(remappedIDs uniqueID.UniqueIDRemapper, parentFormID string) error {

	srcNumberInputs, err := GetNumberInputs(parentFormID)
	if err != nil {
		return fmt.Errorf("CloneNumberInputs: %v", err)
	}

	for _, srcNumberInput := range srcNumberInputs {
		remappedNumberInputID := remappedIDs.AllocNewOrGetExistingRemappedID(srcNumberInput.NumberInputID)
		remappedFormID, err := remappedIDs.GetExistingRemappedID(srcNumberInput.ParentFormID)
		if err != nil {
			return fmt.Errorf("CloneNumberInputs: %v", err)
		}
		destProperties, err := srcNumberInput.Properties.Clone(remappedIDs)
		if err != nil {
			return fmt.Errorf("CloneNumberInputs: %v", err)
		}
		destNumberInput := NumberInput{
			ParentFormID:  remappedFormID,
			NumberInputID: remappedNumberInputID,
			Properties:    *destProperties}
		if err := saveNumberInput(destNumberInput); err != nil {
			return fmt.Errorf("CloneNumberInputs: %v", err)
		}
	}

	return nil
}

func updateExistingNumberInput(numberInputID string, updatedNumberInput *NumberInput) (*NumberInput, error) {

	if updateErr := common.UpdateFormComponent(numberInputEntityKind, updatedNumberInput.ParentFormID,
		updatedNumberInput.NumberInputID, updatedNumberInput.Properties); updateErr != nil {
		return nil, fmt.Errorf("updateExistingNumberInput: error updating existing number input component: %v", updateErr)
	}

	return updatedNumberInput, nil

}
