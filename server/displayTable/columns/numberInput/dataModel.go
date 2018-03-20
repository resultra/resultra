package numberInput

import (
	"database/sql"
	"fmt"
	"log"
	"resultra/datasheet/server/displayTable/columns/common"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/uniqueID"
	"resultra/datasheet/server/trackerDatabase"
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

func saveNumberInput(destDBHandle *sql.DB, newNumberInput NumberInput) error {
	if saveErr := common.SaveNewTableColumn(destDBHandle, numberInputEntityKind,
		newNumberInput.ParentTableID,
		newNumberInput.NumberInputID, newNumberInput.Properties); saveErr != nil {
		return fmt.Errorf("saveNumberInput: Unable to save number input: %v", saveErr)
	}
	return nil

}

func saveNewNumberInput(trackerDBHandle *sql.DB, params NewNumberInputParams) (*NumberInput, error) {

	if fieldErr := field.ValidateField(trackerDBHandle, params.FieldID, validNumberInputFieldType); fieldErr != nil {
		return nil, fmt.Errorf("saveNewNumberInput: %v", fieldErr)
	}

	properties := newDefaultNumberInputProperties()
	properties.FieldID = params.FieldID

	numberInputID := uniqueID.GenerateUniqueID()
	newNumberInput := NumberInput{ParentTableID: params.ParentTableID,
		NumberInputID: numberInputID,
		ColumnID:      numberInputID,
		Properties:    properties}

	if err := saveNumberInput(trackerDBHandle, newNumberInput); err != nil {
		return nil, fmt.Errorf("saveNewNumberInput: Unable to save text box with params=%+v: error = %v", params, err)
	}

	log.Printf("INFO: API: NewLayout: Created new Layout container: %+v", newNumberInput)

	return &newNumberInput, nil

}

func getNumberInput(trackerDBHandle *sql.DB, parentTableID string, numberInputID string) (*NumberInput, error) {

	numberInputProps := newDefaultNumberInputProperties()
	if getErr := common.GetTableColumn(trackerDBHandle, numberInputEntityKind, parentTableID, numberInputID, &numberInputProps); getErr != nil {
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

func getNumberInputsFromSrc(srcDBHandle *sql.DB, parentTableID string) ([]NumberInput, error) {

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
	if getErr := common.GetTableColumns(srcDBHandle, numberInputEntityKind, parentTableID, addNumberInput); getErr != nil {
		return nil, fmt.Errorf("GetNumberInputs: Can't get number inputs: %v")
	}

	return numberInputs, nil

}

func GetNumberInputs(trackerDBHandle *sql.DB, parentTableID string) ([]NumberInput, error) {
	return getNumberInputsFromSrc(trackerDBHandle, parentTableID)
}

func CloneNumberInputs(cloneParams *trackerDatabase.CloneDatabaseParams, parentTableID string) error {

	srcNumberInputs, err := getNumberInputsFromSrc(cloneParams.SrcDBHandle, parentTableID)
	if err != nil {
		return fmt.Errorf("CloneNumberInputs: %v", err)
	}

	for _, srcNumberInput := range srcNumberInputs {
		remappedNumberInputID := cloneParams.IDRemapper.AllocNewOrGetExistingRemappedID(srcNumberInput.NumberInputID)
		remappedTableID, err := cloneParams.IDRemapper.GetExistingRemappedID(srcNumberInput.ParentTableID)
		if err != nil {
			return fmt.Errorf("CloneNumberInputs: %v", err)
		}
		destProperties, err := srcNumberInput.Properties.Clone(cloneParams)
		if err != nil {
			return fmt.Errorf("CloneNumberInputs: %v", err)
		}
		destNumberInput := NumberInput{
			ParentTableID: remappedTableID,
			NumberInputID: remappedNumberInputID,
			ColumnID:      remappedNumberInputID,
			ColType:       numberInputEntityKind,
			Properties:    *destProperties}
		if err := saveNumberInput(cloneParams.DestDBHandle, destNumberInput); err != nil {
			return fmt.Errorf("CloneNumberInputs: %v", err)
		}
	}

	return nil
}

func updateExistingNumberInput(trackerDBHandle *sql.DB, numberInputID string, updatedNumberInput *NumberInput) (*NumberInput, error) {

	if updateErr := common.UpdateTableColumn(trackerDBHandle, numberInputEntityKind, updatedNumberInput.ParentTableID,
		updatedNumberInput.NumberInputID, updatedNumberInput.Properties); updateErr != nil {
		return nil, fmt.Errorf("updateExistingNumberInput: error updating existing number input component: %v", updateErr)
	}

	return updatedNumberInput, nil

}
