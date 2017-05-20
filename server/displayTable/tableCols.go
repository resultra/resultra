package displayTable

import (
	"fmt"
	"resultra/datasheet/server/displayTable/columns/numberInput"
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

	return tableColData, nil
}
