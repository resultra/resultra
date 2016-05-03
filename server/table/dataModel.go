package table

import (
	"appengine"
	"fmt"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/datastoreWrapper"
)

const TableEntityKind string = "Table"

type Table struct {
	Name string
}

type NewTableParams struct {
	DatabaseID string `json:'databaseID'`
	Name       string `json:"name"`
}

type TableRef struct {
	TableID string `json:"tableID"`
	Name    string `json:"name"`
}

func saveNewTable(appEngContext appengine.Context, params NewTableParams) (*TableRef, error) {

	sanitizedTableName, sanitizeErr := generic.SanitizeName(params.Name)
	if sanitizeErr != nil {
		return nil, sanitizeErr
	}

	// TODO - Validate name is unique

	newTable := Table{Name: sanitizedTableName}

	tableID, insertErr := datastoreWrapper.InsertNewChildEntity(
		appEngContext, params.DatabaseID, TableEntityKind, &newTable)
	if insertErr != nil {
		return nil, insertErr
	}

	tableRef := TableRef{
		TableID: tableID,
		Name:    sanitizedTableName}

	return &tableRef, nil
}

type GetTableListParams struct {
	DatabaseID string `json:"databaseID"` // parent database
}

func getTableList(appEngContext appengine.Context, params GetTableListParams) ([]TableRef, error) {

	var tables []Table
	tablesIDs, getErr := datastoreWrapper.GetAllChildEntities(appEngContext, params.DatabaseID, TableEntityKind, &tables)
	if getErr != nil {
		return nil, fmt.Errorf("GetTableList: Unable to retrieve tables from datastore: datastore error = %v", getErr)
	}

	tableRefs := make([]TableRef, len(tables))
	for tableIter, currTable := range tables {
		tableID := tablesIDs[tableIter]
		tableRefs[tableIter] = TableRef{TableID: tableID, Name: currTable.Name}
	}
	return tableRefs, nil

}
