package displayTable

import (
	"fmt"
)

type TableDisplayInfo struct {
	Table DisplayTable  `json:"table"`
	Cols  TableColsInfo `json:"cols"`
}

func getTableDisplayInfo(tableID string) (*TableDisplayInfo, error) {

	table, err := GetTable(tableID)
	if err != nil {
		return nil, fmt.Errorf("getTableDisplayInfo: %v", err)
	}

	_, colsByID, err := getTableCols(tableID)
	if err != nil {
		return nil, fmt.Errorf("getTableDisplayInfo: %v", err)
	}

	sortedCols := TableColsInfo{}

	for _, colID := range table.Properties.OrderedColumns {
		col, idFound := colsByID[colID]
		if idFound {
			sortedCols = append(sortedCols, col)
			delete(colsByID, colID)
		}
	}
	for _, col := range colsByID {
		sortedCols = append(sortedCols, col)
	}

	// TODO - Sort table columns based upon ordering in the table definition
	tableDisplayInfo := TableDisplayInfo{
		Table: *table,
		Cols:  sortedCols}

	return &tableDisplayInfo, nil

}
