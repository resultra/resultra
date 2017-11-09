package displayTable

import (
	"database/sql"
	"fmt"
)

type TableDisplayInfo struct {
	Table DisplayTable  `json:"table"`
	Cols  TableColsInfo `json:"cols"`
}

func getTableDisplayInfo(trackerDBHandle *sql.DB, tableID string) (*TableDisplayInfo, error) {

	table, err := GetTable(trackerDBHandle, tableID)
	if err != nil {
		return nil, fmt.Errorf("getTableDisplayInfo: %v", err)
	}

	_, colsByID, err := getTableCols(trackerDBHandle, tableID)
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
