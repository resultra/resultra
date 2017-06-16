package displayTable

import (
	"fmt"
	"resultra/datasheet/server/displayTable/columns/checkBox"
	"resultra/datasheet/server/displayTable/columns/datePicker"
	"resultra/datasheet/server/displayTable/columns/numberInput"
	"resultra/datasheet/server/displayTable/columns/rating"
	"resultra/datasheet/server/displayTable/columns/textInput"
	"resultra/datasheet/server/displayTable/columns/toggle"
	"resultra/datasheet/server/displayTable/columns/userSelection"
)

type TableColsInfo []interface{}
type TableColsByID map[string]interface{}

func getTableCols(parentTableID string) (TableColsInfo, TableColsByID, error) {

	tableColData := TableColsInfo{}
	tableColsByID := TableColsByID{}

	numberInputCols, err := numberInput.GetNumberInputs(parentTableID)
	if err != nil {
		return nil, nil, fmt.Errorf("getTableCols: %v", err)
	}
	for _, col := range numberInputCols {
		tableColData = append(tableColData, col)
		tableColsByID[col.ColumnID] = col
	}

	textInputCols, err := textInput.GetTextInputs(parentTableID)
	if err != nil {
		return nil, nil, fmt.Errorf("getTableCols: %v", err)
	}
	for _, col := range textInputCols {
		tableColData = append(tableColData, col)
		tableColsByID[col.ColumnID] = col
	}

	datePickerCols, err := datePicker.GetDatePickers(parentTableID)
	if err != nil {
		return nil, nil, fmt.Errorf("getTableCols: %v", err)
	}
	for _, col := range datePickerCols {
		tableColData = append(tableColData, col)
		tableColsByID[col.ColumnID] = col
	}

	checkBoxCols, err := checkBox.GetCheckBoxes(parentTableID)
	if err != nil {
		return nil, nil, fmt.Errorf("getTableCols: %v", err)
	}
	for _, col := range checkBoxCols {
		tableColData = append(tableColData, col)
		tableColsByID[col.ColumnID] = col
	}

	ratingCols, err := rating.GetRatings(parentTableID)
	if err != nil {
		return nil, nil, fmt.Errorf("getTableCols: %v", err)
	}
	for _, col := range ratingCols {
		tableColData = append(tableColData, col)
		tableColsByID[col.ColumnID] = col
	}

	toggleCols, err := toggle.GetToggles(parentTableID)
	if err != nil {
		return nil, nil, fmt.Errorf("getTableCols: %v", err)
	}
	for _, col := range toggleCols {
		tableColData = append(tableColData, col)
		tableColsByID[col.ColumnID] = col
	}

	userSelectionCols, err := userSelection.GetUserSelections(parentTableID)
	if err != nil {
		return nil, nil, fmt.Errorf("getTableCols: %v", err)
	}
	for _, col := range userSelectionCols {
		tableColData = append(tableColData, col)
		tableColsByID[col.ColumnID] = col
	}

	return tableColData, tableColsByID, nil
}
