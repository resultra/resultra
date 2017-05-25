package numberInput

import (
	"fmt"
	"log"
	"resultra/datasheet/server/displayTable/columns/common"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/uniqueID"
)

const numberInputEntityKind string = "numberInput"

type NumberInput struct {
	ParentTableID string                `json:"parentTableID"`
	NumberInputID string                `json:"numberInputID"`
	ColType       string                `json:"colType"`
	ColumnID      string                `json:"columnID"`
	Properties    NumberInputProperties `json:"properties"`
}

type NewNumberInputParams struct {
	ParentTableID string `json:"parentTableID"`
	FieldID       string `json:"fieldID"`
}

func validNumberInputFieldType(fieldType string) bool {
	if fieldType == field.FieldTypeNumber {
		return true
	} else {
		return false
	}
}

func saveNumberInput(newNumberInput NumberInput) error {
	if saveErr := common.SaveNewTableColumn(numberInputEntityKind,
		newNumberInput.ParentTableID,
		newNumberInput.NumberInputID, newNumberInput.Properties); saveErr != nil {
		return fmt.Errorf("saveNumberInput: Unable to save number input: %v", saveErr)
	}
	return nil

}

func saveNewNumberInput(params NewNumberInputParams) (*NumberInput, error) {

	if fieldErr := field.ValidateField(params.FieldID, validNumberInputFieldType); fieldErr != nil {
		return nil, fmt.Errorf("saveNewNumberInput: %v", fieldErr)
	}

	properties := newDefaultNumberInputProperties()
	properties.FieldID = params.FieldID

	numberInputID := uniqueID.GenerateSnowflakeID()
	newNumberInput := NumberInput{ParentTableID: params.ParentTableID,
		NumberInputID: numberInputID,
		ColumnID:      numberInputID,
		Properties:    properties}

	if err := saveNumberInput(newNumberInput); err != nil {
		return nil, fmt.Errorf("saveNewNumberInput: Unable to save text box with params=%+v: error = %v", params, err)
	}

	log.Printf("INFO: API: NewLayout: Created new Layout container: %+v", newNumberInput)

	return &newNumberInput, nil

}

func getNumberInput(parentTableID string, numberInputID string) (*NumberInput, error) {

	numberInputProps := newDefaultNumberInputProperties()
	if getErr := common.GetTableColumn(numberInputEntityKind, parentTableID, numberInputID, &numberInputProps); getErr != nil {
		return nil, fmt.Errorf("getNumberInput: Unable to retrieve number input: %v", getErr)
	}

	numberInput := NumberInput{
		ParentTableID: parentTableID,
		NumberInputID: numberInputID,
		ColType:       numberInputEntityKind,
		ColumnID:      numberInputID,
		Properties:    numberInputProps}

	return &numberInput, nil
}

func GetNumberInputs(parentTableID string) ([]NumberInput, error) {

	numberInputs := []NumberInput{}
	addNumberInput := func(numberInputID string, encodedProps string) error {

		numberInputProps := newDefaultNumberInputProperties()
		if decodeErr := generic.DecodeJSONString(encodedProps, &numberInputProps); decodeErr != nil {
			return fmt.Errorf("GetNumberInputes: can't decode properties: %v", encodedProps)
		}

		currNumberInput := NumberInput{
			ParentTableID: parentTableID,
			NumberInputID: numberInputID,
			ColumnID:      numberInputID,
			ColType:       numberInputEntityKind,
			Properties:    numberInputProps}
		numberInputs = append(numberInputs, currNumberInput)

		return nil
	}
	if getErr := common.GetTableColumns(numberInputEntityKind, parentTableID, addNumberInput); getErr != nil {
		return nil, fmt.Errorf("GetNumberInputs: Can't get number inputs: %v")
	}

	return numberInputs, nil

}

func CloneNumberInputs(remappedIDs uniqueID.UniqueIDRemapper, parentTableID string) error {

	srcNumberInputs, err := GetNumberInputs(parentTableID)
	if err != nil {
		return fmt.Errorf("CloneNumberInputs: %v", err)
	}

	for _, srcNumberInput := range srcNumberInputs {
		remappedNumberInputID := remappedIDs.AllocNewOrGetExistingRemappedID(srcNumberInput.NumberInputID)
		remappedTableID, err := remappedIDs.GetExistingRemappedID(srcNumberInput.ParentTableID)
		if err != nil {
			return fmt.Errorf("CloneNumberInputs: %v", err)
		}
		destProperties, err := srcNumberInput.Properties.Clone(remappedIDs)
		if err != nil {
			return fmt.Errorf("CloneNumberInputs: %v", err)
		}
		destNumberInput := NumberInput{
			ParentTableID: remappedTableID,
			NumberInputID: remappedNumberInputID,
			ColumnID:      remappedNumberInputID,
			ColType:       numberInputEntityKind,
			Properties:    *destProperties}
		if err := saveNumberInput(destNumberInput); err != nil {
			return fmt.Errorf("CloneNumberInputs: %v", err)
		}
	}

	return nil
}

func updateExistingNumberInput(numberInputID string, updatedNumberInput *NumberInput) (*NumberInput, error) {

	if updateErr := common.UpdateTableColumn(numberInputEntityKind, updatedNumberInput.ParentTableID,
		updatedNumberInput.NumberInputID, updatedNumberInput.Properties); updateErr != nil {
		return nil, fmt.Errorf("updateExistingNumberInput: error updating existing number input component: %v", updateErr)
	}

	return updatedNumberInput, nil

}
