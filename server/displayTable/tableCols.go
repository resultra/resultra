package displayTable

import (
	"fmt"
	"resultra/datasheet/server/displayTable/columns/checkBox"
	"resultra/datasheet/server/displayTable/columns/datePicker"
	"resultra/datasheet/server/displayTable/columns/numberInput"
	"resultra/datasheet/server/displayTable/columns/rating"
	"resultra/datasheet/server/displayTable/columns/textInput"
)

type TableColsInfo []interface{}

func getTableCols(parentTableID string) (TableColsInfo, error) {

	tableColData := TableColsInfo{}

	numberInputCols, err := numberInput.GetNumberInputs(parentTableID)
	if err != nil {
		return nil, fmt.Errorf("getTableCols: %v", err)
	}
	for _, col := range numberInputCols {
		tableColData = append(tableColData, col)
	}

	textInputCols, err := textInput.GetTextInputs(parentTableID)
	if err != nil {
		return nil, fmt.Errorf("getTableCols: %v", err)
	}
	for _, col := range textInputCols {
		tableColData = append(tableColData, col)
	}

	datePickerCols, err := datePicker.GetDatePickers(parentTableID)
	if err != nil {
		return nil, fmt.Errorf("getTableCols: %v", err)
	}
	for _, col := range datePickerCols {
		tableColData = append(tableColData, col)
	}

	checkBoxCols, err := checkBox.GetCheckBoxes(parentTableID)
	if err != nil {
		return nil, fmt.Errorf("getTableCols: %v", err)
	}
	for _, col := range checkBoxCols {
		tableColData = append(tableColData, col)
	}

	ratingCols, err := rating.GetRatings(parentTableID)
	if err != nil {
		return nil, fmt.Errorf("getTableCols: %v", err)
	}
	for _, col := range ratingCols {
		tableColData = append(tableColData, col)
	}

	return tableColData, nil
}
