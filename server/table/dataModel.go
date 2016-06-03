package table

import (
	"appengine"
	"fmt"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/datastoreWrapper"
	"resultra/datasheet/server/generic/uniqueID"
)

const TableEntityKind string = "Table"

type Table struct {
	ParentDatabaseID string `json:"parentDatabaseID"`
	TableID          string `json:"tableID"`
	Name             string `json:"name"`
}

const tableParentIDFieldName string = "ParentDatabaseID"

type NewTableParams struct {
	DatabaseID string `json:'databaseID'`
	Name       string `json:"name"`
}

func saveNewTable(appEngContext appengine.Context, params NewTableParams) (*Table, error) {

	sanitizedTableName, sanitizeErr := generic.SanitizeName(params.Name)
	if sanitizeErr != nil {
		return nil, sanitizeErr
	}

	// TODO - Validate name is unique

	if err := uniqueID.ValidatedWellFormedID(params.DatabaseID); err != nil {
		return nil, err
	}

	newTable := Table{ParentDatabaseID: params.DatabaseID, TableID: uniqueID.GenerateUniqueID(), Name: sanitizedTableName}

	insertErr := datastoreWrapper.InsertNewRootEntity(appEngContext, TableEntityKind, &newTable)
	if insertErr != nil {
		return nil, insertErr
	}

	return &newTable, nil
}

type GetTableListParams struct {
	DatabaseID string `json:"databaseID"` // parent database
}

func getTableList(appEngContext appengine.Context, params GetTableListParams) ([]Table, error) {

	var tables []Table

	getErr := datastoreWrapper.GetAllChildEntitiesWithParentUUID(appEngContext, params.DatabaseID,
		TableEntityKind, tableParentIDFieldName, &tables)

	if getErr != nil {
		return nil, fmt.Errorf("GetTableList: Unable to retrieve tables from datastore: datastore error = %v", getErr)
	}

	return tables, nil

}
