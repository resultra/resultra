package table

import (
	"appengine"
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
