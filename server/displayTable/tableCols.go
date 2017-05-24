package displayTable

import (
	"fmt"
	"resultra/datasheet/server/displayTable/columns/datePicker"
	"resultra/datasheet/server/displayTable/columns/numberInput"
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

	return tableColData, nil
}
