package table

import (
	"appengine"
	"resultra/datasheet/server/dataModel"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/datastoreWrapper"
)

var tableChildParentEntityRel = datastoreWrapper.ChildParentEntityRel{
	ParentEntityKind: dataModel.DatabaseEntityKind,
	ChildEntityKind:  dataModel.TableEntityKind}

type Table struct {
	Name string
}

type NewTableParams struct {
	ParentDatabaseID datastoreWrapper.UniqueRootID `json:'parentDatabaseID'`
	Name             string                        `json:"name"`
}

type TableRef struct {
	datastoreWrapper.UniqueIDHeader
	Name string `json:"name"`
}

func saveNewTable(appEngContext appengine.Context, params NewTableParams) (*TableRef, error) {

	sanitizedTableName, sanitizeErr := generic.SanitizeName(params.Name)
	if sanitizeErr != nil {
		return nil, sanitizeErr
	}

	// TODO - Validate name is unique

	newTable := Table{Name: sanitizedTableName}

	tableID, insertErr := datastoreWrapper.InsertNewChildEntity(
		appEngContext, params.ParentDatabaseID.ObjectID, tableChildParentEntityRel, &newTable)
	if insertErr != nil {
		return nil, insertErr
	}

	tableRef := TableRef{
		UniqueIDHeader: datastoreWrapper.NewUniqueIDHeader(params.ParentDatabaseID.ObjectID, tableID),
		Name:           sanitizedTableName}

	return &tableRef, nil
}
