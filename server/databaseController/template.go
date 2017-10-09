package databaseController

import (
	"fmt"
	"resultra/datasheet/server/alert"
	"resultra/datasheet/server/calcField"
	"resultra/datasheet/server/displayTable"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/form"
	"resultra/datasheet/server/formLink"
	"resultra/datasheet/server/generic/uniqueID"
	"resultra/datasheet/server/global"
	"resultra/datasheet/server/itemList"
	"resultra/datasheet/server/trackerDatabase"
	"resultra/datasheet/server/valueList"
)

func cloneFields(remappedIDs uniqueID.UniqueIDRemapper, srcDatabaseID string) error {

	getFieldParams := field.GetFieldListParams{ParentDatabaseID: srcDatabaseID}
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

	remappedDatabaseID, err := remappedIDs.GetExistingRemappedID(srcDatabaseID)
	if err != nil {
		return fmt.Errorf("cloneFields: %v", err)
	}

	for _, currField := range fields {

		clonedField := currField
		clonedField.ParentDatabaseID = remappedDatabaseID

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

			clonedFormulaText, err := calcField.ClonePreprocessedFormula(srcDatabaseID,
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

	return nil

}

type SaveTemplateParams struct {
	SourceDatabaseID string `json:"sourceDatabaseID"`
	NewTemplateName  string `json:"newTemplateName"`
}

func saveDatabaseToTemplate(params SaveTemplateParams) (*trackerDatabase.Database, error) {

	remappedIDs := uniqueID.UniqueIDRemapper{}

	templateDB, err := trackerDatabase.CloneDatabase(remappedIDs, params.NewTemplateName, params.SourceDatabaseID)
	if err != nil {
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

	if err := displayTable.CloneTables(remappedIDs, params.SourceDatabaseID); err != nil {
		return nil, fmt.Errorf("copyDatabaseToTemplate: %v", err)
	}

	// Item lists have a form as a property, so they must be cloned after the forms, ensuring
	// the form IDs have already been remapped.
	if err := itemList.CloneItemLists(remappedIDs, params.SourceDatabaseID); err != nil {
		return nil, fmt.Errorf("copyDatabaseToTemplate: %v", err)
	}

	if err := formLink.CloneFormLinks(remappedIDs, params.SourceDatabaseID); err != nil {
		return nil, fmt.Errorf("copyDatabaseToTemplate: %v", err)
	}

	if err := valueList.CloneValueLists(remappedIDs, params.SourceDatabaseID); err != nil {
		return nil, fmt.Errorf("copyDatabaseToTemplate: %v", err)
	}

	if err := alert.CloneAlerts(remappedIDs, params.SourceDatabaseID); err != nil {
		return nil, fmt.Errorf("copyDatabaseToTemplate: %v", err)
	}

	return templateDB, nil

}
