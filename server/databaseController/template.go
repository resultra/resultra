// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package databaseController

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/resultra/resultra/server/alert"
	"github.com/resultra/resultra/server/calcField"
	"github.com/resultra/resultra/server/common/databaseWrapper"
	"github.com/resultra/resultra/server/common/userAuth"
	"github.com/resultra/resultra/server/dashboard"
	"github.com/resultra/resultra/server/displayTable"
	"github.com/resultra/resultra/server/field"
	"github.com/resultra/resultra/server/form"
	"github.com/resultra/resultra/server/formLink"
	"github.com/resultra/resultra/server/generic/uniqueID"
	"github.com/resultra/resultra/server/global"
	"github.com/resultra/resultra/server/itemList"
	"github.com/resultra/resultra/server/trackerDatabase"
	"github.com/resultra/resultra/server/userRole"
	"github.com/resultra/resultra/server/valueList"
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

func CloneIntoNewTrackerDatabase(cloneParams *trackerDatabase.CloneDatabaseParams) (*trackerDatabase.Database, error) {

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
	return CloneIntoNewTrackerDatabase(&cloneParams)

}

type GetTemplateListParams struct {
	IncludeInactive bool `json:"includeInactive"`
	CurrUserOnly    bool `json:"currUserOnly"`
}

type UserTemplateTrackerDatabaseInfo struct {
	DatabaseID      string `json:"databaseID"`
	DatabaseName    string `json:"databaseName"`
	Description     string `json:"description"`
	IsActive        bool   `json:"isActive"`
	CreatedByUserID string `json:"createdByUserID"`
}

func getCurrentUserTemplateTrackers(params GetTemplateListParams, trackerDBHandle *sql.DB, req *http.Request) ([]UserTemplateTrackerDatabaseInfo, error) {

	currUserID, userErr := userAuth.GetCurrentUserID(req)
	if userErr != nil {
		return nil, fmt.Errorf("getCurrentUserTemplateTrackers: can't get current user: %v", userErr)
	}

	templateInfo := []UserTemplateTrackerDatabaseInfo{}

	rows, queryErr := trackerDBHandle.Query(
		`SELECT database_id, name, description, is_active,created_by_user_id FROM databases WHERE 
			is_template='1'`)
	if queryErr != nil {
		return nil, fmt.Errorf("getCurrentUserTrackingDatabases: Failure querying database: %v", queryErr)
	}
	defer rows.Close()

	for rows.Next() {

		var desc sql.NullString
		var currTemplateInfo UserTemplateTrackerDatabaseInfo
		if scanErr := rows.Scan(&currTemplateInfo.DatabaseID,
			&currTemplateInfo.DatabaseName,
			&desc,
			&currTemplateInfo.IsActive,
			&currTemplateInfo.CreatedByUserID); scanErr != nil {
			return nil, fmt.Errorf("getCurrentUserTemplateTrackers: Failure querying database: %v", scanErr)
		}
		if desc.Valid {
			currTemplateInfo.Description = desc.String
		}

		includeActive := false
		if currTemplateInfo.IsActive == false {
			if params.IncludeInactive {
				includeActive = true
			}
		} else {
			includeActive = true
		}

		includeUser := false
		if params.CurrUserOnly == true {
			if currUserID == currTemplateInfo.CreatedByUserID {
				includeUser = true
			}
		} else {
			includeUser = true
		}

		if includeActive && includeUser {
			templateInfo = append(templateInfo, currTemplateInfo)
		}

	}

	return templateInfo, nil

}
