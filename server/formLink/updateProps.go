// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package formLink

import (
	"database/sql"
	"fmt"
	"resultra/tracker/server/form"
	"resultra/tracker/server/generic/stringValidation"
	"resultra/tracker/server/generic/uniqueID"
	"resultra/tracker/server/record"
)

type FormLinkIDInterface interface {
	getFormLinkID() string
}

type FormLinkIDHeader struct {
	FormLinkID string `json:"formLinkID"`
}

func (idHeader FormLinkIDHeader) getFormLinkID() string {
	return idHeader.FormLinkID
}

type FormLinkPropUpdater interface {
	FormLinkIDInterface
	updateProps(trackerDBHandle *sql.DB, button *FormLink) error
}

func updateFormLinkProps(trackerDBHandle *sql.DB, propUpdater FormLinkPropUpdater) (*FormLink, error) {

	// Retrieve the bar chart from the data store
	linkForUpdate, getErr := GetFormLink(trackerDBHandle, propUpdater.getFormLinkID())
	if getErr != nil {
		return nil, fmt.Errorf("updateFormLinkProps: Unable to get existing button: %v", getErr)
	}

	if propUpdateErr := propUpdater.updateProps(trackerDBHandle, linkForUpdate); propUpdateErr != nil {
		return nil, fmt.Errorf("updateFormLinkProps: Unable to update existing form link properties: %v", propUpdateErr)
	}

	updatedLink, updateErr := updateExistingFormLink(trackerDBHandle, linkForUpdate)
	if updateErr != nil {
		return nil, fmt.Errorf(
			"updateFormLinkProps: Unable to update existing form link properties: datastore update error =  %v", updateErr)
	}

	return updatedLink, nil
}

type FormLinkDefaultValParams struct {
	FormLinkIDHeader
	DefaultValues []record.DefaultFieldValue `json:"defaultValues"`
}

func (updateParams FormLinkDefaultValParams) updateProps(trackerDBHandle *sql.DB, linkForUpdate *FormLink) error {

	if validateErr := record.ValidateWellFormedDefaultValues(trackerDBHandle, updateParams.DefaultValues); validateErr != nil {
		return fmt.Errorf("updateProps: invalid default value(s): %v", validateErr)
	}

	linkForUpdate.Properties.DefaultValues = updateParams.DefaultValues

	return nil
}

type FormLinkNameParams struct {
	FormLinkIDHeader
	NewName string `json:"newName"`
}

func (updateParams FormLinkNameParams) updateProps(trackerDBHandle *sql.DB, linkForUpdate *FormLink) error {

	if !stringValidation.WellFormedItemLabel(updateParams.NewName) {
		return fmt.Errorf("Can't update form link name: invalid name: %v", updateParams.NewName)
	}

	linkForUpdate.Name = updateParams.NewName

	return nil
}

type FormLinkFormParams struct {
	FormLinkIDHeader
	FormID string `json:"formID"`
}

func (updateParams FormLinkFormParams) updateProps(trackerDBHandle *sql.DB, linkForUpdate *FormLink) error {

	newForm, getErr := form.GetForm(trackerDBHandle, updateParams.FormID)
	if getErr != nil {
		return fmt.Errorf("Update form properties: Unable to get form specified as new form: %v", getErr)
	}

	oldForm, getErr := form.GetForm(trackerDBHandle, linkForUpdate.FormID)
	if getErr != nil {
		return fmt.Errorf("FormLinkFormParams.updateProps: Unable to get existing form before setting new form: %v", getErr)
	}

	if newForm.ParentDatabaseID != oldForm.ParentDatabaseID {
		return fmt.Errorf("FormLinkFormParams.updateProps: Database mismatch for new form: %v", newForm.ParentDatabaseID)
	}

	linkForUpdate.FormID = updateParams.FormID

	return nil
}

type FormLinkIncludeInSidebarParams struct {
	FormLinkIDHeader
	IncludeInSidebar bool `json:"includeInSidebar"`
}

func (updateParams FormLinkIncludeInSidebarParams) updateProps(trackerDBHandle *sql.DB, linkForUpdate *FormLink) error {

	linkForUpdate.IncludeInSidebar = updateParams.IncludeInSidebar

	return nil
}

type FormLinkEnableSharedLinkParams struct {
	FormLinkIDHeader
}

func (updateParams FormLinkEnableSharedLinkParams) updateProps(trackerDBHandle *sql.DB, linkForUpdate *FormLink) error {

	if len(linkForUpdate.SharedLinkID) <= 0 {
		// Generate a new shared link ID if there insn't one already (from previously enabling the shared link)
		linkForUpdate.SharedLinkID = uniqueID.GenerateUniqueID()
	}
	linkForUpdate.SharedLinkEnabled = true

	return nil
}

type FormLinkDisableSharedLinkParams struct {
	FormLinkIDHeader
}

func (updateParams FormLinkDisableSharedLinkParams) updateProps(trackerDBHandle *sql.DB, linkForUpdate *FormLink) error {

	linkForUpdate.SharedLinkEnabled = false

	return nil
}
