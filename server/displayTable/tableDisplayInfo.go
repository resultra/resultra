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

	cols, err := getTableCols(tableID)
	if err != nil {
		return nil, fmt.Errorf("getTableDisplayInfo: %v", err)
	}

	// TODO - Sort table columns based upon ordering in the table definition
	tableDisplayInfo := TableDisplayInfo{
		Table: *table,
		Cols:  cols}

	return &tableDisplayInfo, nil

}
