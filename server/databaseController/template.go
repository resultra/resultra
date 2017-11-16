package databaseController

import (
	"database/sql"
	"fmt"
	"net/http"

	"resultra/datasheet/server/alert"
	"resultra/datasheet/server/calcField"
	"resultra/datasheet/server/common/databaseWrapper"
	"resultra/datasheet/server/dashboard"
	"resultra/datasheet/server/displayTable"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/form"
	"resultra/datasheet/server/formLink"
	"resultra/datasheet/server/generic/uniqueID"
	"resultra/datasheet/server/generic/userAuth"
	"resultra/datasheet/server/global"
	"resultra/datasheet/server/itemList"
	"resultra/datasheet/server/trackerDatabase"
	"resultra/datasheet/server/userRole"
	"resultra/datasheet/server/valueList"
)

func cloneFields(cloneParams *trackerDatabase.CloneDatabaseParams) error {

	getFieldParams := field.GetFieldListParams{ParentDatabaseID: cloneParams.SourceDatabaseID}
	fields, err := field.GetAllFieldsFromSrc(cloneParams.SrcDBHandle, getFieldParams)
	if err != nil {
		return fmt.Errorf("cloneFields: %v", err)
	}

	// Since calculated fields can reference other fields by ID, cloning the fields
	// requires a 2-pass algorithm to first remap just the field IDs, then clnoe the
	// the fields themselves with the remapped IDs already in place.
	for _, currField := range fields {
		_, err := cloneParams.IDRemapper.AllocNewRemappedID(currField.FieldID)
		if err != nil {
			return fmt.Errorf("cloneFields: Duplicate mapping for field ID = %v (err=%v)",
				currField.FieldID, err)
		}
	}

	remappedDatabaseID, err := cloneParams.IDRemapper.GetExistingRemappedID(cloneParams.SourceDatabaseID)
	if err != nil {
		return fmt.Errorf("cloneFields: %v", err)
	}

	for _, currField := range fields {

		clonedField := currField
		clonedField.ParentDatabaseID = remappedDatabaseID

		// There's no guarantee regarding the order of fields IDs being re-mapped.
		// So, the re-mapped field ID just needs to be remapped if it isn't already created.
		remappedFieldID := cloneParams.IDRemapper.AllocNewOrGetExistingRemappedID(currField.FieldID)
		clonedField.FieldID = remappedFieldID

		if currField.IsCalcField {
			clonedEqn, err := calcField.CloneEquation(cloneParams.IDRemapper, currField.CalcFieldEqn)
			if err != nil {
				return fmt.Errorf("cloneFields: %v", err)
			}
			clonedField.CalcFieldEqn = clonedEqn

			clonedFormulaText, err := calcField.ClonePreprocessedFormula(cloneParams.SrcDBHandle, cloneParams.SourceDatabaseID,
				cloneParams.IDRemapper, currField.PreprocessedFormulaText)
			if err != nil {
				return fmt.Errorf("cloneFields: %v", err)
			}
			clonedField.PreprocessedFormulaText = clonedFormulaText

			if _, err := field.CreateNewFieldFromRawInputs(cloneParams.DestDBHandle, clonedField); err != nil {
				return fmt.Errorf("cloneFields: failure saving cloned field: %v", err)
			}

		} else {
			if _, err := field.CreateNewFieldFromRawInputs(cloneParams.DestDBHandle, clonedField); err != nil {
				return fmt.Errorf("cloneFields: failure saving cloned field: %v", err)
			}
		}

	}

	return nil

}

func cloneIntoNewTrackerDatabase(cloneParams *trackerDatabase.CloneDatabaseParams) (*trackerDatabase.Database, error) {

	clonedDB, err := trackerDatabase.CloneDatabase(cloneParams)
	if err != nil {
		return nil, fmt.Errorf("copyDatabaseToTemplate: %v", err)
	}

	if err := global.CloneGlobals(cloneParams); err != nil {
		return nil, fmt.Errorf("copyDatabaseToTemplate: %v", err)
	}

	if err := cloneFields(cloneParams); err != nil {
		return nil, fmt.Errorf("copyDatabaseToTemplate: %v", err)
	}

	if err := form.CloneForms(cloneParams); err != nil {
		return nil, fmt.Errorf("copyDatabaseToTemplate: %v", err)
	}

	if err := dashboard.CloneDashboards(cloneParams); err != nil {
		return nil, fmt.Errorf("copyDatabaseToTemplate: %v", err)
	}

	if err := displayTable.CloneTables(cloneParams); err != nil {
		return nil, fmt.Errorf("copyDatabaseToTemplate: %v", err)
	}

	// Item lists have a form as a property, so they must be cloned after the forms, ensuring
	// the form IDs have already been remapped.
	if err := itemList.CloneItemLists(cloneParams); err != nil {
		return nil, fmt.Errorf("copyDatabaseToTemplate: %v", err)
	}

	if err := formLink.CloneFormLinks(cloneParams); err != nil {
		return nil, fmt.Errorf("copyDatabaseToTemplate: %v", err)
	}

	if err := valueList.CloneValueLists(cloneParams); err != nil {
		return nil, fmt.Errorf("copyDatabaseToTemplate: %v", err)
	}

	if err := alert.CloneAlerts(cloneParams); err != nil {
		return nil, fmt.Errorf("copyDatabaseToTemplate: %v", err)
	}

	if err := userRole.CloneRoles(cloneParams); err != nil {
		return nil, fmt.Errorf("copyDatabaseToTemplate: %v", err)
	}
	if err := userRole.CloneListPrivs(cloneParams); err != nil {
		return nil, fmt.Errorf("copyDatabaseToTemplate: %v", err)
	}
	if err := userRole.CloneDashboardPrivs(cloneParams); err != nil {
		return nil, fmt.Errorf("copyDatabaseToTemplate: %v", err)
	}
	if err := userRole.CloneNewItemLinkPrivs(cloneParams); err != nil {
		return nil, fmt.Errorf("copyDatabaseToTemplate: %v", err)
	}
	if err := userRole.CloneAlertPrivs(cloneParams); err != nil {
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

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(req)
	if dbErr != nil {
		return nil, dbErr
	}

	cloneParams := trackerDatabase.CloneDatabaseParams{
		SourceDatabaseID: params.SourceDatabaseID,
		NewName:          params.NewTemplateName,
		IsTemplate:       true,
		CreatedByUserID:  userID,
		SrcDBHandle:      trackerDBHandle,
		DestDBHandle:     trackerDBHandle,
		IDRemapper:       uniqueID.UniqueIDRemapper{}}
	return cloneIntoNewTrackerDatabase(&cloneParams)

}

type UserTemplateTrackerDatabaseInfo struct {
	DatabaseID   string `json:"databaseID"`
	DatabaseName string `json:"databaseName"`
	Description  string `json:"description"`
}

func getCurrentUserTemplateTrackers(trackerDBHandle *sql.DB, req *http.Request) ([]UserTemplateTrackerDatabaseInfo, error) {

	templateInfo := []UserTemplateTrackerDatabaseInfo{}

	rows, queryErr := trackerDBHandle.Query(
		`SELECT databases.database_id, databases.name, databases.description FROM databases WHERE 
			is_template='1'`)
	if queryErr != nil {
		return nil, fmt.Errorf("getCurrentUserTrackingDatabases: Failure querying database: %v", queryErr)
	}

	for rows.Next() {

		var desc sql.NullString
		var currTemplateInfo UserTemplateTrackerDatabaseInfo
		if scanErr := rows.Scan(&currTemplateInfo.DatabaseID,
			&currTemplateInfo.DatabaseName,
			&desc); scanErr != nil {
			return nil, fmt.Errorf("getCurrentUserTemplateTrackers: Failure querying database: %v", scanErr)
		}
		if desc.Valid {
			currTemplateInfo.Description = desc.String
		}

		templateInfo = append(templateInfo, currTemplateInfo)
	}

	return templateInfo, nil

}
