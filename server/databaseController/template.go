package databaseController

import (
	"fmt"
	"net/http"
	"resultra/datasheet/server/alert"
	"resultra/datasheet/server/calcField"
	"resultra/datasheet/server/displayTable"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/form"
	"resultra/datasheet/server/formLink"
	"resultra/datasheet/server/generic/uniqueID"
	"resultra/datasheet/server/generic/userAuth"
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

func cloneIntoNewTrackerDatabase(cloneParams trackerDatabase.CloneDatabaseParams) (*trackerDatabase.Database, error) {

	remappedIDs := uniqueID.UniqueIDRemapper{}

	clonedDB, err := trackerDatabase.CloneDatabase(remappedIDs, cloneParams)
	if err != nil {
		return nil, fmt.Errorf("copyDatabaseToTemplate: %v", err)
	}

	if err := global.CloneGlobals(remappedIDs, cloneParams.SourceDatabaseID); err != nil {
		return nil, fmt.Errorf("copyDatabaseToTemplate: %v", err)
	}

	if err := cloneFields(remappedIDs, cloneParams.SourceDatabaseID); err != nil {
		return nil, fmt.Errorf("copyDatabaseToTemplate: %v", err)
	}

	if err := form.CloneForms(remappedIDs, cloneParams.SourceDatabaseID); err != nil {
		return nil, fmt.Errorf("copyDatabaseToTemplate: %v", err)
	}

	if err := displayTable.CloneTables(remappedIDs, cloneParams.SourceDatabaseID); err != nil {
		return nil, fmt.Errorf("copyDatabaseToTemplate: %v", err)
	}

	// Item lists have a form as a property, so they must be cloned after the forms, ensuring
	// the form IDs have already been remapped.
	if err := itemList.CloneItemLists(remappedIDs, cloneParams.SourceDatabaseID); err != nil {
		return nil, fmt.Errorf("copyDatabaseToTemplate: %v", err)
	}

	if err := formLink.CloneFormLinks(remappedIDs, cloneParams.SourceDatabaseID); err != nil {
		return nil, fmt.Errorf("copyDatabaseToTemplate: %v", err)
	}

	if err := valueList.CloneValueLists(remappedIDs, cloneParams.SourceDatabaseID); err != nil {
		return nil, fmt.Errorf("copyDatabaseToTemplate: %v", err)
	}

	if err := alert.CloneAlerts(remappedIDs, cloneParams.SourceDatabaseID); err != nil {
		return nil, fmt.Errorf("copyDatabaseToTemplate: %v", err)
	}

	return clonedDB, nil

}

type SaveAsTemplateParams struct {
	SourceDatabaseID string `json:"sourceDatabaseID"`
	NewTemplateName  string `json:"newTemplateName"`
}

func saveExistingDatabaseAsTemplate(req *http.Request, params SaveAsTemplateParams) (*trackerDatabase.Database, error) {
	userID, userErr := userAuth.GetCurrentUserID(req)
	if userErr != nil {
		return nil, userErr
	}
	cloneParams := trackerDatabase.CloneDatabaseParams{
		SourceDatabaseID: params.SourceDatabaseID,
		NewName:          params.NewTemplateName,
		IsTemplate:       true,
		CreatedByUserID:  userID}
	return cloneIntoNewTrackerDatabase(cloneParams)

}
