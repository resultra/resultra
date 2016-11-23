package databaseController

import (
	"fmt"
	"resultra/datasheet/server/database"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/generic/uniqueID"
	"resultra/datasheet/server/table"
)

func cloneFields(remappedIDs map[string]string, srcDatabaseID string) error {

	getTableParams := table.GetTableListParams{DatabaseID: srcDatabaseID}
	tables, err := table.GetTableList(getTableParams)
	if err != nil {
		return fmt.Errorf("cloneFields: %v", err)
	}

	for _, currTable := range tables {

		remappedTableID, foundTableID := remappedIDs[currTable.TableID]
		if !foundTableID {
			return fmt.Errorf("cloneFields: Can't find remapped table ID for table = %v", currTable.TableID)
		}

		getFieldParams := field.GetFieldListParams{ParentTableID: currTable.TableID}
		fields, err := field.GetAllFields(getFieldParams)
		if err != nil {
			return fmt.Errorf("cloneFields: %v", err)
		}

		// Since calculated fields can reference other fields by ID, cloning the fields
		// requires a 2-pass algorithm to first remap just the field IDs, then clnoe the
		// the fields themselves with the remapped IDs already in place.
		for _, currField := range fields {
			remappedIDs[currField.FieldID] = uniqueID.GenerateSnowflakeID()
		}

		for _, currField := range fields {

			clonedField := currField
			clonedField.ParentTableID = remappedTableID
			clonedField.FieldID = remappedIDs[currField.FieldID]

			if currField.IsCalcField {

			} else {
				field.CreateNewFieldFromRawInputs(clonedField)
			}

		}

	}

	return nil

}

type SaveTemplateParams struct {
	SourceDatabaseID string `json:"sourceDatabaseID"`
	NewTemplateName  string `json:"newTemplateName"`
}

func saveDatabaseToTemplate(params SaveTemplateParams) (*database.Database, error) {

	remappedIDs := map[string]string{}

	templateDB, err := database.CloneDatabase(remappedIDs, params.NewTemplateName, params.SourceDatabaseID)
	if err != nil {
		return nil, fmt.Errorf("copyDatabaseToTemplate: %v", err)
	}

	if err := table.CloneTables(remappedIDs, params.SourceDatabaseID); err != nil {
		return nil, fmt.Errorf("copyDatabaseToTemplate: %v", err)
	}

	if err := cloneFields(remappedIDs, params.SourceDatabaseID); err != nil {
		return nil, fmt.Errorf("copyDatabaseToTemplate: %v", err)
	}

	return templateDB, nil

}
