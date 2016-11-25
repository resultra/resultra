package databaseController

import (
	"fmt"
	"resultra/datasheet/server/calcField"
	"resultra/datasheet/server/database"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/form"
	"resultra/datasheet/server/generic/uniqueID"
	"resultra/datasheet/server/global"
	"resultra/datasheet/server/table"
)

func cloneFields(remappedIDs uniqueID.UniqueIDRemapper, srcDatabaseID string) error {

	getTableParams := table.GetTableListParams{DatabaseID: srcDatabaseID}
	tables, err := table.GetTableList(getTableParams)
	if err != nil {
		return fmt.Errorf("cloneFields: %v", err)
	}

	for _, currTable := range tables {

		remappedTableID, err := remappedIDs.GetExistingRemappedID(currTable.TableID)
		if err != nil {
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
			_, err := remappedIDs.AllocNewRemappedID(currField.FieldID)
			if err != nil {
				return fmt.Errorf("cloneFields: Duplicate mapping for field ID = %v (err=%v)",
					currField.FieldID, err)
			}
		}

		for _, currField := range fields {

			clonedField := currField
			clonedField.ParentTableID = remappedTableID

			remappedFieldID, err := remappedIDs.GetExistingRemappedID(currField.FieldID)
			if err != nil {
				return fmt.Errorf("cloneFields: Missing mapping for field ID = %v (err=%v)",
					currField.FieldID, err)
			}
			clonedField.FieldID = remappedFieldID

			if currField.IsCalcField {
				clonedEqn, err := calcField.CloneEquation(remappedIDs, currField.CalcFieldEqn)
				if err != nil {
					return fmt.Errorf("cloneFields: %v", err)
				}
				clonedField.CalcFieldEqn = clonedEqn

				clonedFormulaText, err := calcField.ClonePreprocessedFormula(srcDatabaseID, currTable.TableID,
					remappedIDs, currField.PreprocessedFormulaText)
				if err != nil {
					return fmt.Errorf("cloneFields: %v", err)
				}
				clonedField.PreprocessedFormulaText = clonedFormulaText

				if _, err := field.CreateNewFieldFromRawInputs(clonedField); err != nil {
					return fmt.Errorf("cloneFields: failure saving cloned field: %v", err)
				}

			} else {
				if _, err := field.CreateNewFieldFromRawInputs(clonedField); err != nil {
					return fmt.Errorf("cloneFields: failure saving cloned field: %v", err)
				}
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

	remappedIDs := uniqueID.UniqueIDRemapper{}

	templateDB, err := database.CloneDatabase(remappedIDs, params.NewTemplateName, params.SourceDatabaseID)
	if err != nil {
		return nil, fmt.Errorf("copyDatabaseToTemplate: %v", err)
	}

	if err := table.CloneTables(remappedIDs, params.SourceDatabaseID); err != nil {
		return nil, fmt.Errorf("copyDatabaseToTemplate: %v", err)
	}

	if err := global.CloneGlobals(remappedIDs, params.SourceDatabaseID); err != nil {
		return nil, fmt.Errorf("copyDatabaseToTemplate: %v", err)
	}

	if err := cloneFields(remappedIDs, params.SourceDatabaseID); err != nil {
		return nil, fmt.Errorf("copyDatabaseToTemplate: %v", err)
	}

	if err := form.CloneForms(remappedIDs, params.SourceDatabaseID); err != nil {
		return nil, fmt.Errorf("copyDatabaseToTemplate: %v", err)
	}

	return templateDB, nil

}
